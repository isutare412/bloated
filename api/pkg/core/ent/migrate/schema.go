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
		{Name: "ip", Type: field.TypeString, Size: 50},
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
		{Name: "user_id", Type: field.TypeString, Size: 36},
		{Name: "task", Type: field.TypeString},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "todo_user_id",
				Unique:  false,
				Columns: []*schema.Column{TodosColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BannedIpsTable,
		TodosTable,
	}
)

func init() {
}