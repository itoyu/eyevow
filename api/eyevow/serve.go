package eyevow

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type jmap map[string]interface{}

func render(w http.ResponseWriter, status int, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Write(b)
}

func success(w http.ResponseWriter, v interface{}) {
	render(w, http.StatusOK, v)
}

func failed(w http.ResponseWriter, status int, m string) {
	render(w, status, jmap{
		"code":    status,
		"message": m,
	})
}

func buildUser(doc bson.M) jmap {
	return jmap{
		"id": doc["_id"],
	}
}

func buildVow(ctx context.Context, doc bson.M) jmap {
	rs := db.Collection("users").FindOne(ctx, bson.M{"_id": doc["user"]})
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

func buildVows(ctx context.Context, docs []bson.M) []jmap {
	ss := make([]jmap, len(docs))
	for i, d := range docs {
		ss[i] = buildVow(ctx, d)
	}
	return ss
}

func renderVow(ctx context.Context, w http.ResponseWriter, doc bson.M) {
	success(w, jmap{"vow": buildVow(ctx, doc)})
}

func renderVows(ctx context.Context, w http.ResponseWriter, docs []bson.M) {
	success(w, jmap{"vows": buildVows(ctx, docs)})
}

func bind(r *http.Request, v interface{}) error {
	var b []byte
	var err error
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}

func signup(w http.ResponseWriter, r *http.Request) {

}

func signon(w http.ResponseWriter, r *http.Request) {

}

func authorize(r *http.Request) string {
	return ""
}

var (
	currentUserKey struct{}
)

func currentUser(r *http.Request) primitive.ObjectID {
	return r.Context().Value(currentUserKey).(primitive.ObjectID)
}

func getMyVow(w http.ResponseWriter, r *http.Request) {
	uid := currentUser(r)
	var rs bson.M

	err := db.Collection("vows").FindOne(r.Context(), bson.M{"user": uid}).Decode(&rs)
	switch err {
	case mongo.ErrNoDocuments:
		failed(w, http.StatusNotFound, "not found")
	case nil:
		renderVow(r.Context(), w, rs)
	default:
		panic(err)
	}
}

func getVows(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.Collection("vows").Find(r.Context(), bson.D{})
	if err != nil {
		panic(err)
	}

	var rs []bson.M
	if err = cursor.All(r.Context(), &rs); err != nil {
		panic(err)
	}

	renderVows(r.Context(), w, rs)
}

func postVow(w http.ResponseWriter, r *http.Request) {
	u := currentUser(r)

	in := struct {
		Text string `json:"text"`
	}{}

	if err := bind(r, &in); err != nil {
		failed(w, http.StatusBadRequest, "パラメータエラー")
		return
	}

	vow := bson.M{
		"Text": in.Text,
		"User": u,
	}

	db.Collection("vows").InsertOne(r.Context(), vow)
	renderVow(r.Context(), w, vow)
}

func patchArchive(w http.ResponseWriter, r *http.Request) {

}

func authorized(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			u := authorize(r)

			if u == "" {
				failed(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			id, err := primitive.ObjectIDFromHex(u)
			if err != nil {
				panic(err)
			}

			r = r.WithContext(context.WithValue(r.Context(), currentUserKey, id))
			h.ServeHTTP(w, r)
		},
	)
}

func serve() {
	router := chi.NewRouter()
	router.Post("/signup", signup)
	router.Post("/signon", signon)
	router.Get("/vows", getVows)
	router.Group(func(r chi.Router) {
		r.Use(authorized)

		r.Get("/user/vow", getMyVow)
		r.Patch("/vows/{vow}/archive", patchArchive)
		r.Post("/vows", postVow)
	})

	log.Fatal(http.ListenAndServe(":1256", router))
}
