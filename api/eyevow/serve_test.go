package eyevow

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func testDo(method, path string, query jmap, data interface{}, user primitive.ObjectID, v interface{}) error {
	var body io.Reader
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		body = bytes.NewReader(b)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, path, body)

	if !user.IsZero() {
		auth := newAuth(defaultEnv.Secret)
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

func testGet(path string, user primitive.ObjectID, v interface{}) error {
	return testDo("GET", path, nil, nil, user, v)
}

func testPost(path string, data interface{}, user primitive.ObjectID, v interface{}) error {
	return testDo("POST", path, nil, data, user, v)
}

func testPatch(path string, data interface{}, user primitive.ObjectID, v interface{}) error {
	return testDo("PATCH", path, nil, data, user, v)
}

func TestGetUserVow(t *testing.T) {
	defer cleanup()

	ctx := context.Background()
	ensureTestVows(ctx)
	var out vowOut
	if err := testGet("/user/vow", testUser, &out); err != nil {
		t.Fatal(err)
	}
}

func TestGetVows(t *testing.T) {
	defer cleanup()

	ctx := context.Background()
	ensureTestVows(ctx)

	var out vowsOut

	if err := testGet("/vows", primitive.ObjectID{}, &out); err != nil {
		t.Fatal(err)
	}

	if len(out.Vows) != 2 {
		t.Fatal("invalida length")
	}
}

func TestPostVow(t *testing.T) {
	defer cleanup()
	txt := "testpost"

	var out vowOut
	testPost("/vows", vowIn{
		Text: txt,
	}, testUser, &out)

	if out.Vow.Text != txt {
		t.Fatal("invalid text")
	}
}

func TestPatchArchive(t *testing.T) {
	defer cleanup()
	ctx := context.Background()
	vid := primitive.NewObjectID()
	db.Collection("vows").InsertMany(ctx, []interface{}{
		vow{
			ID:   vid,
			Text: "archiving",
			User: testUser,
		},
	})
	var out vowOut
	testPatch(fmt.Sprintf("/vows/%s/archive", vid), nil, testUser, &out)
	if !out.Vow.Archived {
		t.Fatal("invalid archived status")
	}
}

func TestPutCheer(t *testing.T) {

}

func TestDeleteCheer(t *testing.T) {

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
