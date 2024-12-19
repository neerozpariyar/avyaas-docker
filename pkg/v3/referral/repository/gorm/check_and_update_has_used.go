package gorm

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (repo *Repository) CheckAndUpdateHasUsed(referralCode string) error {
	// Find the referral for the given code
	var referral models.Referral
	if err := repo.db.
		Model(&models.Referral{}).
		Where("code = ?", referralCode).
		First(&referral).
		Error; err != nil {
		return err
	}
	hasUsed := true
	if referral.HasUsed == &hasUsed || (referral.HasLimit != nil && !*referral.HasLimit) {
		return fmt.Errorf("referral code has already been used or does not have a limit")
	}

	// Find the number of subscriptions for the given referral code
	var subscriptionCount int64
	if err := repo.db.
		Model(&models.Subscription{}).
		Where("referral_code = ?", referralCode). // add referral code in subscription model
		Count(&subscriptionCount).
		Error; err != nil {
		return err
	}

	// Check if the subscription count has reached the limit
	if referral.Limit > 0 && uint(subscriptionCount) >= referral.Limit {
		// Update hasUsed to true for the referral itself
		if err := repo.db.
			Model(&models.Referral{}).
			Where("code = ?", referralCode).
			Update("has_used", true).
			Error; err != nil {
			return err
		}
		return fmt.Errorf("referral code limit has been reached")
	}

	return nil
}
