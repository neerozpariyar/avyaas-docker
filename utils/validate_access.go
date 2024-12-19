package utils

import (
	"github.com/spf13/viper"
)

func ValidateAccess(originURL string, roleID int) map[string]string {
	adminOriginURL := viper.GetString("server.adminOriginUrl")
	userOriginURL := viper.GetString("server.userOriginUrl")

	errMap := make(map[string]string)
	errMap["403"] = "forbidden, looks like you are trying to access wrong route"

	if originURL == adminOriginURL && roleID == 4 {
		return errMap
	}

	if originURL == userOriginURL && roleID != 4 {
		return errMap
	}

	return nil
}
