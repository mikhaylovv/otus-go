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
			t.Errorf("unexpected error: %s", err)
		}
	}()

	c, err := ln.Dial()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	return fixture{s, ln, c}
}

func tearDownHandleTest(t *testing.T, f *fixture) {
	if err := f.Conn.Close(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := f.Listener.Close(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestHandlerUnsupportedPath(t *testing.T) {
	f := setUpHandleTest(t)

	req := fasthttp.Request{}
	req.Header.SetMethod("GET")
	req.SetHost("test")
	req.URI().Update("/test")

	if _, err := f.Conn.Write([]byte(req.String())); err != nil {
		t.Fatal(err)
	}

	br := bufio.NewReader(f.Conn)
	var resp fasthttp.Response
	if err := resp.Read(br); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assert.Equal(t, fasthttp.StatusNotFound, resp.StatusCode())
	assert.Equal(t, "Unsupported path", string(resp.Body()))

	tearDownHandleTest(t, &f)
}

func TestHandlerHello(t *testing.T) {
	f := setUpHandleTest(t)

	req := fasthttp.Request{}
	req.Header.SetMethod("GET")
	req.SetHost("test")
	req.URI().Update("/hello")

	if _, err := f.Conn.Write([]byte(req.String())); err != nil {
		t.Fatal(err)
	}

	br := bufio.NewReader(f.Conn)
	var resp fasthttp.Response
	if err := resp.Read(br); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assert.Equal(t, fasthttp.StatusOK, resp.StatusCode())
	assert.Equal(t, "hello, stranger", string(resp.Body()))

	tearDownHandleTest(t, &f)
}