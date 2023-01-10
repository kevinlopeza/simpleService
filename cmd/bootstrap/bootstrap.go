package bootstrap

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"simpleService/internal/process"
	"simpleService/internal/rest"
	"syscall"
)

func Run() {
	port := "90"

	cacheImpl, _ := process.NewCacheImpl()
	handler := rest.NewHandler(cacheImpl)

	mux := http.NewServeMux()

	mux.Handle("/getHolidays/", handler)

	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", port),
	}

	signals := make(chan os.Signal)
	errors := make(chan error)

	signal.Notify(signals, syscall.SIGINT)
	signal.Notify(signals, syscall.SIGTERM)

	go func() {
		log.Info("Server will start now :)")
		errors <- server.ListenAndServe()
	}()

	select {
	case s := <-signals:
		log.Info("Service received signal ", "type", s.String())
		break
	case err := <-errors:
		log.Error("An error occurred while starting ", "error", err.Error())
		break
	}
	log.Info("Server is shutting down now")
}
