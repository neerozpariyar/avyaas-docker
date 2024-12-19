package utils

import (
	"strings"

	"github.com/spf13/viper"
)

// The use of this function will be replaced by GetURLObject in all functions where used
func GetObjectURL(url string) string {
	baseUrl := "https://" + viper.GetString("huaweiOBS.bucketName") + viper.GetString("huaweiOBS.endpoint") + "/"
	urlObject := strings.Split(url, baseUrl)[1]

	return urlObject
}

/*
GetURLObject extracts/separates the OBS URL objectKey from a complete URL. It takes full URL path
and return the objectKey split using fileURLSplitString value.

This function is used in create/update type services.
*/
func GetURLObject(url string) string {
	urlObject := strings.Split(url, viper.GetString("fileURLSplitString"))[1]

	return urlObject
}

/*
GetFileURL constructs and returns a full file URL path pre-pending the OBS bucket name and endpoint
value from config file. It takes the file objectKey and returns full constructed URL path.

This function is used in list type services.
*/
func GetFileURL(urlObjectKey string) string {
	url := "https://" + viper.GetString("huaweiOBS.bucketName") + "." + viper.GetString("huaweiOBS.endpoint") + "/" + urlObjectKey

	return url
}
