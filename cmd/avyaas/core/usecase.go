package core

import (
	"avyaas/internal/config"
	_accountUsecase "avyaas/pkg/v3/account/usecase"
	_authUsecase "avyaas/pkg/v3/auth/usecase"
	_blogUsecase "avyaas/pkg/v3/blog/usecase"
	_bookmarkUsecase "avyaas/pkg/v3/bookmark/usecase"
	_chapterUsecase "avyaas/pkg/v3/content/chapter/usecase"
	_contentUsecase "avyaas/pkg/v3/content/content/usecase"
	_commentUsecase "avyaas/pkg/v3/content/content_comment/usecase"
	_courseUsecase "avyaas/pkg/v3/content/course/usecase"
	_courseGroupUsecase "avyaas/pkg/v3/content/course_group/usecase"
	_noteUsecase "avyaas/pkg/v3/content/note/usecase"
	_subjectUsecase "avyaas/pkg/v3/content/subject/usecase"
	_unitUsecase "avyaas/pkg/v3/content/unit/usecase"
	_discussionUsecase "avyaas/pkg/v3/discussion/usecase"
	_faqUsecase "avyaas/pkg/v3/faq/usecase"
	_feedbackUsecase "avyaas/pkg/v3/feedback/usecase"
	_fileClientUsecase "avyaas/pkg/v3/fileClient/usecase"
	_liveUsecase "avyaas/pkg/v3/live/usecase"
	_liveGroupUsecase "avyaas/pkg/v3/liveGroup/usecase"
	_noticeUsecase "avyaas/pkg/v3/notice/usecase"
	_notificationUsecase "avyaas/pkg/v3/notification/usecase"
	_packageUsecase "avyaas/pkg/v3/package/usecase"
	_packageTypeUsecase "avyaas/pkg/v3/package_type/usecase"
	_paymentUsecase "avyaas/pkg/v3/payment/usecase"
	_pollUsecase "avyaas/pkg/v3/poll/usecase"
	_questionUsecase "avyaas/pkg/v3/question/usecase"
	_questionSetUsecase "avyaas/pkg/v3/question_set/usecase"
	_recordingUsecase "avyaas/pkg/v3/recording/usecase"
	_referralUsecase "avyaas/pkg/v3/referral/usecase"
	_replyUsecase "avyaas/pkg/v3/reply/usecase"
	_serviceUsecase "avyaas/pkg/v3/service/usecase"
	_subscriptionUsecase "avyaas/pkg/v3/subscription/usecase"
	_termsConditionUsecase "avyaas/pkg/v3/terms_condition/usecase"
	_testUsecase "avyaas/pkg/v3/test/usecase"
	_testSeriesUsecase "avyaas/pkg/v3/test_series/usecase"
)

