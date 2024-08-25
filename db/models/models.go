package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username, PasswordHash, Role string
}

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const RefeshTokenValidTime = time.Hour * 72


