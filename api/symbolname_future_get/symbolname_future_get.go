package symbolname_future_get

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type FutureCode string

const (
	NK225      FutureCode = "NK225"
	NK225mini  FutureCode = "NK225mini"
	NK225micro FutureCode = "NK225micro"
	VI         FutureCode = "VI"
)

type Request struct {
	FutureCode FutureCode
	DerivMonth *time.Time
}

type Response struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func Handle(logger *slog.Logger, baseURL string, token string, req Request) (Response, error) {
	logger.Debug("Handle begin")
	clt := resty.New()
	url := fmt.Sprintf("%s/symbolname/future", baseURL)
	year_month := "0"
	if req.DerivMonth != nil {
		year_month = req.DerivMonth.Format("200601")
	}
	result := Response{}
	resp, err := clt.R().SetHeader("X-API-KEY", token).SetQueryParams(map[string]string{"FutureCode": string(req.FutureCode), "DerivMonth": year_month}).SetResult(&result).Get(url)
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
