package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"

	"{{.Mod}}/internal/model" // Assuming model package is at internal/model
)

type {{.Class}}Service struct {
	collection *mongo.Collection
}

func New{{.Class}}Service(db *mongo.Database) *{{.Class}}Service {
	return &{{.Class}}Service{
		collection: db.Collection("{{.Snake}}s"), // Collection name based on service name
	}
}

// Create creates a new {{.Camel}}
func (s *{{.Class}}Service) Create(ctx context.Context, data *model.{{.Class}}) (primitive.ObjectID, error) {
	result, err := s.collection.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, xerrors.Errorf("failed to create {{.Camel}}: %w", err)
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// GetByID retrieves a {{.Camel}} by its ID
func (s *{{.Class}}Service) GetByID(ctx context.Context, id primitive.ObjectID) (*model.{{.Class}}, error) {
	var {{.Camel}} model.{{.Class}}
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&{{.Camel}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, xerrors.Errorf("{{.Camel}} not found: %w", err)
		}
		return nil, xerrors.Errorf("failed to get {{.Camel}} by ID: %w", err)
	}
	return &{{.Camel}}, nil
}

// Update updates an existing {{.Camel}}
func (s *{{.Class}}Service) Update(ctx context.Context, id primitive.ObjectID, data *model.{{.Class}}) error {
	update := bson.M{"$set": data}
	result, err := s.collection.UpdateByID(ctx, id, update)
	if err != nil {
		return xerrors.Errorf("failed to update {{.Camel}}: %w", err)
	}
	if result.ModifiedCount == 0 {
		return xerrors.Errorf("no {{.Camel}} found with ID %s to update", id.Hex())
	}
	return nil
}

// Delete deletes a {{.Camel}} by its ID
func (s *{{.Class}}Service) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return xerrors.Errorf("failed to delete {{.Camel}}: %w", err)
	}
	if result.DeletedCount == 0 {
		return xerrors.Errorf("no {{.Camel}} found with ID %s to delete", id.Hex())
	}
	return nil
}
