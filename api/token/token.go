package token

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Request struct {
	APIPassword string `json:"APIPassword"`
}

type Response struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	ResultCode int    `json:"ResultCode"`
	Token      string `json:"Token"`
}

func Handle(logger *slog.Logger, url string, req Request) (Response, error) {
	logger.Debug("handle begin")
	clt := resty.New()
	result := Response{}
	resp, err := clt.R().SetHeader("Content-Type", "application/json").SetBody(map[string]string{"APIPassword": req.APIPassword}).SetResult(&result).Post(url)
	if err != nil {
		return Response{}, fmt.Errorf("resty.Post failed: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return Response{}, fmt.Errorf("got %d HTTP status code: %s", resp.StatusCode(), resp.Status())
	}
	if result.Code != 0 {
		return Response{}, fmt.Errorf("got %d code: %s", result.Code, result.Message)
	}
	logger.Debug("handle end")
	return result, nil
}
