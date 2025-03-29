package board_get

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type MarketCode int

const (
	Tokyo    MarketCode = 1
	WholeDay MarketCode = 2
	Day      MarketCode = 23
	Night    MarketCode = 24
)

type Request struct {
	Symbol     string
	MarketCode MarketCode
}

type Response struct {
	Code         int     `json:"Code"`
	Message      string  `json:"Message"`
	Symbol       string  `json:"Symbol"`
	SymbolName   string  `json:"SymbolName"`
	CurrentPrice float64 `json:"CurrentPrice"`
	IV           float64 `json:"IV"`
	Gamma        float64 `json:"Gamma"`
	Theta        float64 `json:"Theta"`
	Vega         float64 `json:"Vega"`
	Delta        float64 `json:"Delta"`
}

func Handle(logger *slog.Logger, baseURL string, token string, req Request) (Response, error) {
	logger.Debug("Handle begin")
	clt := resty.New()
	symbol := fmt.Sprintf("%s@%s", req.Symbol, strconv.Itoa(int(req.MarketCode)))
	url := fmt.Sprintf("%s/board/%s", baseURL, symbol)
	result := Response{}
	resp, err := clt.R().SetHeader("X-API-KEY", token).SetResult(&result).Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("resty.Get failed: %w", err)
	}
	logger.Debug("Handle", "resp", resp.String())
	if resp.StatusCode() != http.StatusOK {
		return Response{}, fmt.Errorf("got %d HTTP status code: %s", resp.StatusCode(), resp.Status())
	}
	if result.Code != 0 {
		return Response{}, fmt.Errorf("got %d code: %s", result.Code, result.Message)
	}
	logger.Debug("Handle end")
	return result, nil
}
