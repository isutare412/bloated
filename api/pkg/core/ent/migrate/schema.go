// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BannedIpsColumns holds the columns for the "banned_ips" table.
	BannedIpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "ip", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "country", Type: field.TypeString, Nullable: true, Size: 60},
	}
	// BannedIpsTable holds the schema information for the "banned_ips" table.
	BannedIpsTable = &schema.Table{
		Name:       "banned_ips",
		Columns:    BannedIpsColumns,
		PrimaryKey: []*schema.Column{BannedIpsColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "task", Type: field.TypeString, Size: 3000},
		{Name: "owner_id", Type: field.TypeUUID},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_users_todos",
				Columns:    []*schema.Column{TodosColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true, Size: 256},
		{Name: "user_name", Type: field.TypeString, Nullable: true, Size: 800},
		{Name: "given_name", Type: field.TypeString, Nullable: true, Size: 800},
		{Name: "family_name", Type: field.TypeString, Nullable: true, Size: 800},
		{Name: "photo_url", Type: field.TypeString, Nullable: true},
		{Name: "origin", Type: field.TypeEnum, Enums: []string{"none", "google"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BannedIpsTable,
		TodosTable,
		UsersTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = UsersTable
}
