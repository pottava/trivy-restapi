// Package lib defines helper functions
package lib

import (
	"github.com/kelseyhightower/envconfig"
)

// Config can be set via environment variables
type config struct {
	AccessLog bool `envconfig:"ACCESS_LOG" default:"true"`
}

// Config represents its configurations
var Config *config

func init() {
	cfg := &config{}
	envconfig.MustProcess("trivy", cfg)
	Config = cfg
}
