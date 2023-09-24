package adapters

import "go.uber.org/zap"

type Adapter struct {
	correntlyAPIKey string
	logger          *zap.Logger
}

type Adapters interface {
}

func NewAdapter(correntlyAPIKey string, logger *zap.Logger) Adapter {
	return Adapter{correntlyAPIKey: correntlyAPIKey, logger: logger}
}
