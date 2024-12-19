package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type CustomClaim struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

func GenerateFCMJWT() (token string, err error) {
	expirationTime := time.Now().Add(1 * time.Minute).Unix()
	secret := viper.GetString("notificationService.fcm_ayata_secret")
	customClaims := &CustomClaim{
		Id: viper.GetUint64("notificationService.fcm_id"),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	token, err = jwtToken.SignedString([]byte(secret))
	return
}
