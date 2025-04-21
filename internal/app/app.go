package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/keshvan/go-common-forum/httpserver"
	"github.com/keshvan/go-common-forum/postgres"
	"github.com/keshvan/user-service-sstu-forum/config"
	"github.com/keshvan/user-service-sstu-forum/internal/controller"
	"github.com/keshvan/user-service-sstu-forum/internal/repo"
	"github.com/keshvan/user-service-sstu-forum/internal/usecase"
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
