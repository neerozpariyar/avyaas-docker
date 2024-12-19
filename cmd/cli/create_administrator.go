/*
Package responsible for collecting user input, validating and processing it, and creating an administrator
user with the specified details. It utilizes functions to prompt the user for data, identify missing
information, and ensure password strength. The user information is then stored in the database, and
assign the administrator role to the created user using the authority package.
*/

package main

import (
	"avyaas/internal/config"
	"avyaas/internal/domain/models"
	"avyaas/utils"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	authority "github.com/Ayata-Incorporation/roles_and_permission/cmd/roles_and_permissions"
)

var (
	username   string
	email      string
	phone      string
	password   string
	verified   bool
	allMissing []string
)

func main() {
	var user models.User
	var err error

	// Collect user input and validate it
	getData()
	findMissing()

	// Loop until all required fields are provided
	for len(allMissing) != 0 {
		findMissing()
		getMissing()
	}

	// Validate the strength of the provided user's password
	validatePasswordStrength()

	// Initiate viper configuration for reading config file
	config.ConfigureViper()

	// Initialize the database
	db := config.InitDB(true, false)

	// Begin the transaction instance of database
	transaction := db.Begin()

	// Create a new instance of the authority.Authority struct
	auth := authority.New(authority.Options{DB: db})

	user = models.User{
		Username: username,
		Email:    email,
		Phone:    phone,
		RoleID:   1,
		Verified: true,
	}

	splitEmail := strings.Split(email, "@")
	user.Username = strings.ToLower(splitEmail[0])

	// Hash the user's password before storing it
	user.Password, err = utils.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	println("[+] Processing: Creating Administrator User [+]")
	if err = transaction.Create(&user).Error; err != nil {
		println("[-] Error: There was an error creating the ADMINISTRATOR. Please try again. [-]")
		panic(err)
	}

	// Assign the administrator role to the created user
	if err := auth.AssignRole(user.ID, 1); err != nil {
		println("[-] Error: There was an error assigning the ADMINISTRATOR role to the user. Please try again. [-]")
		println("[-] Processing: Rolling back the transaction. [-]")
		transaction.Rollback()
		panic(err)
	}

	// Assign permission with id "1" which gives all permissionto the created user
	if errList := auth.AssignUserPermissions(user.ID, []uint{1}); len(errList) != 0 {
		panic(errList)
	}

	transaction.Commit()

	fmt.Printf("\n[+] Created Administrator user with username %v [+]\n", username)
}

/*
getData prompts the user to input values for various user-related fields, including full name,
username, email, phone, and password.
*/
func getData() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the value for email:")
	if scanner.Scan() {
		email = scanner.Text()
	}
	println()

	fmt.Println("Enter the value for phone number:")
	if scanner.Scan() {
		phone = scanner.Text()
	}
	println()

	fmt.Println("Enter the value for password:")
	if scanner.Scan() {
		password = scanner.Text()
	}
	println()
}

/*
findMissing evaluates whether required fields (email, phone, password) have been provided with
values. If any of these fields are empty, they are added to the 'allMissing' slice, indicating the
fields that need to be filled in by the user. It identifies and tracks missing information before
prompting the user to input the necessary values.
*/
func findMissing() {
	allMissing = []string{}

	// NOTE: email, phone and password are compulsory

	if email == "" {
		allMissing = append(allMissing, "email")
	}

	if phone == "" {
		allMissing = append(allMissing, "phone")
	}

	if password == "" {
		allMissing = append(allMissing, "password")
	}
}

// getMissing prompts the user to input values for missing fields specified in the 'allMissing' slice.
func getMissing() {
	for _, missed := range allMissing {
		switch missed {
		case "email":
			fmt.Println("Enter the value for email:\t")
			fmt.Scanln(&email)
			println()
		case "phone":
			fmt.Println("Enter the value for phone number:\t")
			fmt.Scanln(&phone)
			println()
		case "password":
			fmt.Println("Enter the value for password:\t")
			fmt.Scanln(&password)
			println()
		}
	}
}

/*
validatePasswordStrength continuously prompts the user to input a password, until a password that
satisfies the strength criteria is provided.
*/
func validatePasswordStrength() {
	for true {
		breakLoop := false

		err := utils.CheckPasswordStrength(password)
		if err == nil {
			breakLoop = true
		} else {
			var choice string
			fmt.Println("[+] Error: Password strength doesn't match the requirement [+]")
			fmt.Printf("\"%v\"\n", password)
			fmt.Println("Do you still want to use this password ? [Y(y/yes) / N(n/no)]")
			if _, err = fmt.Scan(&choice); err != nil {
				fmt.Println(err)
			}

			switch choice {
			case "Y", "y":
				breakLoop = true
			case "N", "n":
				// Prompt for the new password string
				fmt.Println("[+] Enter new Password [+]\nPassword:\t")
				if _, err := fmt.Scan(&password); err != nil {
					fmt.Println(err)
				}
			default:
				fmt.Println("Invalid Choice! Please try again.")
			}
		}

		if breakLoop {
			break
		}
	}
}
