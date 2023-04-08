package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat instance router Gin
	r := gin.Default()

	// Menambahkan route GET
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// Menjalankan server HTTP
	r.Run(":8080")

	now := time.Now()
	fmt.Printf("\n%s ⚡️ You are already connected to Golang!\n\n", now.Format("2006-01-02 15:04:05"))
}
