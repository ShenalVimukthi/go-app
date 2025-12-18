package routes

import (
	"goApp/handlers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	EmployeeGroup := r.Group("employees/")
	{
		EmployeeGroup.GET("/:id", handlers.GetEmployee)
	}
}
