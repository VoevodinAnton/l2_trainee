package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"http/pkg/handler"
	"http/pkg/repository"
	"http/pkg/server"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	db, err := repository.NewPostgresDB()
	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	hndlr := handler.NewHandler(repos)
	router := http.NewServeMux()
	server := server.InitServer(hndlr, router)

	addr := "127.0.0.1:" + webPort
	logrus.WithField("addr", addr).Info("starting server")

	err = server.StartServer(webPort)
	if err != nil {
		return err
	}

	return nil
}
