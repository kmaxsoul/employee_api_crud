package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/start", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})

	r.Run(":9060")

}
