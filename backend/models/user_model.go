package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	NetWorth    float64            `json:"net_worth" bson:"net_worth"`
	Accounts    []string           `json:"accounts" bson:"accounts"`
	CreditScore int                `json:"credit_score" bson:"credit_score"`
	Budget      []string           `json:"budget" bson:"budget"`
}