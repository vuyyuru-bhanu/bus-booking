package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func main() {
    ctx := context.TODO()
    var err error
    clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    collection = client.Database("busbooking").Collection("buses")

    r := mux.NewRouter()
    r.HandleFunc("/api/buses", GetBuses).Methods("GET")
    r.HandleFunc("/api/book", BookBus).Methods("POST")

    http.ListenAndServe(":8080", r)
}

func GetBuses(w http.ResponseWriter, r *http.Request) {
    ctx := context.TODO()
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var buses []bson.M
    if err = cursor.All(ctx, &buses); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(buses)
}

func BookBus(w http.ResponseWriter, r *http.Request) {
    var bookingDetails map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&bookingDetails); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Println("Booking Details:", bookingDetails)
    w.WriteHeader(http.StatusOK)
}

