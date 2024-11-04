package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Run test part of the application",
	}
	cmd.AddCommand(newTestStartChargingCmd(), newTestCollectorCmd())
	return cmd
}

func newTestStartChargingCmd() *cobra.Command {
	var r driver.ServerRegistry
	cmd := &cobra.Command{
		Use:   "start-charging",
		Short: "Run test for start charging",
		Long: "This command will wake up the vehicle and start charging." +
			"Then it will print the charge state of the vehicle." +
			"Please make sure that the vehicle is connected to the charger.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			r, err = driver.NewServerRegistry(cmd.Context())
			if err != nil {
				return err
			}
			return setup(cmd.Context(), r)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), time.Minute)
			defer cancel()

			vin := r.AppConfig().TeslaVIN
			fmt.Println("waking up the vehicle", vin)
			if err := r.VehicleRepository().WakeUp(ctx, vin); err != nil {
				return err
			}
			fmt.Println("waiting for vehicle to wake up...")
			if _, err := r.VehicleRepository().WaitUntilWakedUp(ctx, vin); err != nil {
				return err
			}

			fmt.Println("start charging...")
			if err := r.VehicleRepository().StartCharge(ctx, vin); err != nil {
				return err
			}

			fmt.Println("start charging successfully!")
			return nil
		},
	}
	return cmd
}

func newTestCollectorCmd() *cobra.Command {
	var r driver.ServerRegistry
	cmd := &cobra.Command{
		Use:   "collector",
		Short: "Run test collector part of the application",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			r, err = driver.NewServerRegistry(cmd.Context())
			if err != nil {
				return err
			}
			return setup(cmd.Context(), r)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), time.Minute)
			defer cancel()

			{
				metric, err := r.Collector().GetSurplusPower(ctx)
				if err != nil {
					return err
				}
				fmt.Printf("surplus power: %d W\n", metric.Watt)
			}
			{
				metric, err := r.Collector().GetEVUsagePower(ctx)
				if err != nil {
					return err
				}
				fmt.Printf("EV usage power: %d W\n", metric.Watt)
			}

			return nil
		},
	}
	return cmd
}
