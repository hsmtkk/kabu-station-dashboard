package symbolname_option_get

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type OptionCode string

const (
	NK225op     OptionCode = "NK225op"
	NK225miniop OptionCode = "NK225miniop"
)

type PutOrCall string

const (
	Put  PutOrCall = "P"
	Call PutOrCall = "C"
)

type Request struct {
	OptionCode  OptionCode
	DerivMonth  *time.Time
	PutOrCall   PutOrCall
	StrikePrice int
}

type Response struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func Handle(logger *slog.Logger, baseURL string, token string, req Request) (Response, error) {
	logger.Debug("Handle begin", "baseURL", baseURL, "req", req)
	clt := resty.New()
	url := fmt.Sprintf("%s/symbolname/option", baseURL)
	year_month := "0"
	if req.DerivMonth != nil {
		year_month = req.DerivMonth.Format("200601")
	}
	result := Response{}
	resp, err := clt.R().SetHeader("X-API-KEY", token).SetQueryParams(map[string]string{"OptionCode": string(req.OptionCode), "DerivMonth": year_month, "PutOrCall": string(req.PutOrCall), "StrikePrice": strconv.Itoa(req.StrikePrice)}).SetResult(&result).Get(url)
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
	logger.Debug("Handle end", "result", result)
	return result, nil
}
