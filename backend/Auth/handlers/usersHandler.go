package handlers

import (
	"12factorapp/data"
	"12factorapp/db"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type KeyUser struct{}

type UsersHandler struct {
	logger   *log.Logger
	userRepo db.UserRepo
}

// Injecting the logger makes this code much more testable.
func NewUsersHandler(l *log.Logger, ur db.UserRepo) *UsersHandler {
	return &UsersHandler{l, ur}
}

func isEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
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
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		u.logger.Println("Unable to convert from ascii to integer - input was :", vars["id"])
		return
	}

	user, er := u.userRepo.GetUser(id)

	if er != nil {
		http.Error(rw, er.Error(), http.StatusNotFound)
		u.logger.Println("Unable to find user.", er)
		return
	}

	err = user.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

type Log_user struct {
	Username string
	Password string
}

func (u *UsersHandler) LoginUser(rw http.ResponseWriter, req *http.Request) {
	//username := req.FormValue("username")
	//password := req.FormValue("password")
	//vars := req.Body
	//username := strconv.Quote(vars.)
	//password := strconv.Quote(vars["password"])
	//if err != nil {
	//	http.Error(rw, "Unable to convert username and pass to json", http.StatusInternalServerError)
	//	u.logger.Println("Unable to convert to json :", err)
	//	return
	//}
	//if !isEmpty(username) && !isEmpty(password) {
	//
	//	loggeduser, err := u.LoginUser(username, password)
	//}
	decoder := json.NewDecoder(req.Body)
	var logu Log_user
	err := decoder.Decode(&logu)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}
	u.logger.Println(logu)
	if !isEmpty(logu.Username) && !isEmpty(logu.Password) {

		loggeduser, err := u.userRepo.GetLoginUser(logu.Username, logu.Password)
		if err != nil {
			http.Error(rw, "Unable to login", http.StatusInternalServerError)
			u.logger.Println("Unable to login", err)
			return
		}
		u.logger.Println(loggeduser)
	}
}

func (u *UsersHandler) PostUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyUser{}).(*data.User)
	u.userRepo.AddUser(user)
	rw.WriteHeader(http.StatusCreated)
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

		//TODO nema sku, pa ne radi
		//err = user.Validate()
		//
		//if err != nil {
		//	u.logger.Println("Error validating product", err)
		//	http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
		//	return
		//}

		ctx := context.WithValue(h.Context(), KeyUser{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

//Middleware to centralize general logging and to add the header values for all requests.

func (u *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
