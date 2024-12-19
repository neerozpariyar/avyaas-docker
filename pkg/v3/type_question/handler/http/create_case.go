package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateCaseBasedQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		var requestBody presenter.TypeQuestionPresenter
		caseQuestionID, _ := strconv.Atoi(c.FormValue("caseQuestionID"))
		a := uint(caseQuestionID)
		requestBody.CaseQuestionID = &a
		requestBody.Title = c.FormValue("title")
		desc := c.FormValue("description")
		requestBody.Description = &desc
		requestBody.Type = "CaseBased"
		subID, _ := strconv.Atoi(c.FormValue("subjectID"))
		requestBody.SubjectID = uint(subID)
		forTest, _ := strconv.ParseBool(c.FormValue("forTest"))
		requestBody.ForTest = &forTest
		qID, _ := strconv.ParseUint(c.FormValue("questionSetID"), 10, 64)
		requestBody.QuestionSetID = uint(qID)

		requestBody.Image, _ = c.FormFile("image")
		requestBody.Audio, _ = c.FormFile("audio")

		if requestBody.Image != nil {
			fileType := utils.GetFileType(requestBody.Image.Filename)
			if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
				errMap["fileType"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png allowed", fileType).Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

		}

		if requestBody.Audio != nil {
			fileType := utils.GetFileType(requestBody.Audio.Filename)
			if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
				errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
		}

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Loop over the questions
		for q := 0; ; q++ {

			// Check if the question exists in the form data
			nestedQuestionType := c.FormValue(fmt.Sprintf("questions.%d.type", q))
			if nestedQuestionType == "" || nestedQuestionType == "CaseBased" {
				break
			}
			var question presenter.TypeQuestionPresenter
			fTest, _ := strconv.ParseBool(c.FormValue(fmt.Sprintf("questions.%d.forTest", q)))
			question.ForTest = &fTest
			nMark, _ := strconv.ParseFloat(c.FormValue(fmt.Sprintf("questions.%d.negativeMark", q)), 64)
			question.NegativeMark = &nMark

			nImg, _ := c.FormFile(fmt.Sprintf("questions.%d.image", q))
			nAud, _ := c.FormFile(fmt.Sprintf("questions.%d.audio", q))

			if nImg != nil {
				fileType := utils.GetFileType(nImg.Filename)
				if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
					errMap["fileType"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png allowed", fileType).Error()
					return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
				}
				question.Image = nImg

			}

			if nAud != nil {
				fileType := utils.GetFileType(nImg.Filename)
				if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
					errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
					return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
				}
				question.Audio = nAud

			}

			for i := 0; i < 4; i++ {

				var text string
				var correct bool
				var imageFile *multipart.FileHeader
				var audioFile *multipart.FileHeader
				correct, _ = strconv.ParseBool(c.FormValue(fmt.Sprintf("questions.%d.options.%d.isCorrect", q, i)))
				text = c.FormValue(fmt.Sprintf("questions.%d.options.%d.text", q, i))
				imageFile, _ = c.FormFile(fmt.Sprintf("questions.%d.options.%d.image", q, i))
				audioFile, _ = c.FormFile(fmt.Sprintf("questions.%d.options.%d.audio", q, i))

				var optionData presenter.TypeOptionPresenter

				optionData.Text = text
				optionData.IsCorrect = correct

				if imageFile != nil {
					fileType := utils.GetFileType(imageFile.Filename)
					if fileType != "png" && fileType != "jpg" && fileType != "jpeg" { //validate file type before setting
						errMap["fileType"] = fmt.Errorf("file type of %v not allowed: only IMAGE type: jpeg, jpg & png allowed", fileType).Error()
						return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
					}

					optionData.Image = imageFile
				}

				if audioFile != nil {

					fileType := utils.GetFileType(audioFile.Filename)
					fmt.Printf("fileType: %v\n", fileType)
					if fileType != "mpeg" && fileType != "mp3" { //validate file type before setting
						errMap["file_type"] = fmt.Errorf("file type of %v not allowed: only AUDIO type:mpeg & mp3 allowed", fileType).Error()
						return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
					}

					optionData.Audio = audioFile
				}

				question.Options = append(question.Options, optionData)

			}
			question.Type = nestedQuestionType
			requestBody.Questions = append(requestBody.Questions, question)

		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// if len(requestBody.Options) < 4 {
		// 	errMap["error"] = fmt.Sprintf("Not enough options, provided options: %d", len(requestBody.Options))
		// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// }
		// Validate question type
		if requestBody.Description == nil {
			errMap["caseDescription"] = "Case description is required for CaseBased questions"
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Invoke the usecase for creation of question
		errMap = handler.usecase.CreateTypeQuestion(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))

		}

		return c.JSON(presenter.SuccessResponse())
	}
}
