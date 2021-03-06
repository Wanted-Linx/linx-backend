// Code generated by entc, DO NOT EDIT.

package projectlogfeedback

const (
	// Label holds the string label denoting the projectlogfeedback type in the database.
	Label = "project_log_feedback"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeProjectLog holds the string denoting the project_log edge name in mutations.
	EdgeProjectLog = "project_log"
	// Table holds the table name of the projectlogfeedback in the database.
	Table = "project_log_feedbacks"
	// ProjectLogTable is the table that holds the project_log relation/edge.
	ProjectLogTable = "project_log_feedbacks"
	// ProjectLogInverseTable is the table name for the ProjectLog entity.
	// It exists in this package in order to avoid circular dependency with the "projectlog" package.
	ProjectLogInverseTable = "project_logs"
	// ProjectLogColumn is the table column denoting the project_log relation/edge.
	ProjectLogColumn = "project_log_project_log_feedback"
)

// Columns holds all SQL columns for projectlogfeedback fields.
var Columns = []string{
	FieldID,
	FieldAuthor,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "project_log_feedbacks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_log_project_log_feedback",
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
