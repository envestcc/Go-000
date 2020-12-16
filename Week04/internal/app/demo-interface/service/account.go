package service

import (
	"Week04/internal/pkg/dao"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IService interface {
	ListAccount(ctx *gin.Context)
}

type service struct {
	db dao.IAccount
}

func NewService(db dao.IAccount) IService {
	return &service{
		db: db,
	}
}

func (s *service) ListAccount(c *gin.Context) {
	s.db.List(context.Background())

	c.String(http.StatusOK, "")
}
