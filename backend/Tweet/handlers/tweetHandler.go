package handlers

import (
	"Tweet/data"
	"Tweet/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type KeyUser struct{}

type TweetsHandler struct {
	logger    *log.Logger
	tweetRepo db.TweetRepo
}

// NewUsersHandler Injecting the logger makes this code much more testable.
func NewTweetsHandler(l *log.Logger, ur db.TweetRepo) *TweetsHandler {
	return &TweetsHandler{l, ur}
}

func (t *TweetsHandler) GetTweets(rw http.ResponseWriter, h *http.Request) {
	users := t.tweetRepo.GetTweets()
	err := users.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Println("Unable to convert to json :", err)
		return
	}
}

func (t *TweetsHandler) GetTweet(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	tweet, er := t.tweetRepo.GetTweet(id)

	if er != nil {
		http.Error(rw, er.Error(), http.StatusNotFound)
		t.logger.Println("Unable to find tweet.", er)
		return
	}

	err := tweet.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Println("Unable to convert to json :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (t *TweetsHandler) CreateTweet(rw http.ResponseWriter, h *http.Request) {
	tweet := h.Context().Value(KeyUser{}).(*data.Tweet)
	result, err := t.tweetRepo.CreateTweet(tweet)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if result == true {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
	rw.Write([]byte("406 - Not acceptable"))
}

//Middleware to try and decode the incoming body. When decoded we run the validation on it just to check if everything is okay
//with the model. If anything is wrong we terminate the execution and the code won't even hit the handler functions.
//With a key we bind what we read to the context of the current request. Later we use that key to get to the read value.

func (t *TweetsHandler) MiddlewareTweetsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		tweet := &data.Tweet{}
		err := tweet.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			t.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyUser{}, tweet) //Ne znam sta je KeyUser{}
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (t *TweetsHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		t.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (t *TweetsHandler) LikeTweet(writer http.ResponseWriter, request *http.Request) {

}

func (t *TweetsHandler) GetTweetsByUser(writer http.ResponseWriter, request *http.Request) {

}

func isEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
}
