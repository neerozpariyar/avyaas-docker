package utils

import (
	"fmt"
	"log"
	"math"
	"reflect"

	"gorm.io/gorm"
)

/*
Paginate is a higher-order function that handles the pagination based on the given page.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - func(db *gorm.DB) *gorm.DB: A GORM callback function that sets the offset and limit for pagination
    based on the provided page number and the page size configured in the application settings.
*/
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	// Return a GORM callback function for pagination based on the provided page number
	return func(db *gorm.DB) *gorm.DB {
		/* Calculate offset value, which refers to the number of rows from the beginning of the
		result set that should be skipped */
		offset := (page - 1) * pageSize

		// Set the offset and limit for pagination in the GORM query
		return db.Offset(offset).Limit(pageSize)
	}
}

/*
GetTotalPage calculates the total number of pages for a paginated query based on the total count of records.

Parameters:
  - model: An instance of the model representing the database table.
  - db: A GORM database instance used to execute the count query.
  - pageSize: An integer representing the number of records per page.

Returns:
  - totalPage: A float64 representing the total number of pages for pagination. It is calculated by
    dividing the total count of records by the provided page size. If any error occurs during the
    count query, it returns -1.
*/

func GetTotalPage(model interface{}, db *gorm.DB, pageSize int) (totalPage float64) {
	var count int64

	// Execute a count query on the specified model to get the total number of records
	err := db.Model(&model).Count(&count).Error
	if err != nil {
		log.Printf("err : %v\n", err.Error())
		return -1
	}

	/* Calculate the total number of pages by dividing the total count of records by the page size
	i.e. rounds the given number up to the nearest greatest integer i.e. 3.33 will equivalent to 4 */
	totalPage = math.Ceil(float64(count) / float64(pageSize))

	return totalPage
}

/*
GetTotalPageByConditionModel calculates the total number of pages for a paginated query based on a
specified condition.

Parameters:
  - model: An instance of the model representing the database model.
  - conditionData: A map specifying the conditions for the query.
  - operator: A boolean value indicating the logical relationship between conditions (true for AND,
    false for OR).
  - conditionOperator: An array of condition operators to apply to the query.
  - db: A GORM database instance used to execute the count query.
  - pageSize: An integer representing the number of records per page.

Returns:
  - totalPage: A float64 representing the total number of pages for pagination. If any error occurs
    during the count query, it returns 0.
*/
func GetTotalPageByConditionModel(model interface{}, conditionData map[string]interface{}, operator bool, conditionOperator []string, db *gorm.DB, pageSize int) (totalPage float64) {
	var count int64

	// Execute a count query on the specified model with the specified conditions to get the total number of records
	if err := db.Model(model).Debug().Where(ParseColumnCondition(conditionData, operator, conditionOperator)).Count(&count).Error; err != nil {
		fmt.Printf("err: %v\n", err)
		return 0
	}

	// Calculate the total number of pages
	totalPage = math.Ceil(float64(count) / float64(pageSize))

	return totalPage
}

/*
GetTotalPageByConditionTable calculates the total number of pages for a paginated query based on a
specified condition in the specified database table.

Parameters:
  - table: An instance of the table representing the database table.
  - conditionData: A map specifying the conditions for the query.
  - operator: A boolean value indicating the logical relationship between conditions (true for AND,
    false for OR).
  - conditionOperator: An array of condition operators to apply to the query.
  - db: A GORM database instance used to execute the count query.

Returns:
  - totalPage: A float64 representing the total number of pages for pagination. If any error occurs
    during the count query, it returns 0.
*/
func GetTotalPageByConditionTable(table string, conditionData map[string]interface{}, operator bool, conditionOperator []string, db *gorm.DB, pageSize int) (totalPage float64) {
	var count int64

	// Execute a count query on the specified model with the specified conditions to get the total number of records
	if err := db.Table(table).Debug().Where(ParseColumnCondition(conditionData, operator, conditionOperator)).Count(&count).Error; err != nil {
		fmt.Printf("err: %v\n", err)
		return 0
	}

	// Calculate the total number of pages
	totalPage = math.Ceil(float64(count) / float64(pageSize))

	return totalPage
}

/*
ParseColumnCondition generates a GORM-compatible SQL condition string based on the provided conditions.

Parameters:
  - conditionData: A map specifying column-value pairs for the conditions.
  - operator: A boolean value indicating the logical relationship between conditions (true for AND, false for OR).
  - conditionOperator: An array of condition operators to apply to each condition.

Returns:
  - condition: A string representing the SQL condition generated based on the provided conditions.
*/
func ParseColumnCondition(conditionData map[string]interface{}, operator bool, conditionOperator []string) string {
	var condition string

	// Set the logical operator string based on the specified logical relationship
	operatorString := "or"
	if operator {
		operatorString = "and"
	}

	i := 0
	for k, v := range conditionData {
		// Add the logical operator if it's not the first condition
		if len(condition) > 0 {
			condition += " " + operatorString + " "
		}

		// Determine the type of the condition value
		vType := reflect.TypeOf(v)

		// Add the condition to the string, formatting it based on the type of the condition value
		if vType.Kind() == reflect.String {
			condition += fmt.Sprintf("%s %s '%v'", k, conditionOperator[i], v)
		} else {
			condition += fmt.Sprintf("%s %s %v", k, conditionOperator[i], v)
		}

		i += 1
	}

	return condition
}
