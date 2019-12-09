package eyevow

import (
	"bytes"
	"context"
	"encoding/base64"
	"eyevow/infra/bucket"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	"github.com/teris-io/shortid"
	"github.com/vincent-petithory/dataurl"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userIconRadius = 100
)

var defaultEnv = parseEnv()
var defaultBucket bucket.Bucket
var defaultAssets http.FileSystem

var (
	currentUserKey struct{}
)

/*
ユーザー登録(デバッグ)
*/
func signup(w http.ResponseWriter, r *http.Request) {
	var rs *mongo.InsertOneResult
	var err error
	var f http.File

	if f, err = defaultAssets.Open("/user.jpg"); err != nil {
		panic(err)
	}

	defer f.Close()

	sid := shortid.MustGenerate()
	icon := fmt.Sprintf("user/%s.jpg", sid)

	if err = defaultBucket.Put(icon, f); err != nil {
		panic(err)
	}

	u := &user{
		Name: "ゲスト",
		Icon: defaultBucket.URL(icon),
	}

	if rs, err = db.Collection("users").InsertOne(r.Context(), u); err != nil {
		panic(err)
	}
	u.ID = rs.InsertedID.(primitive.ObjectID)
	tk := defaultAuth.Publish(u.ID)

	renderToken(r.Context(), w, tk, u)
}

/*
シングルサインオン
*/
func signon(w http.ResponseWriter, r *http.Request) {

}

/*
自分情報取得
*/
func getSelf(w http.ResponseWriter, r *http.Request) {
	uid := currentUser(r)
	var u user
	if err := db.Collection("users").FindOne(r.Context(), bson.M{"_id": uid}).Decode(&u); err != nil {
		panic(err)
	}
	renderUser(r.Context(), w, &u)
}

/*
自分情報編集
*/
func patchUser(w http.ResponseWriter, r *http.Request) {
	uid := currentUser(r)
	var u user
	if err := db.Collection("users").FindOne(r.Context(), bson.M{"_id": uid}).Decode(&u); err != nil {
		panic(err)
	}
	var in userIn
	if err := bind(r, &in); err != nil {
		failed(w, http.StatusBadRequest, "parameter error")
	}

	if in.Name != "" {
		u.Name = in.Name
	}

	if in.Icon != "" {
		du, err := dataurl.DecodeString(in.Icon)
		if err != nil {
			failed(w, http.StatusBadRequest, "parameter error")
		}
		im, _, err := image.Decode(bytes.NewReader(du.Data))
		if err != nil {
			failed(w, http.StatusBadRequest, "invalid image format")
		}
		art := image.NewRGBA(image.Rect(0, 0, 200, 200))
		draw.Draw(art, art.Bounds(), image.NewUniform(color.White), image.ZP, draw.Over)
		im = imaging.Fill(im, userIconRadius*2, userIconRadius*2, imaging.Center, imaging.Lanczos)
		draw.Draw(
			art,
			im.Bounds().Add(
				art.Bounds().Size().Sub(im.Bounds().Size()).Div(2),
			),
			im,
			image.ZP,
			draw.Over,
		)
		bf := &bytes.Buffer{}
		if err := jpeg.Encode(bf, art, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			panic(err)
		}
		sid := shortid.MustGenerate()
		key := fmt.Sprintf("user/%s.jpg", sid)
		if err := bucket.PutAll(defaultBucket, key, bf.Bytes()); err != nil {
			panic(err)
		}
		u.Icon = defaultBucket.URL(key)
	}

	_, err := db.Collection("users").ReplaceOne(r.Context(), bson.M{"_id": u.ID}, &u)
	if err != nil {
		panic(err)
	}

	renderUser(r.Context(), w, &u)
}

/*
自分の誓い確認
*/
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
	誓い一覧
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

/*
誓い取得
*/
func getVow(w http.ResponseWriter, r *http.Request) {
	s := chi.URLParam(r, "vow")
	vid, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		failed(w, http.StatusBadRequest, "invalid url parameter")
	}

	var v vow
	if err := db.Collection("vows").FindOne(r.Context(), bson.M{"_id": vid}).Decode(&v); err != nil {
		panic(err)
	}
	renderVow(r.Context(), w, &v)
}

/*
誓いを作成
*/
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
		Text:     in.Text,
		User:     u,
		Archived: false,
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

/*
応援する
*/
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

/*
応援を消す
*/
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

func currentUser(r *http.Request) primitive.ObjectID {
	return r.Context().Value(currentUserKey).(primitive.ObjectID)
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

func mux() http.Handler {
	router := chi.NewRouter()
	router.Get("/", spec)
	router.Post("/signup", signup)
	router.Post("/signon", signon)
	router.Get("/vows", getVows)
	router.Get("/vows/{vow}", getVow)
	router.Group(func(r chi.Router) {
		r.Use(authorized)
		r.Get("/user", getSelf)
		r.Patch("/user", patchUser)
		r.Get("/user/vow", getMyVow)
		r.Patch("/vows/{vow}/archive", patchArchive)
		r.Put("/vows/{vow}/cheer", putCheer)
		r.Delete("/vows/{vow}/cheer", deleteCheer)
		r.Post("/vows", postVow)
	})

	return router
}

func spec(w http.ResponseWriter, r *http.Request) {
	f, err := defaultAssets.Open("/api.yml")
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	d := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(d, b)

	html := fmt.Sprintf(`<!DOCTYPE html>
		<html>
			<head>
				<title>ReDoc</title>
				<!-- needed for adaptive design -->
				<meta charset="utf-8"/>
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
				<!--
				ReDoc doesn't change outer page styles
				-->
				<style>
					body {
						margin: 0;
						padding: 0;
					}
				</style>
			</head>
			<body>
				<redoc spec-url='data:text/yaml;base64,%s'></redoc>
				<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
			</body>
		</html>
	`, d)
	w.Write([]byte(html))
}

func serve() {
	api := mux()
	m := http.NewServeMux()
	m.Handle("/api/", http.StripPrefix("/api", api))
	m.Handle("/blob/", http.StripPrefix("/blob/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd, err := defaultBucket.Get(r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer rd.Close()
		if _, err := io.Copy(w, rd); err != nil {
			panic(err)
		}
	})))

	log.Fatal(http.ListenAndServe(":1256", m))
}

func init() {
	u := defaultEnv.URL
	var err error

	if u, err = u.Parse("/blob/"); err != nil {
		panic(err)
	}

	defaultBucket = bucket.NewLocal(defaultEnv.Bucket, u.String())
	if defaultAssets, err = fs.New(); err != nil {
		panic(err)
	}
}
