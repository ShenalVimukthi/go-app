package services

import (
	"goApp/database"
	"goApp/mappers"
	"goApp/models/request"
	"goApp/models/response"

	// "goApp/models/db"
	"errors"
	"fmt"
)

// create custom errors

var(
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists = errors.New("user already exists")
	ErrCreateEmployeeFailed = errors.New("create employee failed")
)


// mock data

// functions to do CRUD operations on employee data
func GetEmployee(id int) (*response.EmployeeResponse, error) {

	// sending employee data inside the database querying section
	ReceivedEmp, err := database.GetEmployeesById(id)

	// handling errors
	if err != nil {
		return nil, err
	}

	// emp nil error handling
	if ReceivedEmp == nil {
		return nil, fmt.Errorf("GetEmployeesById failed for id %d: %w", id, ErrUserNotFound)
	}

	// feeding the data into response model using mapper
	emp := mappers.ToEmployeeResponse(ReceivedEmp)

	return emp, nil

}

func CreateEmployee(employee *request.CreateEmployee) (*response.EmployeeResponse, error) {

	// converting data to database model
	dbFormat:=mappers.ToDataBaseModel(employee)

	// sending the employee data to table
	receivedEmp,err:=database.CreateEmployee(dbFormat)

	// error handling
	if err!=nil{
		return nil,fmt.Errorf("CreatEmployee failed: %w",ErrCreateEmployeeFailed)
	}

	// reconverting the db model to response model
	emp:=mappers.ToEmployeeResponse(receivedEmp)

	return emp,nil


}

// func UpdateEmployee()(){

// }

// func DeleteEmployee(id int )(*response.EmployeeResponse,error){
// }
