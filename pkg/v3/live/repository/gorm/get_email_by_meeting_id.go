package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) GetEmailByMeetingID(meetingID int64) (string, error) {
	var live models.Live
	if err := repo.db.First(&live, "meeting_id=?", meetingID).Error; err != nil {
		return "", fmt.Errorf("failed to fetch email: %v", err)
	}
	return live.Email, nil
}
