// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Wanted-Linx/linx-backend/api/ent/student"
	"github.com/Wanted-Linx/linx-backend/api/ent/user"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// University holds the value of the "university" field.
	University string `json:"university,omitempty"`
	// ProfileLink holds the value of the "profile_link" field.
	ProfileLink *string `json:"profile_link,omitempty"`
	// ProfileImage holds the value of the "profile_image" field.
	ProfileImage *string `json:"profile_image,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges        StudentEdges `json:"edges"`
	user_student *int
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			values[i] = new(sql.NullInt64)
		case student.FieldName, student.FieldUniversity, student.FieldProfileLink, student.FieldProfileImage:
			values[i] = new(sql.NullString)
		case student.ForeignKeys[0]: // user_student
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Student", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case student.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case student.FieldUniversity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field university", values[i])
			} else if value.Valid {
				s.University = value.String
			}
		case student.FieldProfileLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_link", values[i])
			} else if value.Valid {
				s.ProfileLink = new(string)
				*s.ProfileLink = value.String
			}
		case student.FieldProfileImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_image", values[i])
			} else if value.Valid {
				s.ProfileImage = new(string)
				*s.ProfileImage = value.String
			}
		case student.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_student", value)
			} else if value.Valid {
				s.user_student = new(int)
				*s.user_student = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Student entity.
func (s *Student) QueryUser() *UserQuery {
	return (&StudentClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return (&StudentClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", university=")
	builder.WriteString(s.University)
	if v := s.ProfileLink; v != nil {
		builder.WriteString(", profile_link=")
		builder.WriteString(*v)
	}
	if v := s.ProfileImage; v != nil {
		builder.WriteString(", profile_image=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student

func (s Students) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
