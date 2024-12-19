package presenter

type FileListReq struct {
	Page     int
	PageSize int
	// BucketName string `json:"bucketName"` //if only bucket name provided  OBS returns descriptions for some or all objects (a maximum of 1000 objects) in the bucket
	// Prefix     string `json:"prefix"`     //Limits the response to object keys that begin with the specified prefix
	// Marker     string `json:"marker"`     //Indicates the object key to start with when listing objects in a bucket
	// MaxKeys    int    `json:"max-keys"`   //max number of object keys returned;ranges from 1 to 1000;default is 1000
	// Delimiter  string `json:"delimiter"`  //A character or a sequence of character used to group object keys

	Search   string `json:"search"`   //search for a specific file
	Service  string `json:"service"`  //Name of the service/folder
	IsActive bool   `json:"isActive"` //if true, only active files will be returned
}
type FileListRes struct {
	// ListBucketResult []string `xml:"ListBucketResult>ListBucketResult"`
	ID       uint   `json:"id"`
	Title    string `json:"title"`    //Name of the bucket
	IsActive bool   `json:"isActive"` //if true, only active files will be returned

	// Key  string `json:"key"`  //Name of an object
	// Size             int64    `json:"size"` //Size of an object
	URL string `json:"url"` //URL of an object
}
type DeleteObjectReq struct {
	Ids []uint `json:"ids"`
}
