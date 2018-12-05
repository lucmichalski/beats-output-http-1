package http

import (
	"time"

	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/transport"
	"github.com/elastic/beats/libbeat/publisher"
)

type client struct {
	Client   *transport.Client
	observer outputs.Observer
	timeout  time.Duration
}

func newClient(
	tc *transport.Client,
	observer outputs.Observer,
	timeout time.Duration,
) *client {
	return &client{
		Client:   tc,
		observer: observer,
		timeout:  timeout,
	}
}

func (*client) Connect() error {
	return nil
}

func (*client) Close() error {
	return nil
}

func (*client) Publish(batch publisher.Batch) error {
	return nil
}

func (*client) String() string {
	return ""
}
