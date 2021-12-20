package debug

import "github.com/valyala/fasthttp"

const (
	defaultExpvarURL    = "/debug/vars"
	defaultExpvarMethod = fasthttp.MethodGet

	defaultPprofURL    = "/debug/pprof/{profile:*}"
	defaultPprofMethod = fasthttp.MethodGet
)
