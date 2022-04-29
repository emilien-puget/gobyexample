package inter_anon

import (
	"context"
	"database/sql"
)

type Dao struct {
	s *sql.DB
}

func NewDao(s *sql.DB) *Dao {
	return &Dao{s: s}
}

func (d Dao) GetSomething(ctx context.Context) (string, error) {
	return "", nil
}

func (d Dao) GetSomethingElse(ctx context.Context) (string, error) {
	return "", nil
}
