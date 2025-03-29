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

func Handle(logger *slog.Logger, baseURL string, req Request) (Response, error) {
	logger.Debug("Handle begin")
	clt := resty.New()
	url := fmt.Sprintf("%s/token", baseURL)
	result := Response{}
	resp, err := clt.R().SetHeader("Content-Type", "application/json").SetBody(map[string]string{"APIPassword": req.APIPassword}).SetResult(&result).Post(url)
	if err != nil {
		return Response{}, fmt.Errorf("resty.Post failed: %w", err)
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
