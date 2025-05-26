package persontype

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

func (s *Service) CreatePersonType(typeName, actionBy string) (int64, error) {
    personType := PersonType{Type: typeName}
    if err := s.repo.Create(&personType); err != nil {
        return 0, err
    }
    if _, err := s.auditService.LogAction(1, "PersonType", actionBy); err != nil {
        // Log error but don’t fail
    }
    return personType.ID, nil
}

func (s *Service) GetPersonTypeByID(id int64) (*PersonType, error) {
    return s.repo.GetByID(id)
}

func (s *Service) GetPersonTypeByName(typeName string) (*PersonType, error) {
    return s.repo.GetByType(typeName)
}

func (s *Service) GetAllPersonTypes() ([]PersonType, error) {
    return s.repo.GetAll()
}

func (s *Service) UpdatePersonType(id int64, typeName, actionBy string) error {
    personType, err := s.repo.GetByID(id)
    if err != nil {
        return err
    }
    personType.Type = typeName
    if err := s.repo.Update(personType); err != nil {
        return err
    }
    if _, err := s.auditService.LogAction(2, "PersonType", actionBy); err != nil {
        // Log error
    }
    return nil
}

func (s *Service) DeletePersonType(id int64, actionBy string) error {
    if err := s.repo.Delete(id); err != nil {
        return err
    }
    if _, err := s.auditService.LogAction(3, "PersonType", actionBy); err != nil {
        // Log error
    }
    return nil
}