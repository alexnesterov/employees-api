package model

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Sex            string    `json:"sex"`
	Age            int       `json:"age"`
	Salary         int       `json:"salary"`
	DepartmentCode *string   `json:"department_code"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
