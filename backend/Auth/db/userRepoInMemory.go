package db

import (
	"12factorapp/data"
	"errors"
	"log"
)

type UserRepoInMemory struct {
	logger *log.Logger
}

// NewInMemoryRepo Constructor
func NewInMemoryRepo(l *log.Logger) (UserRepoInMemory, error) {
	return UserRepoInMemory{l}, nil
}

func (ur *UserRepoInMemory) GetUsers() data.Users {
	ur.logger.Println("{UserRepoInMemory} - getting all users")
	return userList
}

func (ur *UserRepoInMemory) GetUser(id int) (data.User, error) {
	ur.logger.Println("{UserRepoInMemory} - getting user ", id)
	for _, u := range userList {
		if u.ID == id {
			return *u, nil
		}
	}
	user := data.User{}
	return user, errors.New("couldn't find user")
}

func (ur *UserRepoInMemory) AddUser(u *data.User) {
	ur.logger.Println("{UserRepoInMemory} - adding user", u.ID)
	u.ID = ur.getNextID()
	userList = append(userList, u)
}

func (ur *UserRepoInMemory) getNextID() int {
	max := 0

	for _, currentProd := range ur.GetUsers() {
		if currentProd.ID > max {
			max = currentProd.ID
		}
	}

	return max + 1
}

//In memory data
var userList = []*data.User{
	&data.User{
		ID:        1,
		Username:  "uname",
		FirstName: "fName",
		LastName:  "lName",
		Address:   "abc123",
		Gender:    true,
		Age:       19,
		Password:  "12345",
	},
	&data.User{
		ID:        2,
		Username:  "uname2",
		FirstName: "fName2",
		LastName:  "lName2",
		Address:   "abc122",
		Gender:    true,
		Age:       27,
		Password:  "12345",
	},
}
