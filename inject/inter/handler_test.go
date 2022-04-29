package inter

import (
	"context"
	"testing"

	"github.com/emilien-puget/gobyexample/inject/inter/testdata"
)

func TestHandler_handleFunc(t *testing.T) {
	h := Handler{
		dao: testdata.DaoMock{GetSomethingFunc: func(ctx context.Context) (string, error) {
			return "", nil
		}},
	}
	h.handleFunc(context.Background())
}
