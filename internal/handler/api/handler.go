package api

import (
	"sagara_backend_test/config"
	"sagara_backend_test/internal/handler/api/controller"
	"sagara_backend_test/internal/usecases"
	"sagara_backend_test/lib/log"
	"sagara_backend_test/lib/router"
)

type Options struct {
	Cfg        config.MainConfig
	WardrobeUc usecases.WardrobeUseCases
}

type Handler struct {
	opts        *Options
	listenErrCh chan error
	myRouter    *router.FastRouter
}

func New(opts *Options) *Handler {
	handler := &Handler{opts: opts}
	handler.myRouter = controller.New(&controller.Options{
		Prefix:         opts.Cfg.API.BasePath,
		Port:           opts.Cfg.Server.Port,
		ReadTimeout:    opts.Cfg.Server.ReadTimeout,
		WriteTimeout:   opts.Cfg.Server.WriteTimeout,
		RequestTimeout: opts.Cfg.API.APITimeout,
		EnableSwagger:  opts.Cfg.API.EnableSwagger,
		WardrobeUc:     opts.WardrobeUc,
	}).RegisterRoute()

	return handler
}

func (h *Handler) Run() {
	log.Infof("API Listening on %d", h.opts.Cfg.Server.Port)
	h.listenErrCh <- h.myRouter.StartServe()
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
