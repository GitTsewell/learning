package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type demo struct{}

func (d *demo) GetDemo(ctx context.Context, in *GetDemoReq) (*GetDemoRsp, error) {
	log.Println("demo in:", in.Message)
	var rsp GetDemoRsp
	rsp.Code = 200
	rsp.Message = "hello"
	return &rsp, nil
}

func (d *demo) GetDemoWait(ctx context.Context, in *GetDemoReq) (*GetDemoRsp, error) {
	log.Println("demo wait in:", in.Message)
	var rsp GetDemoRsp
	rsp.Code = 200
	rsp.Message = "hello"
	time.Sleep(time.Second * 10)
	log.Println("demo wait out:", in.Message)
	return &rsp, nil
}

func (d *demo) GetDemoStream(rect *GetDemoReq, stream DemoService_GetDemoStreamServer) error {
	var err error
	ti := time.NewTicker(time.Second)
	for {
		select {
		case <-ti.C:
			err = stream.Send(&GetDemoRsp{Code: 200, Message: "message"})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func newServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	RegisterDemoServiceServer(s, &demo{})
	return s.Serve(lis)
}
