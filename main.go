package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	router.GET("/hello-world", helloWorld)

	router.Run("localhost:8080")
}

func helloWorld(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{ "hello": "world!" })
}
