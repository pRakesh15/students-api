package student_router

import (
	"github.com/gorilla/mux"
	handler "github.com/pRakesh15/student-api/internal/students/api/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Create a subrouter for "/api/v1"
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Student routs routes
	apiRouter.HandleFunc("/students", handler.CreateUser).Methods("POST")

	return router
}
