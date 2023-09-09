package server

import (
	"log"
	"net"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app/handler"
	greetService "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/domain/greet"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/greet"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/server/middleware"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/ports"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	supa "github.com/nedpals/supabase-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RunGRPCServer(supabaseClient *supa.Client, port string, zapLogger *zap.Logger) {

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			zapLogger.Error(err.Error())
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	authMiddleware := middleware.New(supabaseClient)

	greetPort := ports.NewGRPCServer(app.Application{
		Greet: handler.NewGreetService(greetService.Service{}),
	})
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(logger.InterceptorLogger(zapLogger), opts...),
		recovery.UnaryServerInterceptor(recoveryOpts...),
		auth.UnaryServerInterceptor(authMiddleware.Middleware),
	))
	greet.RegisterGreetServiceServer(grpcServer, greetPort)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("RUNNING GRPC SERVER ON PORT: " + port)
	grpcServer.Serve(listen)
}
