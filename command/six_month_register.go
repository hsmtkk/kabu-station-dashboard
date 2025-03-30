package command

import (
	"encoding/csv"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"

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
	records, err := register()
	if err != nil {
		log.Fatal(err)
	}
	if err := saveCSV(records); err != nil {
		log.Fatal(err)
	}
}

func register() ([][]string, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := misc.RequiredEnvVar("KABU_STATION_API_PASSWORD")
	apiClient, err := api.New(logger, apiPassword)
	if err != nil {
		return nil, err
	}
	if err := apiClient.UnregisterAllPut(); err != nil {
		return nil, err
	}
	utility := util.New(logger, apiClient)
	atm, err := utility.AtTheMoney()
	if err != nil {
		return nil, err
	}
	firstMonth, err := utility.FirstMonth()
	if err != nil {
		return nil, err
	}

	records := [][]string{}
	month := firstMonth
	for i := 0; i < 6; i++ {
		misc.Interval()
		putSymbol, err := apiClient.SymbolnameOptionGet(symbolname_option_get.Request{OptionCode: symbolname_option_get.NK225op, DerivMonth: &month, PutOrCall: symbolname_option_get.Put, StrikePrice: atm})
		if err != nil {
			return nil, err
		}
		callSymbol, err := apiClient.SymbolnameOptionGet(symbolname_option_get.Request{OptionCode: symbolname_option_get.NK225op, DerivMonth: &month, PutOrCall: symbolname_option_get.Call, StrikePrice: atm})
		if err != nil {
			return nil, err
		}
		req := register_put.Request{
			Symbols: []register_put.SymbolExchange{
				{Symbol: putSymbol.Symbol, Exchange: board_get.WholeDay},
				{Symbol: callSymbol.Symbol, Exchange: board_get.WholeDay},
			},
		}
		if _, err := apiClient.RegisterPut(req); err != nil {
			return nil, err
		}
		records = append(records, []string{month.Format("2006-01"), putSymbol.Symbol, callSymbol.Symbol})
		month = month.AddDate(0, 1, 0)
	}
	return records, nil
}

func saveCSV(records [][]string) error {
	todayDir := misc.MakeTodayDataDirectory()
	filePath := path.Join(todayDir, "six_month_symbol.csv")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("os.Create failed %s: %w", filePath, err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	csvHeader := []string{"#year_month", "put_symbol", "call_symbol"}
	writer.Write(csvHeader)
	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("csv.WriteAll failed: %w", err)
	}
	return nil
}
