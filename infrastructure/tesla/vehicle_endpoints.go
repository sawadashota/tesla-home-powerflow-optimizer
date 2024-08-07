// Package tesla
// Description: This file contains the implementation of the Tesla API endpoints for vehicle data.
// see https://developer.tesla.com/docs/fleet-api?shell#vehicle-endpoints
package tesla

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

func (c *Client) GetVehicle(ctx context.Context, vin string) (*model.VehicleSummary, error) {
	url := fmt.Sprintf("%s/api/1/vehicles/%s", c.r.TeslaOAuthConfig().APIHost, vin)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	c.r.Logger().Info("requesting the Tesla API", slog.String("url", req.URL.String()), slog.String("method", req.Method))
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return nil, failure.New(model.ErrCodeUnauthorized, failure.Message("unauthorized"))
		case http.StatusNotFound:
			return nil, failure.New(model.ErrCodeNotFound, failure.Messagef("vehicle (%s) not found", c.r.AppConfig().TeslaVIN))
		default:
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, failure.New(model.ErrCodeInternalServer, failure.Messagef("unexpected status code: %d, failed to read response body: %s", resp.StatusCode, err))
			}
			return nil, failure.New(model.ErrCodeInternalServer, failure.Messagef("unexpected status code: %d, response body: %s", resp.StatusCode, string(body)))
		}
	}

	var v VehicleResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &model.VehicleSummary{
		VIN:   v.Response.Vin,
		State: model.VehicleState(v.Response.State),
	}, nil
}

func (c *Client) GetVehicleData(ctx context.Context, vin string) (*model.VehicleData, error) {
	url := fmt.Sprintf("%s/api/1/vehicles/%s/vehicle_data", c.r.TeslaOAuthConfig().APIHost, vin)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	c.r.Logger().Info("requesting the Tesla API", slog.String("url", req.URL.String()), slog.String("method", req.Method))
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusRequestTimeout:
			return nil, failure.New(model.ErrCodeDeviceOffline, failure.Message("device is offline"))
		case http.StatusNotFound:
			return nil, failure.New(model.ErrCodeNotFound, failure.Messagef("vehicle (%s) not found", c.r.AppConfig().TeslaVIN))
		case http.StatusUnauthorized:
			return nil, failure.New(model.ErrCodeUnauthorized, failure.Message("unauthorized"))
		default:
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, failure.New(model.ErrCodeInternalServer, failure.Messagef("unexpected status code: %d, failed to read response body: %s", resp.StatusCode, err))
			}
			return nil, failure.New(model.ErrCodeInternalServer, failure.Messagef("unexpected status code: %d, response body: %s", resp.StatusCode, string(body)))
		}
	}

	var v VehicleDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &model.VehicleData{
		VIN:   v.Response.Vin,
		State: model.VehicleState(v.Response.State),
		ChargeState: model.VehicleChargeState{
			VIN:                     v.Response.Vin,
			BatteryLevel:            v.Response.ChargeState.BatteryLevel,
			BatteryRange:            v.Response.ChargeState.BatteryRange,
			ChargeAmps:              v.Response.ChargeState.ChargeAmps,
			ChargeCurrentRequest:    v.Response.ChargeState.ChargeCurrentRequest,
			ChargeCurrentRequestMax: v.Response.ChargeState.ChargeCurrentRequestMax,
			ChargeEnableRequest:     v.Response.ChargeState.ChargeEnableRequest,
			ChargeLimitSoc:          v.Response.ChargeState.ChargeLimitSoc,
			ChargePortDoorOpen:      v.Response.ChargeState.ChargePortDoorOpen,
			ChargePortLatch:         v.Response.ChargeState.ChargePortLatch,
			ChargerActualCurrent:    v.Response.ChargeState.ChargerActualCurrent,
			ChargerVoltage:          v.Response.ChargeState.ChargerVoltage,
			ChargingState:           model.ChargingState(v.Response.ChargeState.ChargingState),
			MinutesToFullCharge:     v.Response.ChargeState.MinutesToFullCharge,
			Timestamp:               time.Unix(v.Response.ChargeState.Timestamp, 0),
			UsableBatteryLevel:      v.Response.ChargeState.UsableBatteryLevel,
		},
	}, nil
}

func (c *Client) WaitUntilWakedUp(ctx context.Context, vin string) (*model.VehicleData, error) {
	ticker := time.NewTicker(8 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			c.r.Logger().Info("checking if the vehicle is online", slog.String("vin", vin))
			data, err := c.GetVehicleData(ctx, vin)
			if err != nil {
				if failure.Is(err, model.ErrCodeDeviceOffline) {
					continue
				}
				return nil, err
			}
			return data, nil
		}
	}
}

func (c *Client) WakeUp(ctx context.Context, vin string) error {
	url := fmt.Sprintf("%s/api/1/vehicles/%s/wake_up", c.r.TeslaOAuthConfig().APIHost, vin)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	c.r.Logger().Info("requesting the Tesla API", slog.String("url", req.URL.String()), slog.String("method", req.Method))
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return failure.New(model.ErrCodeUnauthorized, failure.Message("unauthorized"))
		case http.StatusNotFound:
			return failure.New(model.ErrCodeNotFound, failure.Messagef("vehicle (%s) not found", c.r.AppConfig().TeslaVIN))
		default:
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return failure.New(model.ErrCodeInternalServer, failure.Messagef("unexpected status code: %d, failed to read response body: %s", resp.StatusCode, err))
			}
			return failure.New(model.ErrCodeInternalServer, failure.Messagef("unexpected status code: %d, response body: %s", resp.StatusCode, string(body)))
		}
	}

	return nil
}
