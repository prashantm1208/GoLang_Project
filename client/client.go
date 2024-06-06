package client

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerDetails struct {
	ID       int     `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Position string  `json:"position" bson:"position"`
	Salary   float64 `json:"salary" bson:"salary"`
}

var mu sync.RWMutex

func InsertCustomer(customer CustomerDetails) (err error) {
	mu.Lock()
	defer mu.Unlock()
	session := Connect()

	collection := session.Database("Customer").Collection("details")

	result, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		log.Fatal("Error while inserting record: ", err)
		return err
	}

	log.Println(result)
	return nil
}

func GetCustomerByID(customid int) (CustomerDetails, error) {
	mu.Lock()
	defer mu.Unlock()
	session := Connect()

	collection := session.Database("Customer").Collection("details")

	// Convert th ID string into the ObjectID
	objID, err := primitive.ObjectIDFromHex(string(customid))

	if err != nil {
		return CustomerDetails{}, err
	}

	filter := bson.M{"_id": objID}

	cur := collection.FindOne(context.TODO(), filter)
	var details CustomerDetails

	err2 := cur.Decode(&details)
	if err2 != nil {
		log.Fatal("error occured while getting all records: ", err2)
	}

	return details, nil
}

func UpdateCustomer(emp CustomerDetails) (err error) {
	mu.Lock()
	defer mu.Unlock()
	session := Connect()

	collection := session.Database("Customer").Collection("details")

	result, err := collection.UpdateOne(context.TODO(), bson.M{"id": emp.ID}, bson.M{"$set": emp})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func DeleteCustomerByID(customid int) (err error) {
	mu.Lock()
	defer mu.Unlock()
	session := Connect()

	collection := session.Database("Customer").Collection("details")

	// Convert th ID string into the ObjectID
	objID, err := primitive.ObjectIDFromHex(string(customid))

	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		log.Println("Not able to delete the record: ", result)
	} else {
		log.Println("Successfully deleted...", result)
	}
	return nil
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("mongoDB connection error: ", err)
	}

	fmt.Println("DB connected")

	return client
}

// func Connect() *mongo.Client {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal("MongoDB connection error: ", err)
// 	}

// 	fmt.Println("DB connected")
// 	return client
// }
