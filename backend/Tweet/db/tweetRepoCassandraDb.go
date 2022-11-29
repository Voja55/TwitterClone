package db

import (
	"Tweet/data"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
)

type TweetRepoCassandraDb struct {
	logger  *log.Logger
	session *gocql.Session
}

// NoSQL: Constructor which reads db configuration from environment
func NewTweetRepoDB(logger *log.Logger) (*TweetRepoCassandraDb, error) {
	dburi := os.Getenv("CASS_DB_URI")

	// Connect to default keyspace
	cluster := gocql.NewCluster(dburi)
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	// Create 'student' keyspace
	err = session.Query(
		fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s
					WITH replication = {
						'class' : 'SimpleStrategy',
						'replication_factor' : %d
					}`, "tweet", 1)).Exec()
	if err != nil {
		logger.Println(err)
	}
	session.Close()

	// Connect to student keyspace
	cluster.Keyspace = "tweet"
	cluster.Consistency = gocql.One
	session, err = cluster.CreateSession()
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	// Return repository with logger and DB session
	return &TweetRepoCassandraDb{
		session: session,
		logger:  logger,
	}, nil
}

// Disconnect from database
func (sr *TweetRepoCassandraDb) CloseSession() {
	sr.session.Close()
}

// Create tables
func (sr *TweetRepoCassandraDb) CreateTables() {
	sr.logger.Println("Creating tables...")
	err := sr.session.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s 
					(tweet_id UUID, user_id UUID, text text, 
					PRIMARY KEY (tweet_id)) `,
			"tweet_by_user")).Exec()
	if err != nil {
		sr.logger.Println(err)
	}

	err = sr.session.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s 
					(tweet_id UUID, user_id UUID, liked boolean, 
					PRIMARY KEY (tweet_id, user_id))`,
			"user_by_tweet")).Exec()
	if err != nil {
		sr.logger.Println(err)
	}
}

func (t TweetRepoCassandraDb) GetTweets() (data.Tweets, error) {
	t.logger.Println("Getting tweets...")
	scanner := t.session.Query(`SELECT tweet_id, user_id, text FROM tweet_by_user`).Iter().Scanner()

	var results data.Tweets
	for scanner.Next() {
		var tweet data.Tweet
		err := scanner.Scan(&tweet.TweetId, &tweet.UserId, &tweet.Text)
		if err != nil {
			t.logger.Println(err)
			return nil, err
		}
		results = append(results, &tweet)
	}
	if err := scanner.Err(); err != nil {
		t.logger.Println(err)
		return nil, err
	}
	return results, nil
}

func (t *TweetRepoCassandraDb) GetLikes(id gocql.UUID) (data.Likes, error) {
	t.logger.Printf("Getting likes for tweet%v\n", id)
	scanner := t.session.Query(`SELECT tweet_id, user_id, liked FROM user_by_tweet WHERE tweet_id = ?`, id).Iter().Scanner()

	var results data.Likes
	for scanner.Next() {
		var like data.Like
		err := scanner.Scan(&like.TweetId, &like.UserId, &like.Liked)
		if err != nil {
			t.logger.Println(err)
			return nil, err
		}
		results = append(results, &like)
	}
	if err := scanner.Err(); err != nil {
		t.logger.Println(err)
		return nil, err
	}
	return results, nil
}

func (t *TweetRepoCassandraDb) CreateTweet(p *data.Tweet) (bool, error) {
	t.logger.Println("Creating tweet...")
	p.TweetId = gocql.TimeUUID()
	err := t.session.Query(
		`INSERT INTO tweet_by_user (tweet_id, user_id, text) 
		VALUES (?, ?, ?)`,
		p.TweetId, p.UserId, p.Text).Exec()
	if err != nil {
		t.logger.Println(err)
		return false, err
	}
	t.logger.Printf("Created tweet with id: %v\n", p.TweetId)
	return true, nil
}

func (t *TweetRepoCassandraDb) LikeTweet(tweetId gocql.UUID, userId gocql.UUID, liked bool) bool {
	t.logger.Printf("Liking tweet with id:  %v\n", tweetId)
	p := data.Like{}
	p.TweetId = tweetId
	p.UserId = userId
	p.Liked = liked

	//Checking if like exists and modifying its value if it already exists
	err := t.session.Query(
		`UPDATE user_by_tweet
				SET liked = ?
				WHERE tweet_id = ? AND user_id = ?`,
		p.Liked, p.TweetId, p.UserId).Exec()

	if err != nil {
		t.logger.Println("Like doesnt exist")
		t.logger.Println(err)
	}

	if err == nil {
		t.logger.Printf("Liked tweet with id: %v\n", p.TweetId)
		return true
	}

	//Creating new like
	err = t.session.Query(
		`INSERT INTO user_by_tweet (tweet_id, user_id, liked) 
		VALUES (?, ?, ?)`,
		p.TweetId, p.UserId, p.Liked).Exec()

	if err != nil {
		t.logger.Println("Error when creating new like")
		t.logger.Println(err)
		return false
	}

	t.logger.Printf("Liked tweet with id: %v\n", p.TweetId)
	return true
}

func (t *TweetRepoCassandraDb) GetTweetsByUser(id int) (data.Tweets, error) {
	//TODO implement me
	panic("implement me")
}
