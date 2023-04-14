package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"refactoring/config"
	v1 "refactoring/internl/controller/http/v1"
	"refactoring/internl/repository"
	"refactoring/internl/usecase"
	"refactoring/pkg/httpserver"
	"time"

	"syscall"
)

func Run(cfg *config.Config) {
	stroe, err := repository.NewStore(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	usecase := usecase.NewManager(stroe)
	handler := v1.NewRouter(usecase)

	httpServer := httpserver.New(handler,
		httpserver.Port(cfg.HTTP.Port),
		httpserver.ReadTimeout(15*time.Second),
		httpserver.WriteTimeout(60*time.Second),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Fatalln(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		log.Fatalln(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
