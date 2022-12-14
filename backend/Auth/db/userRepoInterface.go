package db

import "12factorapp/data"

type UserRepo interface {
	GetUsers() data.Users
	AddUser(p *data.User)
	GetUser(id int) (data.User, error)
}
