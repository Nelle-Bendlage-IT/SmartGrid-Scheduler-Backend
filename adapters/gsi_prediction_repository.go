package adapters

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/common"
	pricePrediction "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const gsiPredictionDBName = "gsipredictions"
const gsiPredictionTableName = "data"
const currentGsiPredicitionQuery = "SELECT * FROM type::table($table) WHERE zip_code = $zipcode AND start_timestamp < $timestamp"

type dbGsiPredictionResponse = struct {
	Result []*pricePrediction.GSIPrediction
}

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

func (a *Adapter) GetCurrentGSIPrediction(zipCode string) ([]*pricePrediction.GSIPrediction, error) {
	_, err := a.db.Use(namespace, gsiPredictionDBName)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, common.DBSwitchNSDBError
	}
	currentTimestamp := timestamppb.Now()

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
		"table":     gsiPredictionTableName,
	}
	result, err := a.db.Query(currentGsiPredicitionQuery, queryVars)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, common.ErrFailedToCreateGSI
	}

	resultJSON, err := json.Marshal(result)
	fmt.Println(string(resultJSON))
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	var predictions []dbGsiPredictionResponse
	err = json.Unmarshal(resultJSON, &predictions)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}
	return predictions[0].Result, nil
}
