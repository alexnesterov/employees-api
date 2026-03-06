package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EmployeeDTO struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Sex            string  `json:"sex"`
	Age            int     `json:"age"`
	Salary         int     `json:"salary"`
	DepartmentCode *string `json:"department_code"`
}

type DepartmentDTO struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CreateEmployeeReq struct {
	Name           string  `json:"name"`
	Sex            string  `json:"sex"`
	Age            int     `json:"age"`
	Salary         int     `json:"salary"`
	DepartmentCode *string `json:"department_code,omitempty"`
}

type UpdateEmployeeReq struct {
	Name           *string `json:"name,omitempty"`
	Sex            *string `json:"sex,omitempty"`
	Age            *int    `json:"age,omitempty"`
	Salary         *int    `json:"salary,omitempty"`
	DepartmentCode *string `json:"department_code,omitempty"`
}

type CreateDepartmentReq struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type apiResponse[T any] struct {
	Data T `json:"data"`
}

type APIClient struct {
	baseURL string
	http    *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		http:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *APIClient) ListEmployees() ([]*EmployeeDTO, error) {
	resp, err := c.http.Get(c.baseURL + "/employees")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result apiResponse[[]*EmployeeDTO]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (c *APIClient) CreateEmployee(req CreateEmployeeReq) error {
	b, _ := json.Marshal(req)
	resp, err := c.http.Post(c.baseURL+"/employees", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", resp.Status)
	}
	return nil
}

func (c *APIClient) UpdateEmployee(id string, req UpdateEmployeeReq) error {
	b, _ := json.Marshal(req)
	r, err := http.NewRequest(http.MethodPut, c.baseURL+"/employees/"+id, bytes.NewReader(b))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", resp.Status)
	}
	return nil
}

func (c *APIClient) DeleteEmployee(id string) error {
	r, err := http.NewRequest(http.MethodDelete, c.baseURL+"/employees/"+id, nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", resp.Status)
	}
	return nil
}

func (c *APIClient) ListDepartments() ([]*DepartmentDTO, error) {
	resp, err := c.http.Get(c.baseURL + "/departments")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result apiResponse[[]*DepartmentDTO]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (c *APIClient) CreateDepartment(req CreateDepartmentReq) error {
	b, _ := json.Marshal(req)
	resp, err := c.http.Post(c.baseURL+"/departments", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", resp.Status)
	}
	return nil
}

func (c *APIClient) DeleteDepartment(code string) error {
	r, err := http.NewRequest(http.MethodDelete, c.baseURL+"/departments/"+code, nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", resp.Status)
	}
	return nil
}
