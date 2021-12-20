package debug

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp/expvarhandler"
	"github.com/valyala/fasthttp/pprofhandler"
)

// Register registers the debug plugin to the server with the given configuration.
func Register(s *atreugo.Atreugo, cfg Config) {
	if cfg.Expvar.Method == "" {
		cfg.Expvar.Method = defaultExpvarMethod
	}

	if cfg.Expvar.URL == "" {
		cfg.Expvar.URL = defaultExpvarURL
	}

	if cfg.Pprof.Method == "" {
		cfg.Pprof.Method = defaultPprofMethod
	}

	if cfg.Pprof.URL == "" {
		cfg.Pprof.URL = defaultPprofURL
	}

	s.RequestHandlerPath(cfg.Expvar.Method, cfg.Expvar.URL, expvarhandler.ExpvarHandler)
	s.RequestHandlerPath(cfg.Pprof.Method, cfg.Pprof.URL, pprofhandler.PprofHandler)
}
