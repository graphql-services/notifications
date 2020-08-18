package notifications

import (
	"context"
	"time"

	"github.com/graphql-services/go-saga/graphqlorm"
	"github.com/graphql-services/notifications/gen"
)

// Client ...
type Client struct {
	gc *graphqlorm.ORMClient
}

// NewClient ...
func NewClient(URL string) *Client {
	client := graphqlorm.NewClient(URL)
	return &Client{client}
}

type Notification gen.Notification

type CreateNotificationInput struct {
	Subject     string    `json:"message"`
	Message     string    `json:"message"`
	Seen        bool      `json:"seen"`
	Principal   *string   `json:"principal"`
	Channel     *string   `json:"channel"`
	Reference   *string   `json:"reference"`
	ReferenceID *string   `json:"referenceID"`
	Date        time.Time `json:"date"`
}

// InviteUser invite user with given Email. If user with given email exists, it just return without any invitation.
func (c *Client) CreateNotification(ctx context.Context, input CreateNotificationInput) (result graphqlorm.MutationResult, err error) {
	return c.gc.CreateEntity(ctx, graphqlorm.CreateEntityOptions{
		Entity: "Notification",
		Input:  input,
	})
}
