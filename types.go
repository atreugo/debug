package debug

// Config is the configuration for the Debug plugin.
type Config struct {
	Expvar ExpvarConfig
	Pprof  PprofConfig
}

// ExpvarConfig is the expvar configuration.
type ExpvarConfig struct {
	// Method is the HTTP method to use for the metrics endpoint.
	//
	// Default: "GET"
	Method string

	// URL is the HTTP URL to scrape for expvar metrics.
	//
	// Default: "/debug/vars"
	URL string
}

// PprofConfig is the pprof configuration.
type PprofConfig struct {
	// Method is the HTTP method to use for the metrics endpoint.
	//
	// Default: "GET"
	Method string

	// URL is the HTTP URL to scrape for pprof metrics.
	//
	// Default: "/debug/pprof/{profile:*}"
	URL string
}
