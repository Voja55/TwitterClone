package db

import "Profile/data"

type UserRepo interface {
	GetProfile(id int) (data.Profile, error)
	CreateProfile(p *data.Profile) bool
}
