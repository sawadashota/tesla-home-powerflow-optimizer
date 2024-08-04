// Package tesla
// https://developer.tesla.com/docs/fleet-api?shell#vehicle-commands
package tesla

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/teslamotors/vehicle-command/pkg/account"
	"github.com/teslamotors/vehicle-command/pkg/protocol"
	"github.com/teslamotors/vehicle-command/pkg/vehicle"
)

func (c *Client) StartCharge(ctx context.Context, vin string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	car, err := c.startCommandSession(ctx, vin)
	if err != nil {
		return err
	}
	c.r.Logger().Info("starting charge", slog.String("vin", car.VIN()))
	if err := car.ChargeStart(ctx); err != nil {
		if protocol.MayHaveSucceeded(err) {
			return fmt.Errorf("command sent, but client could not confirm receipt: %w", err)
		}
		return fmt.Errorf("failed to start charge: %w", err)
	}
	return nil
}

func (c *Client) StopCharge(ctx context.Context, vin string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	car, err := c.startCommandSession(ctx, vin)
	if err != nil {
		return err
	}
	c.r.Logger().Info("stopping charge", slog.String("vin", car.VIN()))
	if err := car.ChargeStop(ctx); err != nil {
		if protocol.MayHaveSucceeded(err) {
			return fmt.Errorf("command sent, but client could not confirm receipt: %w", err)
		}
		return fmt.Errorf("failed to stop charge: %w", err)
	}
	return nil
}

func (c *Client) SetChargeLimit(ctx context.Context, vin string, percent int32) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	car, err := c.startCommandSession(ctx, vin)
	if err != nil {
		return err
	}
	c.r.Logger().Info("setting charge limit", slog.String("vin", car.VIN()))
	if err := car.ChangeChargeLimit(ctx, percent); err != nil {
		if protocol.MayHaveSucceeded(err) {
			return fmt.Errorf("command sent, but client could not confirm receipt: %w", err)
		}
		return fmt.Errorf("failed to change charge limit: %w", err)
	}
	return nil
}

func (c *Client) SetChargeAmps(ctx context.Context, vin string, amps int32) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	car, err := c.startCommandSession(ctx, vin)
	if err != nil {
		return err
	}
	c.r.Logger().Info("setting charge amps", slog.String("vin", car.VIN()))
	if err := car.SetChargingAmps(ctx, amps); err != nil {
		if protocol.MayHaveSucceeded(err) {
			return fmt.Errorf("command sent, but client could not confirm receipt: %w", err)
		}
		return fmt.Errorf("failed to set charge amps: %w", err)
	}
	return nil
}

func (c *Client) startCommandSession(ctx context.Context, vin string) (*vehicle.Vehicle, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get a token: %w", err)
	}

	acct, err := account.New(token.AccessToken, userAgent)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize a new account: %w", err)
	}
	car, err := acct.GetVehicle(ctx, vin, c.r.TeslaOAuthConfig().PrivateKey(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get a vehicle: %w", err)
	}

	c.r.Logger().Info("connecting to the vehicle", slog.String("vin", car.VIN()))
	if err := car.Connect(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to the vehicle: %w", err)
	}
	c.r.Logger().Info("starting a new session", slog.String("vin", car.VIN()))
	if err := car.StartSession(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to to perform handshake with vehicle: %w", err)
	}
	return car, nil
}
