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

func (t *TweetRepoCassandraDb) LikeTweet(id string, username string) bool {
	//tweet, err := t.GetTweet(id)
	//if err != nil {
	//	return false
	//}
	//idx := -1
	//for i, user := range tweet.Likes {
	//	if user == username {
	//		idx = i
	//		break
	//	}
	//}
	//if idx == -1 {
	//	tweet.Likes = append(tweet.Likes, username)
	//} else {
	//	tweet.Likes[idx] = tweet.Likes[len(tweet.Likes)-1]
	//	tweet.Likes[len(tweet.Likes)-1] = ""
	//	tweet.Likes = tweet.Likes[:len(tweet.Likes)-1]
	//}
	//
	//coll := t.getCollection()
	//filter := bson.D{{"id", id}}
	//newtweet, err := tweet.ToBson()
	//if err != nil {
	//	t.logger.Println(err)
	//	return false
	//}
	//_, err = coll.ReplaceOne(context.TODO(), filter, newtweet)
	//if err != nil {
	//	t.logger.Println(err)
	//	return false
	//}
	panic("implement me")
	return true
}

func (t *TweetRepoCassandraDb) GetTweetsByUser(id int) (data.Tweets, error) {
	//TODO implement me
	panic("implement me")
}
