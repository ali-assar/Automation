package db

import (
	"backend/internal/core/action"
	"backend/internal/core/admin"
	"backend/internal/core/credentials"
	"backend/internal/core/person"
	"backend/internal/core/role"
)

func Migrate() error {
	return db.AutoMigrate(
		&admin.Admin{},
		&person.Person{},
		&credentials.Credentials{},
		&role.Role{},
		&action.Action{},
	)
}
