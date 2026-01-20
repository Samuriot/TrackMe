package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccountLabel     string             `json:"account_label" bson:"account_label"`
	AccountType      string             `json:"account_type" bson:"account_type"`
	AccountNumber    string             `json:"account_number" bson:"account_number"`
	RoutingNumber    string             `json:"routing_number" bson:"routing_number"`
	CurrentBalance   float64            `json:"current_balance" bson:"current_balance"`
	AvailableBalance float64            `json:"available_balance" bson:"available_balance"`
	InterestRate     float64            `json:"interest_rate" bson:"interest_rate"`
	AcquiredInterest float64            `json:"acquired_interest" bson:"acquired_interest"`
}