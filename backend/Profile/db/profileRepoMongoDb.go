package db

import (
	"Profile/data"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type ProfileRepoMongoDb struct {
	logger *log.Logger
	client *mongo.Client
}

// NoSQL: Constructor which reads db configuration from environment
func NewProfileRepoDB(ctx context.Context, logger *log.Logger) (*ProfileRepoMongoDb, error) {
	uri := os.Getenv("MONGO_DB_URI")
	logger.Println(uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &ProfileRepoMongoDb{
		client: client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (p *ProfileRepoMongoDb) Disconnect(ctx context.Context) error {
	err := p.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (p *ProfileRepoMongoDb) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := p.client.Ping(ctx, readpref.Primary())
	if err != nil {
		p.logger.Println(err)
	}

	// Print available databases
	databases, err := p.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		p.logger.Println(err)
	}
	fmt.Println(databases)
}

func (p ProfileRepoMongoDb) CreateProfile(dp *data.Profile) bool {
	p.logger.Println("Creating profile...")
	coll := p.getCollection()

	profile, err := dp.ToBson()
	result, err := coll.InsertOne(context.TODO(), profile)
	if err != nil {
		p.logger.Println(err)
		return false
	}

	p.logger.Printf("Registered user : %v\n", result.InsertedID)
	return true
}

func (p ProfileRepoMongoDb) GetProfile(username string) (data.Profile, error) {
	p.logger.Printf("Getting user ", username)
	var result data.Profile

	coll := p.getCollection()
	filter := bson.D{{"username", username}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		p.logger.Println(err)
		return result, errors.New("couldn't find profile")
	}

	return result, nil
}

func (p *ProfileRepoMongoDb) getCollection() *mongo.Collection {
	db := p.client.Database("myDB")
	collection := db.Collection("profiles")
	return collection
}
