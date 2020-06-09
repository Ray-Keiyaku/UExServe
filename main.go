package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/init", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"init": "success",
		})
	})
	r.Run(":8080")
}
