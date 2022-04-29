package testdata

import "context"

type DaoMock struct {
	GetSomethingFunc     func(ctx context.Context) (string, error)
	GetSomethingElseFunc func(ctx context.Context) (string, error)
}

func (m DaoMock) GetSomething(ctx context.Context) (string, error) {
	return m.GetSomethingFunc(ctx)
}

func (m DaoMock) GetSomethingElse(ctx context.Context) (string, error) {
	return m.GetSomethingElseFunc(ctx)
}
