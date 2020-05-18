package grpc

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func getDemo(address string) (*GetDemoRsp, error) {
	// grpc.WithInsecure 表示跳过证书认证
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := NewDemoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	rsp, err := c.GetDemo(ctx, &GetDemoReq{Message: "message"})
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func getDemoStream(address string) {
	// grpc.WithInsecure 表示跳过证书认证
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := NewDemoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stream, err := c.GetDemoStream(ctx, &GetDemoReq{Message: "message"})
	if err != nil {
		panic(err)
	}

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("err : %v", err)
		}
		log.Println(feature)
	}
}
