package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Vote struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	UserID       uint       `json:"userID"`
	DiscussionID uint       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"discussionID"`
	Discussion   Discussion ` json:"-"`
	HasLiked     bool       `json:"hasLiked" gorm:"default:false"`
}

func (m *Vote) AfterSave(tx *gorm.DB) (err error) {
	if m.Discussion.ID != m.DiscussionID {
		m.Discussion.ID = m.DiscussionID

	}
	err = tx.Clauses(clause.Returning{}).Debug().Exec(`update discussions set vote_count =  
	(select COUNT(id) from votes where (votes.discussion_id=discussions.Id and votes.has_liked=1))`).Error
	if err != nil {
		return err
	}
	return err

}
