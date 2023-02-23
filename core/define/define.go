package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var CodeExpire = 600

var PageSize = 20

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
