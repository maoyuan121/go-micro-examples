package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	post "github.com/micro/examples/blog/posts/proto/posts"
)

type Post struct{}

func (e *Post) Handle(ctx context.Context, msg *post.Post) error {
	log.Info("Handler Received message: ", msg)
	return nil
}

func Handler(ctx context.Context, msg *post.Post) error {
	log.Info("Handler Received message: ", msg)
	return nil
}
