package unregister_all_put

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Handle(logger *slog.Logger, baseURL string, token string) error {
	logger.Debug("Handle begin")
	clt := resty.New()
	url := fmt.Sprintf("%s/unregister/all", baseURL)
	result := Response{}
	resp, err := clt.R().SetHeader("X-API-KEY", token).SetResult(&result).Put(url)
	if err != nil {
		return fmt.Errorf("resty.Put failed: %w", err)
	}
	logger.Debug("Handle", "resp", resp.String())
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("got %d HTTP status code: %s", resp.StatusCode(), resp.Status())
	}
	if result.Code != 0 {
		return fmt.Errorf("got %d code: %s", result.Code, result.Message)
	}
	logger.Debug("Handle end")
	return nil
}
