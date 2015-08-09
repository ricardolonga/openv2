package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/openv2/domain"
	"net/http"
	"strings"
	"log"
)

type Response struct {
	UserMessage string       `json:"userMessage"`
	User        *domain.User `json:"user"`
}

func main() {
	eventRepository := &domain.EventsRepository{}

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	//	router.Use(HeadersRequired())

	router.GET("/events", GetEvents(eventRepository))
	router.POST("/events", CreateEvent(eventRepository))
	router.PUT("/events/:id/checkin", Checkin(eventRepository))
	router.PUT("/events/:id/checkout", Checkout(eventRepository))
	router.GET("/events/:id/members", GetMembers(eventRepository))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
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

func GetEvents(eventRepository *domain.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Query("name")

		log.Printf("Name: %s\n", name)

		if name == "" {
			c.JSON(200, eventRepository.GetAll())
			return
		}

		c.JSON(200, eventRepository.GetByName(name))
	}
}

func CreateEvent(eventRepository *domain.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := &domain.Event{}

		if c.BindJSON(&event) == nil {
			if event.Owner.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "Event not contains a owner."})
				return
			}

			c.JSON(http.StatusOK, eventRepository.Save(event))
		}
	}
}

func Checkin(eventRepository *domain.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := eventRepository.Get(c.Param("id"))

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "Event not found."})
		}

		var user domain.User

		if c.BindJSON(&user) == nil {
			if user.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "User not contains a email."})
				return
			}

			event.Members = append(event.Members, user)

			c.JSON(http.StatusOK, eventRepository.Save(event))
		}
	}
}

func Checkout(eventRepository *domain.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := eventRepository.Get(c.Param("id"))

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "Event not found."})
		}

		var user domain.User

		if c.BindJSON(&user) == nil {
			if user.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "User not contains a email."})
				return
			}

			members := make([]domain.User, 0)

			for _, member := range event.Members {
				if strings.EqualFold(member.Email, user.Email) {
					continue
				}

				members = append(members, member)
			}

			event.Members = members

			c.JSON(http.StatusOK, eventRepository.Save(event))
		}
	}
}

func GetMembers(eventRepository *domain.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := eventRepository.Get(c.Param("id"))

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "Event not found."})
		}

		skill := c.Query("skill")

		log.Printf("Skill: %s\n", skill)

		if skill == "" {
			c.JSON(200, event.Members)
			return
		}

		members := make([]domain.User, 0)

		for _, user := range event.Members {
			for _, s := range user.Skills {
				if strings.EqualFold(skill, s) {
					members = append(members, user)
				}
			}
		}

		c.JSON(200, members)
	}
}
