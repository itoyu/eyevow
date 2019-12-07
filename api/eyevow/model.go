package eyevow

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Icon string             `bson:"icon" json:"icon"`
}

type vow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Text     string             `bson:"text"`
	User     primitive.ObjectID `bson:"user"`
	Archived bool               `bson:"archived"`
}
