// package file

// import (
// 	"bytes"
// 	"fmt"
// 	"mime/multipart"
// 	"net/http"

// 	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
// )

// /*
// UploadFile initiates a multipart upload for a given file and uploads its parts to an OBS (Object
// Storage Service) bucket.

// Parameters:
//   - objDir: The target directory within the bucket (objDir)
//   - file: A FileHeader representing the file to upload.

// Returns:
//   - response: A FileUploadSuccessResponse struct containing information about the uploaded file.
//   - err: An error encountered during the process, if any.
// */
// func UploadFile(objDir string, file *multipart.FileHeader) (*FileUploadSuccessResponse, error) {
// 	// Obtain an OBS client and buffer containing file data from the UploadFileHandler function.
// 	obsClient, buffer, err := UploadFileHandler(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Generate the key for the file within the bucket.
// 	key := objDir + "/" + file.Filename

// 	// Initialize the parameters for initiating a multipart upload.
// 	input := &obs.InitiateMultipartUploadInput{}
// 	input.Bucket = config.BucketName
// 	input.Key = key

// 	// Initiate the multipart upload and obtain the upload ID.
// 	output, err := obsClient.InitiateMultipartUpload(input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	uploadID := output.UploadId

// 	// Configure chunk size for uploading parts.
// 	partNumber := 1
// 	chunkSize := 5 * 1024 * 1024

// 	var parts []obs.Part

// 	// Upload file parts iteratively.
// 	for start := 0; start < buffer.Len(); start += chunkSize {
// 		end := start + chunkSize
// 		if end > buffer.Len() {
// 			end = buffer.Len()
// 		}

// 		partInput := &obs.UploadPartInput{
// 			Bucket:     config.BucketName,
// 			Key:        key,
// 			UploadId:   uploadID,
// 			PartNumber: partNumber,
// 			Body:       bytes.NewReader(buffer.Bytes()[start:end]),
// 		}

// 		// Upload the part and obtain ETag for verification.
// 		partOutput, err := obsClient.UploadPart(partInput)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Append part information to the list.
// 		parts = append(parts, obs.Part{PartNumber: partNumber, ETag: partOutput.ETag})
// 		partNumber++
// 	}

// 	// Prepare input parameters for completing the multipart upload.
// 	completeInput := &obs.CompleteMultipartUploadInput{
// 		Bucket:   config.BucketName,
// 		Key:      key,
// 		UploadId: uploadID,
// 		Parts:    parts,
// 	}

// 	// Complete the multipart upload.
// 	_, err = obsClient.CompleteMultipartUpload(completeInput)
// 	if err != nil {
// 		fmt.Printf("Error completing multipart upload: %v\n", err)
// 		return nil, err
// 	}

// 	// Generate the file URL based on the OBS bucket, endpoint, and key.
// 	fileURL := fmt.Sprintf("https://%s.%s/%s", config.BucketName, config.Endpoint, key)

// 	// Prepare the success response with file information.
// 	response := &FileUploadSuccessResponse{
// 		Filename: file.Filename,
// 		FileType: http.DetectContentType(buffer.Bytes()),
// 		Url:      fileURL,
// 	}

//		return response, nil
//	}
package file

