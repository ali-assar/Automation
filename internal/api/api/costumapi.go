package api

import (
	"fmt"
	"net/http"
	"time"

	"backend/internal/core/bloodgroup" // Assuming BloodGroup model package
	"backend/internal/core/contactinfo"
	"backend/internal/core/education" // Assuming Education model package
	educationlevel "backend/internal/core/educationLevel"
	"backend/internal/core/familyinfo"
	"backend/internal/core/gender"
	"backend/internal/core/militarydetails"
	"backend/internal/core/person"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/physicalstatus"
	"backend/internal/core/rank" // Assuming Rank model package
	"backend/internal/core/religion"
	"backend/internal/core/skills"

	"github.com/gin-gonic/gin"
)

type StaticTablesResponse struct {
	BloodGroups      []bloodgroup.BloodGroup         `json:"blood_groups"`
	Religions        []religion.Religion             `json:"religions"`
	PersonTypes      []persontype.PersonType         `json:"person_types"`
	Ranks            []rank.Rank                     `json:"ranks"`
	EducationLevel   []educationlevel.EducationLevel `json:"education_level"`
	Gender           []gender.Gender                 `json:"gender"`
	PhysicalStatuses []physicalstatus.PhysicalStatus `json:"physical_statuses"`
}

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

		// Fetch education level
		educationlevel, err := s.EducationLevelService.GetAllEducations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch education level: " + err.Error()})
			return
		}

		// Fetch physical Status level
		gender, err := s.GenderService.GetAllGenders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch gender: " + err.Error()})
			return
		}

		physStatuses, err := s.PhysicalStatusService.GetAllPhysicalStatuses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch physical statuses: " + err.Error()})
			return
		}

		// Construct response
		response := StaticTablesResponse{
			BloodGroups:      bloodGroups,
			Religions:        religions,
			PersonTypes:      personTypes,
			Ranks:            ranks,
			EducationLevel:   educationlevel,
			Gender:           gender,
			PhysicalStatuses: physStatuses,
		}

		c.JSON(http.StatusOK, response)
	}
}

