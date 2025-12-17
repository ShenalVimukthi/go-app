package services

import(
	// "goApp/models/request"
	"goApp/models/response"
	"goApp/mappers"
	"goApp/database"
	// "goApp/models/db"
	"fmt"
	"errors"
)

// create custom errors 
var ErrUserNotFound = errors.New("User Not Found")
var ErrUserExists = errors.New("User Already Exists")


// mock data

// functions to do CRUD operations on employee data
func GetEmployee(id int)(*response.EmployeeResponse,error){


	// sending employee data inside the database querying section
	ReceivedEmp,err:=database.GetEmployeesById(id)

	// handling errors 
	if err!=nil{
		return nil,err
	}

	// emp nil error handling
	if ReceivedEmp==nil{
		return nil,fmt.Errorf("GetEmployeesById failed for id %d: %w", id, ErrUserNotFound)
	}

	// feeding the data into response model using mapper
	emp:=mappers.ToEmployeeResponse(ReceivedEmp)
	

	return emp,nil


}

// func CreateEmployee(employee *request.CreateEmployee)(*response.EmployeeResponse,error){
// 	var dbVersion *db.Employee

// }

// func UpdateEmployee()(){

// }

// func DeleteEmployee(id int )(*response.EmployeeResponse,error){
// }