package militarydetails

import (
	"backend/internal/core/audit"
	"backend/internal/core/rank"

	"gorm.io/gorm"
)


func SeedMilitaryDetails(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	rankRepo := rank.NewRepository(db)
	militaryDetails := []struct {
		RankName            string
		ServiceStartDate    int64
		ServiceDispatchDate int64
		ServiceUnit         int64
		BattalionUnit       int64
		CompanyUnit         int64
	}{
		{
			RankName:            "Private",
			ServiceStartDate:    1577836800, // 2020-01-01
			ServiceDispatchDate: 1654041600, // 2022-06-01
			ServiceUnit:         1,
			BattalionUnit:       1,
			CompanyUnit:         1,
		},
		{
			RankName:            "Sergeant",
			ServiceStartDate:    1609459200, // 2021-01-01
			ServiceDispatchDate: 1685577600, // 2023-06-01
			ServiceUnit:         2,
			BattalionUnit:       2,
			CompanyUnit:         2,
		},
	}

	for _, md := range militaryDetails {
		rank, err := rankRepo.GetByID(1) // Assuming Private=1, Sergeant=2
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		var existing MilitaryDetails
		err = db.Where("rank_id = ? AND deleted_at = 0", rank.ID).First(&existing).Error
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(&existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		serviceStartDate := md.ServiceStartDate
		serviceDispatchDate := md.ServiceDispatchDate
		serviceUint := md.ServiceUnit
		battalionUnit := md.BattalionUnit
		companyUnit := md.CompanyUnit

		detail := &MilitaryDetails{
			RankID:              rank.ID,
			ServiceStartDate:    &serviceStartDate,
			ServiceDispatchDate: &serviceDispatchDate,
			ServiceUnit:         &serviceUint,
			BattalionUnit:       &battalionUnit,
			CompanyUnit:         &companyUnit,
			DeletedAt:           0,
		}
		if err := repo.Create(detail); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "MilitaryDetails", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
