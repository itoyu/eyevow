package eyevow

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dummyUserIDs = []primitive.ObjectID{
	primitive.NewObjectID(),
	primitive.NewObjectID(),
	primitive.NewObjectID(),
	primitive.NewObjectID(),
	primitive.NewObjectID(),
}

var dummyUsers = []interface{}{
	bson.M{
		"_id": dummyUserIDs[0],
	},
	bson.M{
		"_id": dummyUserIDs[1],
	},
	bson.M{
		"_id": dummyUserIDs[2],
	},
	bson.M{
		"_id": dummyUserIDs[3],
	},
	bson.M{
		"_id": dummyUserIDs[4],
	},
}

var dummyVows = []interface{}{
	bson.M{
		"user":        dummyUserIDs[0],
		"text":        "01 I going to make a go of this project!",
		"cheer_count": 10,
		"archived":    true,
	},
	bson.M{
		"user":        dummyUserIDs[1],
		"text":        "02 I going to make a go of this project!",
		"cheer_count": 20,
		"archived":    false,
	},
	bson.M{
		"user":        dummyUserIDs[0],
		"text":        "03 I going to make a go of this project!",
		"cheer_count": 30,
		"archived":    true,
	},
	bson.M{
		"user":        dummyUserIDs[2],
		"text":        "04 I going to make a go of this project!",
		"cheer_count": 0,
		"archived":    true,
	},
	bson.M{
		"user":        dummyUserIDs[3],
		"text":        "05 I going to make a go of this project!",
		"cheer_count": 0,
		"archived":    false,
	},
	bson.M{
		"user":        dummyUserIDs[4],
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
