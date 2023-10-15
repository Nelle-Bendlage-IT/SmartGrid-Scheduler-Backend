package common

import "errors"

var (
	DBSwitchNSDBError               = errors.New("FAILED_TO_SWITCH_NAMESPACE_AND_DB")
	LocalPricePredictionCreateError = errors.New("FAILED_TO_CREATE_LOCAL_PRICE_PREDICTION")
)
