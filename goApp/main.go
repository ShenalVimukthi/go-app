package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main(){

    // loading env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// starting gin server
	r:=gin.Default()

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
