package main

import (
	"context"
	"fmt"
	"log" // os.Exit(1) on Error

	// Official 'mongo-go-driver' packages
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type MongoFields struct { 
	FieldStr string `json:"Field Str"` 
	FieldInt int `json:"Field Int"` 
	FieldBool bool `json:"Field Bool"` 
	} 
var collection *mongo.Collection
var ctx = context.TODO()
 

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Api Response")
    fmt.Println("Endpoint Hit: homePage")
}
func logit(w http.ResponseWriter, r *http.Request){

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("YOURMONGOURL"))
	 
	
	// Connect to the MongoDB and return Client instance 
	if err != nil {
	    fmt.Println("Connection Error")
}

fmt.Println(client);

usersCollection := client.Database("AppLogs").Collection("MicroDB")
now := time.Now()      // current local time
sec := now.Unix()
user := bson.D{{"fullName", "User 2"}, {"age", 3} , {"createtime",sec}}
// insert the bson object using InsertOne()
result, err := usersCollection.InsertOne(context.TODO(), user)
// check for errors in the insertion
if err != nil {
	fmt.Println("Insert Error"+ err.Error())
}
// display the id of the newly inserted object
fmt.Println(result.InsertedID)

// inse
	
	// Declare a MongoDB struct instance for the document's fields and data 
 
	 
	
	  
}

func handleRequests() {
    http.HandleFunc("/", homePage)
	http.HandleFunc("/logit", logit)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("API is Online")
    handleRequests()
}


 