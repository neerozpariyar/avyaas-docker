package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

/*
this function validates and maps the Questions and options
Validation includes:

 1. Image and audio extension checking
    2.Counts correct options for mcq(1), more than one for multiple choice questions

 3. checks if true or false question's answer is given

 4. checks that fill in the blanks only has text as an option

 5. Invokes create/update casebased questions

    Mapping includes:

 1. Mapping the data need for the option
*/
func (handler *handler) ValidateAndMapQuestionOptions(requestBody *presenter.CreateUpdateQuestionRequest, isCreate bool, c *fiber.Ctx) (*presenter.CreateUpdateQuestionRequest, map[string]string) {
	errMap := make(map[string]string)

	// Parse the image file, if available
	imageFile, _ := c.FormFile("image")
	if imageFile != nil {
		if !utils.IsValidImageFile(imageFile) {
			errMap["image"] = "Invalid image file type. Only jpeg, jpg & png allowed."
			return &presenter.CreateUpdateQuestionRequest{}, errMap
		}
		requestBody.Image = imageFile
	}

	// Parse the audio file, if available
	audioFile, _ := c.FormFile("audio")
	if audioFile != nil {
		if !utils.IsValidAudioFile(audioFile) {
			errMap["audio"] = "Invalid audio file type. Only mpeg & mp3 allowed."
			return &presenter.CreateUpdateQuestionRequest{}, errMap
		}
		requestBody.Audio = audioFile
	}

	// Validation specific to question types
	switch requestBody.Type {
	case "MCQ":
		if utils.CountCorrectOptions(requestBody.Options) != 1 {
			errMap["error"] = "MCQ should have exactly one correct answer"
			return &presenter.CreateUpdateQuestionRequest{}, errMap
		}
	case "TrueFalse":
		if requestBody.IsTrue == nil {
			errMap["error"] = "TrueFalse question should have isTrue field"
			return &presenter.CreateUpdateQuestionRequest{}, errMap
		}
	case "FillInTheBlanks":
		if utils.HasMediaOptions(requestBody.Options) {
			errMap["options"] = "FillInTheBlanks questions must have options of text only"
			return &presenter.CreateUpdateQuestionRequest{}, errMap
		}
	case "MultiAnswer":
		if utils.CountCorrectOptions(requestBody.Options) < 2 {
			errMap["error"] = "MultiAnswer should have at least two correct answers"
			return &presenter.CreateUpdateQuestionRequest{}, errMap
		}
	case "CaseBased":
		if isCreate {
			err := handler.CreateCaseBasedQuestion()(c)
			if err != nil {
				errMap["error"] = err.Error()
				return &presenter.CreateUpdateQuestionRequest{}, errMap
			}
		} else {
			err := handler.UpdateCaseBasedQuestion()(c)
			if err != nil {
				errMap["error"] = err.Error()
				return &presenter.CreateUpdateQuestionRequest{}, errMap
			}
		}

	default:
		errMap["question_type"] = fmt.Sprintf("Invalid question type: %s", requestBody.Type)
		return &presenter.CreateUpdateQuestionRequest{}, errMap
	}

	return requestBody, nil
}
