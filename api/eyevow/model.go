package eyevow

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Icon string             `bson:"icon"`
}

type vow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Text     string             `bson:"text"`
	User     primitive.ObjectID `bson:"user"`
	Archived bool               `bson:"archived"`
}
