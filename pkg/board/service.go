package board

import (
	"context"

	"github.com/willwchan/commuter-app/internal"
)

type Service interface {
	Get(ctx context.Context) ([]internal.BoardPost, error)
	Post(ctx context.Context, post *internal.BoardPost) (string, error)
	//	Delete(ctx context.Context, id string) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
