package db

import (
	"Tweet/data"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type TweetRepoMongoDb struct {
	logger *log.Logger
	client *mongo.Client
}

// NoSQL: Constructor which reads db configuration from environment
func NewTweetRepoDB(ctx context.Context, logger *log.Logger) (*TweetRepoMongoDb, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &TweetRepoMongoDb{
		client: client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (t *TweetRepoMongoDb) Disconnect(ctx context.Context) error {
	err := t.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (t *TweetRepoMongoDb) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := t.client.Ping(ctx, readpref.Primary())
	if err != nil {
		t.logger.Println(err)
	}

	// Print available databases
	databases, err := t.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		t.logger.Println(err)
	}
	fmt.Println(databases)
}

func (t *TweetRepoMongoDb) getCollection() *mongo.Collection {
	db := t.client.Database("myDB")
	collection := db.Collection("tweets")
	return collection
}

func (t TweetRepoMongoDb) GetTweets() data.Tweets {
	t.logger.Println("Getting tweets...")
	coll := t.getCollection()
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		t.logger.Println(err)
	}
	var results []*data.Tweet
	if err = cursor.All(context.TODO(), &results); err != nil {
		t.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			t.logger.Println(err)
		}
		t.logger.Printf("%s\n", output)
	}

	return results
}

func (t *TweetRepoMongoDb) GetTweet(id string) (data.Tweet, error) {
	t.logger.Printf("Getting tweet ", id)
	var result data.Tweet

	coll := t.getCollection()
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		t.logger.Println(err)
		return result, errors.New("couldn't find tweet")
	}

	return result, nil
}

func (t *TweetRepoMongoDb) CreateTweet(p *data.Tweet) (bool, error) {
	t.logger.Println("Creating tweet...")
	coll := t.getCollection()
	id := uuid.New()
	p.ID = id.String()
	tweet, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), tweet)
	if err != nil {
		t.logger.Println(err)
		return false, err
	}

	t.logger.Printf("Created tweet with _id: %v\n", result.InsertedID)
	return true, nil
}

func (t *TweetRepoMongoDb) LikeTweet(id string, username string) bool {
	tweet, err := t.GetTweet(id)
	if err != nil {
		return false
	}
	idx := -1
	for i, user := range tweet.Likes {
		if user == username {
			idx = i
			break
		}
	}
	if idx == -1 {
		tweet.Likes = append(tweet.Likes, username)
	} else {
		tweet.Likes[idx] = tweet.Likes[len(tweet.Likes) -1]
		tweet.Likes[len(tweet.Likes) -1] = ""
		tweet.Likes = tweet.Likes[:len(tweet.Likes) -1]
	}

	coll := t.getCollection()
	filter := bson.D{{"id", id}}
	newtweet, err := tweet.ToBson()
	if err != nil {
		t.logger.Println(err)
		return false
	}
	_, err = coll.ReplaceOne(context.TODO(), filter, newtweet)
	if err != nil {
		t.logger.Println(err)
		return false
	}

	return true
}

func (t *TweetRepoMongoDb) GetTweetsByUser(id int) (data.Tweets, error) {
	//TODO implement me
	panic("implement me")
}
