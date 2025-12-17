package mappers

import(
	"goApp/models/db"
	"goApp/models/response"
)


func ToEmployeeResponse(emp *db.Employee)(*response.EmployeeResponse){

	// handling empty data 
	if emp==nil{
		return nil
	}

	// returning the mapped output
	return &response.EmployeeResponse{
		
		Id: emp.Id,
		Name: emp.Name,
		Tel: emp.Tel,
		Age: emp.Age,
		Dept: emp.Dept,
		
	}


}