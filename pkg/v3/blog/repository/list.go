package repository

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (repo *Repository) ListBlog(requestBody presenter.BlogListReq) ([]models.Blog, float64, error) {
	var blogs []models.Blog

	baseQuery := repo.db.Debug().Model(&models.Blog{})

	if requestBody.CourseID != 0 {
		baseQuery = baseQuery.Where("course_id = ?", requestBody.CourseID)
		if requestBody.Search != "" {
			baseQuery = baseQuery.Where("title like ? OR course_id like ?", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%")
		}
	}

	if requestBody.SubjectID != 0 {
		baseQuery = baseQuery.Where("subject_id = ?", requestBody.SubjectID)
		if requestBody.Search != "" {
			baseQuery = baseQuery.Where("title like ? OR subject_id like ?", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%")
		}
	}

	if requestBody.Search != "" {
		baseQuery = baseQuery.Where("title like ? OR course_id like ? OR subject_id like?", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%", "%"+requestBody.Search+"%")
	}

	totalPage := utils.GetTotalPage(models.Blog{}, baseQuery, requestBody.PageSize)

	if totalPage < float64(requestBody.Page) {
		requestBody.Page = int(totalPage)
	}
	err := baseQuery.Scopes(utils.Paginate(requestBody.Page, requestBody.PageSize)).Find(&blogs).Order("id").Error

	return blogs, totalPage, err
}
