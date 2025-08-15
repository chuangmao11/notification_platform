package service

import (
	"context"

	"github.com/chuangmao11/notification_platform/internal/service/notification/domain"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, key string) (domain.Notification, error)
}
