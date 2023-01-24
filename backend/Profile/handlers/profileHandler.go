package handlers

import (
	"Profile/data"
	"Profile/db"
	"Profile/validation"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type KeyProfile struct{}

type ProfileHandler struct {
	logger      *log.Logger
	profileRepo db.ProfileRepo
}

func NewProfileHandler(l *log.Logger, ur db.ProfileRepo) *ProfileHandler {
	return &ProfileHandler{l, ur}
}

func (p *ProfileHandler) GetProfile(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	username := vars["username"]

	user, er := p.profileRepo.GetProfile(username)

	if er != nil {
		http.Error(rw, er.Error(), http.StatusNotFound)
		p.logger.Println("Unable to find user.", er)
		return
	}

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Println("Unable to convert to json :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *ProfileHandler) CreateNormalProfile(rw http.ResponseWriter, h *http.Request) {
	profile := h.Context().Value(KeyProfile{}).(*data.Profile)

	if profile.FirstName == "" || profile.LastName == "" || profile.Address == "" ||
		profile.Username == "" || profile.Age == 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if validation.ValidateChar(profile.FirstName) && validation.ValidateChar(profile.LastName) &&
		validation.ValidateChar(profile.Address) {

		_, err := p.profileRepo.GetProfile(profile.Username)
		if err == nil {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}

		status := p.profileRepo.CreateProfile(profile)
		if status != true {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}
		if status == true {
			rw.WriteHeader(http.StatusAccepted)
			return
		}

	}

	rw.WriteHeader(http.StatusNotAcceptable)
	rw.Write([]byte("406 - Not acceptable"))
}

func (p *ProfileHandler) CreateBusinessProfile(rw http.ResponseWriter, h *http.Request) {
	profile := h.Context().Value(KeyProfile{}).(*data.Profile)

	if profile.CompanyName == "" || profile.WebSite == "" || profile.Email == "" ||
		profile.Username == "" {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if validation.ValidateChar(profile.CompanyName) && validation.ValidateCharAndNum(profile.WebSite) {

		_, err := p.profileRepo.GetProfile(profile.Username)
		if err == nil {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}

		status := p.profileRepo.CreateProfile(profile)
		if status != true {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}
		if status == true {
			rw.WriteHeader(http.StatusAccepted)
			return
		}
	}

	rw.WriteHeader(http.StatusNotAcceptable)
	rw.Write([]byte("406 - Not acceptable"))
}

func (p *ProfileHandler) MiddlewareUsersValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		profile := &data.Profile{}
		err := profile.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProfile{}, profile)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (p *ProfileHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
