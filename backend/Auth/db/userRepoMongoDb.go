package db

import (
	"12factorapp/data"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepoMongoDb struct {
	logger *log.Logger
	client *mongo.Client
}

// NewUserRepoDB Constructor
func NewUserRepoDB(log *log.Logger, client *mongo.Client) (UserRepoMongoDb, error) {
	return UserRepoMongoDb{log, client}, nil
}

func (u UserRepoMongoDb) GetUsers() data.Users {
	u.logger.Println("Getting users...")
	coll := u.client.Database("myDb").Collection("users")
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		u.logger.Println(err)
	}

	var results []*data.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		u.logger.Println(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			u.logger.Println(err)
		}
		u.logger.Printf("%s\n", output)
	}
	return results
}

func (u UserRepoMongoDb) AddUser(p *data.User) {
	u.logger.Println("Inserting user...")
	coll := u.client.Database("myDB").Collection("users")
	user, err := p.ToBson()

	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
	}

	u.logger.Printf("Inserted user with _id: %v\n", result.InsertedID)
}

func (u UserRepoMongoDb) GetUser(id int) (data.User, error) {
	//TODO implement me
	panic("implement me")
}
