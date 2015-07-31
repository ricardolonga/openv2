package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/openv2/domain"
	"net/http"
)

type Response struct {
	UserMessage string       `json:"userMessage"`
	User        *domain.User `json:"user"`
}

func main() {
	userRepository := &domain.UserRepository{}

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	router.Use(HeadersRequired())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Olá, estou na AWS!")
	})

	router.POST("/login", func(c *gin.Context) {
		user := &domain.User{Name: c.Request.Header.Get("name"), Email: c.Request.Header.Get("email")}
		userRepository.Save(user)
		c.JSON(http.StatusOK, &Response{UserMessage: "Login successful!", User: user})
	})

	err := http.ListenAndServe(":8080", router)
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
