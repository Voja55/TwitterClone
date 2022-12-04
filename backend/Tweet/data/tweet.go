package data

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"io"
)

type Tweet struct {
	TweetId  gocql.UUID `json:"tweetId"`
	Username string     `json:"username"`
	Text     string     `json:"text"`
}

type Like struct {
	TweetId  gocql.UUID `json:"tweetId"`
	Username string     `json:"username"`
	Liked    bool       `json:"liked"`
}

type TweetLikes struct {
	Likes int `json:"likes"`
}

type Tweets []*Tweet

type Likes []*Like

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

func (p *Like) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Like) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Likes) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *TweetLikes) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
