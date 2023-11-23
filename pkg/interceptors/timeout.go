package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

/*
 * @project coffee-store
 * @author nick
 */

func TimeoutUnaryInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
