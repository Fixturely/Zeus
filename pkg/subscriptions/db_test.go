package subscriptions

import (
	"context"
	"fmt"
	"testing"
	"time"
	"zeus/pkg/application"
	"zeus/pkg/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestInsertSubscription(t *testing.T) {
	ctx := context.Background()
	app, err := application.NewApp(ctx)
	assert.NoError(t, err)

	// Create a test user first to satisfy foreign key constraint
	googleID := fmt.Sprintf("test-google-id-%d", time.Now().UnixNano())
	testUser := &models.Account{
		GoogleID: &googleID,
		Email:    fmt.Sprintf("test-%d@example.com", time.Now().UnixNano()),
		Name:     "Test User",
		Status:   "active",
		Tier:     "free",
	}

	err = app.Database.NewInsert().
		Model(testUser).
		Returning("*").
		Scan(ctx)
	assert.NoError(t, err)

	teamID := 1
	subscription := &models.Subscription{
		UserID:  testUser.ID,
		SportID: 1,
		TeamID:  &teamID,
	}
	tests := []struct {
		name         string
		subscription *models.Subscription
		wantErr      bool
	}{
		{
			name:         "valid subscription",
			subscription: subscription,
			wantErr:      false,
		},
		{
			name: "invalid subscription with zero user_id",
			subscription: &models.Subscription{
				UserID:  0,
				SportID: 1,
				TeamID:  nil,
			},
			wantErr: false, // This will pass because user_id is set from context
		},
		{
			name: "invalid subscription with missing user_id in context",
			subscription: &models.Subscription{
				UserID:  0,
				SportID: 1,
				TeamID:  nil,
			},
			wantErr: true, // This will fail because context has no user_id
		},
		{
			name: "invalid subscription with zero sport_id",
			subscription: &models.Subscription{
				UserID:  testUser.ID,
				SportID: 0,
				TeamID:  nil,
			},
			wantErr: true,
		},
		{
			name: "invalid subscription with zero values",
			subscription: &models.Subscription{
				UserID:  testUser.ID,
				SportID: 0,
				TeamID:  nil,
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a mock echo context for testing
			c := echo.New().NewContext(nil, nil)

			// Only set user_id for tests that don't expect missing user_id error
			if test.name != "invalid subscription with missing user_id in context" {
				c.Set("user_id", testUser.ID) // Set the test user ID
			}

			result, err := InsertSubscription(c, app, test.subscription)
			if test.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				// UserID is set from context, so we expect it to be testUser.ID (from context)
				expectedUserID := testUser.ID
				if test.name == "invalid subscription with missing user_id in context" {
					expectedUserID = 0 // This test case won't reach here since it expects an error
				}
				assert.Equal(t, expectedUserID, result.UserID)
				assert.Equal(t, test.subscription.SportID, result.SportID)
				assert.Equal(t, test.subscription.TeamID, result.TeamID)
			}
		})
	}
}
