package repositories

import (
	"context"
	"dune-imperium-service/internal/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PlayerRepository interface {
	Save(ctx context.Context, player *models.Player) error
	GetById(ctx context.Context, id string) (*models.Player, error)
	GetNames(ctx context.Context) ([]string, error)
	GetByNickname(ctx context.Context, nickname string) (*models.Player, error)
}

type playerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(db *mongo.Client) PlayerRepository {
	collection := db.Database("dune").Collection("players")
	return &playerRepository{
		collection: collection,
	}
}

func (r *playerRepository) Save(ctx context.Context, player *models.Player) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, player)
	return err
}

func (r *playerRepository) GetById(ctx context.Context, id string) (*models.Player, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var player models.Player
	filter := bson.M{"id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) GetNames(ctx context.Context) ([]string, error) {
	names := make([]string, 0)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	projection := bson.D{{Key: "nickname", Value: 1}, {Key: "_id", Value: 0}}
	opts := options.Find().SetProjection(projection)

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return names, nil
		}
		return names, fmt.Errorf("failed to execute find query: %w", err)
	}
	defer cursor.Close(ctx)

	var results []struct {
		Nickname string `bson:"nickname"`
	}
	if err = cursor.All(ctx, &results); err != nil {
		return names, fmt.Errorf("failed to decode results: %w", err)
	}

	for _, result := range results {
		names = append(names, result.Nickname)
	}

	return names, nil
}

func (r *playerRepository) GetByNickname(ctx context.Context, nickname string) (*models.Player, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var user models.Player
	filter := bson.M{"nickname": nickname}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
