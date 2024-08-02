package main

import (
	"fmt"
	"os"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
