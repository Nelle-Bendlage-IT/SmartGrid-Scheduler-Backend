package adapters_test

import (
	"testing"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/adapters"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/config"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/db"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateGSIPrediction(t *testing.T) {
	cfg := config.GetConfig()

	db := db.NewSurrealDBClient(cfg.DB.User, cfg.DB.Pass, cfg.DB.URL)
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger(), db)
	err := adapter.CreateGSIPrediction(&gsi_prediction.GSIPrediction{
		StartTimestamp: timestamppb.Now(),
		ZipCode:        12345,
		EndTimestamp:   timestamppb.Now(),
		Solar:          5,
		Wind:           6,
		Gsi:            10.4,
		Co2GStandard:   5,
		Co2GOekostrom:  2,
		Energyprice:    4.2,
	})
	if err != nil {
		t.Fatal(err)
	}

}
