package api

import (
	"log/slog"

	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_future_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_option_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/token"
)

const BASE_URL = "http://localhost:18080/kabusapi"

type Client interface {
	BoardGet(req board_get.Request) (board_get.Response, error)
	SymbolnameFutureGet(req symbolname_future_get.Request) (symbolname_future_get.Response, error)
	SymbolnameOptionGet(req symbolname_option_get.Request) (symbolname_option_get.Response, error)
}

type clientImpl struct {
	logger *slog.Logger
	token  string
}

func New(logger *slog.Logger, apiPassword string) (Client, error) {
	clt := &clientImpl{logger: logger}
	if err := clt.setToken(token.Request{APIPassword: apiPassword}); err != nil {
		return nil, err
	}
	return clt, nil
}

func (c *clientImpl) setToken(req token.Request) error {
	c.logger.Debug("setToken begin")
	resp, err := token.Handle(c.logger, BASE_URL, req)
	if err != nil {
		return err
	}
	c.token = resp.Token
	c.logger.Debug("setToken end")
	return nil
}

func (c *clientImpl) BoardGet(req board_get.Request) (board_get.Response, error) {
	c.logger.Debug("BoardGet begin")
	resp, err := board_get.Handle(c.logger, BASE_URL, c.token, req)
	if err != nil {
		return board_get.Response{}, err
	}
	c.logger.Debug("BoardGet end")
	return resp, nil
}

func (c *clientImpl) SymbolnameFutureGet(req symbolname_future_get.Request) (symbolname_future_get.Response, error) {
	c.logger.Debug("SymbolnameFutureGet begin")
	resp, err := symbolname_future_get.Handle(c.logger, BASE_URL, c.token, req)
	if err != nil {
		return symbolname_future_get.Response{}, err
	}
	c.logger.Debug("SymbolnameFutureGet end")
	return resp, nil
}

func (c *clientImpl) SymbolnameOptionGet(req symbolname_option_get.Request) (symbolname_option_get.Response, error) {
	return symbolname_option_get.Response{}, nil
}
