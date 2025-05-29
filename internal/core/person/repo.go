package person

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(person *Person) error {
	return r.db.Create(person).Error
}

func (r *Repository) GetByID(nationalID string) (*Person, error) {
	var person Person
	if err := r.db.
		Preload("FamilyInfo").
		Preload("ContactInfo").
		Preload("Skills").
		Preload("Skills.Education").
		Preload("Skills.Education.EducationLevel").
		Preload("PhysicalInfo").
		Preload("PhysicalInfo.BloodGroup").
		Preload("PhysicalInfo.Gender").
		Preload("PhysicalInfo.PhysicalStatus").
		Preload("Religion").
		Preload("PersonType").
		Preload("MilitaryDetails").
		Preload("MilitaryDetails.RankRef").
		First(&person, "national_id_number = ?", nationalID).
		Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *Repository) GetAll() ([]Person, error) {
	var persons []Person
	if err := r.db.Preload("FamilyInfo").
		Preload("ContactInfo").
		Preload("Skills").
		Preload("Skills.Education").
		Preload("Skills.Education.EducationLevel").
		Preload("PhysicalInfo").
		Preload("PhysicalInfo.BloodGroup").
		Preload("PhysicalInfo.Gender").
		Preload("PhysicalInfo.PhysicalStatus").
		Preload("Religion").
		Preload("PersonType").
		Preload("MilitaryDetails").
		Preload("MilitaryDetails.RankRef").
		Where("deleted_at = 0").Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func (r *Repository) Update(person *Person) error {
	return r.db.Save(person).Error
}

func (r *Repository) DeleteSoft(nationalID string) error {
	return r.db.Model(&Person{}).Where("national_id_number = ?", nationalID).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(nationalID string) error {
	return r.db.Where("national_id_number = ?", nationalID).Delete(&Person{}).Error
}

func (r *Repository) SearchByName(firstName, lastName string) ([]Person, error) {
	var persons []Person
	query := r.db.Preload("FamilyInfo").Preload("PhysicalInfo").Preload("ContactInfo").
		Preload("Skills").Preload("Religion").Preload("PersonType").Preload("MilitaryDetails").
		Where("deleted_at = 0")

	if firstName != "" {
		query = query.Where("first_name ILIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name ILIKE ?", "%"+lastName+"%")
	}

	if err := query.Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func (r *Repository) FilterByPersonType(personTypeID int64) ([]Person, error) {
	var persons []Person
	if err := r.db.Preload("FamilyInfo").Preload("PhysicalInfo").Preload("ContactInfo").
		Preload("Skills").Preload("Religion").Preload("PersonType").Preload("MilitaryDetails").
		Where("person_type_id = ? AND deleted_at = 0", personTypeID).Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}
