package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/openv2/entity"
	"github.com/ricardolonga/openv2/repository"
	"log"
	"net/http"
)

func GetUsers(usersRepository *repository.UsersRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, usersRepository.GetAll())
	}
}

func UpdateUser(usersRepository *repository.UsersRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := &entity.User{}

		if c.BindJSON(&user) == nil {
			if user.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "User not contains a email."})
				return
			}

			c.JSON(http.StatusOK, usersRepository.Save(user))
		}
	}
}

func GetUser(usersRepository *repository.UsersRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")

		log.Printf("Email: %s\n", email)

		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"devMessage": "Email is required."})
			return
		}

		user := usersRepository.Get(email)

		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"devMessage": "User not found."})
		}

		c.JSON(200, user)
	}
}
