package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/openv2/entity"
	"github.com/ricardolonga/openv2/repository"
	"log"
	"net/http"
	"strings"
)

func GetEvents(eventRepository *repository.EventsRepository) func(c *gin.Context) {
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

func CreateEvent(eventRepository *repository.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := &entity.Event{}

		if c.BindJSON(&event) == nil {
			if event.Owner.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "Event not contains a owner."})
				return
			}

			c.JSON(http.StatusOK, eventRepository.Save(event))
		}
	}
}

func Checkin(eventRepository *repository.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := eventRepository.Get(c.Param("id"))

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "Event not found."})
		}

		var user entity.User

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

func Checkout(eventRepository *repository.EventsRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := eventRepository.Get(c.Param("id"))

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "Event not found."})
		}

		var user entity.User

		if c.BindJSON(&user) == nil {
			if user.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "User not contains a email."})
				return
			}

			members := make([]entity.User, 0)

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

func GetMembers(eventRepository *repository.EventsRepository, usersRepository *repository.UsersRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		event := eventRepository.Get(c.Param("id"))

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "Event not found."})
			return
		}

		skill := c.Query("skill")

		if skill == "" {
			if event.Members == nil {
				c.JSON(http.StatusOK, make([]entity.User, 0))
				return
			}

			c.JSON(200, event.Members)
			return
		}

		emails := make([]string, 0)

		if event.Members == nil {
			c.JSON(http.StatusOK, make([]entity.User, 0))
			return
		}

		for _, member := range event.Members {
			emails = append(emails, member.Email)
		}

		fullMembers := usersRepository.GetAllByEmails(emails)

		members := make([]entity.User, 0)

		skills := strings.Split(skill, ",")

	user:
		for _, fullUser := range fullMembers {
			for _, userSkill := range fullUser.Skills {
				for _, querySkill := range skills {
					if strings.EqualFold(querySkill, userSkill) {
						members = append(members, fullUser)
						continue user
					}
				}
			}
		}

		c.JSON(http.StatusOK, members)
	}
}
