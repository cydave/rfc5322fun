package main

import (
	"html/template"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	router := gin.Default()
	router.FuncMap["noescape"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/xss", func(c *gin.Context) {
		email := c.Query("email")
		_, err := mail.ParseAddress(email)
		if err != nil {
			c.HTML(http.StatusOK, "xss.tmpl", gin.H{
				"error": true,
				"email": "invalid@email.com",
			})
		} else {
			c.HTML(http.StatusOK, "xss.tmpl", gin.H{
				"error": false,
				"email": email,
			})
		}
	})
	router.Run(":8080")
}
