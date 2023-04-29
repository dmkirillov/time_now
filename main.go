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
			"time":  time.Now().UnixNano(),
			"UTC":   time.Now().UTC(),
			"UTC+3": time.Now().UTC().Add(time.Hour * 3),
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK.")
	})
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
