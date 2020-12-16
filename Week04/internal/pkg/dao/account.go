package dao

import (
	"Week04/internal/pkg/model"
	"context"
	"database/sql"
)

type IAccount interface {
	List(ctx context.Context) ([]model.Account, error)
	Create(ctx context.Context, account *model.Account) error
}

type account struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) IAccount {
	return &account{db}
}

func (a *account) List(ctx context.Context) ([]model.Account, error) {
	return []model.Account{}, nil
}

func (a *account) Create(ctx context.Context, account *model.Account) error {
	return nil
}
