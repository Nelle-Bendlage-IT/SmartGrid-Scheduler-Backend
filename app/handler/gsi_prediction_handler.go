package handler

import (
	gsiService "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/domain/gsi_service"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
)

type GSI struct {
	service gsiService.GSIService
}

func NewGsiService(service gsiService.GSIService) GSI {
	return GSI{service: service}
}

func (g GSI) HandleGetGSIPredicition(zipCode string) (*gsi_prediction.GetGSIPredictionResponse, error) {
	return g.service.GetCurrentGSIPrediction(zipCode)
}
