package eyevow

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var defaultEnv = parseEnv()

var (
	currentUserKey struct{}
)

func signup(w http.ResponseWriter, r *http.Request) {

}

func signon(w http.ResponseWriter, r *http.Request) {

}

func authorize(r *http.Request) primitive.ObjectID {
	v := r.Header.Get("Authorization")
	p := strings.Split(v, " ")
	if len(p) != 2 {
		return primitive.ObjectID{}
	}
	typ := p[0]
	if typ != "Bearer" {
		return primitive.ObjectID{}
	}
	v = p[1]
	return defaultAuth.Validate(v)
}

func currentUser(r *http.Request) primitive.ObjectID {
	return r.Context().Value(currentUserKey).(primitive.ObjectID)
}

func getMyVow(w http.ResponseWriter, r *http.Request) {
	uid := currentUser(r)
	var rs vow

	err := db.Collection("vows").FindOne(r.Context(), bson.M{"user": uid}).Decode(&rs)
	switch err {
	case mongo.ErrNoDocuments:
		failed(w, http.StatusNotFound, "not found")
	case nil:
		renderVow(r.Context(), w, &rs)
	default:
		panic(err)
	}
}

/*
	達成済の誓い一覧
*/
func getVows(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("status")
	f := bson.M{}

	switch s {
	case "archived":
		f = bson.M{"archived": true}
	case "progress":
		f = bson.M{"archived": false}
	}

	cursor, err := db.Collection("vows").Find(r.Context(), f)
	if err != nil {
		panic(err)
	}

	var rs []*vow
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

	vow := &vow{
		Text: in.Text,
		User: u,
	}

	db.Collection("vows").InsertOne(r.Context(), vow)
	renderVow(r.Context(), w, vow)
}

/*
	誓いを達成する
*/
func patchArchive(w http.ResponseWriter, r *http.Request) {
	u := currentUser(r)
	var rs vow
	if err := db.Collection("vows").FindOne(r.Context(), bson.M{"user": u}).Decode(&rs); err != nil {
		panic(err)
	}
	if rs.User != u {
		failed(w, http.StatusForbidden, "permission denied")
		return
	}

	rs.Archived = true

	if _, err := db.Collection("vows").ReplaceOne(r.Context(), bson.M{"_id": rs.ID}, rs); err != nil {
		panic(err)
	}

	renderVow(r.Context(), w, &rs)
}

func putCheer(w http.ResponseWriter, r *http.Request) {
	up := chi.URLParam(r, "vow")
	u := currentUser(r)

	vid, err := primitive.ObjectIDFromHex(up)
	if err != nil {
		failed(w, http.StatusBadRequest, "")
		return
	}

	if _, err = db.Collection("cheers").InsertOne(r.Context(), bson.M{
		"user": u,
		"vow":  vid,
	}); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteCheer(w http.ResponseWriter, r *http.Request) {
	up := chi.URLParam(r, "vow")
	u := currentUser(r)

	vid, err := primitive.ObjectIDFromHex(up)
	if err != nil {
		failed(w, http.StatusBadRequest, "")
		return
	}

	if _, err := db.Collection("cheers").DeleteOne(r.Context(), bson.M{
		"user": u,
		"vow":  vid,
	}); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func authorized(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			u := authorize(r)

			if u.IsZero() {
				failed(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), currentUserKey, u))
			h.ServeHTTP(w, r)
		},
	)
}

func server() http.Handler {
	router := chi.NewRouter()
	router.Post("/signup", signup)
	router.Post("/signon", signon)
	router.Get("/vows", getVows)
	router.Group(func(r chi.Router) {
		r.Use(authorized)

		r.Get("/user/vow", getMyVow)
		r.Patch("/vows/{vow}/archive", patchArchive)
		r.Put("/vows/{vow}/cheer", putCheer)
		r.Delete("/vows/{vow}/cheer", deleteCheer)
		r.Post("/vows", postVow)
	})

	return router
}

func serve() {
	log.Fatal(http.ListenAndServe(":1256", server()))
}
