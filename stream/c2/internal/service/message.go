package service

import (
	"fmt"
	"io"

	v1 "c2/api/c2/v1"
	"c2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type StreamService struct {
	v1.UnimplementedStreamServer
	logger *log.Helper
	uc     *biz.StreamUsecase
}

func NewStreamService(uc *biz.StreamUsecase, logger log.Logger) *StreamService {
	return &StreamService{uc: uc, logger: log.NewHelper(logger)}
}

func (s *StreamService) Conversations(conn v1.Stream_ConversationsServer) error {
	s.logger.Infof("Connecting to stream server %v", conn)
	for {
		req, err := conn.Recv()
		if err == io.EOF {
			// 流结束
			fmt.Println("客户端发送的数据流结束")
			return nil
		}
		if err != nil {
			// 流出现错误
			fmt.Println("接收数据出错:", err)
			return err
		}
		// TODO 处理消息
		fmt.Println(req.GetReq())
		// 返回消息
		err = conn.Send(&v1.StreamResponse{
			Res: "Receive " + req.GetReq(),
		})
		if err != nil {
			// 返回出现出现错误
			return err
		}
	}
}
