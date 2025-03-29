package command

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/hsmtkk/kabu-station-dashboard/api"
	"github.com/hsmtkk/kabu-station-dashboard/misc"
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
	clt, err := api.New(logger, apiPassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", clt)
}
