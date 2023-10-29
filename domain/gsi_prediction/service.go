package gsi_prediction

import (
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/adapters"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/config"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/db"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
)

type Service struct {
}

type GSIService interface {
	GetCurrentGSIPrediction(zipcode string) gsi_prediction.GetGSIPredictionResponse
}

func (svc Service) GetCurrentGSIPrediction(zipcode string) []*gsi_prediction.GSIPrediction {
	cfg := config.GetConfig()
	db := db.NewSurrealDBClient(cfg.DB.User, cfg.DB.Pass, cfg.DB.URL)
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger(), db)
	prediction, err := adapter.GetCurrentGSIPrediction(zipcode)
	if err != nil {
		return nil
	}
	return prediction
}
