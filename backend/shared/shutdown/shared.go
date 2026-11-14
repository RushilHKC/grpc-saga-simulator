package shared

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func CreateServer(addr string, mux http.Handler) *http.Server {

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func RunServer(server *http.Server, shutdownTimeOut time.Duration) error {

	serverErr := make(chan error, 1)

	go func() {
		fmt.Println("starting server")
		if err := server.ListenAndServe(); err != nil {
			serverErr <- err
		}
		close(serverErr)
	}()

	stop := make(chan os.Signal, 1)
	defer close(stop)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		fmt.Println("Error starting server: ", err)
		return err
	case <-stop:
		fmt.Println("Shutdown signal received")
	}

	return shutdownServer(shutdownTimeOut, server)

}

func shutdownServer(shutdownTimeOut time.Duration, server *http.Server) error {

	shutdownContext, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeOut,
	)

	defer cancel()

	if shutdownErr := server.Shutdown(shutdownContext); shutdownErr != nil {
		if CloseErr := server.Close(); CloseErr != nil {
			return errors.Join(CloseErr, shutdownErr)
		}

		return shutdownErr
	}

	fmt.Println("server gracefully closed")
	return nil
}
