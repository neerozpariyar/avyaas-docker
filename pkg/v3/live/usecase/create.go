package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (uCase *usecase) CreateLive(data models.Live) map[string]string {
	errMap := make(map[string]string)

	// Check if there's an existing meeting with conflicting times for the given email
	if data.Type == 2 {
		endDateTime := data.StartTime.Add(time.Minute * time.Duration(data.Duration))
		data.EndDateTime = &endDateTime
	}
	conflictingMeeting, err := uCase.repo.GetConflictingMeeting(data.Email, *data.StartTime, *data.EndDateTime)
	if err != nil {
		errMap["database_error"] = err.Error()
		return errMap
	}
	if len(conflictingMeeting) != 0 {
		errMap["conflicting_meeting"] = "A meeting with conflicting times already exists for the given email"
		return errMap
	}

	// Fetch Zoom access token
	accessToken, err := utils.FetchZoomAccessToken(data.Email)
	if err != nil {
		errMap["zoom_api_error"] = err.Error()
		return errMap
	}

	// Construct request body for creating a meeting
	requestData := struct {
		models.Live
		Settings struct {
			AutoRecording string `json:"auto_recording"`
		} `json:"settings"`
		Recurrence struct {
			Type        int       `json:"type"`
			EndDateTime time.Time `json:"end_date_time"`
		} `json:"recurrence"`
	}{
		Live: data,
		Settings: struct {
			AutoRecording string `json:"auto_recording"`
		}{AutoRecording: "cloud"},
		Recurrence: struct {
			Type        int       `json:"type"`
			EndDateTime time.Time `json:"end_date_time"`
		}{
			Type:        1,
			EndDateTime: *data.EndDateTime},
	}

	zoomRequestBody, err := json.Marshal(requestData)
	if err != nil {
		errMap["json_marshal_error"] = err.Error()
		return errMap
	}

	// Send request to Zoom API to create meeting
	zoomAPIURL := "https://api.zoom.us/v2/users/me/meetings"
	zoomRequest, err := http.NewRequest("POST", zoomAPIURL, bytes.NewBuffer(zoomRequestBody))
	if err != nil {
		errMap["http_request_error"] = err.Error()
		return errMap
	}
	zoomRequest.Header.Set("Content-Type", "application/json")
	zoomRequest.Header.Set("Authorization", "Bearer "+accessToken)

	zoomResponse, err := http.DefaultClient.Do(zoomRequest)
	if err != nil {
		errMap["zoom_api_error"] = err.Error()
		return errMap
	}
	defer zoomResponse.Body.Close()

	// Handle Zoom API response
	if zoomResponse.StatusCode != 201 {
		responseBody, err := io.ReadAll(zoomResponse.Body)
		if err != nil {
			errMap["read_response_error"] = err.Error()
			return errMap
		}
		errMap["zoom_api_error"] = string(responseBody)
		return errMap
	}

	// Extract meeting details from Zoom API response
	var zoomMeeting struct {
		ID       int    `json:"id"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(zoomResponse.Body).Decode(&zoomMeeting); err != nil {
		errMap["json_decode_error"] = err.Error()
		return errMap
	}

	// Save meeting details to local database
	data.MeetingID = zoomMeeting.ID
	data.MeetingPwd = zoomMeeting.Password
	if err := uCase.repo.CreateLive(data); err != nil {
		errMap["database_error"] = err.Error()
		return errMap
	}

	return errMap
}
