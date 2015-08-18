package repository
import (
	"github.com/ricardolonga/openv2/entity"
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