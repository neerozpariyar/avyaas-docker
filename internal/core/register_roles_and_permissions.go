/*
Package that demonstrates the initialization and configuration of an authority package for managing
roles and permissions in the system. The GetAuth function provides access to the global authority
instance, while InitPermissions sets up the authority system by creating a new instance, migrating
necessary database tables, and attempting to register default roles.
*/

package core

import (
	"fmt"

	authority "github.com/Ayata-Incorporation/roles_and_permission/cmd/roles_and_permissions"
	"gorm.io/gorm"
)

var auth *authority.Authority

// GetAuth initializes, retrieves and returns the global authority instance.
func GetAuth(db *gorm.DB) *authority.Authority {
	return authority.New(authority.Options{DB: db})
}

/*
InitPermissions initializes the application's role and permission system using the provided GORM
database connection. It creates a new instance of the authority.Authority struct and migrates the
necessary tables to the database. If any errors occur during the process, an error message is printed.

Parameters:
  - db: A pointer to the GORM database connection.
*/
func InitPermissions(db *gorm.DB) {
	fmt.Println("[+][+] Procssing: Creating new Roles and Permissions [+][+]")

	// Create a new instance of the authority.Authority struct
	auth = authority.New(authority.Options{DB: db})

	// Migrate tables required for the authority package
	authority.MigrateTables(authority.Options{DB: db})

	// Register default permission urls paths
	if errList := registerDefaultPermissionUrls(auth); errList != nil {
		fmt.Println("[-][-] Error: Error Creating Default Permission Urls [-][-]")

		for _, err := range errList {
			println(err)
		}
	}

	// Register default permission data
	if errList := registerDefaultPermissions(auth); errList != nil {
		fmt.Println("[-][-] Error: Error Creating Default Permissions [-][-]")

		for _, err := range errList {
			println(err)
		}
	}

	// Register default permission and url path relation
	if err := registerDefaultUrlToPermission(auth); err != nil {
		fmt.Println("[-][-] Error: Error Registering Default Permission Urls [-][-]")
	}

	// Register default roles
	if errList := registerDefaultRoles(auth); errList != nil {
		fmt.Println("[-][-] Error: Error Creating Default Roles [-][-]")

		for _, err := range errList {
			println(err)
		}
	}

	// Register default roles
	if errList := registerDefaultRolePermissions(auth); errList != nil {
		fmt.Println("[-][-] Error: Error Creating Default Role Permissions [-][-]")

		for _, err := range errList {
			println(err)
		}
	}
}

/*
registerDefaultRoles is a function responsible for registering default static roles. It takes an
instance of the authority.Authority struct, which provides access to roles and permissions services.
The function defines a map of roles, where each key represents a role ID and its corresponding value
is the name of the static role. The function iterates over the map, attempting to create each static
role. Any errors encountered during the creation process are appended to the errList slice, and the
slice is returned.

Parameters:
  - auth: A pointer to the authority.Authority struct, representing the instance that provides roles
    and permissions services.

Returns:
  - []string: A slice of error messages indicating any issues encountered during the registration of
    default static roles. An empty slice signifies successful registration without errors.
*/
func registerDefaultRoles(auth *authority.Authority) []string {
	var errList []string

	// Define a map of roles with role ID as the key and the corresponding name as the value
	roles := map[uint]string{
		1: "administrator",
		2: "manager",
		3: "teacher",
		4: "student",
	}

	// Iterate over the map and attempt to create each static role
	for roleID, roleName := range roles {
		// Create static roles in the authority system
		if err := auth.CreateStaticRole(roleID, roleName); err != nil {
			// Append error message to the errList slice
			errList = append(errList, fmt.Errorf("[-] role with name '%s' already exists [-]", roleName).Error())
		}
	}

	return errList
}

