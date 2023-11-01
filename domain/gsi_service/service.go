package gsiService

import (
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
)

type Service struct {
	repository GSIRepository
}

type GSIService interface {
	GetCurrentGSIPrediction(zipcode string) (*gsi_prediction.GetGSIPredictionResponse, error)
}

func NewService(repository GSIRepository) *Service {
	return &Service{repository: repository}
}

func (svc Service) GetCurrentGSIPrediction(zipcode string) (*gsi_prediction.GetGSIPredictionResponse, error) {
	gsiPrediction, err := svc.repository.ReadCurrentGSIPrediction(zipcode)
	if err != nil {
		return nil, err
	}
	return mapToGSIPredictionResponse(gsiPrediction)
}
