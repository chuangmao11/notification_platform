//go:build e2e

package grpc

import (
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"testing"
)

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
	grpcServer *grpc.Server
	listener   *bufconn.Listener
}
