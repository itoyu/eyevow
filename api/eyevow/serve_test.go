package eyevow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testUser = primitive.NewObjectID()
	sv       = server()
)

func testGet(v interface{}, path string, user primitive.ObjectID) error {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", path, nil)

	if !user.IsZero() {
		auth := newAuth()
		tk := auth.Publish(user)
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tk))
	}

	sv.ServeHTTP(w, r)
	switch w.Result().StatusCode {
	case http.StatusOK:
		json.Unmarshal(w.Body.Bytes(), v)
		return nil
	case http.StatusNoContent:
		return nil
	default:
		var e jmap
		var m interface{}
		var ok bool
		json.Unmarshal(w.Body.Bytes(), &e)
		if m, ok = e["message"]; !ok {
			return errors.New("unknown error")
		}
		return errors.New(m.(string))
	}
}

func TestGetVow(t *testing.T) {
	defer cleanup()

	ctx := context.Background()
	ensureTestVows(ctx)
	var out vowOut
	if err := testGet(&out, "/vow", testUser); err != nil {
		t.Fatal(err)
	}
}

func TestGetVows(t *testing.T) {
	defer cleanup()

	ctx := context.Background()
	ensureTestVows(ctx)

	var out vowsOut

	if err := testGet(&out, "/vows", primitive.ObjectID{}); err != nil {
		t.Fatal(err)
	}
}

func ensureTestUser(ctx context.Context) {
	db.Collection("users").InsertOne(ctx, bson.M{
		"_id": testUser,
	})
}

func ensureTestVows(ctx context.Context) {
	u1 := primitive.NewObjectID()
	db.Collection("users").InsertOne(ctx, bson.M{
		"_id": u1,
	})

	db.Collection("vows").InsertMany(ctx, []interface{}{
		vow{
			Text: "test1",
			User: u1,
		},
		vow{
			Text: "current vow",
			User: testUser,
		},
	})
}

func cleanup() {
	ctx := context.Background()
	db.Collection("users").DeleteMany(ctx, bson.D{})
	db.Collection("vows").DeleteMany(ctx, bson.D{})
	ensureTestUser(ctx)
}

func TestMain(m *testing.M) {
	mongo := ensureDB()
	defer mongo.Disconnect(context.Background())
	db = mongo.Database("eyevow")

	cleanup()
	m.Run()
}
