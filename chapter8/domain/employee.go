package domain

import (
	"employee/adapters/types"
	"time"
)

type Position int

const (
	Undetermined Position = iota
	Junior
	Senior
	Manager
	CEO
)

type Employee struct {
	ID          int           `json:"id"`
	FullName    string        `json:"full_name"`
	Position    Position      `json:"position"`
	Salary      float64       `json:"salary"`
	Joined      time.Time     `json:"joined"`
	OnProbation types.BitBool `json:"on_probation"`
	CreatedAt   *time.Time    `json:"created_at,omitempty"`
}

type EmployeeRequest struct {
	FullName    string        `json:"full_name"`
	Position    Position      `json:"position"`
	Salary      float64       `json:"salary"`
	OnProbation types.BitBool `json:"on_probation"`
}
