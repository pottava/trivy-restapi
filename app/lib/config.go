// Package lib defines helper functions
package lib

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// for compile flags
var (
	version = "dev"
	commit  string
	date    = "---"
)

// Config can be set via environment variables
type config struct {
	CacheDir  string `envconfig:"CACHE_DIRECTORY" default:"/root/.cache"`
	Version   string `envconfig:"VERSION" default:"dev"`
	AccessLog bool   `envconfig:"ACCESS_LOG" default:"true"`
	Debug     bool   `envconfig:"APP_DEBUG" default:"false"`
}

// Config represents its configurations
var Config *config

func init() {
	cfg := &config{}
	envconfig.MustProcess("trivy", cfg)
	if len(version) > 0 && len(commit) > 0 && len(date) > 0 {
		cfg.Version = fmt.Sprintf("%s-%s (built at %s)", version, commit, date)
	}
	Config = cfg
	initLogger()
}
