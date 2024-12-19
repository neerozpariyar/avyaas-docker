package usecase

import (
	"avyaas/utils"
	"fmt"
	"net/http"
)

func (uCase *usecase) DeleteLive(id uint) error {
	// Check if the live meeting exists
	live, err := uCase.repo.GetLiveByID(id)
	if err != nil {
		return err

	}
	email, err := uCase.repo.GetEmailByID(id)
	if err != nil {
		return err
	}
	accessToken, err := utils.FetchZoomAccessToken(email)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.zoom.us/v2/meetings/%d", live.MeetingID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	if err := uCase.repo.DeleteLive(id); err != nil {
		return err
	}

	return nil
}
