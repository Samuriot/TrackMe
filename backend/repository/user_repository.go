package repository

import (
	"context"

	"github.com/samuriot/track-me/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UserRepository defines the interface for user database operations
type UserRepository interface {
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, id primitive.ObjectID, update *models.User) (*models.User, error)
	DeleteUserByID(ctx context.Context, id primitive.ObjectID) error
}

// MongoUserRepository defines the specific MongoDB operations
type MongoUserRepository struct {
	collection *mongo.Collection
}

// MongoUserRepository Factory
func NewMongoUserRepository(db *mongo.Database) UserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User

    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *MongoUserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, cursor.Err()
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	_, err := r.collection.InsertOne(ctx, user)

	return err;

}

func (r *MongoUserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, update *models.User) (*models.User, error) {
	var user models.User

	updatedJSON := bson.M{
		"$set": bson.M{
			"username":     update.Username,
			"email":        update.Email,
			"net_worth":    update.NetWorth,
			"accounts":     update.Accounts,
			"credit_score": update.CreditScore,
			"budget":       update.Budget,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, updatedJSON, opts).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *MongoUserRepository) DeleteUserByID(ctx context.Context, id primitive.ObjectID) error {
	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}