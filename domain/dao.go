package domain

import (
	"context"
	"encoding/json"
	"fiberscurd/utils"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (user *User) Create() (*mongo.InsertOneResult, *utils.Resterr) {
	usersC := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	emailCount, _ := usersC.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if emailCount > 0 {
		return nil, utils.BadRequest("Email Already Register")
	}

	value, err := json.Marshal(User{ID: user.ID, Name: user.Name, Email: user.Email, Password: user.Password})
	if err != nil {
		return nil, utils.InternalErr("JSON ERROR")
	}
	err = redisdb.Set(string(user.Email), value, 0).Err()
	if err != nil {
		return nil, utils.InternalErr("Can't the Set the Values")
	}

	result, err := usersC.InsertOne(ctx, user)

	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}

	return result, nil
}

func (user *User) FindUser() *utils.Resterr {

	value, rediserr := redisdb.Get(string(user.Email)).Result()
	text := string(value)
	bytes := []byte(text)
	json.Unmarshal(bytes, &user)

	if rediserr != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		userC := db.Collection("users")
		filter := bson.M{"email": user.Email}
		err := userC.FindOne(ctx, filter).Decode(&user)

		defer cancel()

		if err != nil {
			return utils.NotFound("Email is Not Found")
		}
	}

	return nil

}

func (user *User) Delete() (*mongo.DeleteResult, *utils.Resterr) {

	redisdb.Del(string(user.Email))

	userC := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"email": user.Email}
	result, err := userC.DeleteOne(ctx, filter)
	defer cancel()
	if result.DeletedCount == 0 {
		return nil, utils.BadRequest("No Record Found")
	}

	if err != nil {
		return nil, utils.NotFound("Email is Not Found")
	}

	return result, nil

}

func (user *User) Update() (*mongo.UpdateResult, *utils.Resterr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	userC := db.Collection("users")

	filter := bson.M{"email": user.Email}

	updateValue := bson.M{"$set": bson.M{"name": user.Name}}

	opts := options.Update().SetUpsert(true)

	result, err := userC.UpdateOne(ctx, filter, updateValue, opts)

	defer cancel()

	if result.ModifiedCount == 0 {
		return nil, utils.BadRequest("not modified")
	}

	if err != nil {
		return nil, utils.InternalErr("Data not Updated")
	}
	redisdb.Del(string(user.Email))

	reerr := userC.FindOne(ctx, filter).Decode(&user)
	if reerr != nil {
		return nil, utils.NotFound("Email is Not Found")
	}

	value, err := json.Marshal(User{ID: user.ID, Name: user.Name, Email: user.Email, Password: user.Password})
	if err != nil {
		return nil, utils.InternalErr("JSON ERROR")
	}
	err = redisdb.Set(string(user.Email), value, 0).Err()
	if err != nil {
		return nil, utils.InternalErr("Can't the Set the Values")
	}

	return result, nil

}
