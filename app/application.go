package app

import "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app/handler"

type Application struct {
	Greet            handler.Greet
	GetGSIPrediction handler.GSI
}
