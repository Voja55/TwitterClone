package data

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"

	"go.mongodb.org/mongo-driver/bson"

)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     ERole  `json:"role"`
}

type Users []*User

func (p *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *User) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func (p *User) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
