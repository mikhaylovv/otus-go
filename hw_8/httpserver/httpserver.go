package httpserver

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// HTTPServer - simple realization of http server, using fasthttp
type HTTPServer struct {
	Logger *zap.Logger
}

func handle(ctx *fasthttp.RequestCtx, lg *zap.Logger) {
	lg.Debug(ctx.Request.String())

	switch string(ctx.Path()) {
	case "/hello":
		ctx.Response.SetBody([]byte("hello, stranger"))
		lg.Info("hello message received")

	default:
		lg.Info("unsupported path err", zap.String("ctx", ctx.String()))
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

// StartListen - starts listen http requests in current goroutine
func (h *HTTPServer) StartListen(addr string) error {
	return fasthttp.ListenAndServe(
		addr,
		func(ctx *fasthttp.RequestCtx) {
			handle(ctx, h.Logger)
		},
	)
}
