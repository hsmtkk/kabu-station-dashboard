package command

import (
	"log"
	"log/slog"
	"os"

	"github.com/hsmtkk/kabu-station-dashboard/api"
	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/register_put"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_option_get"
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
	if err := apiClient.UnregisterAllPut(); err != nil {
		log.Fatal(err)
	}
	utility := util.New(logger, apiClient)
	atm, err := utility.AtTheMoney()
	if err != nil {
		log.Fatal(err)
	}
	firstMonth, err := utility.FirstMonth()
	if err != nil {
		log.Fatal(err)
	}

	month := firstMonth
	for i := 0; i < 6; i++ {
		misc.Interval()
		putSymbol, err := apiClient.SymbolnameOptionGet(symbolname_option_get.Request{OptionCode: symbolname_option_get.NK225op, DerivMonth: &month, PutOrCall: symbolname_option_get.Put, StrikePrice: atm})
		if err != nil {
			log.Fatal(err)
		}
		callSymbol, err := apiClient.SymbolnameOptionGet(symbolname_option_get.Request{OptionCode: symbolname_option_get.NK225op, DerivMonth: &month, PutOrCall: symbolname_option_get.Call, StrikePrice: atm})
		if err != nil {
			log.Fatal(err)
		}
		req := register_put.Request{
			Symbols: []register_put.SymbolExchange{
				{Symbol: putSymbol.Symbol, Exchange: board_get.WholeDay},
				{Symbol: callSymbol.Symbol, Exchange: board_get.WholeDay},
			},
		}
		if _, err := apiClient.RegisterPut(req); err != nil {
			log.Fatal(err)
		}
		month = month.AddDate(0, 1, 0)
	}
}
