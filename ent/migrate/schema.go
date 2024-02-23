// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlueprintsColumns holds the columns for the "blueprints" table.
	BlueprintsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "blueprint_template", Type: field.TypeBytes},
		{Name: "variable_types", Type: field.TypeJSON},
		{Name: "blueprint_provider", Type: field.TypeUUID},
	}
	// BlueprintsTable holds the schema information for the "blueprints" table.
	BlueprintsTable = &schema.Table{
		Name:       "blueprints",
		Columns:    BlueprintsColumns,
		PrimaryKey: []*schema.Column{BlueprintsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blueprints_providers_provider",
				Columns:    []*schema.Column{BlueprintsColumns[7]},
				RefColumns: []*schema.Column{ProvidersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DeploymentsColumns holds the columns for the "deployments" table.
	DeploymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_accessed", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"awaiting", "in_progress", "complete", "failed", "destroyed", "suspended"}},
		{Name: "template_vars", Type: field.TypeJSON},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "deployment_blueprint", Type: field.TypeUUID},
		{Name: "deployment_requester", Type: field.TypeUUID},
	}
	// DeploymentsTable holds the schema information for the "deployments" table.
	DeploymentsTable = &schema.Table{
		Name:       "deployments",
		Columns:    DeploymentsColumns,
		PrimaryKey: []*schema.Column{DeploymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployments_blueprints_blueprint",
				Columns:    []*schema.Column{DeploymentsColumns[9]},
				RefColumns: []*schema.Column{BlueprintsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "deployments_users_requester",
				Columns:    []*schema.Column{DeploymentsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DeploymentNodesColumns holds the columns for the "deployment_nodes" table.
	DeploymentNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"to_deploy", "to_destroy", "to_redeploy", "parent_awaiting", "child_awaiting", "in_progress", "complete", "tainted", "failed", "destroyed"}},
		{Name: "vars", Type: field.TypeJSON},
		{Name: "deployment_node_deployment", Type: field.TypeUUID},
		{Name: "deployment_node_resource", Type: field.TypeUUID},
	}
	// DeploymentNodesTable holds the schema information for the "deployment_nodes" table.
	DeploymentNodesTable = &schema.Table{
		Name:       "deployment_nodes",
		Columns:    DeploymentNodesColumns,
		PrimaryKey: []*schema.Column{DeploymentNodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployment_nodes_deployments_deployment",
				Columns:    []*schema.Column{DeploymentNodesColumns[5]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "deployment_nodes_resources_resource",
				Columns:    []*schema.Column{DeploymentNodesColumns[6]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GrantedPermissionsColumns holds the columns for the "granted_permissions" table.
	GrantedPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "subject_type", Type: field.TypeEnum, Enums: []string{"user", "group"}},
		{Name: "subject_id", Type: field.TypeUUID},
		{Name: "object_type", Type: field.TypeEnum, Enums: []string{"blueprint", "deployment", "group", "permission", "provider", "user"}},
		{Name: "object_id", Type: field.TypeUUID},
		{Name: "action", Type: field.TypeEnum, Enums: []string{"blueprint_list", "blueprint_create", "blueprint_get", "blueprint_update", "blueprint_delete", "blueprint_deploy", "deployment_list", "deployment_create", "deployment_get", "deployment_update", "deployment_delete", "deployment_destroy", "deployment_redeploy", "deployment_power", "deployment_console", "group_list", "group_create", "group_get", "group_update", "group_delete", "permission_list", "permission_get", "permission_grant", "permission_revoke", "provider_list", "provider_create", "provider_get", "provider_update", "provider_delete", "provider_load", "provider_unload", "provider_configure", "user_list", "user_create", "user_get", "user_update", "user_delete", "unknown"}},
		{Name: "granted_permission_user", Type: field.TypeUUID, Nullable: true},
		{Name: "granted_permission_group", Type: field.TypeUUID, Nullable: true},
	}
	// GrantedPermissionsTable holds the schema information for the "granted_permissions" table.
	GrantedPermissionsTable = &schema.Table{
		Name:       "granted_permissions",
		Columns:    GrantedPermissionsColumns,
		PrimaryKey: []*schema.Column{GrantedPermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "granted_permissions_users_user",
				Columns:    []*schema.Column{GrantedPermissionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "granted_permissions_groups_group",
				Columns:    []*schema.Column{GrantedPermissionsColumns[9]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "grantedpermission_subject_type_subject_id_object_type_object_id_action",
				Unique:  true,
				Columns: []*schema.Column{GrantedPermissionsColumns[3], GrantedPermissionsColumns[4], GrantedPermissionsColumns[5], GrantedPermissionsColumns[6], GrantedPermissionsColumns[7]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// ProvidersColumns holds the columns for the "providers" table.
	ProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "display_name", Type: field.TypeString},
		{Name: "provider_git_url", Type: field.TypeString},
		{Name: "provider_version", Type: field.TypeString},
		{Name: "config_bytes", Type: field.TypeBytes},
		{Name: "is_loaded", Type: field.TypeBool, Default: false},
	}
	// ProvidersTable holds the schema information for the "providers" table.
	ProvidersTable = &schema.Table{
		Name:       "providers",
		Columns:    ProvidersColumns,
		PrimaryKey: []*schema.Column{ProvidersColumns[0]},
	}
	// ResourcesColumns holds the columns for the "resources" table.
	ResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"resource", "data"}, Default: "resource"},
		{Name: "key", Type: field.TypeString},
		{Name: "resource_type", Type: field.TypeString},
		{Name: "features", Type: field.TypeJSON, Nullable: true},
		{Name: "object", Type: field.TypeJSON},
		{Name: "resource_blueprint", Type: field.TypeUUID},
	}
	// ResourcesTable holds the schema information for the "resources" table.
	ResourcesTable = &schema.Table{
		Name:       "resources",
		Columns:    ResourcesColumns,
		PrimaryKey: []*schema.Column{ResourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "resources_blueprints_blueprint",
				Columns:    []*schema.Column{ResourcesColumns[8]},
				RefColumns: []*schema.Column{BlueprintsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// DeploymentNodeNextNodesColumns holds the columns for the "deployment_node_next_nodes" table.
	DeploymentNodeNextNodesColumns = []*schema.Column{
		{Name: "deployment_node_id", Type: field.TypeUUID},
		{Name: "prev_node_id", Type: field.TypeUUID},
	}
	// DeploymentNodeNextNodesTable holds the schema information for the "deployment_node_next_nodes" table.
	DeploymentNodeNextNodesTable = &schema.Table{
		Name:       "deployment_node_next_nodes",
		Columns:    DeploymentNodeNextNodesColumns,
		PrimaryKey: []*schema.Column{DeploymentNodeNextNodesColumns[0], DeploymentNodeNextNodesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployment_node_next_nodes_deployment_node_id",
				Columns:    []*schema.Column{DeploymentNodeNextNodesColumns[0]},
				RefColumns: []*schema.Column{DeploymentNodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "deployment_node_next_nodes_prev_node_id",
				Columns:    []*schema.Column{DeploymentNodeNextNodesColumns[1]},
				RefColumns: []*schema.Column{DeploymentNodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ResourceRequiredByColumns holds the columns for the "resource_required_by" table.
	ResourceRequiredByColumns = []*schema.Column{
		{Name: "resource_id", Type: field.TypeUUID},
		{Name: "depends_on_id", Type: field.TypeUUID},
	}
	// ResourceRequiredByTable holds the schema information for the "resource_required_by" table.
	ResourceRequiredByTable = &schema.Table{
		Name:       "resource_required_by",
		Columns:    ResourceRequiredByColumns,
		PrimaryKey: []*schema.Column{ResourceRequiredByColumns[0], ResourceRequiredByColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "resource_required_by_resource_id",
				Columns:    []*schema.Column{ResourceRequiredByColumns[0]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "resource_required_by_depends_on_id",
				Columns:    []*schema.Column{ResourceRequiredByColumns[1]},
				RefColumns: []*schema.Column{ResourcesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlueprintsTable,
		DeploymentsTable,
		DeploymentNodesTable,
		GrantedPermissionsTable,
		GroupsTable,
		ProvidersTable,
		ResourcesTable,
		UsersTable,
		DeploymentNodeNextNodesTable,
		ResourceRequiredByTable,
		UserGroupsTable,
	}
)

func init() {
	BlueprintsTable.ForeignKeys[0].RefTable = ProvidersTable
	DeploymentsTable.ForeignKeys[0].RefTable = BlueprintsTable
	DeploymentsTable.ForeignKeys[1].RefTable = UsersTable
	DeploymentNodesTable.ForeignKeys[0].RefTable = DeploymentsTable
	DeploymentNodesTable.ForeignKeys[1].RefTable = ResourcesTable
	GrantedPermissionsTable.ForeignKeys[0].RefTable = UsersTable
	GrantedPermissionsTable.ForeignKeys[1].RefTable = GroupsTable
	ResourcesTable.ForeignKeys[0].RefTable = BlueprintsTable
	DeploymentNodeNextNodesTable.ForeignKeys[0].RefTable = DeploymentNodesTable
	DeploymentNodeNextNodesTable.ForeignKeys[1].RefTable = DeploymentNodesTable
	ResourceRequiredByTable.ForeignKeys[0].RefTable = ResourcesTable
	ResourceRequiredByTable.ForeignKeys[1].RefTable = ResourcesTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
