package employee

import "employee/domain"

type UseCase struct {
	service *Service
}

func NewUseCase(employeeService *Service) *UseCase {
	return &UseCase{
		service: employeeService,
	}
}

func (useCase *UseCase) FindBy(id int) (*domain.Employee, error) {
	return useCase.service.getBy(id)
}

func (useCase *UseCase) GetAll() (*[]domain.Employee, error) {
	return useCase.service.getAll()
}

func (useCase *UseCase) Create(employeeRequest domain.EmployeeRequest) (*domain.Employee, error) {
	return useCase.service.create(employeeRequest)
}

func (useCase *UseCase) DeleteBy(id int) error {
	return useCase.service.deleteBy(id)
}

func (useCase *UseCase) Update(id int, employee domain.Employee) (*domain.Employee, error) {
	return useCase.service.update(id, employee)
}
