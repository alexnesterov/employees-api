package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type UIHandler struct {
	client *APIClient
	tmpl   *template.Template
}

func NewUIHandler(client *APIClient, tmpl *template.Template) *UIHandler {
	return &UIHandler{client: client, tmpl: tmpl}
}

type PageData struct {
	Employees   []*EmployeeDTO
	Departments []*DepartmentDTO
	DeptByCode  map[string]string
}

type EditRowData struct {
	Employee    *EmployeeDTO
	Departments []*DepartmentDTO
}

func (h *UIHandler) buildPageData() (*PageData, error) {
	employees, err := h.client.ListEmployees()
	if err != nil {
		return nil, err
	}
	departments, err := h.client.ListDepartments()
	if err != nil {
		return nil, err
	}
	deptByCode := make(map[string]string, len(departments))
	for _, d := range departments {
		deptByCode[d.Code] = d.Name
	}
	return &PageData{
		Employees:   employees,
		Departments: departments,
		DeptByCode:  deptByCode,
	}, nil
}

func (h *UIHandler) Index(w http.ResponseWriter, r *http.Request) {
	data, err := h.buildPageData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.tmpl.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Println("template error:", err)
	}
}

func (h *UIHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := CreateEmployeeReq{
		Name: r.FormValue("name"),
		Sex:  r.FormValue("sex"),
	}
	if age, err := strconv.Atoi(r.FormValue("age")); err == nil {
		req.Age = age
	}
	if salary, err := strconv.Atoi(r.FormValue("salary")); err == nil {
		req.Salary = salary
	}
	if dc := r.FormValue("department_code"); dc != "" {
		req.DepartmentCode = &dc
	}

	if err := h.client.CreateEmployee(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderEmployeesTable(w)
}

func (h *UIHandler) EditEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	employees, err := h.client.ListEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var emp *EmployeeDTO
	for _, e := range employees {
		if e.ID == id {
			emp = e
			break
		}
	}
	if emp == nil {
		http.Error(w, "employee not found", http.StatusNotFound)
		return
	}

	departments, err := h.client.ListDepartments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := EditRowData{Employee: emp, Departments: departments}
	if err := h.tmpl.ExecuteTemplate(w, "employee-edit-row", data); err != nil {
		log.Println("template error:", err)
	}
}

func (h *UIHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := UpdateEmployeeReq{}
	if v := r.FormValue("name"); v != "" {
		req.Name = &v
	}
	if v := r.FormValue("sex"); v != "" {
		req.Sex = &v
	}
	if v := r.FormValue("age"); v != "" {
		if age, err := strconv.Atoi(v); err == nil {
			req.Age = &age
		}
	}
	if v := r.FormValue("salary"); v != "" {
		if salary, err := strconv.Atoi(v); err == nil {
			req.Salary = &salary
		}
	}
	if v := r.FormValue("department_code"); v != "" {
		req.DepartmentCode = &v
	}

	if err := h.client.UpdateEmployee(id, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderEmployeesTable(w)
}

func (h *UIHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.client.DeleteEmployee(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderEmployeesTable(w)
}

func (h *UIHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := CreateDepartmentReq{
		Code: r.FormValue("code"),
		Name: r.FormValue("name"),
	}

	if err := h.client.CreateDepartment(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderDepartmentsTable(w)
}

func (h *UIHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("id")

	if err := h.client.DeleteDepartment(code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.renderDepartmentsTable(w)
}

func (h *UIHandler) EmployeesTable(w http.ResponseWriter, r *http.Request) {
	h.renderEmployeesTable(w)
}

func (h *UIHandler) renderEmployeesTable(w http.ResponseWriter) {
	data, err := h.buildPageData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.tmpl.ExecuteTemplate(w, "employees-table", data); err != nil {
		log.Println("template error:", err)
	}
}

func (h *UIHandler) renderDepartmentsTable(w http.ResponseWriter) {
	data, err := h.buildPageData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.tmpl.ExecuteTemplate(w, "departments-table", data); err != nil {
		log.Println("template error:", err)
	}
}
