package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"backend/internal/core/bloodgroup"
	educationlevel "backend/internal/core/educationLevel"
	"backend/internal/core/gender"
	"backend/internal/core/person"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalstatus"
	"backend/internal/core/rank"
	"backend/internal/core/religion"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

// StaticTablesResponse defines the structure of the response for static table data.
type StaticTablesResponse struct {
	BloodGroups      []bloodgroup.BloodGroup         `json:"blood_groups"`
	Religions        []religion.Religion             `json:"religions"`
	PersonTypes      []persontype.PersonType         `json:"person_types"`
	Ranks            []rank.Rank                     `json:"ranks"`
	EducationLevel   []educationlevel.EducationLevel `json:"education_level"`
	Gender           []gender.Gender                 `json:"gender"`
	PhysicalStatuses []physicalstatus.PhysicalStatus `json:"physical_statuses"`
}

// GetStaticTables retrieves all static table data and returns it in a single response.
func GetStaticTables(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch BloodGroups
		bloodGroups, err := s.BloodGroupService.GetAllBloodGroups()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch blood groups: " + err.Error()})
			return
		}

		// Fetch Religions
		religions, err := s.ReligionService.GetAllReligions()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch religions: " + err.Error()})
			return
		}

		// Fetch PersonTypes
		personTypes, err := s.PersonTypeService.GetAllPersonTypes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch person types: " + err.Error()})
			return
		}

		// Fetch Ranks
		ranks, err := s.RankService.GetAllRanks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch ranks: " + err.Error()})
			return
		}

		// Fetch Education Levels
		educationLevels, err := s.EducationLevelService.GetAllEducations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch education levels: " + err.Error()})
			return
		}

		// Fetch Genders
		genders, err := s.GenderService.GetAllGenders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch genders: " + err.Error()})
			return
		}

		// Fetch Physical Statuses
		physicalStatuses, err := s.PhysicalStatusService.GetAllPhysicalStatuses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch physical statuses: " + err.Error()})
			return
		}

		// Construct the response
		response := StaticTablesResponse{
			BloodGroups:      bloodGroups,
			Religions:        religions,
			PersonTypes:      personTypes,
			Ranks:            ranks,
			EducationLevel:   educationLevels,
			Gender:           genders,
			PhysicalStatuses: physicalStatuses,
		}

		c.JSON(http.StatusOK, response)
	}
}

