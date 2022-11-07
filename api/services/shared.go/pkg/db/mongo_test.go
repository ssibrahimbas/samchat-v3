package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMongoModule(t *testing.T) {
	t.Run("Test New Mongo Function", func(t *testing.T) {
		mdb, err := NewMongo("mongodb://localhost:27017", "test")
		if err != nil {
			t.Fatal(err)
		}
		if err := mdb.c.Ping(context.TODO(), readpref.Primary()); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test New Mongo Function With Wrong URI", func(t *testing.T) {
		mdb, err := NewMongo("mongod://localst:27011", "test")
		if mdb != nil {
			t.Fatal("mdb is not nil")
		}
		if err == nil {
			t.Fatal("err should not nil")
		}
	})

	t.Run("Test MongoDB Get Collection Function", func(t *testing.T) {
		mdb, err := NewMongo("mongodb://localhost:27017", "test")
		if err != nil {
			t.Fatal(err)
		}
		col := mdb.GetCollection("test")
		if col == nil {
			t.Fatal("collection is nil")
		}
	})

	t.Run("Test Mongo Transform Id Function", func(t *testing.T) {
		mdb, err := NewMongo("mongodb://localhost:27017", "test")
		if err != nil {
			t.Fatal(err)
		}
		id, err := mdb.TransformId("5c9a2d9e7b4d1e0001f0d6d4")
		if !primitive.IsValidObjectID(id.Hex()) || err != nil {
			t.Fatal("id is not valid")
		}
	})

	t.Run("Test Mongo Transform Id Function With Invalid Id", func(t *testing.T) {
		mdb, err := NewMongo("mongodb://localhost:27017", "test")
		if err != nil {
			t.Fatal(err)
		}
		id, err := mdb.TransformId("5c9a2d9e7b4d1e0001f0d6d")
		if err == nil || id != primitive.NilObjectID {
			t.Fatal("id should be nil")
		}
	})
}
