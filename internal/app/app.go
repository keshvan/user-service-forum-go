package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"user-service/config"
	"user-service/internal/controller"
	"user-service/internal/repo"
	"user-service/internal/usecase"

	"github.com/keshvan/go-sstu-forum/pkg/httpserver"
	"github.com/keshvan/go-sstu-forum/pkg/postgres"
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

	//HTTP-Server
	httpServer := httpserver.New("localhost:3000")
	controller.NewRouter(httpServer.Engine, userUsecase)

	httpServer.Run()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
