package main

import (
	"fmt"
	"francischacko/github.com/go_razorpay_integration/controller"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// var c *gin.Context
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		str, err := controller.Executerazorpay()
		encodedStr := template.HTMLEscapeString(str)
		fmt.Println(str)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": "errrrrrrrrrr",
			})
			return
		}

		data := struct {
			Str string
		}{
			Str: encodedStr,
		}
		tmpl, err := template.ParseFiles("rzp.html")
		if err != nil {
			log.Println("Failed to parse template:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
			})
			return
		}

		err = tmpl.Execute(c.Writer, data)
		if err != nil {

			log.Println("Failed to execute template:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
			})
			return
		}
	})

	router.Run("localhost:5000")
}
