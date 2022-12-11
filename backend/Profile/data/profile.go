package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Profile struct {
	Username    string  `json:"username"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Address     string  `json:"address"`
	Gender      EGender `json:"gender"`
	Age         int     `json:"age"`
	CompanyName string  `json:"companyName"`
	WebSite     string  `json:"webSite"`
	Email       string  `json:"email"`
	Private     bool    `json:"private"`
}

type Profiles []*Profile

func (p *Profiles) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Profile) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Profile) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Profile) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
