package fn

import (
	"context"
	"testing"
)

func TestHandler_handleFunc(t *testing.T) {
	h := Handler{
		daoGetSomething: func(ctx context.Context) (string, error) {
			return "", nil
		},
	}
	h.handleFunc(context.Background())
}
