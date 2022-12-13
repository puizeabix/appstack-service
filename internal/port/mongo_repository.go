package port

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/puizeabix/appstack-service/internal/domain"
)

var (
	ErrObjectIDSerialization = errors.New("Unable to deserialize/serialize ObjectID to string")
)

type mrepository struct {
	Collection mongo.Collection
}

func NewMongoRepository() AppStackRepository {
	return &mrepository{}
}

func (r *mrepository) Create(ctx context.Context, s *domain.AppStack) (string, error) {
	res, err := r.Collection.InsertOne(ctx, s)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", ErrObjectIDSerialization
}

func (r *mrepository) Get(ctx context.Context, id string) (*domain.AppStack, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := r.Collection.FindOne(ctx, bson.M{"_id": objID})

	var app domain.AppStack
	err = res.Decode(&app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (r *mrepository) List(ctx context.Context) ([]domain.AppStack, error) {
	cur, err := r.Collection.Find(ctx, nil)
	apps := []domain.AppStack{}
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var app domain.AppStack
		err := cur.Decode(&app)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)
	return apps, nil
}
