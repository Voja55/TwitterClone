package db

import "12factorapp/data"

type UserRepo interface {
	GetUsers() data.Users
	GetUser(id string) (data.User, error)
	LoginUser(username string, password string) (data.User, error)
	Register(p *data.User) bool
	GetUserByUsername(username string) (data.User, error)
	UpdateUser(u *data.User) bool
}
