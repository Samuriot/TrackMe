package services

import (
	"context"
	"errors"

	"github.com/samuriot/track-me/models"
	"github.com/samuriot/track-me/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrUserNotFound = errors.New("Error: User Not Found")

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return users, err
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	// here would be the required account credentials/validation
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id primitive.ObjectID, user *models.User) (*models.User, error) {
	updatedUser, err := s.repo.UpdateUser(ctx, user.ID, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, err
}

func (s *UserService)DeleteUserByID(ctx context.Context, id primitive.ObjectID) error {
	err := s.repo.DeleteUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}