package adapters

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/common"
	pricePrediction "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/local_price_prediction"
)

const localPricePredictionDBName = "localpricepredictions"
const localPricePredictionTableName = "prices"
const currentLocalPricePredicitionQuery = "SELECT * FROM type::table($table) WHERE zip_code = $zipcode AND time::from::secs(start_timestamp.seconds) > time::now()"

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

func (a *Adapter) ReadCurrentLocalPricePrediction(zipCode string) ([]*pricePrediction.LocalPricePrediction, error) {
	_, err := a.db.Use(namespace, localPricePredictionDBName)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, common.DBSwitchNSDBError
	}
	currentTimestamp := time.Now()

	// Define the query parameters as a map
	// works only with zipcode as an int
	zipCodeInt, err := strconv.ParseInt(zipCode, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	queryVars := map[string]interface{}{
		"zipcode":   zipCodeInt,
		"timestamp": currentTimestamp,
		"table":     localPricePredictionTableName,
	}
	result, err := a.db.Query(currentLocalPricePredicitionQuery, queryVars)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, common.ErrFailedToCreateGSI
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	type dbLocalPricePredictionResponse = struct {
		Result []*pricePrediction.LocalPricePrediction
	}

	var predictions []dbLocalPricePredictionResponse
	err = json.Unmarshal(resultJSON, &predictions)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}
	return predictions[0].Result, nil
}
