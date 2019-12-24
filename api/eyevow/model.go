package eyevow

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Icon string             `bson:"icon" json:"icon"`
}

type self struct {
	user
	TwitterID string `bson:"twitter_id,omitempty" json:"twitter_id,omitempty"`
}

type vow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Type     string             `bson:"type"`
	Text     string             `bson:"text"`
	User     primitive.ObjectID `bson:"user"`
	Archived bool               `bson:"archived"`
}
