package board

import (
	"context"
	"net/http"
	"os"

	"github.com/willwchan/commuter-app/internal"

	"github.com/go-kit/kit/log"
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

func (b *boardService) Post(_ context.Context) (string, error) {
	newPostId := shortuuid.New()
	return newPostId, nil
}

func (b *boardService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking service health...")
	return http.StatusOk, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
