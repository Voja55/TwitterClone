package db

import (
	"12factorapp/data"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
	"math/rand"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoMongoDb struct {
	logger *log.Logger
	client *mongo.Client
}

// NoSQL: Constructor which reads db configuration from environment
func NewUserRepoDB(ctx context.Context, logger *log.Logger) (*UserRepoMongoDb, error) {
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

	return &UserRepoMongoDb{
		client: client,
		logger: logger,
	}, nil
}

func (u *UserRepoMongoDb) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Disconnect from database
func (u *UserRepoMongoDb) Disconnect(ctx context.Context) error {
	err := u.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (u *UserRepoMongoDb) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := u.client.Ping(ctx, readpref.Primary())
	if err != nil {
		u.logger.Println(err)
	}

	// Print available databases
	databases, err := u.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		u.logger.Println(err)
	}
	fmt.Println(databases)
}

func (u UserRepoMongoDb) GetUsers() data.Users {
	u.logger.Println("Getting users...")
	coll := u.getCollection()
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
	coll := u.getCollection()
	id := uuid.New()
	p.ID = id.String()
	hashedPass, err := p.HashPassword(p.Password)
	p.Password = hashedPass
	rand.Seed(time.Now().UnixNano())
	cCode := rand.Intn(999999-100001) + 100000
	p.CCode = cCode

	user, err := p.ToBson()
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		u.logger.Println(err)
		return false
	}

	u.logger.Printf("Registered user with _id: %v\n", result.InsertedID)
	return true
}

func (u UserRepoMongoDb) GetUser(id string) (data.User, error) {
	u.logger.Printf("Getting user ", id)
	var result data.User

	coll := u.getCollection()
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

	coll := u.getCollection()
	//Finding user
	filter := bson.D{{"username", username}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	//Checking password
	resultBool := u.CheckPasswordHash(password, result.Password)
	if resultBool == false {
		return result, errors.New("wrong username or password")
	}

	if err != nil {
		u.logger.Println(err)
		return result, errors.New("wrong username or password")
	}

	return result, nil
}

func (u UserRepoMongoDb) GetUserByUsername(username string) (data.User, error) {
	u.logger.Printf("Getting user ", username)
	var result data.User

	coll := u.getCollection()
	filter := bson.D{{"username", username}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		u.logger.Println(err)
		return result, errors.New("couldn't find user")
	}

	return result, nil
}

func (u *UserRepoMongoDb) UpdateUser(user *data.User) (bool) {
	coll := u.getCollection()
	filter := bson.D{{"id", user.ID}}
	
	_, err := coll.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		u.logger.Println(err)
		return false
	}
	return true
	
}  

func (u *UserRepoMongoDb) getCollection() *mongo.Collection {
	db := u.client.Database("myDB")
	collection := db.Collection("users")
	return collection
}
