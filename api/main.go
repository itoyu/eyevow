package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

type jmap map[string]interface{}

func result(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func buildUser(doc bson.M) jmap {
	return jmap{
		"id": doc["_id"],
	}
}

func buildVow(doc bson.M) jmap {
	rs := db.Collection("users").FindOne(context.Background(), bson.M{"_id": doc["user"]})
	if rs.Err() != nil {
		panic(rs.Err())
	}
	ud := bson.M{}
	if err := rs.Decode(&ud); err != nil {
		panic(err)
	}

	return jmap{
		"id":          doc["_id"],
		"text":        doc["text"],
		"cheer_count": doc["cheer_count"],
		"archived":    doc["archived"],
		"user":        buildUser(ud),
	}
}

func buildVows(docs []bson.M) []jmap {
	ss := make([]jmap, len(docs))
	for i, d := range docs {
		ss[i] = buildVow(d)
	}
	return ss
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func signout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func vows(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cursor, err := db.Collection("vows").Find(r.Context(), bson.D{})
	if err != nil {
		panic(err)
	}

	var rs []bson.M
	if err := cursor.All(r.Context(), &rs); err != nil {
		panic(err)
	}

	result(w, map[string]interface{}{
		"vows": buildVows(rs),
	})
}

func serve() {
	router := httprouter.New()
	router.POST("/signup", signup)
	router.POST("/signout", signout)
	router.GET("/vows/", vows)
	log.Fatal(http.ListenAndServe(":1256", router))
}

var dummyUsers = []interface{}{
	bson.M{
		"_id": "e_taro",
	},
	bson.M{
		"_id": "y_taro",
	},
	bson.M{
		"_id": "v_taro",
	},
	bson.M{
		"_id": "o_taro",
	},
	bson.M{
		"_id": "w_taro",
	},
}

var dummyVows = []interface{}{
	bson.M{
		"user":        "e_taro",
		"text":        "01 I going to make a go of this project!",
		"cheer_count": 10,
		"archived":    true,
	},
	bson.M{
		"id":          2,
		"user":        "y_taro",
		"text":        "02 I going to make a go of this project!",
		"cheer_count": 20,
		"archived":    false,
	},
	bson.M{
		"id":          3,
		"user":        "e_taro",
		"text":        "03 I going to make a go of this project!",
		"cheer_count": 30,
		"archived":    true,
	},
	bson.M{
		"id":          4,
		"user":        "v_taro",
		"text":        "04 I going to make a go of this project!",
		"cheer_count": 0,
		"archived":    true,
	},
	bson.M{
		"id":          5,
		"user":        "o_taro",
		"text":        "05 I going to make a go of this project!",
		"cheer_count": 0,
		"archived":    false,
	},
	bson.M{
		"id":          6,
		"user":        "w_taro",
		"text":        "06 I going to make a go of this project!",
		"cheer_count": 60,
		"archived":    false,
	},
}

func seed() {
	ctx := context.Background()
	db.Collection("users").DeleteMany(ctx, bson.D{})
	db.Collection("vows").DeleteMany(ctx, bson.D{})

	db.Collection("users").InsertMany(ctx, dummyUsers)
	db.Collection("vows").InsertMany(ctx, dummyVows)
}

func ensureDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		panic(err)
	}

	return client
}

func main() {
	mongo := ensureDB()
	defer mongo.Disconnect(context.Background())
	db = mongo.Database("eyevow")

	cmd := "server"

	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "server":
		serve()
	case "seed":
		seed()
	}
}
