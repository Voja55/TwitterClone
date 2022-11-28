package db

import "Tweet/data"

type TweetRepo interface {
	GetTweets() (data.Tweets, error)
	CreateTweet(p *data.Tweet) (bool, error)
	LikeTweet(id string, username string) bool
	GetTweetsByUser(id int) (data.Tweets, error)
}
