package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = 8080

func main() {

	router := gin.Default()
	router.StaticFile("/", "./APIREADME.md")

	launchStr := fmt.Sprintf("localhost:%d", PORT)
	fmt.Println("serving on ", launchStr)

	router.Run(launchStr)
}
