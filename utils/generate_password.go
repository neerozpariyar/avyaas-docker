package utils

import (
	"github.com/sethvargo/go-password/password"
)

func GenerateRandomPassword() (string, error) {
	res, err := password.Generate(15, 10, 0, false, false)
	if err != nil {
		return "", err
	}
	return res, nil
}

func GenerateReferralCode() (string, error) {
	res, err := password.Generate(6, 2, 0, false, false)
	if err != nil {
		return "", err
	}
	return res, nil
}
