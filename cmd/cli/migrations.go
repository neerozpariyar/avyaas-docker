package main

import (
	"avyaas/internal/config"
	"avyaas/internal/core"
	"avyaas/internal/domain/models"

	"log"

	"gorm.io/gorm"
)

func main() {
	// Configure Viper settings for reading config file
	config.ConfigureViper()

	// Initialize the database connection
	db := config.InitDB(true, false)

	// Perform database model migration
	MakeMigrations(db)

	// Initialize core permissions that handles migration and creation of roles and permissions tables
	core.InitPermissions(db)
}

/*
MakeMigrations performs database migrations for all the models using the provided GORM database
connection. It uses GORM's AutoMigrate function to automatically create or update the corresponding
database table based on the model's structure.
Parameters:
  - db: A pointer to the GORM database connection.
*/
func MakeMigrations(db *gorm.DB) {
	println("[+][+] Processing: Migrating User Model [+][+]")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		println("[-][-] Failed: User migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TeacherSubject Model [+][+]")
	if err := db.AutoMigrate(&models.TeacherSubject{}); err != nil {
		println("[-][-] Failed: TeacherSubject migration failed [-][-]")
		log.Fatalln(err)
	}

	// println("[+][+] Processing: Migrating Student Model [+][+]")
	// if err := db.AutoMigrate(&models.Student{}); err != nil {
	// 	println("[-][-] Failed: Student migration failed [-][-]")
	// 	log.Fatalln(err)
	// }

	println("[+][+] Processing: Migrating UserOtp Model [+][+]")
	if err := db.AutoMigrate(&models.UserOtp{}); err != nil {
		println("[-][-] Failed: UserOtp migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating CourseGroup Model [+][+]")
	if err := db.AutoMigrate(&models.CourseGroup{}); err != nil {
		println("[-][-] Failed: CourseGroup migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Course Model [+][+]")
	if err := db.AutoMigrate(&models.Course{}); err != nil {
		println("[-][-] Failed: Course migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Subject Model [+][+]")
	if err := db.AutoMigrate(&models.Subject{}); err != nil {
		println("[-][-] Failed: Subject migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Unit Model [+][+]")
	if err := db.AutoMigrate(&models.Unit{}); err != nil {
		println("[-][-] Failed: Unit migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Chapter Model [+][+]")
	if err := db.AutoMigrate(&models.Chapter{}); err != nil {
		println("[-][-] Failed: Chapter migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Content Model [+][+]")
	if err := db.AutoMigrate(&models.Content{}); err != nil {
		println("[-][-] Failed: Content migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating ChapterContent Model [+][+]")
	if err := db.AutoMigrate(&models.ChapterContent{}); err != nil {
		println("[-][-] Failed: ChapterContent migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Note Model [+][+]")
	if err := db.AutoMigrate(&models.Note{}); err != nil {
		println("[-][-] Failed: Note migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TestSeries Model [+][+]")
	if err := db.AutoMigrate(&models.TestSeries{}); err != nil {
		println("[-][-] Failed: TestSeries migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TestType Model [+][+]")
	if err := db.AutoMigrate(&models.TestType{}); err != nil {
		println("[-][-] Failed: TestType migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Test Model [+][+]")
	if err := db.AutoMigrate(&models.Test{}); err != nil {
		println("[-][-] Failed: Test migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TestQuestionSet Model [+][+]")
	if err := db.AutoMigrate(&models.TestQuestionSet{}); err != nil {
		println("[-][-] Failed: TestQuestionSet migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating QuestionSet Model [+][+]")
	if err := db.AutoMigrate(&models.QuestionSet{}); err != nil {
		println("[-][-] Failed: QuestionSet migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Question Model [+][+]")
	if err := db.AutoMigrate(&models.Question{}); err != nil {
		println("[-][-] Failed: Question migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating QuestionSetQuestion Model [+][+]")
	if err := db.AutoMigrate(&models.QuestionSetQuestion{}); err != nil {
		println("[-][-] Failed: QuestionSetQuestion migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Option Model [+][+]")
	if err := db.AutoMigrate(&models.Option{}); err != nil {
		println("[-][-] Failed: Option migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating File Model [+][+]")
	if err := db.AutoMigrate(&models.File{}); err != nil {
		println("[-][-] Failed: File migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating StudentCourse Model [+][+]")
	if err := db.AutoMigrate(&models.StudentCourse{}); err != nil {
		println("[-][-] Failed: StudentCourse migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating StudentContent Model [+][+]")
	if err := db.AutoMigrate(&models.StudentContent{}); err != nil {
		println("[-][-] Failed: StudentContent migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating StudentLiveGroup Model [+][+]")
	if err := db.AutoMigrate(&models.StudentLiveGroup{}); err != nil {
		println("[-][-] Failed: StudentLiveGroup migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating StudentLive Model [+][+]")
	if err := db.AutoMigrate(&models.StudentLive{}); err != nil {
		println("[-][-] Failed: StudentLive migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Discussion Model [+][+]")
	if err := db.AutoMigrate(&models.Discussion{}); err != nil {
		println("[-][-] Failed: Discussion migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Vote Model [+][+]")
	if err := db.AutoMigrate(&models.Vote{}); err != nil {
		println("[-][-] Failed: Vote migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Reply Model [+][+]")
	if err := db.AutoMigrate(&models.Reply{}); err != nil {
		println("[-][-] Failed: Reply migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Poll Model [+][+]")
	if err := db.AutoMigrate(&models.Poll{}); err != nil {
		println("[-][-] Failed: Poll migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating PollOption Model [+][+]")
	if err := db.AutoMigrate(&models.PollOption{}); err != nil {
		println("[-][-] Failed: PollOption migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating PollVote Model [+][+]")
	if err := db.AutoMigrate(&models.PollVote{}); err != nil {
		println("[-][-] Failed: PollVote migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating PackageType Model [+][+]")
	if err := db.AutoMigrate(&models.PackageType{}); err != nil {
		println("[-][-] Failed: PackageType migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Package Model [+][+]")
	if err := db.AutoMigrate(&models.Package{}); err != nil {
		println("[-][-] Failed: Package migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Service Model [+][+]")
	if err := db.AutoMigrate(&models.Service{}); err != nil {
		println("[-][-] Failed: Service migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating ServiceUrl Model [+][+]")
	if err := db.AutoMigrate(&models.ServiceUrl{}); err != nil {
		println("[-][-] Failed: ServiceUrl migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Referral Model [+][+]")
	if err := db.AutoMigrate(&models.Referral{}); err != nil {
		println("[-][-] Failed: Referral migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating UserReferral Model [+][+]")
	if err := db.AutoMigrate(&models.UserReferral{}); err != nil {
		println("[-][-] Failed: UserReferral migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating ReferralInTransaction Model [+][+]")
	if err := db.AutoMigrate(&models.ReferralInTransaction{}); err != nil {
		println("[-][-] Failed: ReferralInTransaction migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Live Model [+][+]")
	if err := db.AutoMigrate(&models.Live{}); err != nil {
		println("[-][-] Failed: Live migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Recording Model [+][+]")
	if err := db.AutoMigrate(&models.Recording{}); err != nil {
		println("[-][-] Failed: Recording migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating LiveGroup Model [+][+]")
	if err := db.AutoMigrate(&models.LiveGroup{}); err != nil {
		println("[-][-] Failed: LiveGroup migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Payment Model [+][+]")
	if err := db.AutoMigrate(&models.Payment{}); err != nil {
		println("[-][-] Failed: Payment migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Subscription Model [+][+]")
	if err := db.AutoMigrate(&models.Subscription{}); err != nil {
		println("[-][-] Failed: Subscription migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TestResult Model [+][+]")
	if err := db.AutoMigrate(&models.TestResult{}); err != nil {
		println("[-][-] Failed: TestResult migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TestResponse Model [+][+]")
	if err := db.AutoMigrate(&models.TestResponse{}); err != nil {
		println("[-][-] Failed: TestResponse migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating StudentTestSeries Model [+][+]")
	if err := db.AutoMigrate(&models.StudentTestSeries{}); err != nil {
		println("[-][-] Failed: StudentTestSeries migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating StudentTest Model [+][+]")
	if err := db.AutoMigrate(&models.StudentTest{}); err != nil {
		println("[-][-] Failed: StudentTest migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Feedback Model [+][+]")
	if err := db.AutoMigrate(&models.Feedback{}); err != nil {
		println("[-][-] Failed: Feedback migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Bookmark Model [+][+]")
	if err := db.AutoMigrate(&models.Bookmark{}); err != nil {
		println("[-][-] Failed: Bookmark migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Notification Model [+][+]")
	if err := db.AutoMigrate(&models.Notification{}); err != nil {
		println("[-][-] Failed: Notification migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating FCMToken Model [+][+]")
	if err := db.AutoMigrate(&models.FCMToken{}); err != nil {
		println("[-][-] Failed: FCMToken migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Comment Model [+][+]")
	if err := db.AutoMigrate(&models.Comment{}); err != nil {
		println("[-][-] Failed: Comment migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Notice Model [+][+]")
	if err := db.AutoMigrate(&models.Notice{}); err != nil {
		println("[-][-] Failed: Notice migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating FAQ Model [+][+]")
	if err := db.AutoMigrate(&models.FAQ{}); err != nil {
		println("[-][-] Failed: FAQ migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TermsAndCondition Model [+][+]")
	if err := db.AutoMigrate(&models.TermsAndCondition{}); err != nil {
		println("[-][-] Failed: TermsAndCondition migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Student Model [+][+]")
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		println("[-][-] Failed: Student migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Registration Count Model [+][+]")
	if err := db.AutoMigrate(&models.RegistrationCount{}); err != nil {
		println("[-][-] Failed: Registration Count migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Blog Model [+][+]")
	if err := db.AutoMigrate(&models.Blog{}); err != nil {
		println("[-][-] Failed: Blog migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Blog Like Model [+][+]")
	if err := db.AutoMigrate(&models.BlogLike{}); err != nil {
		println("[-][-] Failed: Blog Like migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating Blog comment Model [+][+]")
	if err := db.AutoMigrate(&models.BlogComment{}); err != nil {
		println("[-][-] Failed: Blog comment migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TypeQuestion Model [+][+]")
	if err := db.AutoMigrate(&models.TypeQuestion{}); err != nil {
		println("[-][-] Failed: TypeQuestion migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating TypeOption Model [+][+]")
	if err := db.AutoMigrate(&models.TypeOption{}); err != nil {
		println("[-][-] Failed: TypeOption migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating UnitChapterContent Model [+][+]")
	if err := db.AutoMigrate(&models.UnitChapterContent{}); err != nil {
		println("[-][-] Failed: UnitChapterContent migration failed [-][-]")
		log.Fatalln(err)
	}

	println("[+][+] Processing: Migrating SubjectUnitChapterContent Model [+][+]")
	if err := db.AutoMigrate(&models.SubjectUnitChapterContent{}); err != nil {
		println("[-][-] Failed: SubjectUnitChapterContent migration failed [-][-]")
		log.Fatalln(err)
	}
}
