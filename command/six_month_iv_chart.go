package command

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/hsmtkk/kabu-station-dashboard/api"
	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
	"github.com/hsmtkk/kabu-station-dashboard/misc"
	"github.com/spf13/cobra"
)

var SixMonthIVChart = &cobra.Command{
	Use: "six_month_iv_chart",
	Run: sixMonthIVChart,
}

func sixMonthIVChart(cmd *cobra.Command, args []string) {
	records, err := readSymbols()
	if err != nil {
		log.Fatal(err)
	}
	if err := makeJSON(records); err != nil {
		log.Fatal(err)
	}
}

type CSVRecord struct {
	YearMonth  time.Time
	PutSymbol  string
	CallSymbol string
}

func readSymbols() ([]CSVRecord, error) {
	todayDir := misc.MakeTodayDataDirectory()
	filePath := path.Join(todayDir, "six_month_symbol.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed %s: %w", filePath, err)
	}
	result := []CSVRecord{}
	reader := csv.NewReader(file)
	reader.Comment = '#'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv.ReadAll failed: %w", err)
	}
	for _, record := range records {
		parsed, err := time.Parse("2006-01", record[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", record[0], err)
		}
		result = append(result, CSVRecord{
			YearMonth:  parsed,
			PutSymbol:  record[1],
			CallSymbol: record[2],
		})
	}
	return result, nil
}

type jsonSchema struct {
	YearMonth []string `json:"year_month"`
	PutIV     []string `json:"put_iv"`
	CallIV    []string `json:"call_iv"`
}

func makeJSON(records []CSVRecord) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := misc.RequiredEnvVar("KABU_STATION_API_PASSWORD")
	apiClient, err := api.New(logger, apiPassword)
	if err != nil {
		return err
	}
	result := jsonSchema{}
	for _, record := range records {
		putBoard, err := apiClient.BoardGet(board_get.Request{Symbol: record.PutSymbol, MarketCode: board_get.WholeDay})
		if err != nil {
			return err
		}
		callBoard, err := apiClient.BoardGet(board_get.Request{Symbol: record.CallSymbol, MarketCode: board_get.WholeDay})
		if err != nil {
			return err
		}
		result.YearMonth = append(result.YearMonth, record.YearMonth.Format("2006-01"))
		result.PutIV = append(result.PutIV, fmt.Sprintf("%f", putBoard.IV))
		result.CallIV = append(result.CallIV, fmt.Sprintf("%f", callBoard.IV))
	}

	jsonBytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent failed: %w", err)
	}

	todayDir := misc.MakeTodayDataDirectory()
	filePath := path.Join(todayDir, "six_month_iv.json")
	if err := os.WriteFile(filePath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("os.WriteFile failed %s: %w", filePath, err)
	}
	return nil
}
