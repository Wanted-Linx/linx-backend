// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClubsColumns holds the columns for the "clubs" table.
	ClubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "organization", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "profile_image", Type: field.TypeString, Nullable: true},
		{Name: "profile_link", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "student_club", Type: field.TypeInt, Nullable: true},
	}
	// ClubsTable holds the schema information for the "clubs" table.
	ClubsTable = &schema.Table{
		Name:       "clubs",
		Columns:    ClubsColumns,
		PrimaryKey: []*schema.Column{ClubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "clubs_students_club",
				Columns:    []*schema.Column{ClubsColumns[7]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClubMembersColumns holds the columns for the "club_members" table.
	ClubMembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "club_id", Type: field.TypeInt, Nullable: true},
		{Name: "student_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_club_member", Type: field.TypeInt, Nullable: true},
	}
	// ClubMembersTable holds the schema information for the "club_members" table.
	ClubMembersTable = &schema.Table{
		Name:       "club_members",
		Columns:    ClubMembersColumns,
		PrimaryKey: []*schema.Column{ClubMembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "club_members_clubs_club_member",
				Columns:    []*schema.Column{ClubMembersColumns[1]},
				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "club_members_students_club_member",
				Columns:    []*schema.Column{ClubMembersColumns[2]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "club_members_users_club_member",
				Columns:    []*schema.Column{ClubMembersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "clubmember_club_id_student_id",
				Unique:  true,
				Columns: []*schema.Column{ClubMembersColumns[1], ClubMembersColumns[2]},
			},
		},
	}
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "business_number", Type: field.TypeString},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "profile_image", Type: field.TypeString, Nullable: true},
		{Name: "hompage", Type: field.TypeString, Nullable: true},
		{Name: "user_company", Type: field.TypeInt, Nullable: true},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "companies_users_company",
				Columns:    []*schema.Column{CompaniesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "start_date", Type: field.TypeString},
		{Name: "end_date", Type: field.TypeString},
		{Name: "applying_start_date", Type: field.TypeString},
		{Name: "applying_end_date", Type: field.TypeString},
		{Name: "qualification", Type: field.TypeString},
		{Name: "task_experience", Type: field.TypeString, Nullable: true},
		{Name: "profile_image", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "sponsor_fee", Type: field.TypeInt},
		{Name: "club_project", Type: field.TypeInt, Nullable: true},
		{Name: "company_project", Type: field.TypeInt, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_clubs_project",
				Columns:    []*schema.Column{ProjectsColumns[12]},
				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_companies_project",
				Columns:    []*schema.Column{ProjectsColumns[13]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectClubsColumns holds the columns for the "project_clubs" table.
	ProjectClubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start_date", Type: field.TypeString},
		{Name: "club_id", Type: field.TypeInt, Nullable: true},
		{Name: "project_id", Type: field.TypeInt, Nullable: true},
	}
	// ProjectClubsTable holds the schema information for the "project_clubs" table.
	ProjectClubsTable = &schema.Table{
		Name:       "project_clubs",
		Columns:    ProjectClubsColumns,
		PrimaryKey: []*schema.Column{ProjectClubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_clubs_clubs_project_club",
				Columns:    []*schema.Column{ProjectClubsColumns[2]},
				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_clubs_projects_project_club",
				Columns:    []*schema.Column{ProjectClubsColumns[3]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "projectclub_project_id_club_id",
				Unique:  true,
				Columns: []*schema.Column{ProjectClubsColumns[3], ProjectClubsColumns[2]},
			},
		},
	}
	// ProjectLogsColumns holds the columns for the "project_logs" table.
	ProjectLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "start_date", Type: field.TypeString},
		{Name: "end_date", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "project_project_log", Type: field.TypeInt, Nullable: true},
		{Name: "project_club_project_log", Type: field.TypeInt, Nullable: true},
	}
	// ProjectLogsTable holds the schema information for the "project_logs" table.
	ProjectLogsTable = &schema.Table{
		Name:       "project_logs",
		Columns:    ProjectLogsColumns,
		PrimaryKey: []*schema.Column{ProjectLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_logs_projects_project_log",
				Columns:    []*schema.Column{ProjectLogsColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_logs_project_clubs_project_log",
				Columns:    []*schema.Column{ProjectLogsColumns[8]},
				RefColumns: []*schema.Column{ProjectClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectLogFeedbacksColumns holds the columns for the "project_log_feedbacks" table.
	ProjectLogFeedbacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "author", Type: field.TypeString},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "project_log_project_log_feedback", Type: field.TypeInt, Nullable: true},
	}
	// ProjectLogFeedbacksTable holds the schema information for the "project_log_feedbacks" table.
	ProjectLogFeedbacksTable = &schema.Table{
		Name:       "project_log_feedbacks",
		Columns:    ProjectLogFeedbacksColumns,
		PrimaryKey: []*schema.Column{ProjectLogFeedbacksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_log_feedbacks_project_logs_project_log_feedback",
				Columns:    []*schema.Column{ProjectLogFeedbacksColumns[3]},
				RefColumns: []*schema.Column{ProjectLogsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectLogParticipantsColumns holds the columns for the "project_log_participants" table.
	ProjectLogParticipantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "project_log_project_log_participant", Type: field.TypeInt, Nullable: true},
	}
	// ProjectLogParticipantsTable holds the schema information for the "project_log_participants" table.
	ProjectLogParticipantsTable = &schema.Table{
		Name:       "project_log_participants",
		Columns:    ProjectLogParticipantsColumns,
		PrimaryKey: []*schema.Column{ProjectLogParticipantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_log_participants_project_logs_project_log_participant",
				Columns:    []*schema.Column{ProjectLogParticipantsColumns[2]},
				RefColumns: []*schema.Column{ProjectLogsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "university", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
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
	// TaskTypesColumns holds the columns for the "task_types" table.
	TaskTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "club_task_type", Type: field.TypeInt, Nullable: true},
		{Name: "company_task_type", Type: field.TypeInt, Nullable: true},
		{Name: "project_task_type", Type: field.TypeInt, Nullable: true},
		{Name: "student_task_type", Type: field.TypeInt, Nullable: true},
	}
	// TaskTypesTable holds the schema information for the "task_types" table.
	TaskTypesTable = &schema.Table{
		Name:       "task_types",
		Columns:    TaskTypesColumns,
		PrimaryKey: []*schema.Column{TaskTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_types_clubs_task_type",
				Columns:    []*schema.Column{TaskTypesColumns[2]},
				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "task_types_companies_task_type",
				Columns:    []*schema.Column{TaskTypesColumns[3]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "task_types_projects_task_type",
				Columns:    []*schema.Column{TaskTypesColumns[4]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "task_types_students_task_type",
				Columns:    []*schema.Column{TaskTypesColumns[5]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
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
		ClubsTable,
		ClubMembersTable,
		CompaniesTable,
		ProjectsTable,
		ProjectClubsTable,
		ProjectLogsTable,
		ProjectLogFeedbacksTable,
		ProjectLogParticipantsTable,
		StudentsTable,
		TaskTypesTable,
		UsersTable,
	}
)

func init() {
	ClubsTable.ForeignKeys[0].RefTable = StudentsTable
	ClubMembersTable.ForeignKeys[0].RefTable = ClubsTable
	ClubMembersTable.ForeignKeys[1].RefTable = StudentsTable
	ClubMembersTable.ForeignKeys[2].RefTable = UsersTable
	CompaniesTable.ForeignKeys[0].RefTable = UsersTable
	ProjectsTable.ForeignKeys[0].RefTable = ClubsTable
	ProjectsTable.ForeignKeys[1].RefTable = CompaniesTable
	ProjectClubsTable.ForeignKeys[0].RefTable = ClubsTable
	ProjectClubsTable.ForeignKeys[1].RefTable = ProjectsTable
	ProjectLogsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectLogsTable.ForeignKeys[1].RefTable = ProjectClubsTable
	ProjectLogFeedbacksTable.ForeignKeys[0].RefTable = ProjectLogsTable
	ProjectLogParticipantsTable.ForeignKeys[0].RefTable = ProjectLogsTable
	StudentsTable.ForeignKeys[0].RefTable = UsersTable
	TaskTypesTable.ForeignKeys[0].RefTable = ClubsTable
	TaskTypesTable.ForeignKeys[1].RefTable = CompaniesTable
	TaskTypesTable.ForeignKeys[2].RefTable = ProjectsTable
	TaskTypesTable.ForeignKeys[3].RefTable = StudentsTable
}
