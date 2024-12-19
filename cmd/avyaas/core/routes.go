package core

import (
	_accountHandler "avyaas/pkg/v3/account/handler/http"
	_authHandler "avyaas/pkg/v3/auth/handler/http"
	_bookmarkHandler "avyaas/pkg/v3/bookmark/handler/http"
	_chapterHandler "avyaas/pkg/v3/content/chapter/handler/http"
	_contentHandler "avyaas/pkg/v3/content/content/handler/http"
	_courseHandler "avyaas/pkg/v3/content/course/handler/http"
	_courseGroupHandler "avyaas/pkg/v3/content/course_group/handler/http"
	_noteHandler "avyaas/pkg/v3/content/note/handler/http"
	_subjectHandler "avyaas/pkg/v3/content/subject/handler/http"
	_unitHandler "avyaas/pkg/v3/content/unit/handler/http"
	_discussionHandler "avyaas/pkg/v3/discussion/handler/http"
	_feedbackHandler "avyaas/pkg/v3/feedback/handler/http"
	_fileClientHandler "avyaas/pkg/v3/fileClient/handler/http"
	_liveHandler "avyaas/pkg/v3/live/handler/http"
	_liveGroupHandler "avyaas/pkg/v3/liveGroup/handler/http"
	_notificationHandler "avyaas/pkg/v3/notification/handler/http"
	_packageHandler "avyaas/pkg/v3/package/handler/http"
	_packageTypeHandler "avyaas/pkg/v3/package_type/handler/http"
	_paymentHandler "avyaas/pkg/v3/payment/handler/http"
	_pollHandler "avyaas/pkg/v3/poll/handler/http"
	_questionHandler "avyaas/pkg/v3/question/handler/http"
	_questionSetHandler "avyaas/pkg/v3/question_set/handler/http"
	_recordingHandler "avyaas/pkg/v3/recording/handler/http"
	_referralHandler "avyaas/pkg/v3/referral/handler/http"
	_replyHandler "avyaas/pkg/v3/reply/handler/http"
	_serviceHandler "avyaas/pkg/v3/service/handler/http"
	_testHandler "avyaas/pkg/v3/test/handler/http"
	_testSeriesHandler "avyaas/pkg/v3/test_series/handler/http"

	"github.com/gofiber/fiber/v2"
)

/*
initRoutes initializes routes for the all endpoints on the provided Fiber router. It sets up the
necessary handlers and routes for handling operations.
*/
func InitRoutes(authRoute fiber.Router, protectedRoutes fiber.Router, usecases *Usecases) {
	// Initialize routes for the server
	_authHandler.New(authRoute, usecases.authUsecase)
	_accountHandler.New(protectedRoutes, usecases.accountUsecase)
	_courseGroupHandler.New(protectedRoutes, usecases.courseGroupUsecase)
	_courseHandler.New(protectedRoutes, usecases.courseUsecase)
	_subjectHandler.New(protectedRoutes, usecases.subjectUsecase)
	_unitHandler.New(protectedRoutes, usecases.unitUsecase)
	_chapterHandler.New(protectedRoutes, usecases.chapterUsecase)
	_contentHandler.New(protectedRoutes, usecases.contentUsecase)
	_testHandler.New(protectedRoutes, usecases.testUsecase)
	_questionSetHandler.New(protectedRoutes, usecases.questionSetUsecase)
	_questionHandler.New(protectedRoutes, usecases.questionUsecase)
	_noteHandler.New(protectedRoutes, usecases.noteUsecase)
	_discussionHandler.New(protectedRoutes, usecases.discussionUsecase)
	_replyHandler.New(protectedRoutes, usecases.replyUsecase)
	_pollHandler.New(protectedRoutes, usecases.pollUsecase)
	_packageTypeHandler.New(protectedRoutes, usecases.packageTypeUsecase)
	_packageHandler.New(protectedRoutes, usecases.packageUsecase, usecases.packageTypeUsecase)
	_serviceHandler.New(protectedRoutes, usecases.serviceUsecase)
	_referralHandler.New(protectedRoutes, usecases.referralUsecase)
	_paymentHandler.New(protectedRoutes, usecases.paymentUsecase)
	_liveGroupHandler.New(protectedRoutes, usecases.liveGroupUsecase)
	_liveHandler.New(protectedRoutes, usecases.liveUsecase)
	_recordingHandler.New(protectedRoutes, usecases.recordingUsecase)
	_feedbackHandler.New(protectedRoutes, usecases.feedbackUsecase)
	_notificationHandler.New(protectedRoutes, usecases.notificationUsecase)
	_testSeriesHandler.New(protectedRoutes, usecases.testSeriesUsecase)
	_fileClientHandler.New(protectedRoutes, usecases.fileClientUsecase)
	// _typeQuestionHandler.New(protectedRoutes, usecases.typeQuestionUsecase)
	_bookmarkHandler.New(protectedRoutes, usecases.bookmarkUsecase)
}
