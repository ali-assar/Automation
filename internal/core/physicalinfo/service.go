package physicalinfo

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

func (s *Service) CreatePhysicalInfo(height, weight int, eyeColor string, descriptionOfHealth *string, bloodGroupID, genderID, physicalStatusID int64, actionBy string) (int64, error) {
	physicalInfo := PhysicalInfo{
		Height:              height,
		Weight:              weight,
		EyeColor:            eyeColor,
		BloodGroupID:        bloodGroupID,
		GenderID:            genderID,
		PhysicalStatusID:    physicalStatusID,
		DescriptionOfHealth: descriptionOfHealth, // Directly use *string
		DeletedAt:           0,
	}
	if err := s.repo.Create(&physicalInfo); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "PhysicalInfo", actionBy); err != nil {
		// Log error but don’t fail
	}
	return physicalInfo.ID, nil
}
func (s *Service) GetPhysicalInfoByID(id int64) (*PhysicalInfo, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllPhysicalInfos() ([]PhysicalInfo, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdatePhysicalInfo(id int64, updates map[string]interface{}, actionBy string) error {
	physicalInfo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(physicalInfo).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "PhysicalInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePhysicalInfo(id int64, actionBy string) error {
	physicalInfo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if physicalInfo.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "PhysicalInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}
