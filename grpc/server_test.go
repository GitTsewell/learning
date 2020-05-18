package grpc

import "testing"

func TestNewServer(t *testing.T) {
	err := newServer(":10000")
	if err != nil {
		t.Error(err)
	}
}

func TestGetDemo(t *testing.T) {
	rsp, err := getDemo("127.0.0.1:10000")
	if err != nil {
		t.Error(err)
	}
	t.Logf("get demo rsp:%v", rsp)
}

func TestGetDemoStream(t *testing.T) {
	getDemoStream("127.0.0.1:10000")
}
