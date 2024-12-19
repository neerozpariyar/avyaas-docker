package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListObjects(req *presenter.FileListReq) ([]presenter.FileListRes, int, error) {
	// Create a new OBS client

	// obsClient, err := obs.New(viper.GetString("huaweiOBS.accessKeyID"), viper.GetString("huaweiOBS.secretAccessKey"), viper.GetString("huaweiOBS.endpoint"))
	// if err != nil {
	// 	return nil, err
	// }

	// // Create the input parameters for the GET Bucket request
	// input := &obs.ListObjectsInput{
	// 	Bucket: viper.GetString("huaweiOBS.bucketName"),
	// }

	// // Set optional parameters if specified
	// // if req.Prefix != "" {
	// // 	input.Prefix = req.Prefix
	// // }
	// // if req.Marker != "" {
	// // 	input.Marker = req.Marker
	// // }
	// // if req.MaxKeys > 0 {
	// // 	input.MaxKeys = int(req.MaxKeys)
	// // }
	// // if req.Delimiter != "" {
	// // 	input.Delimiter = req.Delimiter
	// // }

	// // Send the GET Bucket request
	// output, err := obsClient.ListObjects(input)
	// if err != nil {
	// 	return nil, err
	// }
	// // Process the response and return the list of objects
	// var fileList []presenter.FileListRes
	// for _, object := range output.Contents {
	// 	fileList = append(fileList, presenter.FileListRes{
	// 		Name: object.Key,
	// 		URL:  fmt.Sprintf("https://%s/%s/%s", viper.GetString("huaweiOBS.endpoint"), viper.GetString("huaweiOBS.bucketName"), object.Key),
	// 	})
	// }
	files, totalPage, err := u.repo.ListObjects(req)
	if err != nil {
		return nil, int(totalPage), err
	}
	for i := range files { //appending the URL to the endpoint to get the whole URI
		files[i].URL = utils.GetFileURL(files[i].URL)

	}
	return files, int(totalPage), nil
}
