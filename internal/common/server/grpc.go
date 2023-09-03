package server

import (
	"log"
	"net"

	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/app/handler"
	greetService "github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/domain/greet"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/genproto/greet"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/internal/common/logger"
	"github.com/Nelle-Bendlage-IT/SmartGrid-Scheduler-Backend/ports"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RunGRPCServer() {
	zapLogger := logger.GetLogger()

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			zapLogger.Error(err.Error())
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}
	greetPort := ports.NewGRPCServer(app.Application{
		Greet: handler.NewGreetService(greetService.Service{}),
	})
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(logger.InterceptorLogger(zapLogger), opts...),
		recovery.UnaryServerInterceptor(recoveryOpts...),
	))
	greet.RegisterGreetServiceServer(grpcServer, greetPort)

	listen, err := net.Listen("tcp", ":3007")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("RUNNING GRPC SERVER ON PORT: 3007")
	grpcServer.Serve(listen)
}
