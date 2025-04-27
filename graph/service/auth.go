package service

import (
	"context"
	"go-graphql-blog/graph/database"
	"go-graphql-blog/graph/model"
	"go-graphql-blog/graph/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

const USER_COLLECTION = "users"

func (u *UserService) Register(input model.NewUser) string {
	bs, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	var password string = string(bs)

	var user model.User = model.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(USER_COLLECTION)

	res, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return ""
	}

	var userId string = res.InsertedID.(primitive.ObjectID).Hex()

	token, err:= utils.GenerateNewAccessToke(userId)

	if err != nil {
		return ""
	}
	return token
}
