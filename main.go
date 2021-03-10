package main

import (
	"finego/handler"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/ping", handler.HelloWorld)
	r.Run()
}
