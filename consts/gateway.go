package consts

import "time"

const (
	AccessTokenHeader           = "access_token"
	RefreshTokenHeader          = "refresh_token"
	HeaderForwardedProto        = "X-Forwarded-Proto"
	MaxAge                      = 3600 * 24
	AccessTokenExpiredDuration  = 24 * time.Hour
	RefreshTokenExpiredDuration = 10 * 24 * time.Hour
)

const UserInfoKey = "user_info_key"
