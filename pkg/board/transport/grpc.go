package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	post          grpctransport.Handler
	serviceStatus grpctransport.Handler
}
