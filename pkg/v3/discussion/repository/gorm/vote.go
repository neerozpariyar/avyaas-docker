package gorm

import (
	"avyaas/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

// LikeOrUnlikeDiscussion allows a user to like or unlike a vote.
// func (repo *repository) LikeOrUnlikeDiscussion(discussionID, userID uint) error {
// 	var vote models.Vote
// 	var discussion models.Discussion
// 	fmt.Printf("REEEEEEEEEEEPPPPPOOOOOOOOOdiscussionID: %v\n", discussionID)
// 	fmt.Printf("RepouserrrrrrrrrID: %v\n", userID)

// 	if err := repo.db.First(&vote, userID).Error; err != nil {
// 		return err

// 	}
// 	if vote.Vote {

// 		vote.Vote = false
// 		discussion.VoteCount--
// 	} else {
// 		vote.Vote = true
// 		discussion.VoteCount++
// 	}

// 	if err := repo.db.Save(&vote).Error; err != nil {
// 		return err
// 	}

//		return nil
//	}
func (repo *Repository) LikeOrUnlikeDiscussion(discussionID, userID uint) error {
	var vote models.Vote
	// var discussion models.Discussion
	// Check if the user has already voted
	if err := repo.db.Where("user_id = ? AND discussion_id = ?", userID, discussionID).First(&vote).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newVote := models.Vote{
				UserID:       userID,
				DiscussionID: discussionID,
				HasLiked:     true,
			}

			if err := repo.db.Create(&newVote).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// Vote toggle

	if err := repo.db.Save(&models.Vote{
		ID:           vote.ID,
		UserID:       userID,
		DiscussionID: discussionID,
		HasLiked:     !vote.HasLiked,
	}).Error; err != nil {
		return err
	}

	// if err := repo.db.Debug().Exec(`update discussions set vote_count =
	// (select COUNT(id) from votes where votes.discussion_id=discussions.Id)`).Model(&models.Discussion{}).Error; err != nil {
	// 	return err
	// }

	return nil
}
