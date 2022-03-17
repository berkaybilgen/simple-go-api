package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

func getUser(c *gin.Context) {
	user := User{c.Param("name")}
	c.IndentedJSON(http.StatusOK, user)
}

func main() {
	routes := gin.Default()
	routes.GET("/test/:name", getUser)
	routes.Run("localhost:8081")
}
