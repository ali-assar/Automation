package person

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

func (s *Service) CreatePerson(person *Person, actionBy string) (string, error) {
	person.DeletedAt = 0 // Initialize as not deleted
	if err := s.repo.Create(person); err != nil {
		return "", err
	}
	if _, err := s.auditService.LogAction(1, "Person", actionBy); err != nil {
		// Log error but don’t fail
	}
	return person.NationalIDNumber, nil
}

func (s *Service) GetPersonByID(nationalID string) (*Person, error) {
	return s.repo.GetByID(nationalID)
}

func (s *Service) GetAllPersons() ([]Person, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdatePerson(nationalID string, updates map[string]interface{}, actionBy string) error {
	// Prevent updating critical fields
	delete(updates, "national_id_number")
	delete(updates, "deleted_at")

	person, err := s.repo.GetByID(nationalID)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(person).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Person", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePerson(nationalID string, actionBy string) error {
	person, err := s.repo.GetByID(nationalID)
	if err != nil {
		return err
	}

	if person.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(nationalID); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Person", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePersonHard(nationalID string, actionBy string) error {
	if err := s.repo.DeleteHard(nationalID); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Person", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) SearchPersonsByName(firstName, lastName string, actionBy string) ([]Person, error) {
	persons, err := s.repo.SearchByName(firstName, lastName)
	if err != nil {
		return nil, err
	}
	// Optional: Log search action if required
	if _, err := s.auditService.LogAction(4, "Person", actionBy); err != nil {
		// Log error
	}
	return persons, nil
}

func (s *Service) FilterPersonsByPersonType(personTypeID int64, actionBy string) ([]Person, error) {
	persons, err := s.repo.FilterByPersonType(personTypeID)
	if err != nil {
		return nil, err
	}
	if _, err := s.auditService.LogAction(4, "Person", actionBy); err != nil {
		// Log error
	}
	return persons, nil
}

func (s *Service) UpdateContactInfo(nationalID string, contactInfoID int64, actionBy string) error {
	person, err := s.repo.GetByID(nationalID)
	if err != nil {
		return err
	}

	person.ContactInfoID = contactInfoID
	if err := s.repo.Update(person); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Person", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) UpdateMilitaryDetails(nationalID string, militaryDetailsID int64, actionBy string) error {
	person, err := s.repo.GetByID(nationalID)
	if err != nil {
		return err
	}

	person.MilitaryDetailsID = militaryDetailsID
	if err := s.repo.Update(person); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Person", actionBy); err != nil {
		// Log error
	}
	return nil
}
