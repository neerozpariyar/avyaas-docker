package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func (uCase *usecase) MeetingSDKKey(key string, meetingID int64, role int, secret string) (string, error) {
	_, err := uCase.repo.GetLiveByMeetingID(meetingID)
	if err != nil {
		return "", err
	}
	iat := time.Now().Unix()
	exp := iat + 3600*2 // 2 hours from issuance
	email, err := uCase.repo.GetEmailByMeetingID(meetingID)
	var zoomAccount string
	switch email {
	case "abc@gmail.com":
		zoomAccount = "zoom1"
	case "xyz@gmail.com":
		zoomAccount = "zoom2"
	default:
		return "", fmt.Errorf("no Zoom account found for email: %s", email)
	}
	// Load credentials for the selected Zoom account
	creds := viper.GetStringMapString(zoomAccount)
	if creds == nil {
		return "", fmt.Errorf("credentials not found for Zoom account: %s", zoomAccount)
	}
	if err != nil {
		return "", err
	}
	appKey := creds["app_key"]
	secret = creds["app_secret"]
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"appKey":   appKey, //appkey for app, sdkkey for web--clientID of SDK key same thing
		"mn":       meetingID,
		"role":     role,
		"iat":      iat,
		"exp":      exp,
		"tokenExp": exp,
	})
	return token.SignedString([]byte(secret))
}
