package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/keshvan/go-common-forum/postgres"
	"github.com/keshvan/user-service-sstu-forum/config"
	authgrpc "github.com/keshvan/user-service-sstu-forum/internal/grpc"
	"github.com/keshvan/user-service-sstu-forum/internal/repo"
	"github.com/keshvan/user-service-sstu-forum/internal/usecase"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	//Repository
	pg, err := postgres.New(cfg.PG_URL)
	if err != nil {
		log.Fatalf("app - Run - postgres.New")
	}
	defer pg.Close()

	userRepo := repo.New(pg)

	//Usecase
	userUsecase := usecase.New(userRepo)

	//GRPC-Server
	grpcServer := grpc.NewServer()
	authgrpc.Register(grpcServer, userUsecase)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("app - Run - net.Listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(l); err != nil {
			log.Fatalf("app - Run - grpcServer.Serve: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
