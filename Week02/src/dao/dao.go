package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Dao struct {
}

func New() *Dao {
	return &Dao{}
}

func (d *Dao) Get(id uint64) error {

	if err := d.mockQuery(id); err != nil {
		return errors.Wrapf(err, "id=%d", id)
	}

	return nil
}

func (d *Dao) mockQuery(id uint64) error {
	switch id {
	case 0:
		return sql.ErrConnDone
	case 1:
		return sql.ErrNoRows
	default:
		return nil
	}
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
