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
	"net/url"
	"testing"

	_ "eyevow/statik"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testUser = primitive.NewObjectID()
	sv       = mux()
)

func testDo(method, path string, query url.Values, data interface{}, user primitive.ObjectID, v interface{}) error {
	var body io.Reader

	if query != nil {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

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

func testGet(path string, query url.Values, user primitive.ObjectID, v interface{}) error {
	return testDo("GET", path, query, nil, user, v)
}

func testPost(path string, data interface{}, user primitive.ObjectID, v interface{}) error {
	return testDo("POST", path, nil, data, user, v)
}

func testPatch(path string, data interface{}, user primitive.ObjectID, v interface{}) error {
	return testDo("PATCH", path, nil, data, user, v)
}

func testPut(path string, data interface{}, user primitive.ObjectID, v interface{}) error {
	return testDo("PUT", path, nil, data, user, v)
}

func testDelete(path string, user primitive.ObjectID) error {
	return testDo("DELETE", path, nil, nil, user, nil)
}

func TestSignup(t *testing.T) {
	defer cleanup()
	var o userOut
	if err := testPost("/signup", nil, primitive.ObjectID{}, &o); err != nil {
		panic(err)
	}
}

func TestGetSelf(t *testing.T) {
	defer cleanup()
	var o userOut
	if err := testGet("/user", nil, testUser, &o); err != nil {
		panic(err)
	}
}

func TestGetUserVow(t *testing.T) {
	defer cleanup()
	ctx := context.Background()

	db.Collection("vows").InsertMany(ctx, []interface{}{
		vow{
			Text: "current vow",
			User: testUser,
		},
	})

	var out vowOut
	if err := testGet("/user/vow", nil, testUser, &out); err != nil {
		t.Fatal(err)
	}
}

func TestGetVows(t *testing.T) {
	defer cleanup()

	ctx := context.Background()

	db.Collection("vows").InsertMany(ctx, []interface{}{
		vow{
			Text: "progress vow",
			User: testUser,
		},
		vow{
			Text:     "archived vow",
			User:     testUser,
			Archived: true,
		},
	})

	var out vowsOut

	t.Run("all", func(t *testing.T) {
		if err := testGet("/vows", nil, primitive.ObjectID{}, &out); err != nil {
			t.Fatal(err)
		}

		if len(out.Vows) != 2 {
			t.Fatal("invalid length")
		}
	})

	t.Run("progress", func(t *testing.T) {
		if err := testGet("/vows", url.Values{"status": {"progress"}}, primitive.ObjectID{}, &out); err != nil {
			t.Fatal(err)
		}

		if len(out.Vows) != 1 {
			t.Fatal("invalid length")
		}
	})

	t.Run("archived", func(t *testing.T) {
		if err := testGet("/vows", url.Values{"status": {"archived"}}, primitive.ObjectID{}, &out); err != nil {
			t.Fatal(err)
		}

		if len(out.Vows) != 1 {
			t.Fatal("invalid length")
		}
	})
}

func TestPostVow(t *testing.T) {
	defer cleanup()
	txt := "testpost"

	var out vowOut
	testPost("/vows", vowIn{
		Type: "image",
		Text: txt,
	}, testUser, &out)

	if out.Vow.Type == "image" && out.Vow.Text != txt {
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
	defer cleanup()
	vow := primitive.NewObjectID()

	db.Collection("vows").InsertOne(context.Background(), bson.M{
		"_id":  vow,
		"user": testUser,
	})

	if err := testPut(fmt.Sprintf("/vows/%s/cheer", vow.Hex()), nil, testUser, nil); err != nil {
		panic(err)
	}
	var rs vowOut
	if err := testGet(fmt.Sprintf("/vows/%s", vow.Hex()), nil, testUser, &rs); err != nil {
		panic(err)
	}
	if rs.Vow.CheerCount != 1 {
		t.Fatal("invalid cheer length")
	}
}

func TestDeleteCheer(t *testing.T) {
	defer cleanup()
	vow := primitive.NewObjectID()

	db.Collection("vows").InsertOne(context.Background(), bson.M{
		"_id":  vow,
		"user": testUser,
	})

	if err := testPut(fmt.Sprintf("/vows/%s/cheer", vow.Hex()), nil, testUser, nil); err != nil {
		panic(err)
	}

	if err := testDelete(fmt.Sprintf("/vows/%s/cheer", vow.Hex()), testUser); err != nil {
		panic(err)
	}

	var rs vowOut
	if err := testGet(fmt.Sprintf("/vows/%s", vow.Hex()), nil, testUser, &rs); err != nil {
		panic(err)
	}
	if rs.Vow.CheerCount != 0 {
		t.Fatal("invalid cheer length")
	}
}

func ensureTestUser(ctx context.Context) {
	db.Collection("users").InsertOne(ctx, bson.M{
		"_id": testUser,
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
