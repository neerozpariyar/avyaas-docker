package gorm

import (
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) GetSubjectHeirarchy(id uint, userID uint) ([]presenter.SubjectHeirarchyDetails, error) {

	var queryResults []presenter.SubjectHeirarchy
	var response []presenter.SubjectHeirarchyDetails

	relationPrimaryKeys, err := repo.GetRelationsBySubjectID(id)
	if err != nil {
		return nil, err
	}

	query := `SELECT 
	ucc.id AS unit_chapter_content_id,
	u.id AS unit_id, u.title AS unit_title,
	ch.id AS chapter_id, ch.title AS chapter_title,
	c.id AS content_id, c.title AS content_title, c.is_premium AS content_is_premium, c.content_type AS content_type, c.length AS content_length, sc.paid AS is_paid
	FROM unit_chapter_contents ucc
	LEFT JOIN 
		units u ON ucc.unit_id = u.id
	LEFT JOIN 
		chapters ch ON ucc.chapter_id = ch.id
	LEFT JOIN 
		contents c ON ucc.content_id = c.id
	LEFT JOIN
		student_contents sc ON sc.user_id = ? and sc.content_id = c.id
	WHERE 
		ucc.id IN (?);`

	err = repo.db.Raw(query, userID, relationPrimaryKeys).Scan(&queryResults).Error

	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {
		singleDetail := presenter.SubjectHeirarchyDetails{
			ID: queryResult.UnitChapterContentID,
			Unit: presenter.FilterDetail{
				ID:    queryResult.UnitID,
				Title: queryResult.UnitTitle,
			},
			Chapter: presenter.FilterDetail{
				ID:    queryResult.ChapterID,
				Title: queryResult.ChapterTitle,
			},
			Content: presenter.FilterDetail{
				ID:               queryResult.ContentID,
				Title:            queryResult.ContentTitle,
				ContentIsPremium: queryResult.ContentIsPremium,
				ContentType:      queryResult.ContentType,
				ContentLength:    queryResult.ContentLength,
				Paid:             queryResult.IsPaid,
			},
		}

		response = append(response, singleDetail)
	}
	return response, nil
}
