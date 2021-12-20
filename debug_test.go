package debug

import (
	"testing"

	"github.com/atreugo/httptest/assert"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func assertResponse(t *testing.T, s *atreugo.Atreugo, req *fasthttp.Request) {
	t.Helper()

	assert.Server(t, req, s, func(resp *fasthttp.Response) {
		if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
			t.Errorf("status code %d, want %d", statusCode, fasthttp.StatusOK)
		}

		if len(resp.Body()) == 0 {
			t.Error("body is empty")
		}
	})
}

func Test_Register_Expvar(t *testing.T) {
	s := atreugo.New(atreugo.Config{})
	Register(s, Config{})

	req := new(fasthttp.Request)
	req.SetRequestURI(defaultExpvarURL)
	req.Header.SetMethod(defaultExpvarMethod)

	assertResponse(t, s, req)
}

func Test_Register_Pprof(t *testing.T) {
	s := atreugo.New(atreugo.Config{})
	Register(s, Config{})

	req := new(fasthttp.Request)
	req.SetRequestURI("/debug/pprof/heap")
	req.Header.SetMethod(defaultPprofMethod)

	assertResponse(t, s, req)
}
