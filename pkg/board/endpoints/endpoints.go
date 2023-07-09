package endpoints

import (
	"context"
	"errors"

	"github.com/willwchan/commuter-app/internal"
	"github.com/willwchan/commuter-app/pkg/board"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	PostEndpoint          endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc board.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		PostEndpoint:          MakePostEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func MakeGetEndpoint(svc board.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		posts, err := svc.Get(ctx)
		if err != nil {
			return GetResponse{posts, err.Error()}, nil
		}
		return GetResponse{posts, ""}, nil
	}
}

func MakePostEndpoint(svc board.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		postId, err := svc.Post(ctx, req.Post)
		if err != nil {
			return PostResponse{PostId: postId, Err: err.Error()}, nil
		}
		return PostResponse{PostId: postId, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc board.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context) ([]internal.BoardPost, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{})
	if err != nil {
		return []internal.BoardPost{}, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.BoardPost{}, errors.New(getResp.Err)
	}
	return getResp.Posts, nil
}

func (s *Set) Post(ctx context.Context, post *internal.BoardPost) (string, error) {
	resp, err := s.PostEndpoint(ctx, PostRequest{Post: post})
	if err != nil {
		return "", err
	}
	postResp := resp.(PostResponse)
	if postResp.Err != "" {
		return "", errors.New(postResp.Err)
	}
	return postResp.PostId, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}
