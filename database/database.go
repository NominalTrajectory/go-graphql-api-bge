package database

import (
	"context"
	"log"
	"time"

	"github.com/NominalTrajectory/go-graphql-api-bge/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client}
}

func (db *DB) AddDevice(input *model.NewDevice) *model.Device {
	collection := db.client.Database("bge-hillis-army").Collection("devices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Device{
		ID:             res.InsertedID.(primitive.ObjectID).Hex(),
		Title:          input.Title,
		Description:    input.Description,
		Specifications: input.Specifications,
	}
}

func (db *DB) FindDeviceById(id string) *model.Device {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("bge-hillis-army").Collection("devices")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}

	device := model.Device{}
	res.Decode(&device)
	return &device
}

func (db *DB) GetAllDevices() []*model.Device {
	collection := db.client.Database("bge-hillis-army").Collection("devices")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var devices []*model.Device

	for cur.Next(ctx) {
		var device *model.Device
		err := cur.Decode(&device)
		if err != nil {
			log.Fatal(err)
		}
		devices = append(devices, device)
	}

	return devices
}
