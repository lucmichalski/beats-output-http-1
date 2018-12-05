package http

import (
	"time"

	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/transport"
)

// backoff ...
type backoff struct {
	Init time.Duration
	Max  time.Duration
}

type config struct {
	Protocol         string                `config:"protocol"`
	Port             int                   `config:"port"`
	Path             string                `config:"path"`
	Params           map[string]string     `config:"parameters"`
	Username         string                `config:"username"`
	Password         string                `config:"password"`
	ProxyURL         string                `config:"proxy_url"`
	LoadBalance      bool                  `config:"loadbalance"`
	CompressionLevel int                   `config:"compression_level" validate:"min=0, max=9"`
	TLS              *outputs.TLSConfig    `config:"tls"`
	Proxy            transport.ProxyConfig `config:",inline"`
	MaxRetries       int                   `config:"max_retries"`
	BulkMaxSize      int                   `config:"bulk_max_size"`
	Timeout          time.Duration         `config:"timeout"`
	Backoff          backoff               `config:"backoff"`
}

var defaultConfig = config{
	Protocol: "https",
	Port:     443,
	// Path:     "",
	// Params: "",
	// ProxyURL:         "",
	// Username:         "",
	// Password:         "",
	Timeout: 90 * time.Second,
	Backoff: backoff{
		Init: 1 * time.Second,
		Max:  60 * time.Second,
	},
	CompressionLevel: 0,
	// TLS:              nil,
	// Proxy: transport.ProxyConfig{
	// 	URL:          "",
	// 	LocalResolve: false,
	// },
	MaxRetries:  3,
	BulkMaxSize: 1,
	LoadBalance: true,
}

func newConfig() *config {
	c := defaultConfig
	return &c
}