/*
registerDefaultPermissionUrls is a function responsible for registering default static permission URLs.
It takes an instance of the authority.Authority struct, which provides access to roles and permissions
services. The function defines a map of URL paths, where each key represents a URL ID and its corresponding
value is the path of the static permission URL. The function iterates over the map, attempting to create
each static permission URL. Any errors encountered during the creation process are appended to the errList
slice, and the slice is returned.

Parameters:
  - auth: A pointer to the authority.Authority struct, representing the instance that provides roles
    and permissions services.

Returns:
  - []string: A slice of error messages indicating any issues encountered during the registration of
    default static permission URLs. An empty slice signifies successful registration without errors.
*/
func registerDefaultPermissionUrls(auth *authority.Authority) []string {
	var errList []string

	// Define a map of URL paths with URL ID as the key and the corresponding path as the value
	urlPaths := map[uint]string{
		1:   "/any/",
		2:   "/teacher/create/",
		3:   "/teacher/list/",
		4:   "/teacher/update/",
		5:   "/teacher/delete/",
		6:   "/student/list/",
		7:   "/student/update/",
		8:   "/account/change-password/",
		9:   "/course-group/create/",
		10:  "/course-group/list/",
		11:  "/course-group/update/",
		12:  "/course-group/delete/",
		13:  "/course/create/",
		14:  "/course/list/",
		15:  "/course/details/",
		16:  "/course/update/",
		17:  "/course/delete/",
		18:  "/course/enroll/",
		19:  "/course/enrolled/",
		20:  "/subject/create/",
		21:  "/subject/list/",
		22:  "/subject/details/",
		23:  "/subject/update/",
		24:  "/subject/delete/",
		25:  "/unit/create/",
		26:  "/unit/list/",
		27:  "/unit/update/",
		28:  "/unit/delete/",
		29:  "/unit/update-position/",
		30:  "/chapter/create/",
		31:  "/chapter/list/",
		32:  "/chapter/update/",
		33:  "/chapter/delete/",
		34:  "/chapter/update-position/",
		35:  "/content/create/",
		36:  "/content/list/",
		37:  "/content/details/",
		38:  "/content/update/",
		39:  "/content/delete/",
		40:  "/content/assign/",
		41:  "/content/update-position/",
		42:  "/content/update-progress/",
		43:  "/content/mark-completed/",
		44:  "/note/create/",
		45:  "/note/list/",
		46:  "/note/update/",
		47:  "/note/delete/",
		48:  "/test-series/create/",
		49:  "/test-series/list/",
		50:  "/test-series/update/",
		51:  "/test-series/delete/",
		52:  "/test-type/create/",
		53:  "/test-type/list/",
		54:  "/test-type/update/",
		55:  "/test-type/delete/",
		56:  "/test/create/",
		57:  "/test/list/",
		58:  "/test/details/",
		59:  "/test/update/",
		60:  "/test/delete/",
		61:  "/test/assign-question-set/",
		62:  "/test/update-status/",
		63:  "/test/submit/",
		64:  "/test/leaderboard/",
		65:  "/test/result/",
		66:  "/question-set/create/",
		67:  "/question-set/list/",
		68:  "/question-set/details/",
		69:  "/question-set/update/",
		70:  "/question-set/delete/",
		71:  "/question-set/assign-questions/",
		72:  "/question/create/",
		73:  "/question/list/",
		74:  "/question/update/",
		75:  "/question/delete/",
		76:  "/discussion/create/",
		77:  "/discussion/list/",
		78:  "/discussion/update/",
		79:  "/discussion/delete/",
		80:  "/discussion/vote/",
		81:  "/reply/create/",
		82:  "/reply/list/",
		83:  "/reply/update/",
		84:  "/reply/delete/",
		85:  "/poll/create/",
		86:  "/poll/list/",
		87:  "/poll/update/",
		88:  "/poll/delete/",
		89:  "/poll/vote/",
		90:  "/feedback/create/",
		91:  "/feedback/list/",
		92:  "/feedback/update/",
		93:  "/feedback/delete/",
		94:  "/live-group/create/",
		95:  "/live-group/list/",
		96:  "/live-group/update/",
		97:  "/live-group/delete/",
		98:  "/live/create/",
		99:  "/live/msdk/",
		100: "/live/list/",
		101: "/live/update/",
		102: "/live/delete/",
		103: "/recording/create/",
		104: "/recording/list/",
		105: "/recording/update/",
		106: "/recording/delete/",
		107: "/notification/create/",
		108: "/notification/list/",
		109: "/notification/update/",
		110: "/notification/delete/",
		111: "/notification/publish/",
		112: "/notification/add-fcm-token/",
		113: "/package-type/create/",
		114: "/package-type/list/",
		115: "/package-type/update/",
		116: "/package-type/delete/",
		117: "/package/create/",
		118: "/package/list/",
		119: "/package/update/",
		120: "/package/delete/",
		121: "/package/subscribe/",
		122: "/service/create/",
		123: "/service/list/",
		124: "/service/update/",
		125: "/service/delete/",
		126: "/referral/create/",
		127: "/referral/list/",
		128: "/referral/update/",
		129: "/referral/delete/",
		130: "/referral/apply/",
		131: "/payment/verify/",
		132: "/payment/list/",
		133: "/payment/add-maunal-payment/",
		134: "/bookmark/create/",
		135: "/bookmark/list/",
		136: "/bookmark/details/",
		137: "/bookmark/delete/",
		138: "/subscription/list/",
		139: "/payment/initiate-khalti/",
		140: "/content-comment/create/",
		141: "/content-comment/list/",
		142: "/content-comment/update/",
		143: "/content-comment/delete/",
	}

	// Iterate over the map and attempt to create each static permission URL
	for urlID, urlPath := range urlPaths {
		err := auth.CreateStaticPermissionUrl(urlID, urlPath)
		if err != nil {
			// Append error message to the errList slice
			errList = append(errList, fmt.Errorf("[-] permission url with path '%s' already exists [-]", urlPath).Error())
		}
	}

	return errList
}

