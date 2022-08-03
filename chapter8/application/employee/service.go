package employee

import (
	"employee/domain"
	"employee/ports"
	"time"
)

type Service struct {
	repository ports.EmployeeRepository
}

func NewService(repo ports.EmployeeRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) getBy(id int) (*domain.Employee, error) {
	return service.repository.Get(id)
}

func (service *Service) getAll() (*[]domain.Employee, error) {
	return service.repository.GetAll()
}

func (service *Service) create(employeeRequest domain.EmployeeRequest) (*domain.Employee, error) {
	var employee domain.Employee
	employee.FullName = employeeRequest.FullName
	employee.Position = employeeRequest.Position
	employee.Salary = employeeRequest.Salary
	employee.Joined = time.Now().UTC()
	employee.OnProbation = employeeRequest.OnProbation

	return service.repository.Create(employee)
}

func (service *Service) deleteBy(id int) error {
	return service.repository.Delete(id)
}

func (service *Service) update(id int, employee domain.Employee) (*domain.Employee, error) {
	return service.repository.Update(id, employee)
}
