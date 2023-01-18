package repo

import "go.mongodb.org/mongo-driver/mongo"

var client *mongo.Client

type DataEntry struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

type Models struct {
	DataEntry DataEntry
}

func New(c *mongo.Client) Models {
	client = c

	return Models{
		DataEntry: DataEntry{},
	}
}
