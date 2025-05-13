package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func addAuthorizationToken(ctx context.Context) context.Context {
	md := metadata.Pairs("authorization", "Bearer ")
	return metadata.NewOutgoingContext(ctx, md)
}

func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = addAuthorizationToken(ctx)

	return invoker(ctx, method, req, reply, cc, opts...)
}

func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	ctx = addAuthorizationToken(ctx)

	return streamer(ctx, desc, cc, method, opts...)
}
