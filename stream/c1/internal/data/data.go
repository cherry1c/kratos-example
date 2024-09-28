package data

import (
	v1 "c1/api/c2/v1"
	"c1/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewC2Repo)

// Data .
type Data struct {
	// TODO wrapped database client
	c2 v1.StreamClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	c2, c2clean := NewC2Client(c.C2)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		c2clean()
	}
	return &Data{c2: c2}, cleanup, nil
}

func NewC2Client(c *conf.Data_Client) (v1.StreamClient, func()) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.GetEndpoint()),
		grpc.WithMiddleware(recovery.Recovery()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	log.Infof("c2 client connected")
	client := v1.NewStreamClient(conn)
	return client, func() { conn.Close() }
}
