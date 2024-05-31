package notifications

import (
	"context"

	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/notifications/types"
)

// Store defines the API between the notifications system and the storage.
// This abstract is in place so that we can intercept the direct database interactions, or (later) swap out these calls
// with dRPC calls should we want to split the notifiers out into their own component for high availability/throughput.
// TODO: don't use database types here
type Store interface {
	AcquireNotificationMessages(ctx context.Context, params database.AcquireNotificationMessagesParams) ([]database.AcquireNotificationMessagesRow, error)
}

// Renderer is responsible for substituting any variable content in a given template with Labels.
type Renderer interface {
	Provider

	Render(template string, input types.Labels) (string, error)
}

// Dispatcher is responsible for delivering a notification to a given receiver.
type Dispatcher interface {
	Provider
	
	// Validate ensures that Labels which are required for the Send operation to complete accurately are present.
	Validate(input types.Labels) bool
	// Send delivers the notification.
	Send(ctx context.Context, input types.Labels) error
}
