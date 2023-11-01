package ports

import (
	"context"
	"fmt"
	"strconv"

	supa "github.com/nedpals/supabase-go"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/greet"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/gsi_prediction"
)

type GRPCService struct {
	app app.Application
}

func NewGRPCServer(app app.Application) GRPCService {
	return GRPCService{app}
}

func (g GRPCService) GetGreet(ctx context.Context, req *greet.GetGreetRequest) (*greet.GetGreetResponse, error) {
	user := ctx.Value("user").(*supa.User)
	fmt.Println(user)
	return &greet.GetGreetResponse{Message: fmt.Sprint(g.app.Greet.HandleGetGreet(req.User))}, nil
}

func (g GRPCService) GetGSIPrediction(ctx context.Context, req *gsi_prediction.GetGSIPredictionsRequest) (*gsi_prediction.GetGSIPredictionResponse, error) {
	zipCode := strconv.FormatUint(uint64(req.Zipcode), 10)

	response, err := g.app.GetGSIPrediction.HandleGetGSIPredicition(zipCode)
	if err != nil {
		return nil, err
	}
	return response, nil
}
