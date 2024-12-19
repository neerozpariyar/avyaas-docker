package core

import (
	"avyaas/internal/config"
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
	_termsConditionRepo "avyaas/pkg/v3/terms_condition/repository"
	_testRepo "avyaas/pkg/v3/test/repository/gorm"
	_testSeriesRepo "avyaas/pkg/v3/test_series/repository/gorm"
)

func InitRepositories(server *config.Server) *Repositories {
	allRepos := &Repositories{
		courseGroupRepo:  _courseGroupRepo.New(server.DB),
		chapterRepo:      _chapterRepo.New(server.DB),
		noteRepo:         _noteRepo.New(server.DB),
		replyRepo:        _replyRepo.New(server.DB),
		pollRepo:         _pollRepo.New(server.DB),
		packageTypeRepo:  _packageTypeRepo.New(server.DB),
		serviceRepo:      _serviceRepo.New(server.DB),
		referralRepo:     _referralRepo.New(server.DB),
		feedbackRepo:     _feedbackRepo.New(server.DB),
		bookmarkRepo:     _bookmarkRepo.New(server.DB),
		notificationRepo: _notificationRepo.New(server.DB),
		fileClientRepo:   _fileClientRepo.New(server.DB),
		liveGroupRepo:    _liveGroupRepo.New(server.DB),
		recordingRepo:    _recordingRepo.New(server.DB),
		questionRepo:     _questionRepo.New(server.DB),
		accountRepo:      _accountRepo.New(server.DB),
		testSeriesRepo:   _testSeriesRepo.New(server.DB),
		commentRepo:      _commentRepo.New(server.DB),
		faqRepo:          _faqRepo.New(server.DB),
		noticeRepo:       _noticeRepo.New(server.DB),
		blogRepo:         _blogRepo.New(server.DB),
	}

	allRepos.authRepo = _authRepo.New(server.DB, allRepos.accountRepo, allRepos.courseRepo)
	allRepos.contentRepo = _contentRepo.New(server.DB, allRepos.noteRepo)
	allRepos.questionSetRepo = _questionSetRepo.New(server.DB, allRepos.questionRepo)
	allRepos.paymentRepo = _paymentRepo.New(server.DB, allRepos.accountRepo)
	allRepos.liveRepo = _liveRepo.New(server.DB, allRepos.accountRepo)
	allRepos.discussionRepo = _discussionRepo.New(server.DB, allRepos.accountRepo)
	allRepos.courseRepo = _courseRepo.New(server.DB, allRepos.accountRepo, allRepos.courseGroupRepo)
	allRepos.packageRepo = _packageRepo.New(server.DB, allRepos.courseRepo, allRepos.packageTypeRepo)
	allRepos.testRepo = _testRepo.New(server.DB, allRepos.accountRepo, allRepos.courseRepo, allRepos.subjectRepo, allRepos.questionRepo)
	allRepos.subscriptionRepo = _subscriptionRepo.New(server.DB, allRepos.accountRepo)
	allRepos.termsAndConditionRepo = _termsConditionRepo.New(server.DB)
	allRepos.subjectRepo = _subjectRepo.New(server.DB, allRepos.courseRepo)
	allRepos.unitRepo = _unitRepo.New(server.DB, allRepos.subjectRepo)

	return allRepos
}
