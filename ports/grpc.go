package ports

import (
	"context"
	"fmt"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/greet"
)

type GRPCService struct {
	app app.Application
}

func NewGRPCServer(app app.Application) GRPCService {
	return GRPCService{app}
}

func (g GRPCService) GetGreet(ctx context.Context, req *greet.GetGreetRequest) (*greet.GetGreetResponse, error) {
	fmt.Println(req.GetUser())
	return &greet.GetGreetResponse{Message: fmt.Sprint(g.app.Greet.HandleGetGreet(req.User))}, nil
}
