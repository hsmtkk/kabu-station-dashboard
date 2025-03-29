package main

import (
	"log"

	"github.com/hsmtkk/kabu-station-dashboard/command"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}
	cmd.AddCommand(command.SixMonthRegister)
	cmd.AddCommand(command.SixMonthIVChart)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
