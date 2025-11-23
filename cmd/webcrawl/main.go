package main

import (
	"log"

route "github.com/quandev306/webcrawl/api/route"
"github.com/quandev306/webcrawl/bootstrap"
"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
