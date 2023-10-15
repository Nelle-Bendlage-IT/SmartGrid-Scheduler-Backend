package adapters

import (
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/common"
	pricePrediction "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
)

const gsiPredictionDBName = "gsipredictions"
const gsiPredictionTableName = "data"

func (a *Adapter) CreateGSIPrediction(prediction *pricePrediction.GSIPrediction) error {
	_, err := a.db.Use(namespace, gsiPredictionDBName)
	if err != nil {
		a.logger.Error(err.Error())
		return common.DBSwitchNSDBError
	}

	// SKIP data here with _
	_, err = a.db.Create(gsiPredictionTableName, prediction)
	if err != nil {
		a.logger.Error(err.Error())
		return common.ErrFailedToCreateGSI
	}

	return nil
}
