package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tesla-home-powerflow-optimizer",
	Short: "Automatically start, stop, and adjust the amperage of your Tesla Home Charging system based on surplus power availability. Maximize your energy efficiency and reduce costs by optimizing your home charging process with real-time power management.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newServeCommand(), newAuthenticateCommand(), newTestCmd())
}
