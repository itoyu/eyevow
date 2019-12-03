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

type user struct {
	ID primitive.ObjectID `bson:"_id"`
}

type vow struct {
	ID         primitive.ObjectID `bson:"_id"`
	Text       string             `bson:"text"`
	User       primitive.ObjectID `bson:"user"`
	CheerCount int                `bson:"cheer_count"`
	Archived   bool               `bson:"archived"`
}

type userData struct {
	ID string `json:"id"`
}

type vowData struct {
	ID         string    `json:"id"`
	Text       string    `json:"text"`
	User       *userData `json:"user"`
	CheerCount int       `json:"cheer_count"`
	Archived   bool      `json:"archived"`
}

type vowOut struct {
	Vow *vowData `json:"vow"`
}

type vowsOut struct {
	Vows []*vowData `json:"vows"`
}

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

func buildUser(u *user) *userData {
	return &userData{
		ID: u.ID.String(),
	}
}

func buildVow(ctx context.Context, vow *vow) *vowData {
	rs := db.Collection("users").FindOne(ctx, bson.M{"_id": vow.User})
	if rs.Err() != nil {
		panic(rs.Err())
	}
	var ud user

	if err := rs.Decode(&ud); err != nil {
		panic(err)
	}

	return &vowData{
		ID:         vow.ID.String(),
		Text:       vow.Text,
		CheerCount: vow.CheerCount,
		Archived:   vow.Archived,
		User:       buildUser(&ud),
	}
}

func buildVows(ctx context.Context, vows []*vow) []*vowData {
	ss := make([]*vowData, len(vows))
	for i, d := range vows {
		ss[i] = buildVow(ctx, d)
	}
	return ss
}

func renderVow(ctx context.Context, w http.ResponseWriter, vow *vow) {
	success(w, vowOut{Vow: buildVow(ctx, vow)})
}

func renderVows(ctx context.Context, w http.ResponseWriter, vows []*vow) {
	success(w, vowsOut{Vows: buildVows(ctx, vows)})
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

type auth struct{}

func newAuth() *auth {
	return &auth{}
}

func (a *auth) Publish(uid primitive.ObjectID) string {
	return ""
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
	var rs *vow

	err := db.Collection("vows").FindOne(r.Context(), bson.M{"user": uid}).Decode(rs)
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

func server() http.Handler {
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

	return router
}

func serve() {
	log.Fatal(http.ListenAndServe(":1256", server()))
}
