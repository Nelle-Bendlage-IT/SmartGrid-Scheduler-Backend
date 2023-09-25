package adapters_test

import (
	"context"
	"testing"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/adapters"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/config"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
	"github.com/stretchr/testify/assert"
)

func TestGetGSIPrediction(t *testing.T) {
	cfg := config.GetConfig()
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger())
	resp, err := adapter.GetGSIPrediction(context.TODO(), "48155")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, resp)
}

func TestGetLocalPricePrediction(t *testing.T) {
	cfg := config.GetConfig()
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger())
	resp, err := adapter.GetLocalPricePrediction(context.TODO(), "48155")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, resp)
}

func TestGetBestHourPrediction(t *testing.T) {
	cfg := config.GetConfig()
	adapter := adapters.NewAdapter(cfg.CorrentlyAPIKey, logger.GetLogger())
	resp, err := adapter.GetBestHourForEnergyConsumption(context.TODO(), "48155", "6")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, resp)
}
