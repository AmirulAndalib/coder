package notifications

import (
	"context"

	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/notifications/dispatch"
	"github.com/coder/coder/v2/coderd/notifications/types"
	"github.com/google/uuid"
)

// Store defines the API between the notifications system and the storage.
// This abstract is in place so that we can intercept the direct database interactions, or (later) swap out these calls
// with dRPC calls should we want to split the notifiers out into their own component for high availability/throughput.
// TODO: don't use database types here
type Store interface {
	AcquireNotificationMessages(ctx context.Context, params database.AcquireNotificationMessagesParams) ([]database.AcquireNotificationMessagesRow, error)
	BulkMarkNotificationMessagesSent(ctx context.Context, arg database.BulkMarkNotificationMessagesSentParams) (int64, error)
	BulkMarkNotificationMessagesFailed(ctx context.Context, arg database.BulkMarkNotificationMessagesFailedParams) (int64, error)
	EnqueueNotificationMessage(ctx context.Context, arg database.EnqueueNotificationMessageParams) (database.NotificationMessage, error)
	FetchNewMessageMetadata(ctx context.Context, arg database.FetchNewMessageMetadataParams) (database.FetchNewMessageMetadataRow, error)
}

// Handler is responsible for preparing and delivering a notification by a given method.
type Handler interface {
	NotificationMethod() database.NotificationMethod

	// Dispatcher delivers the notification by a given method.
	Dispatcher(payload types.MessagePayload, title, body string) (dispatch.DeliveryFunc, error)
}

type Enqueuer interface {
	Enqueue(ctx context.Context, userID, templateID uuid.UUID, labels types.Labels, createdBy string, targets ...uuid.UUID) (*uuid.UUID, error)
}
