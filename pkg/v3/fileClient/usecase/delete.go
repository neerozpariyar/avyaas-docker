package usecase

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/spf13/viper"
)

func (uCase *usecase) DeleteObjects(id []uint) error {
	// Create a new OBS client
	obsClient, err := obs.New(viper.GetString("huaweiOBS.accessKeyID"), viper.GetString("huaweiOBS.secretAccessKey"), viper.GetString("huaweiOBS.endpoint"))
	if err != nil {
		return err
	}
	// Fetch the URL from the given ID and set it as the object key
	objectKey, err := uCase.repo.GetURLsByID(id)
	if err != nil {
		return err
	}
	// Create the input parameters for the DELETE Object request
	input := &obs.DeleteObjectsInput{
		Bucket: viper.GetString("huaweiOBS.bucketName"),
		Objects: func() []obs.ObjectToDelete {
			var objects []obs.ObjectToDelete
			for _, key := range objectKey {
				objects = append(objects, obs.ObjectToDelete{Key: key, VersionId: ""})
			}
			return objects
		}(),
	}

	// Send the DELETE Object request
	deletedObjects, err := obsClient.DeleteObjects(input)
	if err != nil {
		return err
	}
	// Print the deleted objects
	for _, deletedObject := range deletedObjects.Deleteds {
		fmt.Printf("Deleted object: %s\n", deletedObject.Key)
	}
	if _, err := uCase.repo.GetObjectsByID(id); err != nil {
		return err
	}

	return uCase.repo.DeleteObjects(id)
}
