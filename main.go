package main

import (
	"fmt"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/config"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/server"
	"github.com/nedpals/supabase-go"
)

func main() {
	cfg := config.GetConfig()
	zapLogger := logger.GetLogger()
	supabaseClient := supabase.CreateClient(cfg.Supabase.URL, cfg.Supabase.Key)
	server.RunGRPCServer(supabaseClient, fmt.Sprint(cfg.Port), zapLogger)
}
