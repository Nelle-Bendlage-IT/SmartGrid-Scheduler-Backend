package adapters_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

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

func TestGetCurrentGSIPrediction(t *testing.T) {
	cfg := config.GetConfig()
	db := db.NewSurrealDBClient(cfg.DB.User, cfg.DB.Pass, cfg.DB.URL)
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger(), db)
	result, err := adapter.ReadCurrentGSIPrediction("12345")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print("123", result)
}

func TestGetCurrentGSIPredictionQuery(t *testing.T) {
	cfg := config.GetConfig()
	db := db.NewSurrealDBClient(cfg.DB.User, cfg.DB.Pass, cfg.DB.URL)
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger(), db)
	currentTime := time.Now()
	// Add one hour to the current time
	oneHourLater := currentTime.Add(time.Hour)
	//Convert the time to a Timestamp
	timestampProto := timestamppb.New(oneHourLater)

	err := adapter.CreateGSIPrediction(&gsi_prediction.GSIPrediction{
		StartTimestamp: timestampProto,
		ZipCode:        12366,
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
	result, err := adapter.ReadCurrentGSIPrediction("12366")
	if err != nil {
		t.Fatal(err)
	}
	if len(result) == 0 {
		t.Fatal(errors.New("FAILED TO RECIEVE CREATED PREDICTION"))
	}
}
