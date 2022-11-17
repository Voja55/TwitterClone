package db

import "Tweet/data"

type TweetRepo interface {
	GetTweets() data.Tweets
	GetTweet(id string) (data.Tweet, error)
	CreateTweet(p *data.Tweet) (bool, error)
	LikeTweet(id string, username string) bool
	GetTweetsByUser(id int) (data.Tweets, error)
}
