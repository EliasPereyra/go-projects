package main

import "github.com/gin-gonic/gin"

func main() {
	// main instance of the web server
	router := gin.Default()

	router.GET("/checkhealth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	router.Run(":4000")
}
