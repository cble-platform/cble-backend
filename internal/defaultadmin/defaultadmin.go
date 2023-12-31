package defaultadmin

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/group"
	"github.com/cble-platform/cble-backend/ent/permissionpolicy"
	"github.com/cble-platform/cble-backend/ent/user"
	"github.com/cble-platform/cble-backend/internal/permissionengine"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func InitializeDefaultAdminUserGroup(ctx context.Context, client *ent.Client, pe *permissionengine.PermissionEngine, cbleConfig *config.Config) error {
	// Ensure the built-in admin group exists
	cbleAdminGroup, err := client.Group.Query().Where(
		group.NameEQ(cbleConfig.Initialization.AdminGroup),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("failed to query for default admin group: %v", err)
	} else if cbleAdminGroup == nil {
		// If it doesn't exist, make it
		cbleAdminGroup, err = client.Group.Create().
			SetName(cbleConfig.Initialization.AdminGroup).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create default admin group: %v", err)
		}
	}

	// Give this admin group every permission
	entPermissions, err := client.Permission.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("failed to query all permissions")
	}
	for _, entPermission := range entPermissions {
		_, err := pe.SetPermissionPolicy(ctx, permissionpolicy.TypeALLOW, entPermission, cbleAdminGroup)
		if err != nil {
			// Log non-fatal errors
			logrus.Errorf("failed to grant permission %s to default admin group", entPermission.Key)
		}
	}

	// Check if there are any admins in existence with the built in admin group
	if anyAdminExists, err := client.User.Query().Where(
		user.And(
			user.HasGroupsWith(
				group.NameEQ(cbleConfig.Initialization.AdminGroup),
			),
		),
	).Exist(ctx); err != nil {
		return fmt.Errorf("failed to query for existing admins: %v", err)
	} else if !anyAdminExists {
		// If there are no admins in admin group, check if the default one exists
		defaultAdminExists, err := client.User.Query().Where(
			user.And(
				user.UsernameEQ(cbleConfig.Initialization.DefaultAdmin.Username),
				user.HasGroupsWith(
					group.NameEQ(cbleConfig.Initialization.AdminGroup),
				),
			),
		).Exist(ctx)
		if err != nil {
			return fmt.Errorf("failed to query for default admin: %v", err)
		}
		if !defaultAdminExists {
			// If the default one doesn't exist, make it

			// Hash the default password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cbleConfig.Initialization.DefaultAdmin.Password), 8)
			if err != nil {
				return fmt.Errorf("failed to hash default admin password: %v", err)
			}
			// Create the default admin
			_, err = client.User.Create().
				SetUsername(cbleConfig.Initialization.DefaultAdmin.Username).
				SetPassword(string(hashedPassword)).
				SetEmail(cbleConfig.Initialization.DefaultAdmin.Email).
				SetFirstName(cbleConfig.Initialization.DefaultAdmin.FirstName).
				SetLastName(cbleConfig.Initialization.DefaultAdmin.LastName).
				AddGroups(cbleAdminGroup).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create default admin: %v", err)
			}
			logrus.Info("Created default admin user")
		}
	}

	return nil
}
