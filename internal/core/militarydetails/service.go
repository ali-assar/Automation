package militarydetails

import (
	"backend/internal/core/audit"
	"fmt"

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

// Implement RankValidator interface
func (s *Service) IsRankUsed(id int64) (bool, error) {
	var count int64
	if err := s.repo.db.Model(&MilitaryDetails{}).Where("rank_id = ? AND deleted_at = 0", id).Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check military details: %v", err)
	}
	return count > 0, nil
}
func (s *Service) CreateMilitaryDetails(rankID int64, serviceStartDate, serviceDispatchDate, serviceUnit, battalionUnit, companyUnit *int64, actionBy string) (int64, error) {

	militaryDetails := MilitaryDetails{
		RankID:              rankID,
		ServiceStartDate:    serviceStartDate,
		ServiceDispatchDate: serviceDispatchDate,
		ServiceUnit:         serviceUnit,
		BattalionUnit:       battalionUnit,
		CompanyUnit:         companyUnit,
		DeletedAt:           0,
	}
	if err := s.repo.Create(&militaryDetails); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "MilitaryDetails", actionBy); err != nil {
		// Log error but don’t fail
	}
	return militaryDetails.ID, nil
}

func (s *Service) GetMilitaryDetailsByID(id int64) (*MilitaryDetails, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllMilitaryDetails() ([]MilitaryDetails, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateMilitaryDetails(id int64, updates map[string]interface{}, actionBy string) error {
	militaryDetails, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(militaryDetails).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "MilitaryDetails", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteMilitaryDetails(id int64, actionBy string) error {
	militaryDetails, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if militaryDetails.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "MilitaryDetails", actionBy); err != nil {
		// Log error
	}
	return nil
}
