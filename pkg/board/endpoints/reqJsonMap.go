package endpoints

import "github.com/willwchan/commuter-app/internal"

type GetRequest struct{}

type GetResponse struct {
	Posts []internal.BoardPost `json:"posts"`
	Err   string               `json:"err,omitempty"`
}

type PostRequest struct {
	Post *internal.BoardPost `json:"post"`
}

type PostResponse struct {
	PostId string `json:"postId"`
	Err    string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
