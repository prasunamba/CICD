package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerfunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "your ping reached here",
	})

}
func main() {
	r := gin.Default()
	r.GET("/ping", handlerfunc)
	r.Run()
}
