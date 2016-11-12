package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {port = "5000"}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/user/:name", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title" : c.Param("name"),
		})
	})

	router.GET("/json/:name", func(c *gin.Context) {
		c.JSON(200, gin.H{
            c.Param("name"): "x65han",
        })
	})

	//Run Server
	router.Run(":" + port)
}
