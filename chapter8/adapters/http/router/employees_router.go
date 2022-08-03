package router

import (
	"employee/application/employee"
	"employee/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type EmployeeRouter struct {
	useCase employee.UseCase
}

func NewEmployeeRouter(useCase employee.UseCase) *EmployeeRouter {
	return &EmployeeRouter{
		useCase: useCase,
	}
}

func (router *EmployeeRouter) Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parts := strings.Split(request.URL.Path, "/")
	partsLen := len(parts)

	if partsLen == 2 || partsLen == 3 && parts[2] == "" {
		if request.Method == "GET" {
			listEmployees(router.useCase, writer, request)
		} else if request.Method == "POST" {
			createEmployee(router.useCase, writer, request)
		}
	} else if partsLen == 3 || partsLen == 4 && parts[3] == "" {
		if request.Method == "GET" {
			searchEmployee(router.useCase, writer, request)
		} else if request.Method == "DELETE" {
			deleteEmployee(router.useCase, writer, request)
		} else if request.Method == "PUT" {
			updateEmployee(router.useCase, writer, request)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func searchEmployee(useCase employee.UseCase, writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	employee, err := useCase.FindBy(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(employee)
}

func listEmployees(useCase employee.UseCase, writer http.ResponseWriter, request *http.Request) {
	employees, err := useCase.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	json.NewEncoder(writer).Encode(employees)
}

func createEmployee(useCase employee.UseCase, writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var employeeRequest domain.EmployeeRequest
	err = json.Unmarshal(body, &employeeRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	employeeCreated, err := useCase.Create(employeeRequest)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(employeeCreated)
}

func deleteEmployee(useCase employee.UseCase, writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = useCase.DeleteBy(id)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func updateEmployee(useCase employee.UseCase, writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.Path, "/")

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var employee domain.Employee
	err = json.Unmarshal(body, &employee)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	employeeUpdated, err := useCase.Update(id, employee)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(writer).Encode(employeeUpdated)
}
