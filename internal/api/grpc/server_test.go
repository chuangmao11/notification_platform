//go:build e2e

package grpc

import (
	"context"
	"net"
	"testing"

	notificationv1 "github.com/chuangmao11/notification_platform/api/proto/gen/notification/v1"
	grpcapi "github.com/chuangmao11/notification_platform/internal/api/grpc"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
	grpcServer *grpc.Server
	listener   *bufconn.Listener
}

func (s *ServerTestSuite) SetupSuite() {
	s.listener = bufconn.Listen(1024 * 1024)

	//启动grpc server
	s.grpcServer = grpc.NewServer()
	notificationv1.RegisterNotificationServiceServer(s.grpcServer, &grpcapi.Server{})

	ready := make(chan struct{})
	go func() {
		close(ready)
		if err := s.grpcServer.Serve(s.listener); err != nil {
			s.NoError(err, "gRpc Server exited")
		}
	}()
	<-ready
}

func (s *ServerTestSuite) TearDownSuite() {
	s.grpcServer.Stop()
}

func (s *ServerTestSuite) newClientConn() *grpc.ClientConn {
	conn, err := grpc.NewClient(
		"passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	s.NoError(err)
	return conn
}

func (s *ServerTestSuite) TestSendNotification() {
	t := s.T()
	testCases := []struct {
		name     string
		before   func(t *testing.T)
		after    func(t *testing.T)
		req      *notificationv1.SendNotificationRequest
		wantResp *notificationv1.SendNotificationResponse
		wantErr  error
	}{
		{
			name:     "SMS_立即发送_成功",
			before:   func(t *testing.T) {},
			after:    func(t *testing.T) {},
			req:      &notificationv1.SendNotificationRequest{},
			wantResp: &notificationv1.SendNotificationResponse{},
			wantErr:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.before(t)

			conn := s.newClientConn()
			defer conn.Close()
			client := notificationv1.NewNotificationServiceClient(conn)
			resp, err := client.SendNotification(context.Background(), tc.req)
			if tc.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, tc.wantErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.True(t, proto.Equal(tc.wantResp, resp))
			}
			tc.after(t)
		})
	}
}
