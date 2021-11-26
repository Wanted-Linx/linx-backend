// Code generated by entc, DO NOT EDIT.

package project

import (
	"time"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldApplyingStartDate holds the string denoting the applying_start_date field in the database.
	FieldApplyingStartDate = "applying_start_date"
	// FieldApplyingEndDate holds the string denoting the applying_end_date field in the database.
	FieldApplyingEndDate = "applying_end_date"
	// FieldQualification holds the string denoting the qualification field in the database.
	FieldQualification = "qualification"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldSponsorFee holds the string denoting the sponsor_fee field in the database.
	FieldSponsorFee = "sponsor_fee"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// EdgeClub holds the string denoting the club edge name in mutations.
	EdgeClub = "club"
	// EdgeProjectClub holds the string denoting the project_club edge name in mutations.
	EdgeProjectClub = "project_club"
	// EdgeProjectLog holds the string denoting the project_log edge name in mutations.
	EdgeProjectLog = "project_log"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "projects"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_project"
	// ClubTable is the table that holds the club relation/edge.
	ClubTable = "projects"
	// ClubInverseTable is the table name for the Club entity.
	// It exists in this package in order to avoid circular dependency with the "club" package.
	ClubInverseTable = "clubs"
	// ClubColumn is the table column denoting the club relation/edge.
	ClubColumn = "club_project"
	// ProjectClubTable is the table that holds the project_club relation/edge.
	ProjectClubTable = "project_clubs"
	// ProjectClubInverseTable is the table name for the ProjectClub entity.
	// It exists in this package in order to avoid circular dependency with the "projectclub" package.
	ProjectClubInverseTable = "project_clubs"
	// ProjectClubColumn is the table column denoting the project_club relation/edge.
	ProjectClubColumn = "project_id"
	// ProjectLogTable is the table that holds the project_log relation/edge.
	ProjectLogTable = "project_logs"
	// ProjectLogInverseTable is the table name for the ProjectLog entity.
	// It exists in this package in order to avoid circular dependency with the "projectlog" package.
	ProjectLogInverseTable = "project_logs"
	// ProjectLogColumn is the table column denoting the project_log relation/edge.
	ProjectLogColumn = "project_project_log"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldContent,
	FieldStartDate,
	FieldEndDate,
	FieldApplyingStartDate,
	FieldApplyingEndDate,
	FieldQualification,
	FieldCreatedAt,
	FieldSponsorFee,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "projects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"club_project",
	"company_project",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
