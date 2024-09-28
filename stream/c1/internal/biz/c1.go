package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type C2Repo interface {
	Conversations(ctx context.Context)
}

type C1Usecase struct {
	c2Repo C2Repo
	log    *log.Helper
}

func NewC1Usecase(c2Repo C2Repo, logger log.Logger) *C1Usecase {
	return &C1Usecase{c2Repo: c2Repo, log: log.NewHelper(logger)}
}

func (uc *C1Usecase) Conversations(ctx context.Context) {
	uc.log.WithContext(ctx).Infof("Conversations")
	uc.c2Repo.Conversations(ctx)
}
