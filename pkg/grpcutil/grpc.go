package grpcutil

import (
	"context"
	"fmt"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
)

func UserId(ctx context.Context) int64 {
	session, ok := ctx.Value(entity.UserClaimsContextKey).(*entity.UserClaims)
	if !ok {
		fmt.Println("failed to retrieve user from context")
		return 0
	}

	return session.UId
}
