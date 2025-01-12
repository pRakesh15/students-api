package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Printf("server started at %s", cfg.HTTPServer.Address)

	//check the error if there is any error ...
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}

}
