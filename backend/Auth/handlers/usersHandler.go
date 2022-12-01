package handlers

import (
	"12factorapp/data"
	"12factorapp/db"
	"12factorapp/validation"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type KeyUser struct{}

type UsersHandler struct {
	logger   *log.Logger
	userRepo db.UserRepo
}

type LogUser struct {
	Username string
	Password string
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type Jwt struct {
	Token string `json:"jwt"`
}

type PasswordReset struct {
	Token string `json:"token"`
	Password string `json:"password"`
}

var jwtKey = []byte("secret_key")

// NewUsersHandler Injecting the logger makes this code much more testable.
func NewUsersHandler(l *log.Logger, ur db.UserRepo) *UsersHandler {
	return &UsersHandler{l, ur}
}

func (u *UsersHandler) GetUsers(rw http.ResponseWriter, h *http.Request) {
	users := u.userRepo.GetUsers()
	err := users.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}

}

func (u *UsersHandler) GetUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]


	user, er := u.userRepo.GetUser(id)

	if er != nil {
		http.Error(rw, er.Error(), http.StatusNotFound)
		u.logger.Println("Unable to find user.", er)
		return
	}

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *UsersHandler) LoginUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var logged data.User
	err := decoder.Decode(&logged)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}
	u.logger.Println(logged)
	if validation.ValidateUsername(logged.Username) && validation.ValidatePassword(logged.Password) {

		_, err := u.userRepo.LoginUser(logged.Username, logged.Password)
		if err != nil {
			http.Error(rw, "Unable to login", http.StatusInternalServerError)
			u.logger.Println("Unable to login", err)
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("401 - Unauthorized"))
			return
		}

		expirationTime := time.Now().Add(time.Minute * 5)

		claims := &Claims{
			Username: logged.Username,
			Role:     string(logged.Role),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		var userResponse Jwt
		userResponse.Token = tokenString
		jsonUser, err := json.Marshal(userResponse)
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(jsonUser)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
	rw.Write([]byte("406 - Not acceptable"))
}

func (u *UsersHandler) Register(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyUser{}).(*data.User)

		if validation.ValidateUsername(user.Username) && validation.ValidatePassword(user.Password) && validation.ValidateRole(string(user.Role)) {
			_, err := u.userRepo.GetUserByUsername(user.Username)
			if err == nil {
				rw.WriteHeader(http.StatusNotAcceptable)
				rw.Write([]byte("406 - Not acceptable"))
				return
			}
			if u.userRepo.Register(user) {
				_, err = SendMail(user.Email, "Confirmation code", strconv.Itoa(user.CCode))
				if err != nil {
					u.logger.Println("Faild to email", err)
				}
				rw.WriteHeader(http.StatusAccepted)
				rw.Write([]byte("202 - Accepted"))
				return
			}

		if u.userRepo.Register(user) == true {
			rw.WriteHeader(http.StatusAccepted)
			rw.Write([]byte("202 - Accepted"))
			return
		}
	}

	rw.WriteHeader(http.StatusNotAcceptable)
	rw.Write([]byte("406 - Not acceptable"))
}

func (u *UsersHandler) Confirm(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data data.User
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}
	u.logger.Println(data)

	user, err := u.userRepo.GetUser(data.ID)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		u.logger.Println("Unable to find user.", err)
		return
	}
	
	if user.CCode != data.CCode {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	user.CCode = 0
	if u.userRepo.UpdateUser(&user) == false {
		rw.WriteHeader(http.StatusNotAcceptable)
	} 
	rw.WriteHeader(http.StatusAccepted)

}

func (u *UsersHandler) RequestResetPassword(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data data.User
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}
	u.logger.Println(data)

	user, err := u.userRepo.GetUserByEmail(data.Email)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		u.logger.Println("Unable to find user.", err)
		return
	}
	//TODO zameni base64 sa nekom vrstom sifrovanja
	//TODO Link ka frontu da bude env promenljiva
	encoded := base64.StdEncoding.EncodeToString([]byte(user.Email))
	_, err1 := SendMail(user.Email, "Reset password", "https://localhost:4200/resetPass?id=" + encoded)
	if err1 != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}

func (u *UsersHandler) ResetPassword(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data PasswordReset
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}
	u.logger.Println(data)

	decoded, err := base64.StdEncoding.DecodeString(data.Token)
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		u.logger.Println("Unable to decode token", err)
		return
	}

	user, err := u.userRepo.GetUserByEmail(string(decoded))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		u.logger.Println("Unable to find user.", err)
		return
	}

	hashedPass, err := user.HashPassword(data.Password)
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	user.Password = hashedPass
	if !u.userRepo.UpdateUser(&user) {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}

//Middleware to try and decode the incoming body. When decoded we run the validation on it just to check if everything is okay
//with the model. If anything is wrong we terminate the execution and the code won't even hit the handler functions.
//With a key we bind what we read to the context of the current request. Later we use that key to get to the read value.

func (u *UsersHandler) MiddlewareUsersValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &data.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyUser{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func isEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
}
