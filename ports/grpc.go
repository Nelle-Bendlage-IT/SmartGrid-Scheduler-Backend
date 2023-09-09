package ports

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"

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
	user := ctx.Value("user").(*supa.User)
	fmt.Println(user)
	return &greet.GetGreetResponse{Message: fmt.Sprint(g.app.Greet.HandleGetGreet(req.User))}, nil
}
