package eyevow

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ensureDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		panic(err)
	}

	return client
}

func Exec() {
	mongo := ensureDB()
	defer mongo.Disconnect(context.Background())
	db = mongo.Database("eyevow")

	cmd := "server"

	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "server":
		serve()
	case "seed":
		seed()
	}

}
