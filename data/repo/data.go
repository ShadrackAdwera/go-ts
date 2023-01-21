package repo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type DataEntry struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
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

func (d *DataEntry) Insert(entry DataEntry) (*mongo.InsertOneResult, error) {
	collection := client.Database("data").Collection("data")

	result, err := collection.InsertOne(context.TODO(), DataEntry{
		Title:       entry.Title,
		Description: entry.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Println("Error creating product:", err)
		return nil, err
	}

	return result, nil
}

func (d *DataEntry) GetData() ([]*DataEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("data").Collection("data")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*DataEntry

	// iterate through the cursor, decode into entry struct
	for cursor.Next(ctx) {
		var item DataEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Print("Error decoding product into slice:", err)
			return nil, err
		} else {
			categories = append(categories, &item)
		}
	}

	return categories, nil
}

func (p *DataEntry) GetById(id string) (*DataEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("products").Collection("products")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entry DataEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (d *DataEntry) UpdateData(id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	collection := client.Database("data").Collection("data")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: d.Title},
			{Key: "description", Value: d.Description},
			{Key: "updated_at", Value: time.Now()},
		}},
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DataEntry) Delete(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	collection := client.Database("data").Collection("data")

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": docId})

	if err != nil {
		return nil, err
	}

	return result, nil

}
