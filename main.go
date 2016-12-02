package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to defaul port on server
	port := os.Getenv("PORT")
	// If not specified, use port 5000
	if port == "" {port = "5000"}

	// Initialize Gin server
	router := gin.New()
	// Initialize Gin logger
	router.Use(gin.Logger())
	// Load all html in templates directory
	router.LoadHTMLGlob("templates/*.html")
	// REST API
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

	router.GET("/info", func(c *gin.Context) {
		kids := []interface{}{
			"Selena",
			"Montegomery",
		}
		section := "Name"
		f := map[string]interface{}{
		    section: "Wednesday",
		    "Age":  6,
		    "Parents": []interface{}{
		        "Gomez",
		        "Morticia",
		    },
			"Kids":kids,
		}
		c.JSON(200, f)
	})

	//Run Server
	router.Run(":" + port)
}
