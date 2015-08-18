package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/openv2/controller"
	"github.com/ricardolonga/openv2/entity"
	"github.com/ricardolonga/openv2/repository"
	"net/http"
)

type Response struct {
	UserMessage string       `json:"userMessage"`
	User        *entity.User `json:"user"`
}

func main() {
	usersRepository := &repository.UsersRepository{}
	eventRepository := &repository.EventsRepository{}

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	//	router.Use(HeadersRequired())

	router.GET("/events", controller.GetEvents(eventRepository))
	router.POST("/events", controller.CreateEvent(eventRepository))
	router.PUT("/events/:id/checkin", controller.Checkin(eventRepository))
	router.PUT("/events/:id/checkout", controller.Checkout(eventRepository))
	router.GET("/events/:id/members", controller.GetMembers(eventRepository))

//	router.POST("/users", controller.CreateUser(usersRepository))
	router.PUT("/users", controller.UpdateUser(usersRepository))
	router.GET("/users", controller.GetUsers(usersRepository))

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

//func HeadersRequired() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if _, contain := c.Request.Header["Name"]; !contain {
//			c.JSON(http.StatusBadRequest, &Response{UserMessage: "Name is not defined into headers."})
//			c.Abort()
//		}
//
//		if _, contain := c.Request.Header["Email"]; !contain {
//			c.JSON(http.StatusBadRequest, &Response{UserMessage: "Email is not defined into headers."})
//			c.Abort()
//		}
//	}
//}
