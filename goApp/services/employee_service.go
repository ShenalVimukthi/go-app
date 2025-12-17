package services

import(
	// "goApp/models/request"
	"goApp/models/response"
	"goApp/mappers"
	"goApp/database"
	// "goApp/models/db"
	// "errors"
	// "fmt"
)

// create custom errors 



// mock data

// functions to do CRUD operations on employee data
func GetEmployee(id int)(*response.EmployeeResponse,error){


	// sending emloyee data inside the database querying section
	ReceivedEmp,err:=database.GetEmployeesById(id)

	// handling errors 
	if err!=nil{
		return nil,err
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