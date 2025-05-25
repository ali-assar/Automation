package admin

import (
	"backend/internal/audit"
	"backend/internal/person"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	personRepo := person.NewRepository(db)
	admins := []struct {
		NationalIDNumber string
		UserName         string
		Password         string
		RoleID           int64
	}{
		{
			NationalIDNumber: "012345678",
			UserName:         "admin1",
			Password:         "admin123",
			RoleID:           1, // SuperAdmin
		},
		{
			NationalIDNumber: "987654321",
			UserName:         "admin2",
			Password:         "admin456",
			RoleID:           2, // Admin
		},
	}

	for _, a := range admins {
		existingPerson, err := personRepo.GetByID(a.NationalIDNumber)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if err == gorm.ErrRecordNotFound {
			person := &person.Person{
				NationalIDNumber: a.NationalIDNumber,
				FirstName:        a.UserName,
				LastName:         "Admin",
				DeletedAt:        0,
			}
			if err := personRepo.Create(person); err != nil {
				return err
			}
		} else if existingPerson.DeletedAt != 0 {
			existingPerson.DeletedAt = 0
			if err := personRepo.Update(existingPerson); err != nil {
				return err
			}
		}

		existing, err := repo.GetByUsername(a.UserName)
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		hashPassword, err := security.HashPassword(a.Password)
		if err != nil {
			return err
		}

		admin := &Admin{
			ID:               uuid.New(),
			NationalIDNumber: a.NationalIDNumber,
			UserName:         a.UserName,
			HashPassword:     hashPassword,
			RoleID:           a.RoleID,
			DeletedAt:        0,
		}
		if err := repo.Create(admin); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "Admin", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
