package api

import (
	"log/slog"

	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/register_put"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_future_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_option_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/token"
	"github.com/hsmtkk/kabu-station-dashboard/api/unregister_all_put"
)

const BASE_URL = "http://localhost:18080/kabusapi"

type Client interface {
	BoardGet(req board_get.Request) (board_get.Response, error)
	RegisterPut(req register_put.Request) (register_put.Response, error)
	SymbolnameFutureGet(req symbolname_future_get.Request) (symbolname_future_get.Response, error)
	SymbolnameOptionGet(req symbolname_option_get.Request) (symbolname_option_get.Response, error)
	UnregisterAllPut() error
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

func (c *clientImpl) RegisterPut(req register_put.Request) (register_put.Response, error) {
	c.logger.Debug("RegisterPut begin")
	resp, err := register_put.Handle(c.logger, BASE_URL, c.token, req)
	if err != nil {
		return register_put.Response{}, err
	}
	c.logger.Debug("RegisterPut end")
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
	c.logger.Debug("SymbolnameOptionGet begin")
	resp, err := symbolname_option_get.Handle(c.logger, BASE_URL, c.token, req)
	if err != nil {
		return symbolname_option_get.Response{}, err
	}
	c.logger.Debug("SymbolnameOptionGet end")
	return resp, nil
}

func (c *clientImpl) UnregisterAllPut() error {
	c.logger.Debug("UnregisterAllPut begin")
	if err := unregister_all_put.Handle(c.logger, BASE_URL, c.token); err != nil {
		return err
	}
	c.logger.Debug("UnregisterAllPut end")
	return nil
}