// custom data type of type []string, just for removing reduntant data type warning message
type permissionData []string

/*
registerDefaultPermissions is a function responsible for registering default static permissions. It
takes an instance of the authority.Authority struct, which provides access to roles and permissions
services. The function defines a map of permission data, where each key represents a permission ID
and its corresponding value includes the permission name and description. The function iterates over
the map, attempting to create each static permission. Any errors encountered during the creation
process are appended to the errList slice, and the slice is returned.

Parameters:
  - auth: A pointer to the authority.Authority struct, representing the instance that provides roles
    and permissions services.

Returns:
  - []string: A slice of error messages indicating any issues encountered during the registration of
    default static permissions. An empty slice signifies successful registration without errors.
*/
func registerDefaultPermissions(auth *authority.Authority) []string {
	var errList []string

	/* Define a map of permission data with permission ID as the key and corresponding name and
	description as the value */
	permissions := map[uint]permissionData{
		1:   {"Any", "Full access to all features"},
		2:   {"Create teacher", "Create a new teacher"},
		3:   {"List teachers", "View a list of all teachers"},
		4:   {"Update teacher", "Update a teacher information"},
		5:   {"Delete teacher", "Delete a teacher"},
		6:   {"List students", "View a list of all students"},
		7:   {"Update student", "Update a student information"},
		8:   {"Change password", "Change account password"},
		9:   {"Create course group", "Create a new course group"},
		10:  {"List course groups", "View a list of all course groups"},
		11:  {"Update course group", "Update a course group information"},
		12:  {"Delete course group", "Delete a course group"},
		13:  {"Create course", "Create a new course"},
		14:  {"List courses", "View a list of all courses"},
		15:  {"View course details", "View detailed information about a course including all the contents inside"},
		16:  {"Update course", "Update a course information"},
		17:  {"Delete course", "Delete a course"},
		18:  {"Enroll in course", "Enroll in a course"},
		19:  {"List enrolled courses", "View a list of student enrolled courses"},
		20:  {"Create subject", "Create a new subject"},
		21:  {"List subjects", "View a list of all subjects"},
		22:  {"View subject details", "View detailed information about a subject"},
		23:  {"Update subject", "Update subject information"},
		24:  {"Delete subject", "Delete a subject"},
		25:  {"Create unit", "Create a new unit"},
		26:  {"List units", "View a list of all units"},
		27:  {"Update unit", "Update unit information"},
		28:  {"Delete unit", "Delete a unit"},
		29:  {"Update unit position", "Update position of the units in a course"},
		30:  {"Create chapter", "Create a new chapter"},
		31:  {"List chapters", "View a list of all chapters"},
		32:  {"Update chapter", "Update chapter information"},
		33:  {"Delete chapter", "Delete a chapter"},
		34:  {"Update chapter position", "Update position of the chapters in a unit"},
		35:  {"Create content", "Create a new content"},
		36:  {"List contents", "View a list of all contents"},
		37:  {"View content details", "View detailed information about a content"},
		38:  {"Update content", "Update content information"},
		39:  {"Delete content", "Delete content"},
		40:  {"Assign content", "Assign content(s) to a chapter"},
		41:  {"Update content position", "Update the position of the contents in a chapter"},
		42:  {"Update content progress", "Update the progress of a content"},
		43:  {"Mark content as completed", "Mark a content as completed"},
		44:  {"Create note", "Create a new note"},
		45:  {"List notes", "View a list of all notes"},
		46:  {"Update note", "Update note information"},
		47:  {"Delete note", "Delete a note"},
		48:  {"Create test series", "Create a new test series"},
		49:  {"List test serieses", "View a list of all test serieses"},
		50:  {"Update test series", "Update test series information"},
		51:  {"Delete test series", "Delete a test series"},
		52:  {"Create test type", "Create a new test type"},
		53:  {"List test types", "View a list of all test types"},
		54:  {"Update test type", "Update test type information"},
		55:  {"Delete test type", "Delete a test type"},
		56:  {"Create test", "Create a new test"},
		57:  {"List tests", "View a list of all tests"},
		58:  {"View test details", "View detailed information about a test"},
		59:  {"Update test", "Update test information"},
		60:  {"Delete test", "Delete a test"},
		61:  {"Assign question set", "Assign question set to a test"},
		62:  {"Update test status", "Update test visibility status"},
		63:  {"Submit Test", "Submit Test"},
		64:  {"Check Test Leaderboard", "Check Test Leaderboard"},
		65:  {"See Test Result", "See Test Result"},
		66:  {"Create question set", "Create a new question set"},
		67:  {"List question sets", "View a list of all question sets"},
		68:  {"View question set details", "View question set detail information"},
		69:  {"Update question set", "Update question set information"},
		70:  {"Delete question set", "Delete a question set"},
		71:  {"Assign question(s) to question set", "Assign question(s) to a question set"},
		72:  {"Create question", "Create a new question"},
		73:  {"List questions", "View a list of all questions"},
		74:  {"Update question", "Update question information"},
		75:  {"Delete question", "Delete a question"},
		76:  {"Create discussion", "Create a new discussion"},
		77:  {"List discussions", "View a list of all discussions"},
		78:  {"Update discussion", "Update discussion information"},
		79:  {"Delete discussion", "Delete a discussion"},
		80:  {"Vote in discussion", "Vote in a discussion"},
		81:  {"Create reply", "Create a new reply in the discussion"},
		82:  {"List replies", "View a list of all replies in the discussion"},
		83:  {"Update reply", "Update reply information"},
		84:  {"Delete reply", "Delete a reply from the discussion"},
		85:  {"Create poll", "Create a new poll"},
		86:  {"List polls", "View a list of all polls"},
		87:  {"Update poll", "Update poll information"},
		88:  {"Delete poll", "Delete a poll"},
		89:  {"Vote in poll", "Vote in a poll"},
		90:  {"Create feedback", "Create a new feedback"},
		91:  {"List feedbacks", "View a list of all feedbacks"},
		92:  {"Update feedback", "Update feedback information"},
		93:  {"Delete feedback", "Delete a feedback"},
		94:  {"Create live group", "Create a new live group"},
		95:  {"List live groups", "View a list of all live groups"},
		96:  {"Update live group", "Update live group information"},
		97:  {"Delete live group", "Delete a live group"},
		98:  {"Create live meeting", "Create a new live meeting"},
		99:  {"Create MSDK key for live meeting", "Create a MSDK key for zoom meeting"},
		100: {"List live meetings", "View a list of all live meetings"},
		101: {"Update live meeting", "Update live meeting information"},
		102: {"Delete live meeting", "Delete a live meeting"},
		103: {"Create recording", "Create a new recording"},
		104: {"List recordings", "View a list of all recordings"},
		105: {"Update recording", "Update recording information"},
		106: {"Delete recording", "Delete a recording"},
		107: {"Create notification", "Create a new notification"},
		108: {"List notifications", "View a list of all notifications"},
		109: {"Update notification", "Update notification information"},
		110: {"Delete notification", "Delete a notification"},
		111: {"Publish notification", "Publish a notification to send to the user(s)"},
		112: {"Add FCM token", "Register a new FCM token"},
		113: {"Create package type", "Create a new package type"},
		114: {"List package types", "View a list of all package types"},
		115: {"Update package type", "Update package type information"},
		116: {"Delete package type", "Delete a package type"},
		117: {"Create package", "Create a new package"},
		118: {"List packages", "View a list of all packages"},
		119: {"Update package", "Update package information"},
		120: {"Delete package", "Delete a package"},
		121: {"Subscribe package", "Subscribe to a package"},
		122: {"Create service", "Create a new service"},
		123: {"List services", "View a list of all services"},
		124: {"Update service", "Update service information"},
		125: {"Delete service", "Delete a service"},
		126: {"Create referral", "Create a new referral"},
		127: {"List referrals", "View a list of all referrals"},
		128: {"Update referral", "Update referral information"},
		129: {"Delete referral", "Delete a referral"},
		130: {"Apply referral code", "Apply a referral code during subscription"},
		131: {"Verify payment", "Verify a payment"},
		132: {"List payments", "View a list of all payments"},
		133: {"Add manual payment", "Add a manual payment for student subscription"},
		134: {"Create bookmark", "Create a new bookmark"},
		135: {"List bookmark", "View a list of all bookmarks"},
		136: {"View bookmark details", "View detailed information about a bookmark"},
		137: {"Delete bookmark", "Delete a bookmark"},
		138: {"List subscriptions", "View a list of all subscriptions"},
		139: {"Initiate Khalti payment", "Initiate Khalti payment"},
		140: {"Create content comment", "Create a new content comment"},
		141: {"List content comments", "View a list of all content comments"},
		142: {"Update content comment", "Update content comment information"},
		143: {"Delete content comment", "Delete a content comment"},
	}

	// Iterate over the map and attempt to create each static permission
	for permissionID, permissionData := range permissions {
		err := auth.CreateStaticPermission(permissionID, permissionData)
		if err != nil {
			// Append the error to the errList slice
			errList = append(errList, fmt.Errorf("[-] permission with name '%s' already exists [-]", permissions[permissionID]).Error())
		}
	}

	return errList
}

/*
registerDefaultUrlToPermission is a function responsible for assigning default static URLs to permissions.
It takes an instance of the authority.Authority struct, which provides access to roles and permissions
services.The function iterates over a specified range of integers and assigns each value as both the
role ID and permission ID. The purpose of this function is to register default static URLs to
corresponding permissions, ensuring that specific roles have access to predefined URLs. The last
error encountered during the assignment process is returned.

Parameters:
- auth: A pointer to the authority.Authority struct, representing the instance that provides roles and permissions services.

Returns:
  - error: An error indicating any issues encountered during the assignment of default static URLs to permissions.
    A nil error signifies successful assignment.
*/
func registerDefaultUrlToPermission(auth *authority.Authority) error {
	var err error
	for i := 1; i <= 143; i++ {
		err = auth.AssignStaticUrlToPermission(uint(i), uint(i))
	}

	return err
}

func registerDefaultRolePermissions(auth *authority.Authority) []string {
	// var errList []string

	errList := auth.AssignRolePermissions(1, []uint{1})

	return errList
}
