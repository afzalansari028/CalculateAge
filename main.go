package main

import (
	"fmt"
	"practice/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")

	router := gin.Default()

	router.POST("/reverse", handler.API)
	router.POST("/age", handler.FindAge)

	router.Run()
}
