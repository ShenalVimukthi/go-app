package mappers

import (
	"goApp/models/db"
	"goApp/models/request"
	"goApp/models/response"
)

func ToEmployeeResponse(emp *db.Employee) *response.EmployeeResponse {

	// handling empty data
	if emp == nil {
		return nil
	}

	// returning the mapped output
	return &response.EmployeeResponse{

		Id:    emp.Id,
		Name:  emp.Name,
		Email: emp.Email,
		Tel:   emp.Tel,
		Age:   emp.Age,
		Dept:  emp.Dept,
		CreatedAt: emp.CreatedAt.String(),
	}

}

func ToDataBaseModel(emp *request.CreateEmployee) *db.Employee {

	// error handling
	if emp == nil {
		return nil
	}

	return &db.Employee{

		Name:  emp.Name,
		Email: emp.Email,
		Tel:   emp.Tel,
		Age:   emp.Age,
		Dept:  emp.Dept,
	}

}

func ToDataBaseModelForUpdate(emp *request.UpdateEmployee) *db.Employee {

	// error handling
	if emp == nil {
		return nil
	}

	return &db.Employee{

		Name:  emp.Name,
		Email: emp.Email,
		Tel:   emp.Tel,
		Age:   emp.Age,
		Dept:  emp.Dept,
	}

}
