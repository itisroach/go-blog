package models

import "time"

type JWTRefreshRequest struct {
	Token string  `binding:"required"`
}

type JWTResponse struct {
	Access  string
	Refresh string
}

type RefreshToken struct {
	Token     string
	Username  string `gorm:"primarykey"`
	ExpiresAt time.Time
}

type TokenClaims struct {
	Username string
	ExpTime  time.Time
	Type     string
}