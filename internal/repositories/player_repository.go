package repositories

import (
	"context"
	"dune-imperium-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PlayerRepository interface {
	Save(player *models.Player) error
	GetById(id string) (*models.Player, error)
	GetAllNames() ([]string, error)
	GetByNickname(nickname string) (*models.Player, error)
}

type playerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(db *mongo.Client) PlayerRepository {
	collection := db.Database("dune").Collection("results")
	return &playerRepository{
		collection: collection,
	}
}

func (r *playerRepository) Save(player *models.Player) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, player)
	return err
}

func (r *playerRepository) GetById(id string) (*models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var player models.Player
	filter := bson.M{"id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) GetAllNames() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.D{{Key: "nickname", Value: 1}, {Key: "_id", Value: 0}}
	opts := options.Find().SetProjection(projection)

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var names []string
	for cursor.Next(ctx) {
		var result struct {
			Nickname string `bson:"nickname"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		names = append(names, result.Nickname)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return names, nil
}

func (r *playerRepository) GetByNickname(nickname string) (*models.Player, error) {
	var user models.Player
	filter := bson.M{"nickname": nickname}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
