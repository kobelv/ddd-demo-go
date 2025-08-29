package interfaces

import (
	"context"
	"net/http"

	"ddd-demo-go/infrastructure/common/logit"
	ih "ddd-demo-go/interfaces/http"
)

type app struct {
	ctx        context.Context
	logger     logit.LoggerInterface
	conf       *AppConf
	httpServer *ih.Server
}

func (a *app) StartServers() error {
	return a.httpServer.Start()
}

func (a *app) BeforeShutdown() {
	a.logger.Cleanup()
}

func newHTTPServer(conf *AppConf, handler http.Handler) (*ih.Server, error) {
	return ih.NewServer(&ih.Conf{
		Addr:         conf.HTTPServer.Addr,
		ReadTimeout:  conf.HTTPServer.ReadTimeout,
		WriteTimeout: conf.HTTPServer.WriteTimeout,
		IdleTimeout:  conf.HTTPServer.IdleTimeout,
	}, handler)
}
