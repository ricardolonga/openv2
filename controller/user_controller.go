package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardolonga/openv2/entity"
	"github.com/ricardolonga/openv2/repository"
	"net/http"
)

func GetUsers(usersRepository *repository.UsersRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, usersRepository.GetAll())
	}
}

//func CreateUser(usersRepository *repository.UsersRepository) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		user := &entity.User{}
//
//		if c.BindJSON(&user) == nil {
//			if user.Email == "" {
//				c.JSON(http.StatusBadRequest, gin.H{"devMessage": "User not contains a email."})
//				return
//			}
//
//			c.JSON(http.StatusOK, usersRepository.Save(user))
//		}
//	}
//}

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
