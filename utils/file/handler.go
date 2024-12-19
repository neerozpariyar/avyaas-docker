package file

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/spf13/viper"
)

/*
OBSConfig represents the configuration parameters required to initialize an OBS (Object Storage
Service) client.
*/
type OBSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
	BucketName      string
}

var config OBSConfig

/*
NewOBSClient initializes a new OBS client using the provided configuration data.
It returns the initialized OBS client and an error, if any.
*/
func NewOBSClient(config OBSConfig) (*obs.ObsClient, error) {
	obsClient, err := obs.New(config.AccessKeyID, config.SecretAccessKey, config.Endpoint)

	return obsClient, err
}

/*
UploadFileHandler processes a multipart file header, reads its content into a buffer,
and initializes an OBS client using the configured parameters.

Returns:
  - obsClinet: The pointer of initialized OBS client, the file content buffer.
  - err: An error, if any.
*/
func UploadFileHandler(file *multipart.FileHeader) (*obs.ObsClient, *bytes.Buffer, error) {
	// Replace the empty spaces in the filename with "-"
	file.Filename = Slugify(file.Filename)

	// Open the file and create a buffer to read its content.
	fileData, err := file.Open()
	if err != nil {
		fileData.Close()
		return nil, nil, err
	}

	// Ensure the file is closed when the function exits.
	defer fileData.Close()

	// Read the file data into a buffer.
	buffer := new(bytes.Buffer)
	if _, err = buffer.ReadFrom(fileData); err != nil {
		return nil, nil, err
	}

	// Configure OBS parameters from the Viper configuration.
	config = OBSConfig{
		AccessKeyID:     viper.GetString("huaweiOBS.accessKeyID"),
		SecretAccessKey: viper.GetString("huaweiOBS.secretAccessKey"),
		Endpoint:        viper.GetString("huaweiOBS.endpoint"),
		BucketName:      viper.GetString("huaweiOBS.bucketName"),
	}

	// Create an OBS client based on the loaded configuration.
	obsClient, err := NewOBSClient(config)
	if err != nil {
		return nil, nil, err
	}

	// Return the OBS client, buffer containing file data, and no error.
	return obsClient, buffer, nil
}

func GetSignedURL(objectKey string) (string, error) {
	// Configure OBS parameters from the Viper configuration.
	config = OBSConfig{
		AccessKeyID:     viper.GetString("huaweiOBS.accessKeyID"),
		SecretAccessKey: viper.GetString("huaweiOBS.secretAccessKey"),
		Endpoint:        viper.GetString("huaweiOBS.endpoint"),
		BucketName:      viper.GetString("huaweiOBS.bucketName"),
	}

	// Create an OBS client based on the loaded configuration.
	obsClient, err := NewOBSClient(config)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
		return "", err
	}

	getObjectInput := &obs.CreateSignedUrlInput{}
	getObjectInput.Method = obs.HttpMethodGet
	getObjectInput.Bucket = config.BucketName
	getObjectInput.Key = objectKey
	getObjectInput.Expires = 1800

	// Create a signed URL for downloading an object.
	getObjectOutput, err := obsClient.CreateSignedUrl(getObjectInput)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return getObjectOutput.SignedUrl, nil
}

func GetFileLength(objectKey string) (*FileLengthResponse, error) {
	config = OBSConfig{
		AccessKeyID:     viper.GetString("huaweiOBS.accessKeyID"),
		SecretAccessKey: viper.GetString("huaweiOBS.secretAccessKey"),
		Endpoint:        viper.GetString("huaweiOBS.endpoint"),
		BucketName:      viper.GetString("huaweiOBS.bucketName"),
	}

	// Create an OBS client based on the loaded configuration.
	obsClient, err := NewOBSClient(config)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
		return nil, err
	}

	input := &obs.GetObjectMetadataInput{
		Bucket: config.BucketName,
		Key:    objectKey,
	}

	mData, err := obsClient.GetObjectMetadata(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode == http.StatusNotFound {
			// File does not exist
			return nil, err
		}
		return nil, err
	}

	durationStr, ok := mData.Metadata["video_duration"]
	if !ok {
		return nil, fmt.Errorf("video duration not found in metadata")
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return nil, err
	}

	response := &FileLengthResponse{
		FileName: strings.Split(objectKey, "/")[1],
		Url:      "https://" + viper.GetString("huaweiOBS.bucketName") + "." + viper.GetString("huaweiOBS.endpoint") + "/" + objectKey,
		Length:   duration,
	}

	return response, nil
}
