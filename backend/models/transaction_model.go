package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Transaction struct {
		ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name string `json:"name" bson:"name"`
		AccountNumber string `json:"account_number" bson:"account_number"`
		Category string `json:"category" bson:"category"`
		Type string `json:"type" bson:"type"`
		BudgetID primitive.ObjectID `json:"budget_id" bson:"budget_id,omitempty"` 
		Amount float32 `json:"amount" bson:"amount"`
		PointsRewarded float32 `json:"points_rewarded" bson:"points_rewarded"`
		TransactionDate time.Time `json:"transaction_date" bson:"transaction_date"`
		TransactionPosted time.Time `json:"transaction_posted" bson:"transaction_posted"`
		Description string `json:"description" bson:"description"`
}