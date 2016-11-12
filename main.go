package main

import (
	"net/http"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {port = "5000"}

	// {{template "header.tmpl.html"}}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/:name", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		name := c.Params.ByName("name")
		fmt.Println("Params: ", name);
	})

	//Run Server
	router.Run(":" + port)
}
