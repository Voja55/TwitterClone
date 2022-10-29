package db

import (
	"12factorapp/data"
	"log"
)

type UserRepoDB struct {
	log *log.Logger
	//db *gorm.DB
}

// NewUserRepoDB Constructor
func NewUserRepoDB(log *log.Logger) (UserRepoDB, error) {
	return UserRepoDB{log}, nil
}

func (u UserRepoDB) GetUsers() data.Users {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoDB) AddUser(p *data.User) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepoDB) GetUser(id int) (data.User, error) {
	//TODO implement me
	panic("implement me")
}
