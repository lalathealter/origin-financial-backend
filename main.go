package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lalathealter/originfin/controllers"
)

const PORT = 8080

func main() {

	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.POST("/risks", controllers.HandleRisksCalculation)
	router.StaticFile("/", "./APIREADME.md")

	launchStr := fmt.Sprintf("localhost:%d", PORT)
	fmt.Println("serving on ", launchStr)

	router.Run(launchStr)
}
