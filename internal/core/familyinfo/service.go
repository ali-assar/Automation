package familyinfo

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

type Service struct {
	repo         *Repository
	auditService audit.ActionLogger
}

func NewService(db *gorm.DB, auditService audit.ActionLogger) *Service {
	return &Service{
		repo:         NewRepository(db),
		auditService: auditService,
	}
}

func (s *Service) CreateFamilyInfo(fatherDetails, motherDetails, childsDetails, husbandDetails string, actionBy string) (int64, error) {
	familyInfo := FamilyInfo{
		FatherDetails:  fatherDetails,
		MotherDetails:  motherDetails,
		ChildsDetails:  childsDetails,
		HusbandDetails: husbandDetails,
		DeletedAt:      0,
	}
	if err := s.repo.Create(&familyInfo); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "FamilyInfo", actionBy); err != nil {
		// Log error but don’t fail
	}
	return familyInfo.ID, nil
}

func (s *Service) GetFamilyInfoByID(id int64) (*FamilyInfo, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllFamilyInfos() ([]FamilyInfo, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateFamilyInfo(id int64, updates map[string]interface{}, actionBy string) error {
	familyInfo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(familyInfo).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "FamilyInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteFamilyInfo(id int64, actionBy string) error {
	familyInfo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if familyInfo.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "FamilyInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}
