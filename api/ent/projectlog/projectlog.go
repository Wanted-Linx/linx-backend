// Code generated by entc, DO NOT EDIT.

package projectlog

import (
	"time"
)

const (
	// Label holds the string label denoting the projectlog type in the database.
	Label = "project_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeProjectClub holds the string denoting the project_club edge name in mutations.
	EdgeProjectClub = "project_club"
	// EdgeProjectLogParticipant holds the string denoting the project_log_participant edge name in mutations.
	EdgeProjectLogParticipant = "project_log_participant"
	// EdgeProjectLogFeedback holds the string denoting the project_log_feedback edge name in mutations.
	EdgeProjectLogFeedback = "project_log_feedback"
	// Table holds the table name of the projectlog in the database.
	Table = "project_logs"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "project_logs"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_project_log"
	// ProjectClubTable is the table that holds the project_club relation/edge.
	ProjectClubTable = "project_logs"
	// ProjectClubInverseTable is the table name for the ProjectClub entity.
	// It exists in this package in order to avoid circular dependency with the "projectclub" package.
	ProjectClubInverseTable = "project_clubs"
	// ProjectClubColumn is the table column denoting the project_club relation/edge.
	ProjectClubColumn = "project_club_project_log"
	// ProjectLogParticipantTable is the table that holds the project_log_participant relation/edge.
	ProjectLogParticipantTable = "project_log_participants"
	// ProjectLogParticipantInverseTable is the table name for the ProjectLogParticipant entity.
	// It exists in this package in order to avoid circular dependency with the "projectlogparticipant" package.
	ProjectLogParticipantInverseTable = "project_log_participants"
	// ProjectLogParticipantColumn is the table column denoting the project_log_participant relation/edge.
	ProjectLogParticipantColumn = "project_log_project_log_participant"
	// ProjectLogFeedbackTable is the table that holds the project_log_feedback relation/edge.
	ProjectLogFeedbackTable = "project_log_feedbacks"
	// ProjectLogFeedbackInverseTable is the table name for the ProjectLogFeedback entity.
	// It exists in this package in order to avoid circular dependency with the "projectlogfeedback" package.
	ProjectLogFeedbackInverseTable = "project_log_feedbacks"
	// ProjectLogFeedbackColumn is the table column denoting the project_log_feedback relation/edge.
	ProjectLogFeedbackColumn = "project_log_project_log_feedback"
)

// Columns holds all SQL columns for projectlog fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAuthor,
	FieldContent,
	FieldStartDate,
	FieldEndDate,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "project_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_project_log",
	"project_club_project_log",
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
