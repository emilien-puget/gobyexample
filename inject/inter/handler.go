package inter

import "context"

type Handler struct {
	dao DaoI
}

type DaoI interface {
	GetSomething(ctx context.Context) (string, error)
}

func NewHandler(dao *Dao) *Handler {
	return &Handler{dao: dao}
}

func (h Handler) handleFunc(ctx context.Context) {
	something, err := h.dao.GetSomething(ctx)
	if err != nil {
		return
	}
	_ = something
}
