// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "university", Type: field.TypeString},
		{Name: "interested_type", Type: field.TypeString, Nullable: true},
		{Name: "profile_link", Type: field.TypeString, Nullable: true},
		{Name: "profile_image", Type: field.TypeString, Nullable: true},
		{Name: "user_student", Type: field.TypeInt, Nullable: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_users_student",
				Columns:    []*schema.Column{StudentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"company", "student"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StudentsTable,
		UsersTable,
	}
)

func init() {
	StudentsTable.ForeignKeys[0].RefTable = UsersTable
}
