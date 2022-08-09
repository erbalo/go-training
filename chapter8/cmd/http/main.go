package main

import (
	"employee/adapters/repository/mysql"
	"employee/application/employee"

	httpAdapter "employee/adapters/http"
)

func main() {
	repository := mysql.NewRepository()
	service := employee.NewService(repository)
	useCase := employee.NewUseCase(service)

	server := httpAdapter.NewServer(*useCase)
	server.ConfigureRoutes()
	server.Run()
}
