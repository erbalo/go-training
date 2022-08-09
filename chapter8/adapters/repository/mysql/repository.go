package mysql

import (
	"database/sql"
	"employee/adapters/errors"
	"employee/adapters/types"
	"employee/domain"
	"employee/ports"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	queryGetEmployees   = "select * from employees;"
	queryGetEmployee    = "select * from employees where id = ?;"
	queryCreateEmployee = "insert into employees(full_name, position, salary, joined, on_probation) values (?,?,?,?,?);"
	queryDeleteEmployee = "delete from employees where id = ?;"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() ports.EmployeeRepository {
	db, err := sql.Open("mysql", "go-training-user:123456@tcp(localhost:3306)/go-training?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Repository{
		db: db,
	}
}

// TODO pagination?
func (repository *Repository) GetAll() (*[]domain.Employee, error) {
	statement, err := repository.db.Prepare(queryGetEmployees)
	if err != nil {
		log.Println("error when trying to prepare the statement", err)
		return nil, errors.NewInternal("error when trying to prepare the statement")
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Println("error when trying to get employees", err)
		return nil, errors.NewInternal("error when trying to get employees")
	}
	defer rows.Close()

	results := make([]domain.Employee, 0)

	for rows.Next() {
		var employee domain.Employee
		err = rows.Scan(&employee.ID, &employee.FullName, &employee.Position, &employee.Salary, &employee.Joined, &employee.OnProbation, &employee.CreatedAt)

		if err != nil {
			log.Println("error when tying to get employee", err)
			return nil, errors.NewInternal("error when tying to get employee")
		}

		results = append(results, employee)
	}

	return &results, nil
}

func (repository *Repository) Get(id int) (*domain.Employee, error) {
	statement, err := repository.db.Prepare(queryGetEmployee)
	if err != nil {
		log.Println("error when trying to prepare the statement", err)
		return nil, errors.NewInternal("error when trying to prepare the statement")
	}
	defer statement.Close()

	result := statement.QueryRow(id)

	var employee domain.Employee
	err = result.Scan(&employee.ID, &employee.FullName, &employee.Position, &employee.Salary, &employee.Joined, &employee.OnProbation, &employee.CreatedAt)
	if err != nil {
		log.Println("error when trying to get user by id", err)
		return nil, errors.NewInternal("error when trying to get user by id")
	}

	return &employee, nil
}

func (repository *Repository) Create(employee domain.Employee) (*domain.Employee, error) {
	statement, err := repository.db.Prepare(queryCreateEmployee)
	if err != nil {
		log.Println("error when trying to prepare the statement", err)
		return nil, errors.NewInternal("error when trying to prepare the statement")
	}
	defer statement.Close()

	result, err := statement.Exec(employee.FullName, employee.Position, employee.Salary, employee.Joined.Format(types.DatetimeFormat), employee.OnProbation)
	if err != nil {
		log.Println("error when trying to save the employee", err)
		return nil, errors.NewInternal("error when trying to save the employee")
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("error when finding last instert id", err)
		return nil, errors.NewInternal("error when finding last instert id")
	}

	employee.ID = int(lastId)

	return &employee, nil
}

// TODO adds update logic
func (repository *Repository) Update(id int, employee domain.Employee) (*domain.Employee, error) {
	return nil, nil
}

func (repository *Repository) Delete(id int) error {
	statement, err := repository.db.Prepare(queryDeleteEmployee)
	if err != nil {
		log.Println("error when trying to prepare the statement", err)
		return errors.NewInternal("error when trying to prepare the statement")
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		log.Println("error when trying to delete user", err)
		return errors.NewInternal("error when trying to delete user")
	}

	return nil
}
