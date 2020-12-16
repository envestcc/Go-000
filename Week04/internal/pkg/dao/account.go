package dao

import (
	"context"
	"database/sql"

	"gitlab.p1staff.com/backend/ttx-core/app/core-restapi/model"
)

type IDaoAccount interface {
	List(ctx context.Context) ([]model.Account, error)
	Create(ctx context.Context, account *model.Account) error
}

type daoAccount struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) IDaoAccount {
	return &daoAccount{db}
}

func (a *daoAccount) List(ctx context.Context) ([]model.Account, error) {
	return []model.Account{}, nil
}

func (a *daoAccount) Create(ctx context.Context, account *model.Account) error {
	return nil
}
