package repository

import (
	"context"
	"inframanager/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRepository interface {
	FindById(ctx context.Context, id int) (*models.Player, error)
}

type playerRepository struct {
	client *mongo.Client
}

func (r *playerRepository) FindById(ctx context.Context, id int) (*models.Player, error) {
	var player models.Player

	err := r.client.Collection("players").FindOne(ctx, bson.M{"_id": id}).Decode(&player)

	if err != nil {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &player, nil
}
