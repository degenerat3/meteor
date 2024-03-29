// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActionsColumns holds the columns for the "actions" table.
	ActionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "mode", Type: field.TypeString},
		{Name: "args", Type: field.TypeString},
		{Name: "queued", Type: field.TypeBool},
		{Name: "responded", Type: field.TypeBool},
		{Name: "result", Type: field.TypeString, Default: "N/A"},
		{Name: "host_actions", Type: field.TypeInt, Nullable: true},
	}
	// ActionsTable holds the schema information for the "actions" table.
	ActionsTable = &schema.Table{
		Name:       "actions",
		Columns:    ActionsColumns,
		PrimaryKey: []*schema.Column{ActionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "actions_hosts_actions",
				Columns: []*schema.Column{ActionsColumns[7]},

				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BotsColumns holds the columns for the "bots" table.
	BotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "interval", Type: field.TypeInt},
		{Name: "delta", Type: field.TypeInt},
		{Name: "last_seen", Type: field.TypeInt},
		{Name: "host_bots", Type: field.TypeInt, Nullable: true},
	}
	// BotsTable holds the schema information for the "bots" table.
	BotsTable = &schema.Table{
		Name:       "bots",
		Columns:    BotsColumns,
		PrimaryKey: []*schema.Column{BotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bots_hosts_bots",
				Columns: []*schema.Column{BotsColumns[5]},

				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "desc", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// HostsColumns holds the columns for the "hosts" table.
	HostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hostname", Type: field.TypeString, Unique: true},
		{Name: "interface", Type: field.TypeString},
		{Name: "last_seen", Type: field.TypeInt},
	}
	// HostsTable holds the schema information for the "hosts" table.
	HostsTable = &schema.Table{
		Name:        "hosts",
		Columns:     HostsColumns,
		PrimaryKey:  []*schema.Column{HostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupMembersColumns holds the columns for the "group_members" table.
	GroupMembersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "host_id", Type: field.TypeInt},
	}
	// GroupMembersTable holds the schema information for the "group_members" table.
	GroupMembersTable = &schema.Table{
		Name:       "group_members",
		Columns:    GroupMembersColumns,
		PrimaryKey: []*schema.Column{GroupMembersColumns[0], GroupMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "group_members_group_id",
				Columns: []*schema.Column{GroupMembersColumns[0]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "group_members_host_id",
				Columns: []*schema.Column{GroupMembersColumns[1]},

				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActionsTable,
		BotsTable,
		GroupsTable,
		HostsTable,
		UsersTable,
		GroupMembersTable,
	}
)

func init() {
	ActionsTable.ForeignKeys[0].RefTable = HostsTable
	BotsTable.ForeignKeys[0].RefTable = HostsTable
	GroupMembersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupMembersTable.ForeignKeys[1].RefTable = HostsTable
}
