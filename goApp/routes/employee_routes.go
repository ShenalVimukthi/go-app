package routes

import (
	"goApp/handlers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	EmployeeGroup := r.Group("employees/")
	{
		EmployeeGroup.GET("/:id", handlers.GetEmployee)
		EmployeeGroup.POST("/create_employee", handlers.CreateEmployee)
		EmployeeGroup.PUT("/:id", handlers.UpdateEmployee)
		EmployeeGroup.DELETE("/:id", handlers.DeleteEmployee)
	}
}
