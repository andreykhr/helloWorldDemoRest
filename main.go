package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()

	router.GET("/employee/:id", handler.GetEmployee)
	router.POST("/employee", handler.CreateEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()
}
