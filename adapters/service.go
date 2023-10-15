package adapters

import (
	pricePrediction "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/local_price_prediction"
	"github.com/surrealdb/surrealdb.go"
	"go.uber.org/zap"
)

const namespace = "main"

type Adapter struct {
	correntlyAPIKey string
	logger          *zap.Logger
	db              *surrealdb.DB
}

type Adapters interface {
	CreateLocalPricePrediction(prediction *pricePrediction.LocalPricePrediction) error
}

func NewAdapter(correntlyAPIKey string, logger *zap.Logger, db *surrealdb.DB) *Adapter {
	return &Adapter{correntlyAPIKey: correntlyAPIKey, logger: logger, db: db}
}
