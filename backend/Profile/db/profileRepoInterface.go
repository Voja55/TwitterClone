package db

import "Profile/data"

type ProfileRepo interface {
	GetProfile(id string) (data.Profile, error)
	CreateProfile(p *data.Profile) bool
}
