package util

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/hsmtkk/kabu-station-dashboard/api"
	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_future_get"
)

type Utility interface {
	AtTheMoney() (int, error)
	FirstMonth() (time.Time, error)
}

type utilityImpl struct {
	logger    *slog.Logger
	apiClient api.Client
}

func New(logger *slog.Logger, apiClient api.Client) Utility {
	return &utilityImpl{logger, apiClient}
}

func (u *utilityImpl) AtTheMoney() (int, error) {
	u.logger.Debug("AtTheMoney begin")
	symbolResp, err := u.apiClient.SymbolnameFutureGet(symbolname_future_get.Request{FutureCode: symbolname_future_get.NK225mini, DerivMonth: nil})
	if err != nil {
		return 0, err
	}
	boardResp, err := u.apiClient.BoardGet(board_get.Request{Symbol: symbolResp.Symbol, MarketCode: board_get.WholeDay})
	if err != nil {
		return 0, err
	}
	result := boardResp.CurrentPrice
	rounded := int(result/250) * 250
	u.logger.Debug("AtTheMoney end", "result", rounded)
	return rounded, nil
}

func (u *utilityImpl) FirstMonth() (time.Time, error) {
	u.logger.Debug("FirstMonth begin")
	symbolResp, err := u.apiClient.SymbolnameFutureGet(symbolname_future_get.Request{FutureCode: symbolname_future_get.NK225mini, DerivMonth: nil})
	if err != nil {
		return time.Time{}, err
	}
	elems := strings.Split(symbolResp.SymbolName, " ")
	if len(elems) != 2 {
		return time.Time{}, fmt.Errorf("failed to parse symbolname: %s", symbolResp.SymbolName)
	}
	result, err := time.Parse("06/01", elems[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %s", elems[1])
	}
	u.logger.Debug("FirstMonth end", "result", result)
	return result, nil
}
