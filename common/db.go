package common

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)







func ConnectDb () *mongo.Client {
	var db_pass string = viper.Get("MONGO_PASS").(string)
	var db_user string = viper.Get("MONGO_USER").(string)
	var uri string = fmt.Sprintf("mongodb://%s:%s@localhost:27017", db_user, db_pass)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
  fmt.Println("Connected to Mongo........")
  return client
}