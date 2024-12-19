package core

import (
	interfaces "avyaas/internal/domain/interfaces"
	_accountRepo "avyaas/pkg/v3/account/repository/gorm"
	_authRepo "avyaas/pkg/v3/auth/repository/gorm"
	_blogRepo "avyaas/pkg/v3/blog/repository"
	_bookmarkRepo "avyaas/pkg/v3/bookmark/repository/gorm"
	_chapterRepo "avyaas/pkg/v3/content/chapter/repository/gorm"
	_contentRepo "avyaas/pkg/v3/content/content/repository/gorm"
	_commentRepo "avyaas/pkg/v3/content/content_comment/repository/gorm"
	_courseRepo "avyaas/pkg/v3/content/course/repository/gorm"
	_courseGroupRepo "avyaas/pkg/v3/content/course_group/repository/gorm"
	_noteRepo "avyaas/pkg/v3/content/note/repository/gorm"
	_subjectRepo "avyaas/pkg/v3/content/subject/repository/gorm"
	_unitRepo "avyaas/pkg/v3/content/unit/repository/gorm"
	_discussionRepo "avyaas/pkg/v3/discussion/repository/gorm"
	_faqRepo "avyaas/pkg/v3/faq/repository"
	_feedbackRepo "avyaas/pkg/v3/feedback/repository/gorm"
	_fileClientRepo "avyaas/pkg/v3/fileClient/repository/gorm"
	_liveRepo "avyaas/pkg/v3/live/repository/gorm"
	_liveGroupRepo "avyaas/pkg/v3/liveGroup/repository/gorm"
	_noticeRepo "avyaas/pkg/v3/notice/repository"
	_notificationRepo "avyaas/pkg/v3/notification/repository/gorm"
	_packageRepo "avyaas/pkg/v3/package/repository/gorm"
	_packageTypeRepo "avyaas/pkg/v3/package_type/repository/gorm"
	_paymentRepo "avyaas/pkg/v3/payment/repository/gorm"
	_pollRepo "avyaas/pkg/v3/poll/repository/gorm"
	_questionRepo "avyaas/pkg/v3/question/repository/gorm"
	_questionSetRepo "avyaas/pkg/v3/question_set/repository/gorm"
	_recordingRepo "avyaas/pkg/v3/recording/repository/gorm"
	_referralRepo "avyaas/pkg/v3/referral/repository/gorm"
	_replyRepo "avyaas/pkg/v3/reply/repository/gorm"
	_serviceRepo "avyaas/pkg/v3/service/repository/gorm"
	_subscriptionRepo "avyaas/pkg/v3/subscription/repository/gorm"
	_termsAndConditionRepo "avyaas/pkg/v3/terms_condition/repository"
	_testRepo "avyaas/pkg/v3/test/repository/gorm"
	_testSeriesRepo "avyaas/pkg/v3/test_series/repository/gorm"
)

// usecases struct represents a collection of use case dependencies
type Usecases struct {
	authUsecase              interfaces.AuthUsecase
	accountUsecase           interfaces.AccountUsecase
	courseGroupUsecase       interfaces.CourseGroupUsecase
	testUsecase              interfaces.TestUsecase
	questionSetUsecase       interfaces.QuestionSetUsecase
	questionUsecase          interfaces.QuestionUsecase
	courseUsecase            interfaces.CourseUsecase
	subjectUsecase           interfaces.SubjectUsecase
	unitUsecase              interfaces.UnitUsecase
	chapterUsecase           interfaces.ChapterUsecase
	contentUsecase           interfaces.ContentUsecase
	noteUsecase              interfaces.NoteUsecase
	discussionUsecase        interfaces.DiscussionUsecase
	replyUsecase             interfaces.ReplyUsecase
	pollUsecase              interfaces.PollUsecase
	packageTypeUsecase       interfaces.PackageTypeUsecase
	packageUsecase           interfaces.PackageUsecase
	serviceUsecase           interfaces.ServiceUsecase
	referralUsecase          interfaces.ReferralUsecase
	paymentUsecase           interfaces.PaymentUsecase
	liveGroupUsecase         interfaces.LiveGroupUsecase
	liveUsecase              interfaces.LiveUsecase
	recordingUsecase         interfaces.RecordingUsecase
	feedbackUsecase          interfaces.FeedbackUsecase
	notificationUsecase      interfaces.NotificationUsecase
	testSeriesUsecase        interfaces.TestSeriesUsecase
	fileClientUsecase        interfaces.FileClientUsecase
	bookmarkUsecase          interfaces.BookmarkUsecase
	commentUsecase           interfaces.ContentCommentUsecase
	subscriptionUsecase      interfaces.SubscriptionUsecase
	noticeUsecase            interfaces.NoticeUsecase
	faqUsecase               interfaces.FAQUsecase
	termsAndConditionUsecase interfaces.TermsAndConditionUsecase
	blogUsecase              interfaces.BlogUsecase
	// typeQuestionUsecase      interfaces.TypeQuestionUsecases
}

type Repositories struct {
	courseGroupRepo       *_courseGroupRepo.Repository
	subjectRepo           *_subjectRepo.Repository
	unitRepo              *_unitRepo.Repository
	chapterRepo           *_chapterRepo.Repository
	contentRepo           *_contentRepo.Repository
	noteRepo              *_noteRepo.Repository
	replyRepo             *_replyRepo.Repository
	pollRepo              *_pollRepo.Repository
	packageTypeRepo       *_packageTypeRepo.Repository
	packageRepo           *_packageRepo.Repository
	serviceRepo           *_serviceRepo.Repository
	referralRepo          *_referralRepo.Repository
	paymentRepo           *_paymentRepo.Repository
	feedbackRepo          *_feedbackRepo.Repository
	notificationRepo      *_notificationRepo.Repository
	fileClientRepo        *_fileClientRepo.Repository
	liveGroupRepo         *_liveGroupRepo.Repository
	liveRepo              *_liveRepo.Repository
	recordingRepo         *_recordingRepo.Repository
	questionRepo          *_questionRepo.Repository
	questionSetRepo       *_questionSetRepo.Repository
	accountRepo           *_accountRepo.Repository
	discussionRepo        *_discussionRepo.Repository
	courseRepo            *_courseRepo.Repository
	testRepo              *_testRepo.Repository
	testSeriesRepo        *_testSeriesRepo.Repository
	authRepo              *_authRepo.Repository
	bookmarkRepo          *_bookmarkRepo.Repository
	commentRepo           *_commentRepo.Repository
	subscriptionRepo      *_subscriptionRepo.Repository
	noticeRepo            *_noticeRepo.Repository
	faqRepo               *_faqRepo.Repository
	termsAndConditionRepo *_termsAndConditionRepo.Repository
	blogRepo              *_blogRepo.Repository
	// typeQuestionRepo      *_tQuestionRepo.Repository
}
