package db

import (
	"12factorapp/data"
	"context"
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	coll := u.client.Database("myDB").Collection("users")
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

func (u UserRepoMongoDb) Register(p *data.User) bool {
	u.logger.Println("Registering user...")
	coll := u.client.Database("myDB").Collection("users")
	user, err := p.ToBson()

	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Registered user with _id: %v\n", result.InsertedID)
	return true
}

func (u UserRepoMongoDb) GetUser(id int) (data.User, error) {
	u.logger.Printf("Getting user ", id)
	var result data.User

	coll := u.client.Database("myDB").Collection("users")
	filter := bson.D{{"id", id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		u.logger.Println(err)
		return result, errors.New("couldn't find user")
	}

	return result, nil
}

func (u UserRepoMongoDb) LoginUser(username string, password string) (data.User, error) {
	u.logger.Printf("Checking user...")
	var result data.User

	//u bazi ludilo mozga, treba zajednicka kolekcija da se napravi za sve tipove radi auth
	coll := u.client.Database("myDB").Collection("users")
	filter := bson.D{{"username", username}, {"password", password}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		u.logger.Println(err)
		return result, errors.New("wrong username or password")
	}

	return result, nil
}
