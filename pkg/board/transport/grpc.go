package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/willwchan/commuter-app/api/v1/pb/board"
	"github.com/willwchan/commuter-app/internal"
	"github.com/willwchan/commuter-app/pkg/board/endpoints"
)

type grpcServer struct {
	get           grpctransport.Handler
	post          grpctransport.Handler
	serviceStatus grpctransport.Handler
	board.UnimplementedBoardServer
}

func NewGrpcServer(ep endpoints.Set) board.BoardServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGrpcGetRequest,
			decodeGrpcGetResponse,
		),
		post: grpctransport.NewServer(
			ep.PostEndpoint,
			decodeGrpcPostRequest,
			decodeGrpcPostResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGrpcServiceStatusRequest,
			decodeGrpcServiceStatusResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *board.GetRequest) (*board.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*board.GetReply), nil
}

func (g *grpcServer) Post(ctx context.Context, r *board.PostRequest) (*board.PostReply, error) {
	_, rep, err := g.post.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*board.PostReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *board.ServiceStatusRequest) (*board.ServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*board.ServiceStatusReply), nil
}

func decodeGrpcGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.GetRequest{}, nil
}

func decodeGrpcPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*board.PostRequest)
	bp := &internal.BoardPost{
		Content: req.Board.Content,
		Author:  req.Board.Author,
	}
	return endpoints.PostRequest{Post: bp}, nil
}

func decodeGrpcServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGrpcGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*board.GetReply)
	var posts []internal.BoardPost
	for _, d := range reply.Boardposts {
		post := internal.BoardPost{
			Content: d.Content,
			Author:  d.Author,
		}
		posts = append(posts, post)
	}
	return endpoints.GetResponse{Posts: posts, Err: reply.Err}, nil
}

func decodeGrpcPostResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*board.PostReply)
	return endpoints.PostResponse{PostId: reply.PostId, Err: reply.Err}, nil
}

func decodeGrpcServiceStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*board.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
