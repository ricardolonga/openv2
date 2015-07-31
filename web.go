package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"openv2/domain"
	"os"
)

type Response struct {
	UserMessage string       `json:"userMessage"`
	User        *domain.User `json:"user"`
}

func main() {
	userRepository := &domain.UserRepository{}

	router := gin.Default()

	router.Use(HeadersRequired())

	router.GET("/", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Olá!"))
	})

	router.POST("/login", func(c *gin.Context) {
		user := &domain.User{Name: c.Request.Header.Get("name"), Email: c.Request.Header.Get("email")}
		userRepository.Save(user)
		c.JSON(http.StatusOK, &Response{UserMessage: "Login successful!", User: user})
	})

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), router)
	if err != nil {
		panic(err)
	}
}

func HeadersRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, contain := c.Request.Header["Name"]; !contain {
			c.JSON(http.StatusBadRequest, &Response{UserMessage: "Name is not defined into headers."})
			c.Abort()
		}

		if _, contain := c.Request.Header["Email"]; !contain {
			c.JSON(http.StatusBadRequest, &Response{UserMessage: "Email is not defined into headers."})
			c.Abort()
		}
	}
}
