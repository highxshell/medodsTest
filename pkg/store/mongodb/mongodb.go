package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"os"
	"time"
)

type Storage struct {
	DB *mongo.Client
}

func New() (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO")))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	slog.Info("connected to mongodb")
	return &Storage{DB: client}, nil
}

func (s *Storage) OpenCollection(dbName, collectionName string) *mongo.Collection {
	return s.DB.Database(dbName).Collection(collectionName)
}
