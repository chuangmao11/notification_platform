package grpc

import (
	"context"
	notificationv1 "github.com/chuangmao11/notification_platform/api/proto/gen/notification/v1"
)

type Server struct {
	notificationv1.UnimplementedCallbackServiceServer
}

func (s *Server) SendNotification(ctx context.Context, req *notificationv1.SendNotificationRequest) (*notificationv1.SendNotificationResponse, error) {
	return &notificationv1.SendNotificationResponse{}, nil
}

func (s *Server) SendNotificationAsync(ctx context.Context, req *notificationv1.SendNotificationAsyncRequest) (*notificationv1.SendNotificationAsyncResponse, error) {
	panic("implement me")
}

func (s *Server) BatchSendNotifications(ctx context.Context, req *notificationv1.BatchSendNotificationsRequest) (*notificationv1.BatchSendNotificationsResponse, error) {
	panic("implement me")
}

func (s *Server) BatchSendNotificationsAsync(ctx context.Context, req *notificationv1.BatchSendNotificationsAsyncRequest) (*notificationv1.BatchSendNotificationsAsyncResponse, error) {
	panic("implement me")
}

func (s *Server) QueryNotification(ctx context.Context, req *notificationv1.QueryNotificationRequest) (*notificationv1.QueryNotificationResponse, error) {
	panic("implement me")
}
