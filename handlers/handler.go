package handlers

import "github.com/gin-gonic/gin"

// ExampleHandler is an example handler function.
func ExampleHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello from the handler!"})
}