// FullPersonRequest defines the structure of the JSON request for creating a full person record.
type FullPersonRequest struct {
	NationalIDNumber string `json:"national_id_number" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	BirthDate        string `json:"birth_date" binding:"required"`

	FamilyInfo struct {
		FatherDetails  string  `json:"father_details" binding:"required"` // required JSON string
		MotherDetails  string  `json:"mother_details" binding:"required"` // required JSON string
		ChildsDetails  string  `json:"childs_details"`                    // optional JSON string
		HusbandDetails *string `json:"husband_details"`                   // optional JSON string
	} `json:"family_info"`

	ContactInfo struct {
		Address              string `json:"address" binding:"required"`
		PhoneNumber          string `json:"phone_number" binding:"required"`
		EmergencyPhoneNumber string `json:"emergency_phone_number" binding:"required"`
		LandlinePhone        string `json:"landline_phone" binding:"required"`
		EmailAddress         string `json:"email_address"` // optional
		SocialMedia          string `json:"social_media"`  // optional JSON string
	} `json:"contact_info"`

	Skills struct {
		Education struct {
			EducationLevelID int64  `json:"education_level_id" binding:"required"`
			Description      string `json:"description"` // optional
			University       string `json:"university"`  // optional
			StartDate        int64  `json:"start_date"`  // optional
			EndDate          int64  `json:"end_date"`    // optional
		} `json:"education"`
		Languages         string `json:"languages"`          // optional JSON string
		SkillsDescription string `json:"skills_description"` // optional
		Certificates      string `json:"certificates"`       // optional
	} `json:"skills"`

	PhysicalInfo struct {
		BloodGroupID        int64  `json:"blood_group_id" binding:"required"`
		Height              int    `json:"height" binding:"required"`
		Weight              int    `json:"weight" binding:"required"`
		EyeColor            string `json:"eye_color" binding:"required"`
		GenderID            int64  `json:"gender_id" binding:"required"`
		PhysicalStatusID    int64  `json:"physical_status_id" binding:"required"`
		DescriptionOfHealth string `json:"description_of_health"` // optional
	} `json:"physical_info"`

	Religion struct {
		ReligionID int64 `json:"religion_id" binding:"required"`
	} `json:"religion"`

	PersonType struct {
		PersonTypeID int64 `json:"person_type_id" binding:"required"`
	} `json:"person_type"`

	MilitaryDetails struct {
		RankID              int64 `json:"rank_id" binding:"required"`
		ServiceStartDate    int64 `json:"service_start_date"`    // optional
		ServiceDispatchDate int64 `json:"service_dispatch_date"` // optional
		ServiceUnit         int64 `json:"service_unit"`          // optional
		BattalionUnit       int64 `json:"battalion_unit"`        // optional
		CompanyUnit         int64 `json:"company_unit"`          // optional
	} `json:"military_details"`
}

// CreateFullPerson handles the creation of a full person record with related entities.
func CreateFullPerson(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req FullPersonRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload: " + err.Error()})
			return
		}

		// Check for X-Action-By header
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}

		// Parse birth date
		birthDate, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birth_date format, expected YYYY-MM-DD: " + err.Error()})
			return
		}

		// **Family Info**: Validate and create
		fatherDetails, err := utils.ValidateRequiredJSON(req.FamilyInfo.FatherDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "father_details must be non-empty valid JSON: " + err.Error()})
			return
		}
		motherDetails, err := utils.ValidateRequiredJSON(req.FamilyInfo.MotherDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mother_details must be non-empty valid JSON: " + err.Error()})
			return
		}
		childsDetails, err := utils.SafeJSONPtr(req.FamilyInfo.ChildsDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "childs_details must be valid JSON: " + err.Error()})
			return
		}
		var husbandDetails *string
		if req.FamilyInfo.HusbandDetails != nil {
			s := strings.TrimSpace(*req.FamilyInfo.HusbandDetails)
			if s != "" {
				var js json.RawMessage
				if err := json.Unmarshal([]byte(s), &js); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "husband_details must be valid JSON: " + err.Error()})
					return
				}
				husbandDetails = &s
			}
		}
		familyInfoID, err := s.FamilyInfoService.CreateFamilyInfo(fatherDetails, motherDetails, childsDetails, husbandDetails, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create family info: " + err.Error()})
			return
		}

		// **Contact Info**: Handle optional fields and create
		email := utils.NilIfEmpty(req.ContactInfo.EmailAddress)
		var socialMedia *string
		if req.ContactInfo.SocialMedia != "" {
			s := strings.TrimSpace(req.ContactInfo.SocialMedia)
			if s != "" {
				var js json.RawMessage
				if err := json.Unmarshal([]byte(s), &js); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "social_media must be valid JSON: " + err.Error()})
					return
				}
				socialMedia = &s
			}
		}
		contactInfoID, err := s.ContactInfoService.CreateContactInfo(
			req.ContactInfo.Address,
			email,
			socialMedia,
			req.ContactInfo.PhoneNumber,
			req.ContactInfo.EmergencyPhoneNumber,
			req.ContactInfo.LandlinePhone,
			actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create contact info: " + err.Error()})
			return
		}

		// **Education**: Convert optional fields to pointers and create
		edu := req.Skills.Education
		educationID, err := s.EducationService.CreateEducation(
			edu.EducationLevelID,
			utils.NilIfEmpty(edu.Description),
			utils.NilIfEmpty(edu.University),
			utils.Int64PtrIfNonZero(edu.StartDate),
			utils.Int64PtrIfNonZero(edu.EndDate),
			actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create education: " + err.Error()})
			return
		}

		// **Skills**: Validate JSON and create
		sk := req.Skills
		languages, err := utils.SafeJSONPtr(sk.Languages)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "languages must be valid JSON: " + err.Error()})
			return
		}
		skillsID, err := s.SkillsService.CreateSkills(
			educationID,
			languages,
			utils.NilIfEmpty(sk.SkillsDescription),
			utils.NilIfEmpty(sk.Certificates),
			actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create skills: " + err.Error()})
			return
		}

		// **Physical Info**: Convert optional field to pointer and create
		pi := req.PhysicalInfo
		descriptionOfHealth := utils.NilIfEmpty(pi.DescriptionOfHealth)
		physicalInfoID, err := s.PhysicalInfoService.CreatePhysicalInfo(

			pi.Height,
			pi.Weight,
			pi.EyeColor,
			descriptionOfHealth,
			pi.BloodGroupID,
			pi.GenderID,
			pi.PhysicalStatusID,
			actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create physical info: " + err.Error()})
			return
		}

		// **Military Details**: Convert optional fields to pointers and create
		mdet := req.MilitaryDetails
		militaryDetailsID, err := s.MilitaryDetailsService.CreateMilitaryDetails(
			mdet.RankID,
			utils.Int64PtrIfNonZero(mdet.ServiceStartDate),
			utils.Int64PtrIfNonZero(mdet.ServiceDispatchDate),
			utils.Int64PtrIfNonZero(mdet.ServiceUnit),
			utils.Int64PtrIfNonZero(mdet.BattalionUnit),
			utils.Int64PtrIfNonZero(mdet.CompanyUnit),
			actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create military details: " + err.Error()})
			return
		}

		// **Create Person**: Assemble and create the person record
		p := &person.Person{
			NationalIDNumber: req.NationalIDNumber,
			FirstName:        req.FirstName,
			LastName:         req.LastName,
			BirthDate:        birthDate,
			DeletedAt:        0,
		}
		p.SetFamilyInfoID(familyInfoID)
		p.SetContactInfoID(contactInfoID)
		p.SetSkillsID(skillsID)
		p.SetPhysicalInfoID(physicalInfoID)
		p.SetReligionID(req.Religion.ReligionID)
		p.SetPersonTypeID(req.PersonType.PersonTypeID)
		p.SetMilitaryDetailsID(militaryDetailsID)

		personID, err := s.PersonService.CreatePerson(p, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create person: " + err.Error()})
			return
		}

		// Return success response
		c.JSON(http.StatusCreated, gin.H{"national_id_number": personID})
	}
}
