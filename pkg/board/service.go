package board

import (
	"context"

	"github.com/willwchan/commuter-app/internal"
)

type Service interface {
	Get(ctx context.Context) ([]internal.BoardPost, error)
	Post(ctx context.Context, post string) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
