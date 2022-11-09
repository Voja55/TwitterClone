package data

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson"
)

type RegisterUser struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     ERole  `json:"role" validate:"required"`
}

type RegisterUsers []*RegisterUser

func (p *RegisterUsers) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *RegisterUser) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *RegisterUser) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *RegisterUser) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
