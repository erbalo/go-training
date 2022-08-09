package memory

import (
	"employee/domain"
	"employee/ports"
	"errors"
)

type MemoryRepository struct {
	employees []domain.Employee
}

func NewRepository() ports.EmployeeRepository {
	return &MemoryRepository{
		employees: []domain.Employee{},
	}
}

func (repository *MemoryRepository) GetAll() (*[]domain.Employee, error) {
	return &repository.employees, nil
}

func (repository *MemoryRepository) Get(id int) (*domain.Employee, error) {
	for _, employee := range repository.employees {
		if employee.ID == id {
			return &employee, nil
		}
	}

	return nil, errors.New("employee not found")
}

func (repository *MemoryRepository) Create(employee domain.Employee) (*domain.Employee, error) {
	var lastEmployee domain.Employee
	if len(repository.employees) > 0 {
		lastEmployee = repository.employees[len(repository.employees)-1]
	}

	employee.ID = lastEmployee.ID + 1
	repository.employees = append(repository.employees, employee)

	return &employee, nil
}

func (repository *MemoryRepository) Update(id int, employee domain.Employee) (*domain.Employee, error) {
	employeeIndex := -1
	for index, employee := range repository.employees {
		if employee.ID == id {
			employeeIndex = index
			break
		}
	}

	if employeeIndex < 0 {
		// TODO build custom errors
		return nil, errors.New("employee not found")
	}

	employee.ID = id
	repository.employees[employeeIndex] = employee

	return &employee, nil
}

func (repository *MemoryRepository) Delete(id int) error {
	employeeIndex := -1
	for index, employee := range repository.employees {
		if employee.ID == id {
			employeeIndex = index
			break
		}
	}

	if employeeIndex < 0 {
		return errors.New("employee not found")
	}

	repository.employees = append(repository.employees[0:employeeIndex], repository.employees[employeeIndex+1:]...)
	return nil
}
