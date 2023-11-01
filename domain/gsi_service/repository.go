package gsiService

import "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"

type GSIRepository interface {
	ReadCurrentGSIPrediction(zipCode string) ([]*gsi_prediction.GSIPrediction, error)
}
