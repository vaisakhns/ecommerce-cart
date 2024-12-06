package database

import (
	"context"
	"fmt"
	"log"
	"time"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DBset() *mongo.Client {
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil{
		log.Fatal(err)
	}

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	
     if err != nil{
         log.Fatal(err)
	 }

	 err = client.Ping(context.TODO(), nil)

	 if err != nil{ 
		log.Println("faild to connect to mongodb")
		return nil
	 }

	 fmt.Println("successfuly connected to the mongodb")
	 
	 return client
}


var Client *mongo.Client = DBset()


func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var Collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
      
      return Collection
}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection {
	
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)

	return productcollection
}