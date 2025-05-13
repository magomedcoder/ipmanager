package entity

import "context"

type contextKey string

const UserClaimsContextKey contextKey = "userClaims"

type UserLogin struct {
	AccessToken string
	ExpiresAt   int64
}

type UserClaims struct {
	Token    string `json:"token,omitempty"`
	UId      uint   `json:"uid,omitempty"`
	Username string `json:"username,omitempty"`
}

func GetUserClaimsFromContext(ctx context.Context) (*UserClaims, bool) {
	if claims, ok := ctx.Value(UserClaimsContextKey).(*UserClaims); ok {
		return claims, true
	}

	return nil, false
}

type User struct {
	Id       int64
	Username string
	Name     string
	Surname  string
}
