package inter

import "context"

type HandlerB struct {
	dao DaoI
}

type DaoIB interface {
	GetSomething(ctx context.Context) (string, error)
}

func NewHandlerB(dao *Dao) *Handler {
	return &Handler{dao: dao}
}

func (h HandlerB) handleFunc(ctx context.Context) {
	something, err := h.dao.GetSomething(ctx)
	if err != nil {
		return
	}
	_ = something
}
