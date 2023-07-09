package board

import (
	"context"
	"net/http"
	"os"

	"github.com/willwchan/commuter-app/internal"

	"github.com/go-kit/log"
	"github.com/lithammer/shortuuid/v3"
)

type boardService struct{}

func NewService() Service { return &boardService{} }

func (b *boardService) Get(_ context.Context) ([]internal.BoardPost, error) {
	post := internal.BoardPost{
		Content: "content",
		Author:  "user",
	}
	return []internal.BoardPost{post}, nil
}

func (b *boardService) Post(_ context.Context, post *internal.BoardPost) (string, error) {
	newPostId := shortuuid.New()
	return newPostId, nil
}

//func (b *boardService) Delete(_ context.Context, postId string) (string, error) {
//	return postId, nil
//}

func (b *boardService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
