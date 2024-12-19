package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (uCase *usecase) ListBlog(res presenter.BlogListReq) ([]presenter.BlogListRes, int, error) {
	blogs, totalPage, err := uCase.repo.ListBlog(res)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allBlogs []presenter.BlogListRes
	for _, blog := range blogs {
		eachBlog := &presenter.BlogListRes{
			ID:    blog.ID,
			Title: blog.Title,
			Tags:  blog.Tags,
		}
		if blog.Cover != "" {
			url := utils.GetFileURL(blog.Cover)
			eachBlog.Cover = url

		}
		if blog.CourseID != 0 {
			course, err := uCase.courseRepo.GetCourseByID(blog.CourseID)
			if err != nil {
				return nil, 0, err
			}

			courseData := make(map[string]interface{})
			courseData["id"] = course.ID
			courseData["courseID"] = course.CourseID

			eachBlog.Course = courseData
		}

		if blog.SubjectID != 0 {
			subject, err := uCase.subjectRepo.GetSubjectByID(blog.SubjectID)
			if err != nil {
				return nil, 0, err
			}

			subjectData := make(map[string]interface{})
			subjectData["id"] = subject.ID
			subjectData["subjectId"] = subject.SubjectID

			eachBlog.Subject = subjectData
		}

		allBlogs = append(allBlogs, *eachBlog)

	}
	return allBlogs, int(totalPage), err
}
