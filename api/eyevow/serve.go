package eyevow

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func renderVows(w http.ResponseWriter, docs []bson.M) {
	result(w, jmap{"vows": buildVows(docs)})
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func signon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func authorize(r *http.Request) string {
	return ""
}

func myVows(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var uid string
	if uid = authorize(r); uid == "" {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	var cs *mongo.Cursor
	var rs []bson.M
	var err error

	if cs, err = db.Collection("vows").Find(r.Context(), bson.M{"user": uid}); err != nil {
		panic(err)
	}
	if err = cs.All(r.Context(), &rs); err != nil {
		panic(err)
	}

	renderVows(w, rs)
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

	renderVows(w, rs)
}

func serve() {
	router := httprouter.New()
	router.POST("/signup", signup)
	router.POST("/signon", signon)
	router.GET("/vows/", vows)
	router.GET("/vows/:")
	router.GET("/self/vows", myVows)
	log.Fatal(http.ListenAndServe(":1256", router))
}
