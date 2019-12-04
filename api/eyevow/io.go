package eyevow

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type vowIn struct {
	Text string `json:"text"`
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
