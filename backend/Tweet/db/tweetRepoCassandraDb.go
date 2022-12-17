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
					(username text, tweet_id UUID, text text, 
					PRIMARY KEY (username, tweet_id)) `,
			"tweet_by_user")).Exec()
	if err != nil {
		sr.logger.Println(err)
	}

	err = sr.session.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s 
					(tweet_id UUID, username text, liked boolean, 
					PRIMARY KEY (tweet_id, username))`,
			"user_by_tweet")).Exec()
	if err != nil {
		sr.logger.Println(err)
	}
}

func (t TweetRepoCassandraDb) GetTweets() (data.Tweets, error) {
	t.logger.Println("Getting tweets...")
	scanner := t.session.Query(`SELECT username, tweet_id, text FROM tweet_by_user`).Iter().Scanner()

	var results data.Tweets
	for scanner.Next() {
		var tweet data.Tweet
		err := scanner.Scan(&tweet.Username, &tweet.TweetId, &tweet.Text)
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

func (t *TweetRepoCassandraDb) GetTweetsByUser(username string) (data.Tweets, error) {
	t.logger.Println("Getting tweets...")
	scanner := t.session.Query(`SELECT username, tweet_id, text FROM tweet_by_user WHERE username = ?`, username).Iter().Scanner()

	var results data.Tweets
	for scanner.Next() {
		var tweet data.Tweet
		err := scanner.Scan(&tweet.Username, &tweet.TweetId, &tweet.Text)
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

func (t *TweetRepoCassandraDb) GetLikes(id gocql.UUID) (int, error) {
	t.logger.Printf("Getting likes for tweet%v\n", id)
	scanner := t.session.Query(`SELECT tweet_id, username, liked FROM user_by_tweet WHERE tweet_id = ?`, id).Iter().Scanner()

	result := 0
	for scanner.Next() {
		var like data.Like
		err := scanner.Scan(&like.TweetId, &like.Username, &like.Liked)
		t.logger.Printf("tweetId: %v\n username: %v\n liked: %v\n", like.TweetId, like.Username, like.Liked)
		if err != nil {
			t.logger.Println(err)
			return -1, err
		}

		if like.Liked == true {
			result = result + 1
		}
	}
	if err := scanner.Err(); err != nil {
		t.logger.Println(err)
		return -1, err
	}
	return result, nil
}

func (t *TweetRepoCassandraDb) GetLikesUsers(id gocql.UUID) (data.LikesUsers, error) {
	t.logger.Printf("Getting likes for tweet%v\n", id)
	scanner := t.session.Query(`SELECT tweet_id, username, liked FROM user_by_tweet WHERE tweet_id = ?`, id).Iter().Scanner()

	var result []string
	for scanner.Next() {
		var like data.Like
		err := scanner.Scan(&like.TweetId, &like.Username, &like.Liked)
		t.logger.Printf("tweetId: %v\n username: %v\n liked: %v\n", like.TweetId, like.Username, like.Liked)
		if err != nil {
			t.logger.Println(err)
			return nil, err
		}

		if like.Liked == true {
			result = append(result, like.Username)
		}
	}
	if err := scanner.Err(); err != nil {
		t.logger.Println(err)
		return nil, err
	}
	return result, nil
}

func (t *TweetRepoCassandraDb) CreateTweet(p *data.Tweet) (bool, error) {
	t.logger.Println("Creating tweet...")
	p.TweetId = gocql.TimeUUID()
	err := t.session.Query(
		`INSERT INTO tweet_by_user (username, tweet_id, text) 
		VALUES (?, ?, ?)`,
		p.Username, p.TweetId, p.Text).Exec()
	if err != nil {
		t.logger.Println(err)
		return false, err
	}
	t.logger.Printf("Created tweet with id: %v\n", p.TweetId)
	return true, nil
}

func (t *TweetRepoCassandraDb) LikeTweet(tweetId gocql.UUID, username string) bool {
	t.logger.Printf("Liking fTweet with id:  %v\n", tweetId)
	p := data.Like{}
	p.TweetId = tweetId
	p.Username = username

	//Checking if like or dislike
	scanner := t.session.Query(`SELECT tweet_id, username, liked FROM user_by_tweet WHERE tweet_id = ? AND username = ?`,
		p.TweetId, p.Username).Iter().Scanner()

	var fTweet data.Like
	for scanner.Next() {
		err := scanner.Scan(&fTweet.TweetId, &fTweet.Username, &fTweet.Liked)
		if err != nil {
			t.logger.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		t.logger.Println(err)
	}

	p.Liked = !fTweet.Liked

	//Checking if like exists and modifying its value if it already exists
	err := t.session.Query(
		`UPDATE user_by_tweet
				SET liked = ?
				WHERE tweet_id = ? AND username = ?`,
		p.Liked, p.TweetId, p.Username).Exec()

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
		`INSERT INTO user_by_tweet (tweet_id, username, liked) 
		VALUES (?, ?, ?)`,
		p.TweetId, p.Username, p.Liked).Exec()

	if err != nil {
		t.logger.Println("Error when creating new like")
		t.logger.Println(err)
		return false
	}

	t.logger.Printf("Liked tweet with id: %v\n", p.TweetId)
	return true
}
