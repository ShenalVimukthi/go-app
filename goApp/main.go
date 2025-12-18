package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goApp/database"
	"net/http"
	"goApp/routes"
)

func main(){

    // loading env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// ✅ DB connection test
	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("✅ Connected to Supabase")

	// starting gin server
	r:=gin.Default()

	// connecting the routes to the system
	routes.EmployeeRoutes(r)

	// testing setup(fast-check)
	r.GET("/ping",func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{"message":"pong"})
	})

	port:=os.Getenv("PORT")
		if port==" "{
			port="8080"
		}
	

	log.Fatal(r.Run(":"+port))


}
