package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Tweet struct {
	ID      string `json:"id"`
	Author  string `json:"author" validate:"required"`
	Text    string `json:"text" validate:"required"`
	Picture string `json:"picture"`
	//Likes - lista i guess i onda kao usernamovi ljudi koji su lajkovali jer je username jedinstev a lepse je od samog id-a
}

type Tweets []*Tweet

func (p *Tweets) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tweet) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tweet) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Tweet) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
