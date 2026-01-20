package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MinimumSpending  float64            `json:"minimum_spending" bson:"minimum_spending"`
	MaximumSpending  float64            `json:"maximum_spending" bson:"maximum_spending"`
	TargetGoal       float64            `json:"target_goal" bson:"target_goal"`
	StartDate        time.Time          `json:"start_date" bson:"start_date"`
	EndDate          time.Time          `json:"end_date" bson:"end_date"`
	IsMeetingBudget  bool               `json:"is_meeting_budget" bson:"is_meeting_budget"`
}