import (
	"bytes"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

const (
	chunkSize = 40 * 1024 * 1024 // 40 MB per part
	workers   = 10               // Number of parts to upload concurrently
)

// func UploadFile(objDir string, file *multipart.FileHeader) (*FileUploadSuccessResponse, error) {
// 	obsClient, buffer, err := UploadFileHandler(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	key := objDir + "/" + file.Filename
// 	input := &obs.InitiateMultipartUploadInput{}
// 	input.Bucket = config.BucketName
// 	input.Key = key
// 	output, err := obsClient.InitiateMultipartUpload(input)
// 	if err != nil {
// 		return nil, err
// 	}
// 	uploadID := output.UploadId

// 	var parts []obs.Part
// 	var wg sync.WaitGroup
// 	partNumber := 1
// 	start := 0

// 	for start < buffer.Len() {
// 		end := start + chunkSize
// 		if end > buffer.Len() {
// 			end = buffer.Len()
// 		}

// 		wg.Add(1)
// 		go func(start, end, partNumber int) {
// 			defer wg.Done()

// 			partInput := &obs.UploadPartInput{
// 				Bucket:     config.BucketName,
// 				Key:        key,
// 				UploadId:   uploadID,
// 				PartNumber: partNumber,
// 				Body:       bytes.NewReader(buffer.Bytes()[start:end]),
// 			}

// 			partOutput, err := obsClient.UploadPart(partInput)
// 			if err != nil {
// 				fmt.Printf("Error uploading part %d: %v\n", partNumber, err)
// 				return
// 			}

// 			parts = append(parts, obs.Part{PartNumber: partNumber, ETag: partOutput.ETag})
// 		}(start, end, partNumber)

// 		start = end
// 		partNumber++

// 		// Limit the number of parallel uploads
// 		if partNumber%workers == 0 {
// 			wg.Wait()
// 		}
// 	}

// 	wg.Wait()
// 	completeInput := &obs.CompleteMultipartUploadInput{
// 		Bucket:   config.BucketName,
// 		Key:      key,
// 		UploadId: uploadID,
// 		Parts:    parts,
// 	}

// 	_, err = obsClient.CompleteMultipartUpload(completeInput)
// 	if err != nil {
// 		fmt.Printf("Error completing multipart upload: %v\n", err)
// 		return nil, err
// 	}

// 	fileURL := fmt.Sprintf("https://%s.%s/%s", config.BucketName, config.Endpoint, key)
// 	response := &FileUploadSuccessResponse{
// 		Filename: file.Filename,
// 		FileType: http.DetectContentType(buffer.Bytes()),
// 		Url:      fileURL,
// 	}

//		return response, nil
//	}
func UploadFile(objDir string, file *multipart.FileHeader) (*FileUploadSuccessResponse, error) {
	obsClient, buffer, err := UploadFileHandler(file)
	if err != nil {
		return nil, err
	}

	key := objDir + "/" + file.Filename

	// Check if file with the same name exists in OBS
	if exists, err := checkFileExists(obsClient, key); err != nil {
		return nil, err
	} else if exists {
		// Append random integer to the filename
		randomInt := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000)
		fileExtension := filepath.Ext(file.Filename)
		fileName := strings.TrimSuffix(file.Filename, fileExtension)

		fileName += "_" + strconv.Itoa(randomInt) + fileExtension
		key = objDir + "/" + fileName
	}

	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = config.BucketName
	input.Key = key
	output, err := obsClient.InitiateMultipartUpload(input)
	if err != nil {
		return nil, err
	}
	uploadID := output.UploadId

	var parts []obs.Part
	var wg sync.WaitGroup
	partNumber := 1
	start := 0

	for start < buffer.Len() {
		end := start + chunkSize
		if end > buffer.Len() {
			end = buffer.Len()
		}

		wg.Add(1)
		go func(start, end, partNumber int) {
			defer wg.Done()

			partInput := &obs.UploadPartInput{
				Bucket:     config.BucketName,
				Key:        key,
				UploadId:   uploadID,
				PartNumber: partNumber,
				Body:       bytes.NewReader(buffer.Bytes()[start:end]),
			}

			partOutput, err := obsClient.UploadPart(partInput)
			if err != nil {
				fmt.Printf("Error uploading part %d: %v\n", partNumber, err)
				return
			}

			parts = append(parts, obs.Part{PartNumber: partNumber, ETag: partOutput.ETag})
		}(start, end, partNumber)

		start = end
		partNumber++

		// Limit the number of parallel uploads
		if partNumber%workers == 0 {
			wg.Wait()
		}
	}

	wg.Wait()
	completeInput := &obs.CompleteMultipartUploadInput{
		Bucket:   config.BucketName,
		Key:      key,
		UploadId: uploadID,
		Parts:    parts,
	}

	_, err = obsClient.CompleteMultipartUpload(completeInput)
	if err != nil {
		fmt.Printf("Error completing multipart upload: %v\n", err)
		return nil, err
	}

	fileURL := fmt.Sprintf("https://%s.%s/%s", config.BucketName, config.Endpoint, key)
	response := &FileUploadSuccessResponse{
		Filename: file.Filename,
		FileType: http.DetectContentType(buffer.Bytes()),
		Url:      fileURL,
	}

	return response, nil
}

func checkFileExists(obsClient *obs.ObsClient, key string) (bool, error) {
	input := &obs.GetObjectMetadataInput{
		Bucket: config.BucketName,
		Key:    key,
	}

	_, err := obsClient.GetObjectMetadata(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode == http.StatusNotFound {
			// File does not exist
			return false, nil
		}
		return false, err
	}

	// File exists
	return true, nil
}
