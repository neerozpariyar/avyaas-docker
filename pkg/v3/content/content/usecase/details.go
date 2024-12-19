package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"errors"

	"gorm.io/gorm"
)

func (u *usecase) GetContentDetails(contentID, requesterID uint) (*presenter.ContentDetailResponse, map[string]string) {
	errMap := make(map[string]string)

	_, err := u.repo.GetContentByID(contentID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	contentDetails, err := u.repo.GetContentDetails(contentID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	requester, err := u.accountRepo.GetUserByID(requesterID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	// fullURL := utils.GetFileURL(contentDetails.Url)
	// Obtain a signed URL from the "file" package.
	signedURL, err := file.GetSignedURL(contentDetails.Url)
	if err != nil {
		errMap["error"] = err.Error()
	}

	// Save the signedURL for time being to check if the requester is student later
	if requester.RoleID != 4 {
		contentDetails.Url = signedURL
	} else {
		contentDetails.Url = ""
	}

	if requester.RoleID == 4 {
		encryptedSignedUrl, err := utils.GetEncryptedSignedUrlString(signedURL)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		studentContent, err := u.repo.CheckStudentContent(requesterID, contentID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		bookmark, isBookmarked, err := u.bookmarkRepo.GetBookmarkedContentAndCheckIfBookmarked(requesterID, contentID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		contentDetails.IsBookmarked = isBookmarked
		if isBookmarked {
			contentDetails.BookmarkID = bookmark.ID
		}

		if *studentContent.Paid {
			contentDetails.Url = encryptedSignedUrl
		}

		contentDetails.Paid = studentContent.Paid
	}

	if contentDetails.Note != nil {
		fullNoteURL := utils.GetFileURL(contentDetails.Note.File)
		signedNoteURL, err := file.GetSignedURL(fullNoteURL)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		contentDetails.Note.File = signedNoteURL

		if requester.RoleID == 4 {
			encryptedSignedNoteUrl, err := utils.GetEncryptedSignedUrlString(signedNoteURL)
			if err != nil {
				errMap["error"] = err.Error()
				return nil, errMap
			}

			studentContent, err := u.repo.CheckStudentContent(requesterID, contentID)
			if err != nil {
				errMap["error"] = err.Error()
				return nil, errMap
			}

			if *studentContent.Paid {
				contentDetails.Note.File = encryptedSignedNoteUrl
			}
			sContent, err := u.repo.GetContentProgressByContentID(contentID, requesterID)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					errMap["error"] = err.Error()
					return nil, errMap
				}
			}
			contentDetails.Progress = sContent.Progress
			contentDetails.HasCompleted = *sContent.HasCompleted
		}

	}

	return contentDetails, nil
}
