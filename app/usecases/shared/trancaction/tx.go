package trancaction

import (
	"study-pal-backend/ent"
)

type Tx interface {
	Commit() error
	Rollback()
}

type tx struct {
	tx *ent.Tx
}

func NewTx(value *ent.Tx) Tx {
	return &tx{
		tx: value,
	}
}

func (t *tx) Commit() error {
	return t.tx.Commit()
}

func (t *tx) Rollback() {
	t.tx.Rollback()
}
