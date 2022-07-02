package telegram

import (
	"context"
	_ "errors"
	_ "fmt"
	_ "github.com/oklahomer/go-kasumi/logger"
	_ "github.com/oklahomer/go-kasumi/retry"
	"github.com/oklahomer/go-sarah/v4"
)

const (
	// GITTER is a dedicated sarah.BotType for Gitter integration.
	TELEGRAM sarah.BotType = "telegram"
)

// AdapterOption defines a function's signature that Adapter's functional options must satisfy.
type AdapterOption func(adapter *Adapter)

// Adapter is a sarah.Adapter implementation for Gitter.
// This holds REST/Streaming API clients' instances.
type Adapter struct {
	config          *Config	
}

var _ sarah.Adapter = (*Adapter)(nil)

// NewAdapter creates and returns a new Adapter instance.
func NewAdapter(config *Config, options ...AdapterOption) (*Adapter, error) {
	adapter := &Adapter{
		config:          config,		
	}

	for _, opt := range options {
		opt(adapter)
	}

	return adapter, nil
}

// BotType returns a designated BotType for Gitter integration.
func (adapter *Adapter) BotType() sarah.BotType {
	return TELEGRAM
}

// Run fetches all belonging Room information and connects to them.
// New goroutines are activated for each Room to connect, and the interactions run in a concurrent manner.
func (adapter *Adapter) Run(ctx context.Context, enqueueInput func(sarah.Input) error, notifyErr func(error)) {
	
}

// SendMessage lets sarah.Bot send a message to Gitter.
func (adapter *Adapter) SendMessage(ctx context.Context, output sarah.Output) {
	
}