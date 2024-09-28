package service

import (
	v1 "c1/api/c1/v1"
	"c1/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type C1Service struct {
	v1.UnimplementedC1Server
	uc     *biz.C1Usecase
	logger *log.Helper
}

func NewC1Service(uc *biz.C1Usecase, logger log.Logger) *C1Service {
	return &C1Service{uc: uc, logger: log.NewHelper(logger)}
}

func (s *C1Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.uc.Conversations(ctx)
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
