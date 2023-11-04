package adapters_test

import (
	"testing"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/adapters"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/config"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/db"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/local_price_prediction"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateLocalPricePrediction(t *testing.T) {
	cfg := config.GetConfig()
	db := db.NewSurrealDBClient(cfg.DB.User, cfg.DB.Pass, cfg.DB.URL)
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger(), db)
	err := adapter.CreateLocalPricePrediction(&local_price_prediction.LocalPricePrediction{
		StartTimestamp: timestamppb.Now(),
		ZipCode:        12345,
		EndTimestamp:   timestamppb.Now(),
		MarketPrice:    123.0,
		LocalPrice:     123.0,
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetCurrentLocalPricePrediction(t *testing.T) {
	cfg := config.GetConfig()
	db := db.NewSurrealDBClient(cfg.DB.User, cfg.DB.Pass, cfg.DB.URL)
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger(), db)
	result, err := adapter.ReadCurrentLocalPricePrediction("12345")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
