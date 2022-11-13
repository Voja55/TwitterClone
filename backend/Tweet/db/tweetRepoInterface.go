package db

import "Tweet/data"

type TweetRepo interface {
	GetTweets() data.Tweets
	GetTweet(id int) (data.Tweet, error)
	CreateTweet(p *data.Tweet) bool
	LikeTweet(id int) bool
	DislikeTweet(id int) bool
	GetTweetsByUser(id int) (data.Tweets, error)
}
