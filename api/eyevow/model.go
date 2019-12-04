package eyevow

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

type vow struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Text       string             `bson:"text"`
	User       primitive.ObjectID `bson:"user"`
	CheerCount int                `bson:"cheer_count"`
	Archived   bool               `bson:"archived"`
}
