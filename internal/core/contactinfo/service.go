package contactinfo

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

func (s *Service) CreateContactInfo(address, emailAddress, socialMedia, phoneNumber, emergencyPhoneNumber, landlinePhone string, actionBy string) (int64, error) {
	contactInfo := ContactInfo{
		Address:              address,
		PhoneNumber:          phoneNumber,
		EmergencyPhoneNumber: emergencyPhoneNumber,
		LandlinePhone:        landlinePhone,
		EmailAddress:         emailAddress,
		SocialMedia:          socialMedia,
		DeletedAt:            0,
	}
	if err := s.repo.Create(&contactInfo); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "ContactInfo", actionBy); err != nil {
		// Log error but don’t fail
	}
	return contactInfo.ID, nil
}

func (s *Service) GetContactInfoByID(id int64) (*ContactInfo, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetContactInfoByEmail(email string) (*ContactInfo, error) {
	return s.repo.GetByEmail(email)
}

func (s *Service) GetAllContactInfos() ([]ContactInfo, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateContactInfo(id int64, updates map[string]interface{}, actionBy string) error {
	// Prevent updating critical fields
	delete(updates, "id")
	delete(updates, "deleted_at")

	contactInfo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(contactInfo).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "ContactInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteContactInfo(id int64, actionBy string) error {
	contactInfo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if contactInfo.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "ContactInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteContactInfoHard(id int64, actionBy string) error {
	if err := s.repo.DeleteHard(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "ContactInfo", actionBy); err != nil {
		// Log error
	}
	return nil
}