func InitUsecases(server *config.Server, repo Repositories) *Usecases {
	allUsecase := &Usecases{
		authUsecase:         _authUsecase.New(repo.authRepo, repo.accountRepo, repo.courseRepo, server.RedisClient),
		accountUsecase:      _accountUsecase.New(repo.accountRepo, repo.courseRepo, repo.subjectRepo, repo.packageRepo),
		courseGroupUsecase:  _courseGroupUsecase.New(repo.courseGroupRepo, repo.courseRepo),
		courseUsecase:       _courseUsecase.New(repo.courseRepo, repo.courseGroupRepo, repo.packageRepo, repo.paymentRepo, repo.accountRepo, repo.subjectRepo),
		subjectUsecase:      _subjectUsecase.New(repo.subjectRepo, repo.courseRepo, repo.accountRepo, repo.contentRepo, repo.unitRepo),
		unitUsecase:         _unitUsecase.New(repo.unitRepo, repo.courseRepo, repo.subjectRepo, repo.chapterRepo),
		chapterUsecase:      _chapterUsecase.New(repo.chapterRepo, repo.unitRepo, repo.contentRepo),
		contentUsecase:      _contentUsecase.New(repo.contentRepo, repo.courseRepo, repo.chapterRepo, repo.accountRepo, repo.bookmarkRepo, repo.subjectRepo),
		noteUsecase:         _noteUsecase.New(repo.noteRepo, repo.contentRepo),
		discussionUsecase:   _discussionUsecase.New(repo.discussionRepo, repo.courseRepo, repo.subjectRepo, repo.accountRepo),
		replyUsecase:        _replyUsecase.New(repo.replyRepo, repo.courseRepo, repo.discussionRepo, repo.accountRepo),
		pollUsecase:         _pollUsecase.New(repo.pollRepo, repo.accountRepo, repo.courseRepo, repo.subjectRepo),
		referralUsecase:     _referralUsecase.New(repo.referralRepo, repo.accountRepo, repo.courseRepo, repo.packageRepo),
		questionSetUsecase:  _questionSetUsecase.New(repo.questionSetRepo, repo.accountRepo, repo.courseRepo, repo.subjectRepo, repo.questionRepo, repo.bookmarkRepo),
		questionUsecase:     _questionUsecase.New(repo.questionRepo, repo.courseRepo, repo.subjectRepo, repo.questionSetRepo, repo.bookmarkRepo),
		packageTypeUsecase:  _packageTypeUsecase.New(repo.packageTypeRepo, repo.serviceRepo),
		packageUsecase:      _packageUsecase.New(repo.packageRepo, repo.packageTypeRepo, repo.courseRepo, repo.serviceRepo, repo.testSeriesRepo, repo.testRepo, repo.liveGroupRepo, repo.liveRepo, repo.paymentRepo),
		serviceUsecase:      _serviceUsecase.New(repo.serviceRepo),
		liveGroupUsecase:    _liveGroupUsecase.New(repo.liveGroupRepo, repo.courseRepo, repo.packageTypeRepo),
		liveUsecase:         _liveUsecase.New(repo.liveRepo, repo.accountRepo, repo.liveGroupRepo, repo.courseRepo, repo.subjectRepo),
		recordingUsecase:    _recordingUsecase.New(repo.recordingRepo, repo.liveRepo),
		feedbackUsecase:     _feedbackUsecase.New(repo.feedbackRepo, repo.courseRepo),
		bookmarkUsecase:     _bookmarkUsecase.New(repo.bookmarkRepo, repo.contentRepo, repo.questionRepo),
		notificationUsecase: _notificationUsecase.New(repo.notificationRepo, repo.courseRepo, repo.accountRepo),
		testSeriesUsecase:   _testSeriesUsecase.New(repo.testSeriesRepo, repo.courseRepo, repo.packageTypeRepo, repo.packageRepo),
		fileClientUsecase:   _fileClientUsecase.New(repo.fileClientRepo),
		commentUsecase:      _commentUsecase.New(repo.commentRepo, repo.contentRepo, repo.accountRepo),
		subscriptionUsecase: _subscriptionUsecase.New(repo.subscriptionRepo, repo.courseRepo),
		// typeQuestionUsecase:      _typeQuestionUsecase.New(repo.typeQuestionRepo, repo.courseRepo, repo.subjectRepo, repo.questionSetRepo, repo.bookmarkRepo),
		noticeUsecase:            _noticeUsecase.New(repo.noticeRepo, repo.courseRepo),
		faqUsecase:               _faqUsecase.New(repo.faqRepo),
		termsAndConditionUsecase: _termsConditionUsecase.New(repo.termsAndConditionRepo),
		blogUsecase:              _blogUsecase.New(repo.blogRepo, repo.accountRepo, repo.courseRepo, repo.subjectRepo),
	}
	allUsecase.paymentUsecase = _paymentUsecase.New(repo.paymentRepo, repo.accountRepo, repo.courseRepo, allUsecase.packageUsecase, repo.packageRepo)
	allUsecase.testUsecase = _testUsecase.New(repo.testRepo, repo.courseRepo, repo.questionSetRepo, repo.questionRepo, repo.accountRepo, repo.testSeriesRepo, allUsecase.questionSetUsecase)

	return allUsecase
}
