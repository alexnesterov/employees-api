package model

type Employee struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Sex            string  `json:"sex"`
	Age            int     `json:"age"`
	Salary         int     `json:"salary"`
	DepartmentCode *string `json:"department_code"`
}
