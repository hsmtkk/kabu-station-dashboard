package register_put

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
)

type Request struct {
	Symbols []SymbolExchange `json:"Symbols"`
}

type Response struct {
	Code       int              `json:"Code"`
	Message    string           `json:"Message"`
	RegistList []SymbolExchange `json:"RegistList"`
}

type SymbolExchange struct {
	Symbol   string               `json:"Symbol"`
	Exchange board_get.MarketCode `json:"Exchange"`
}

func Handle(logger *slog.Logger, baseURL string, token string, req Request) (Response, error) {
	logger.Debug("Handle begin", "baseURL", baseURL, "req", req)
	clt := resty.New()
	url := fmt.Sprintf("%s/register", baseURL)
	result := Response{}
	resp, err := clt.R().SetHeader("X-API-KEY", token).SetBody(req).SetResult(&result).Put(url)
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
