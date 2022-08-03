package http

import (
	"employee/adapters/http/router"
	"employee/application/employee"
	"log"
	"net/http"
)

type Server struct {
	UseCase employee.UseCase
}

func NewServer(useCase employee.UseCase) *Server {
	return &Server{
		UseCase: useCase,
	}
}

func (server *Server) ConfigureRoutes() {
	employeeRouter := router.NewEmployeeRouter(server.UseCase)

	http.HandleFunc("/employees", employeeRouter.Handler)
	http.HandleFunc("/employees/", employeeRouter.Handler)
}

func (server *Server) Run() {
	log.Println("Running server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
