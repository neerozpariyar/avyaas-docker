package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (uCase *usecase) UpdateLive(data models.Live) map[string]string {
	errMap := make(map[string]string)
	_, err := uCase.repo.GetLiveByID(data.ID)
	if err != nil {
		errMap["live_id_error"] = err.Error()
		return errMap
	}
	metID, err := uCase.repo.GetMeetingIDByLiveID(data.ID)
	if err != nil {
		errMap["meeting_id_error"] = err.Error()
		return errMap
	}

	accessToken, err := utils.FetchZoomAccessToken(data.Email)
	if err != nil {
		errMap["zoom_api_error"] = err.Error()
		return errMap
	}

	// Prepare request data for Zoom update
	requestData := struct {
		Topic      string    `json:"topic"`
		StartTime  time.Time `json:"start_time"`
		Duration   int       `json:"duration"`
		Type       int       `json:"type"`
		Recurrence struct {
			Type        int       `json:"type"`
			EndDateTime time.Time `json:"end_date_time"`
		} `json:"recurrence"`
	}{
		Topic:     data.Topic,
		StartTime: *data.StartTime,
		Duration:  data.Duration,
		Type:      int(data.Type),
	}

	// Set EndDateTime if Type is not 2
	if data.Type != 2 {
		requestData.Recurrence.Type = 1
		requestData.Recurrence.EndDateTime = *data.EndDateTime
	} else {
		requestData.Recurrence.Type = 1
		requestData.Recurrence.EndDateTime = data.StartTime.Add(time.Minute * time.Duration(data.Duration))
	}

	zoomRequestBody, err := json.Marshal(requestData)
	if err != nil {
		errMap["json_marshal_error"] = err.Error()
		return errMap
	}

	url := fmt.Sprintf("https://api.zoom.us/v2/meetings/%d", metID)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(zoomRequestBody))
	if err != nil {
		errMap["http_request_error"] = err.Error()
		return errMap
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errMap["zoom_api_error"] = err.Error()
		return errMap
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusNoContent {
		errMap["zoom_api_error"] = fmt.Sprintf("Unexpected response status: %s", resp.Status)
		// Print Zoom API response body for debugging
		responseBody, _ := io.ReadAll(resp.Body)
		errMap["zoom_api_response_body"] = string(responseBody)
		return errMap
	}

	if err := uCase.repo.UpdateLive(data); err != nil {

		errMap["database_error"] = err.Error()
		return errMap
	}

	return errMap
}
