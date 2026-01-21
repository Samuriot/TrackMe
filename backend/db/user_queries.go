package db

import (
	"context"
	"errors"
	"github.com/samuriot/track-me/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// GetUserByID fetches a user document by ObjectID from the users collection.
func GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User

	collection := client.Database(dbName).Collection("users")
	res := collection.FindOne(ctx, bson.M{"_id": id})

	if res.Err() != nil {
		return user, res.Err()
	}

	if err := res.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

// GetAllUsers returns all users from the users collection.
// For production, consider adding pagination (limit/skip) and filters.
func GetAllUsers(ctx context.Context) ([]models.User, error) {
	collection := client.Database(dbName).Collection("users")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	users := make([]models.User, 0)
	for cur.Next(ctx) {
		var user models.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser inserts a new user document and returns the inserted ID.
func CreateUser(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	collection := client.Database(dbName).Collection("users")

	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("inserted ID is not an ObjectID")
	}
	return oid, nil
}

func UpdateUser(ctx context.Context, id primitive.ObjectID, update models.User) (models.User, error) {
	var user models.User

	collection := client.Database(dbName).Collection("users")

	updateDoc := bson.M{
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
	res := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, updateDoc, opts)

	if res.Err() != nil {
		return user, res.Err()
	}

	if err := res.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

// DeleteUserByID deletes a user by ObjectID.
// Returns mongo.ErrNoDocuments when no user was deleted.
func DeleteUserByID(ctx context.Context, id primitive.ObjectID) error {
	collection := client.Database(dbName).Collection("users")

	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}