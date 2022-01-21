package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func TestContextValue(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	ctx = context.WithValue(ctx, "key2", "value2")
	ctx = context.WithValue(ctx, "key3", "value3")
	v := ctx.Value("key")
	v2 := ctx.Value("key2")
	v3 := ctx.Value("key3")
	fmt.Println(v)
	fmt.Println(v2)
	fmt.Println(v3)
}
