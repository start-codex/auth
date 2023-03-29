package role

import "github.com/start-codex/auth/internal/domain/permission"

type Role struct {
	id          string
	name        string
	description string
	permissions []permission.Permission
}

func (r *Role) ID() string {
	return r.id
}

func (r *Role) Name() string {
	return r.name
}

func (r *Role) Description() string {
	return r.description
}

func (r *Role) Permissions() []permission.Permission {
	return r.permissions
}
