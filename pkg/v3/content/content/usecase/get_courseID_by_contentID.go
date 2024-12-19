package usecase

func (uCase *usecase) GetCourseIDByContentID(contentID uint) (uint, error) {

	return uCase.repo.GetCourseIDByContentID(contentID)
}
