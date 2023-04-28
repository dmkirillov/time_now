package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().UnixNano(),
		})
	})
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
