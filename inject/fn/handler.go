package fn

import "context"

type Handler struct {
	daoGetSomething func(ctx context.Context) (string, error)
}

func NewHandler(dao *Dao) *Handler {
	return &Handler{daoGetSomething: dao.GetSomething}
}

func (h Handler) handleFunc(ctx context.Context) {
	something, err := h.daoGetSomething(ctx)
	if err != nil {
		return
	}
	_ = something
}
