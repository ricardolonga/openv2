package domain

type UserRepository struct {
	users map[string]*User
}

func (this *UserRepository) Save(user *User) {
	if this.users == nil {
		this.users = make(map[string]*User)
	}

	this.users[user.Email] = user
}
