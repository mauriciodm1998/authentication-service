package repositories

import (
	"authentication-service/internal/canonical"
	"authentication-service/internal/config"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rs/zerolog/log"
)

const (
	collection = "users"
	database   = "default"
)

type Repository interface {
	GetUser(ctx context.Context, login canonical.Login) (*canonical.User, error)
	CreateUser(context.Context, canonical.User) error
}

type repository struct {
	collection *mongo.Collection
}

func New() Repository {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Get().Database.ConnectionString))
	if err != nil {
		log.Fatal().Err(err).Msg("error when connect to the database")
	}

	return &repository{
		collection: client.Database(database).Collection(collection),
	}
}

func (r *repository) CreateUser(ctx context.Context, user canonical.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUser(ctx context.Context, login canonical.Login) (*canonical.User, error) {
	var statement primitive.D

	if login.Registration != "" {
		statement = bson.D{{Key: "registration", Value: login.Registration}}
	}

	if login.UserName != "" {
		statement = bson.D{{Key: "username", Value: login.UserName}}
	}

	var user canonical.User

	err := r.collection.FindOne(ctx, statement).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
