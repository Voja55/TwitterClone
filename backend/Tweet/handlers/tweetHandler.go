package handlers

import (
	"Tweet/data"
	"Tweet/db"
	"context"
	"github.com/gocql/gocql"
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
	tweets, err := t.tweetRepo.GetTweets()
	if err != nil {
		http.Error(rw, "Problem with getting tweets from db", http.StatusInternalServerError)
		t.logger.Println("Problem with getting tweets from db :", err)
		return
	}

	err = tweets.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Println("Unable to convert to json :", err)
		return
	}
}

func (t *TweetsHandler) GetLikes(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	if id == "" {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	idUUID, err := gocql.ParseUUID(id)
	likes, err := t.tweetRepo.GetLikes(idUUID)
	if err != nil {
		http.Error(rw, "Problem with getting likes from db", http.StatusInternalServerError)
		t.logger.Println("Problem with getting likes from db :", err)
		return
	}

	tweetLikes := data.TweetLikes{Likes: likes}
	err = tweetLikes.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Println("Unable to convert to json :", err)
		return
	}
}

func (t *TweetsHandler) LikeTweet(rw http.ResponseWriter, h *http.Request) {
	liked := h.Context().Value(KeyUser{}).(*data.Like)
	if liked.Username == "" || liked.TweetId.String() == "" {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	result := t.tweetRepo.LikeTweet(liked.TweetId, liked.Username)
	if result == true {
		rw.WriteHeader(http.StatusAccepted)
		return
	} else {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

}

func (t *TweetsHandler) CreateTweet(rw http.ResponseWriter, h *http.Request) {
	tweet := h.Context().Value(KeyUser{}).(*data.Tweet)
	if tweet.Username == "" || tweet.Text == "" {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	result, err := t.tweetRepo.CreateTweet(tweet)
	if err != nil {
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

func (t *TweetsHandler) MiddlewareLikeValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		tweet := &data.Like{}
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

func (t *TweetsHandler) GetTweetsByUser(writer http.ResponseWriter, request *http.Request) {

}
