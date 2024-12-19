package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListChapter(data presenter.ChapterListRequest) ([]presenter.Chapter, int, error) {
	chm, totalPage, err := u.repo.ListChapter(data)
	if err != nil {
		return nil, int(totalPage), err
	}

	var chapters []presenter.Chapter
	for i := range chm {
		chapter := presenter.Chapter{
			ID:    chm[i].ID,
			Title: chm[i].Title,
			// Contents: chm[i].Contents,
			Position: int(chm[i].Position),
		}

		// unit, err := u.unitRepo.GetUnitByID(chm[i].UnitID)
		// if err != nil {
		// 	return nil, int(totalPage), err
		// }

		// unitData := make(map[string]interface{})
		// unitData["id"] = unit.ID
		// unitData["title"] = unit.Title
		// chapter.Unit = unitData

		chapters = append(chapters, chapter)
	}

	return chapters, int(totalPage), nil
}