type FullPersonRequest struct {
	NationalIDNumber string `json:"national_id_number" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	BirthDate        string `json:"birth_date" binding:"required"`
	FamilyInfo       struct {
		FatherDetails  string `json:"father_details" binding:"required"`
		MotherDetails  string `json:"mother_details" binding:"required"`
		ChildsDetails  string `json:"childs_details" binding:"required"`
		HusbandDetails string `json:"husband_details" binding:"required"`
	} `json:"family_info"`
	ContactInfo struct {
		Address              string `json:"address" binding:"required"`
		PhoneNumber          string `json:"phone_number" binding:"required"`
		EmergencyPhoneNumber string `json:"emergency_phone_number" binding:"required"`
		LandlinePhone        string `json:"landline_phone" binding:"required"`
		EmailAddress         string `json:"email_address" binding:"required,email"`
		SocialMedia          string `json:"social_media" binding:"required"`
	} `json:"contact_info"`
	Skills struct {
		Education struct {
			EducationLevelID int64  `json:"education_level_id" binding:"required"`
			Description      string `json:"description" binding:"required"`
			University       string `json:"university" binding:"required"`
			StartDate        int64  `json:"start_date" binding:"required"`
			EndDate          int64  `json:"end_date" binding:"required"`
		} `json:"education"`
		Languages         string `json:"languages" binding:"required"`
		SkillsDescription string `json:"skills_description" binding:"required"`
		Certificates      string `json:"certificates" binding:"required"`
	} `json:"skills"`
	PhysicalInfo struct {
		BloodGroupID        int64  `json:"blood_group_id" binding:"required"`
		Height              int    `json:"height" binding:"required"`
		Weight              int    `json:"weight" binding:"required"`
		EyeColor            string `json:"eye_color" binding:"required"`
		GenderID            int64  `json:"gender_id" binding:"required"`
		PhysicalStatusID    int64  `json:"physical_status_id" binding:"required"`
		DescriptionOfHealth string `json:"description_of_health" binding:"required"`
	} `json:"physical_info"`
	Religion struct {
		ReligionID int64 `json:"religion_id" binding:"required"`
	} `json:"religion"`
	PersonType struct {
		PersonTypeID int64 `json:"person_type_id" binding:"required"`
	} `json:"person_type"`
	MilitaryDetails struct {
		RankID              int64 `json:"rank_id" binding:"required"`
		ServiceStartDate    int64 `json:"service_start_date" binding:"required"`
		ServiceDispatchDate int64 `json:"service_dispatch_date" binding:"required"`
		ServiceUnit         int64 `json:"service_unit" binding:"required"`
		BattalionUnit       int64 `json:"battalion_unit" binding:"required"`
		CompanyUnit         int64 `json:"company_unit" binding:"required"`
	} `json:"military_details"`
}

func CreateFullPerson(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req FullPersonRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}

		// TODO: add x action by validation

		// Parse birth date
		birthDate, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birth_date format, should be YYYY-MM-DD"})
			return
		}

		// Create FamilyInfo
		familyInfo := &familyinfo.FamilyInfo{
			FatherDetails:  req.FamilyInfo.FatherDetails,
			MotherDetails:  req.FamilyInfo.MotherDetails,
			ChildsDetails:  req.FamilyInfo.ChildsDetails,
			HusbandDetails: req.FamilyInfo.HusbandDetails,
			DeletedAt:      0,
		}
		familyinfoID, err := s.FamilyInfoService.CreateFamilyInfo(familyInfo.FatherDetails, familyInfo.MotherDetails, familyInfo.ChildsDetails, familyInfo.HusbandDetails, actionBy)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create family info: " + err.Error()})
			return
		}

		// Create ContactInfo
		contactInfo := &contactinfo.ContactInfo{
			Address:              req.ContactInfo.Address,
			PhoneNumber:          req.ContactInfo.PhoneNumber,
			EmergencyPhoneNumber: req.ContactInfo.EmergencyPhoneNumber,
			LandlinePhone:        req.ContactInfo.LandlinePhone,
			EmailAddress:         req.ContactInfo.EmailAddress,
			SocialMedia:          req.ContactInfo.SocialMedia,
			DeletedAt:            0,
		}

		contactInfoID, err := s.ContactInfoService.CreateContactInfo(
			contactInfo.Address, contactInfo.EmailAddress, contactInfo.SocialMedia,
			contactInfo.PhoneNumber, contactInfo.EmergencyPhoneNumber, contactInfo.LandlinePhone, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create contact info: " + err.Error()})
			return
		}

		// Create Education (dependency for Skills)
		education := &education.Education{
			EducationLevelID: req.Skills.Education.EducationLevelID,
			Description:      req.Skills.Education.Description,
			University:       req.Skills.Education.University,
			StartDate:        req.Skills.Education.StartDate,
			EndDate:          req.Skills.Education.EndDate,
			DeletedAt:        0,
		}

		educationInfoID, err := s.EducationService.CreateEducation(
			education.EducationLevelID, education.Description,
			education.University, education.StartDate, education.EndDate, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create education: " + err.Error()})
			return
		}

		// Create Skills
		skills := &skills.Skills{
			EducationID:       educationInfoID,
			Languages:         req.Skills.Languages,
			SkillsDescription: req.Skills.SkillsDescription,
			Certificates:      req.Skills.Certificates,
			DeletedAt:         0,
		}

		skillsID, err := s.SkillsService.CreateSkills(
			skills.EducationID, skills.Languages, skills.SkillsDescription, skills.Certificates, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create skills: " + err.Error()})
			return
		}

		// 2) Now build & persist PhysicalInfo with that new status ID:
		physicalInfo := &physicalinfo.PhysicalInfo{
			BloodGroupID:        req.PhysicalInfo.BloodGroupID,
			Height:              req.PhysicalInfo.Height,
			Weight:              req.PhysicalInfo.Weight,
			EyeColor:            req.PhysicalInfo.EyeColor,
			GenderID:            req.PhysicalInfo.GenderID,
			PhysicalStatusID:    req.PhysicalInfo.PhysicalStatusID,
			DescriptionOfHealth: req.PhysicalInfo.DescriptionOfHealth,
			DeletedAt:           0,
		}

		physicalInfoID, err := s.PhysicalInfoService.CreatePhysicalInfo(
			physicalInfo.Height,
			physicalInfo.Weight,
			physicalInfo.EyeColor,
			physicalInfo.DescriptionOfHealth,
			physicalInfo.BloodGroupID,
			physicalInfo.GenderID,
			physicalInfo.PhysicalStatusID,

			actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create physical info: " + err.Error()})
			return
		}

		// Create MilitaryDetails
		militaryDetails := &militarydetails.MilitaryDetails{
			RankID:              req.MilitaryDetails.RankID,
			ServiceStartDate:    req.MilitaryDetails.ServiceStartDate,
			ServiceDispatchDate: req.MilitaryDetails.ServiceDispatchDate,
			ServiceUnit:         req.MilitaryDetails.ServiceUnit,
			BattalionUnit:       req.MilitaryDetails.BattalionUnit,
			CompanyUnit:         req.MilitaryDetails.CompanyUnit,
			DeletedAt:           0,
		}

		militaryDetailsID, err := s.MilitaryDetailsService.CreateMilitaryDetails(
			militaryDetails.RankID, militaryDetails.ServiceStartDate, militaryDetails.ServiceDispatchDate,
			militaryDetails.ServiceUnit, militaryDetails.BattalionUnit, militaryDetails.CompanyUnit, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create military details: " + err.Error()})
			return
		}
		fmt.Println("militaryDetailsid", militaryDetailsID)
		// Create Person
		person := &person.Person{
			NationalIDNumber: req.NationalIDNumber,
			FirstName:        req.FirstName,
			LastName:         req.LastName,
			BirthDate:        birthDate,
			DeletedAt:        0,
		}
		person.SetFamilyInfoID(familyinfoID)
		person.SetContactInfoID(contactInfoID)
		person.SetSkillsID(skillsID)
		person.SetPhysicalInfoID(physicalInfoID)
		person.SetReligionID(req.Religion.ReligionID)
		person.SetPersonTypeID(req.PersonType.PersonTypeID)
		person.SetMilitaryDetailsID(militaryDetailsID)

		strID, err := s.PersonService.CreatePerson(person, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create person: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"national_id_number": strID})
	}
}
