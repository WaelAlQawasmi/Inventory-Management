package controllers

import (
	"InventoryManagement/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const conectionString = "mongodb+srv://waelabuomr_db_user:rLmWeEvZLLXfJGZT@cluster0.qki2dye.mongodb.net/?appName=Cluster0"
const dbName = "inventory_management"
const colName = "items"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(conectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connected to MongoDB!")
}
func insertItem(item models.Item) {
	data, _ := collection.InsertOne(context.Background(), item)
	fmt.Println("Inserted item with ID:", data.InsertedID)

}
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	insertItem(item)
	json.NewEncoder(w).Encode(item)
}

func GetAllItems() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var items []primitive.M
	if err = cur.All(context.Background(), &items); err != nil {
		log.Fatal(err)
	}
	return items
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	items := GetAllItems()
	json.NewEncoder(w).Encode(items)
}
