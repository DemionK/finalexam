package main

import (
	"fmt"
	"net/http"

	"github.com/DemionK/finalexam/customer"
	"github.com/gin-gonic/gin"
)

var PORT = ":2019"

func authMiddleware(c *gin.Context) {
	fmt.Println("before middleware")
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}
	c.Next()
	fmt.Println("after middleware")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(authMiddleware)
	r.POST("/customers", customer.PostHandler)
	r.GET("/customers/:id", customer.GetByIDHandler)
	r.GET("/customers", customer.GetHandler)
	r.PUT("/customers/:id", customer.PutHandler)
	r.DELETE("/customers/:id", customer.DeleteByIDHandler)
	return r
}

func main() {
	customer.InitDatabase()
	r := setupRouter()
	r.Run(PORT)
}
