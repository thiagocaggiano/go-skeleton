package main

import (
	"github.com/gin-gonic/gin")

var (
	Router *gin.Engine
)
func main() {
	configureRouter()
	mapUrlToPing()
	Router.Run(":8080")
}

func configureRouter() {
	// Configure Router
	Router = gin.Default()
}

func mapUrlToPing() {
	// Add health check 
	Router.GET("/ping", func(c *gin.Context) {
        c.String(200,"pong")
    })
}