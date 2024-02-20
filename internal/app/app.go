package app

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/qorganbek/kwaka_test_assignment/internal/config"
	"github.com/qorganbek/kwaka_test_assignment/internal/handler"
	"github.com/qorganbek/kwaka_test_assignment/internal/repository"
	"github.com/qorganbek/kwaka_test_assignment/internal/repository/pgrepo"
	"github.com/qorganbek/kwaka_test_assignment/internal/service"
	"github.com/qorganbek/kwaka_test_assignment/pkg/http_server"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {

	pg := pgrepo.Postgres{
		Port:     cfg.DB.Port,
		Host:     cfg.DB.Host,
		DBName:   cfg.DB.DBName,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		SSLMode:  "disable",
	}

	db, err := pgrepo.NewPostgresDB(pg)

	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.New(services)

	server := http_server.New(handlers.InitRouter(),
		http_server.WithPort(cfg.HTTP.Port),
		http_server.WithReadTimeout(cfg.HTTP.ReadTimeout),
		http_server.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		http_server.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout))

	server.Start()
	fmt.Println("Server started!")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
