package httpserver

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"go.uber.org/zap/zaptest"
	"net"
	"testing"
)


type testLogger struct {
	t *testing.T
}

func (t testLogger) Printf(format string, args ...interface{}) {
	t.t.Logf(format, args...)
}

type fixture struct {
	Server   *fasthttp.Server
	Listener *fasthttputil.InmemoryListener
	Conn     net.Conn
}

func setUpHandleTest(t *testing.T) fixture {
	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			l := zaptest.NewLogger(t)
			handle(ctx, l)
		},
		Logger: &testLogger{},
	}

	ln := fasthttputil.NewInmemoryListener()

	go func() {
		if err := s.Serve(ln); err != nil {
			assert.FailNow(t, "Error serve listener", err)
		}
	}()

	c, err := ln.Dial()
	if err != nil {
		assert.FailNow(t, "Error dial", err)
	}

	return fixture{s, ln, c}
}

func tearDownHandleTest(t *testing.T, f *fixture) {
	if err := f.Conn.Close(); err != nil {
		assert.Fail(t, "Error close connection", err)
	}
	if err := f.Listener.Close(); err != nil {
		assert.Fail(t, "Error close listener", err)
	}
}

func TestHandlerUnsupportedPath(t *testing.T) {
	f := setUpHandleTest(t)
	defer tearDownHandleTest(t, &f)

	req := fasthttp.Request{}
	req.Header.SetMethod("GET")
	req.SetHost("test")
	req.URI().Update("/test")

	if _, err := f.Conn.Write([]byte(req.String())); err != nil {
		assert.Fail(t, "Error write in connection", err)
	}

	br := bufio.NewReader(f.Conn)
	var resp fasthttp.Response
	if err := resp.Read(br); err != nil {
		assert.Fail(t, "Error read connection", err)
	}

	assert.Equal(t, fasthttp.StatusNotFound, resp.StatusCode())
	assert.Equal(t, "Unsupported path", string(resp.Body()))

}

func TestHandlerHello(t *testing.T) {
	f := setUpHandleTest(t)
	defer tearDownHandleTest(t, &f)

	req := fasthttp.Request{}
	req.Header.SetMethod("GET")
	req.SetHost("test")
	req.URI().Update("/hello")

	if _, err := f.Conn.Write([]byte(req.String())); err != nil {
		assert.Fail(t, "Error write in connection", err)
	}

	br := bufio.NewReader(f.Conn)
	var resp fasthttp.Response
	if err := resp.Read(br); err != nil {
		assert.Fail(t, "Error read connection", err)
	}

	assert.Equal(t, fasthttp.StatusOK, resp.StatusCode())
	assert.Equal(t, "hello, stranger", string(resp.Body()))
}