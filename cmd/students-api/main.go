package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pRakesh15/student-api/pkg/config"
)

func main() {

	//load config

	cfg := config.MustLoad()
	//set database

	//set routs
	//this is helps us to create routers...
	router := http.NewServeMux()
	//use the router with functions
	//in future we write the functions in another folder or another package...
	router.HandleFunc("GET /api/v1/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("JAy Shree RAM"))
	})

	//start server
	//start the server by using the config files....
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("server started at %s", cfg.Address)

	//creating channel for

	done := make(chan os.Signal, 1)
	// This line in Go is used for signal handling in an application.
	//  It allows the program to listen for specific operating system signals,
	//   such as when a user interrupts the program (e.g., by pressing Ctrl+C) or when the system sends termination signals.

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//check the error if there is any error ...
	//creating goroutine and function call at a time
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	//this things are called gras full shutdown
	//means if there are some request processing then the server not directly shuting down
	//it will take some time but at this time server note receive  any request

	slog.Info("Shutting down the server")

	//it will wait 5 sec for complete the ongoing req..

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
