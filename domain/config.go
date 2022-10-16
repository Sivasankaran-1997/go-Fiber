package domain

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db      *mongo.Database
	redisdb *redis.Client
)

func ConnDB() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("MONGO_DB")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db = client.Database("fiber_users")
	fmt.Println("Successfuly Mongodb connected to the database.")
}

func RedisConn() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("REDIS")

	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: "",
		DB:       0,
	})

	redisdb = client
	fmt.Println("Successfully Redis Connected to the database")

}
