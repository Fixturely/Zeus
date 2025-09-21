package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Subscription struct {
	bun.BaseModel    `bun:"subscriptions"`
	ID               int       `bun:"id,pk,autoincrement" json:"id"`
	UserID           int       `bun:"user_id,notnull" json:"user_id"`
	TeamID           *int      `bun:"team_id" json:"team_id,omitempty"`
	SportID          int       `bun:"sport_id,notnull" json:"sport_id"`
	IsActive         bool      `bun:"is_active" json:"is_active"`
	BillingStatus    string    `bun:"billing_status,notnull" json:"billing_status"`
	SubscriptionType string    `bun:"subscription_type" json:"subscription_type"`
	DateCreated      time.Time `bun:"date_created" json:"date_created"`
	DateUpdated      time.Time `bun:"date_updated" json:"date_updated"`
	DateDeleted      time.Time `bun:"date_deleted,notnull" json:"date_deleted"`
}
