package ports

import "employee/domain"

type EmployeeRepository interface {
	GetAll() (*[]domain.Employee, error)
	Get(id int) (*domain.Employee, error)
	Create(employee domain.Employee) (*domain.Employee, error)
	Update(id int, employee domain.Employee) (*domain.Employee, error)
	Delete(id int) error
}
