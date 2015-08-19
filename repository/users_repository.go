package repository

import (
	"github.com/ricardolonga/openv2/entity"
	"log"
	"strings"
)

type UsersRepository struct {
	users map[string]*entity.User
}

func (this *UsersRepository) Save(user *entity.User) *entity.User {
	if this.users == nil {
		this.users = make(map[string]*entity.User)
	}

	this.users[user.Email] = user

	return this.users[user.Email]
}

func (this *UsersRepository) GetAll() *[]entity.User {
	if this.users == nil {
		this.users = make(map[string]*entity.User, 0)
	}

	users := make([]entity.User, 0)

	for _, user := range this.users {
		users = append(users, *user)
	}

	return &users
}

func (this *UsersRepository) Get(email string) *entity.User {
	if this.users == nil {
		this.users = make(map[string]*entity.User, 0)
		return nil
	}

	for _, user := range this.users {
		if strings.EqualFold(user.Email, email) {
			return user
		}
	}

	return nil
}

func (this *UsersRepository) GetAllByEmails(emails []string) []entity.User {
	log.Printf("GetAllByEmails - emails: %s\n", emails)

	if emails == nil {
		return make([]entity.User, 0)
	}

	if this.users == nil {
		this.users = make(map[string]*entity.User, 0)
		log.Printf("this.users == nil.\n")
		return make([]entity.User, 0)
	}

	fullUsers := make([]entity.User, 0)

	for _, user := range this.users {
		for _, email := range emails {
			if strings.EqualFold(email, user.Email) {
				fullUsers = append(fullUsers, *user)
				break
			}
		}
	}

	log.Printf("returning %s\n", fullUsers)

	return fullUsers
}
