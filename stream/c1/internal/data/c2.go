package data

import (
	"context"
	"fmt"
	"io"

	v1 "c1/api/c2/v1"
	"c1/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.C2Repo = &c2Repo{}

type c2Repo struct {
	data *Data
	log  *log.Helper
}

func NewC2Repo(data *Data, logger log.Logger) biz.C2Repo {
	return &c2Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *c2Repo) Conversations(ctx context.Context) {
	conn, err := c.data.c2.Conversations(ctx)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := conn.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("receive error: %v", err)
			}
			log.Infof("received message: %v", in.GetRes())
		}
	}()

	for n := 0; n < 5; n++ {
		if err := conn.Send(&v1.StreamRequest{Req: fmt.Sprintf("this is %d", n)}); err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
	}
	if err = conn.CloseSend(); err != nil {
		log.Fatalf("failed to close send: %v", err)
	}
	<-waitc
}
