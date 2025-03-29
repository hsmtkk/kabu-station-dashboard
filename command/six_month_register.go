package command

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/hsmtkk/kabu-station-dashboard/api"
	"github.com/hsmtkk/kabu-station-dashboard/misc"
	"github.com/hsmtkk/kabu-station-dashboard/util"
	"github.com/spf13/cobra"
)

var SixMonthRegister = &cobra.Command{
	Use: "six_month_register",
	Run: sixMonthRegister,
}

func sixMonthRegister(cmd *cobra.Command, args []string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := misc.RequiredEnvVar("KABU_STATION_API_PASSWORD")
	apiClient, err := api.New(logger, apiPassword)
	if err != nil {
		log.Fatal(err)
	}
	utility := util.New(logger, apiClient)
	atm, err := utility.AtTheMoney()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("At the money: %d\n", atm)
	fistMonth, err := utility.FirstMonth()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First month: %v\n", fistMonth)
}
