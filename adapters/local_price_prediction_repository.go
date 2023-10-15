package adapters

import (
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/common"
	pricePrediction "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/local_price_prediction"
)

const localPricePredictionDBName = "localpricepredictions"
const localPricePredictionTableName = "prices"

func (a *Adapter) CreateLocalPricePrediction(prediction *pricePrediction.LocalPricePrediction) error {
	_, err := a.db.Use(namespace, localPricePredictionDBName)
	if err != nil {
		a.logger.Error(err.Error())
		return common.DBSwitchNSDBError
	}

	// SKIP data here with _
	_, err = a.db.Create(localPricePredictionTableName, prediction)
	if err != nil {
		a.logger.Error(err.Error())
		return common.LocalPricePredictionCreateError
	}

	return nil
}
