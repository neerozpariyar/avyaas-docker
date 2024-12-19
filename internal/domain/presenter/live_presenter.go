package presenter

type ListLiveRequest struct {
	UserID      uint   `json:"userID"`
	CourseID    uint   `json:"courseID"`
	LiveGroupID uint   `json:"liveGroupID"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	Search      string `json:"search"`
}

type LiveListResponse struct {
	ID          uint        `json:"id"`
	Topic       string      `json:"topic"`
	LiveGroup   interface{} `json:"liveGroup"`
	Course      interface{} `json:"course"`
	Subject     interface{} `json:"subject"`
	Type        uint        `json:"type"` //only 2 and 8 : scheduled meeting and recurring meeting, given by zoom
	StartTime   string      `json:"start_time"`
	EndDateTime string      `json:"endDateTime"` //only in recurring live
	Duration    int         `json:"duration"`
	IsLive      *bool       `json:"isLive"` //if only true live can be started
	MeetingID   int         `json:"meetingID"`
	MeetingPwd  string      `json:"password"`
	Email       string      `json:"email"`
	IsFree      *bool       `json:"isFree"`
	IsPackage   *bool       `json:"isPackage"`
	Price       int         `json:"price"`
}
