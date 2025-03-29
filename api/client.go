package api

import (
	"fmt"
	"log/slog"

	"github.com/hsmtkk/kabu-station-dashboard/api/board_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/symbolname_option_get"
	"github.com/hsmtkk/kabu-station-dashboard/api/token"
)

const LIVE_PORT = 18080

type Client interface {
	BoardGet(req board_get.Request) (board_get.Response, error)
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

func (c *clientImpl) makeURL(path string) string {
	return fmt.Sprintf("http://localhost:%d/kabusapi%s", LIVE_PORT, path)
}

func (c *clientImpl) setToken(req token.Request) error {
	c.logger.Debug("setToken begin")
	url := c.makeURL("/token")
	resp, err := token.Handle(c.logger, url, req)
	if err != nil {
		return err
	}
	c.token = resp.Token
	c.logger.Debug("setToken end")
	return nil
}

func (c *clientImpl) BoardGet(req board_get.Request) (board_get.Response, error) {
	return board_get.Response{}, nil
}

func (c *clientImpl) SymbolnameOptionGet(req symbolname_option_get.Request) (symbolname_option_get.Response, error) {
	return symbolname_option_get.Response{}, nil
}
