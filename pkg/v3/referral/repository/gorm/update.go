package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateReferral(referral models.Referral) error {
	return repo.db.Debug().Omit("created_at").Save(&models.Referral{
		Timestamp: models.Timestamp{
			ID: referral.ID,
		},
		Title:        referral.Title,
		CourseID:     referral.CourseID,
		UserID:       referral.UserID,
		Code:         referral.Code,
		Type:         referral.Type,
		DiscountType: referral.DiscountType,
		Discount:     referral.Discount,
		HasLimit:     referral.HasLimit,
		HasUsed:      referral.HasUsed,
		Limit:        referral.Limit,
	}).Error
}
