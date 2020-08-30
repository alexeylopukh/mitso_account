package model

type Uni struct {
	FirstName                    string  `json:"first_name"`
	SecondName                   string  `json:"second_name"`
	ThirdName                    string  `json:"third_name"`
	LatinFirstName               string  `json:"latin_first_name"`
	LatinSecondName              string  `json:"latin_second_name"`
	Gender                       string  `json:"gender"`
	DateOfBirth                  int   `json:"date_of_birth"`
	IsForeigner                  bool    `json:"is_foreigner"`
	Nationality                  string  `json:"nationality"`
	FamilyStatus                 string  `json:"family_status"`
	Region                       string  `json:"region"`
	PostIndex                    string  `json:"post_index"`
	BirthPlace                   string  `json:"birth_place"`
	RegisterAddress              string  `json:"register_address"`
	HomeAddress                  string  `json:"home_address"`
	LiveAddress                  string  `json:"live_address"`
	PhoneNumber                  string  `json:"phone_number"`
	HomePhoneNumber              string  `json:"home_phone_number"`
	DocumentType                 string  `json:"document_type"`
	SerialAndPassportNumber      string  `json:"serial_and_passport_number"`
	IssuedBy                     string  `json:"issued_by"`
	WhenIssud                    int64   `json:"when_issud"`
	ValidUntil                   int64   `json:"valid_until"`
	PersonalNomer                string  `json:"personal_nomer"`
	FormStuding                  string  `json:"form_studing"`
	Faculty                      string  `json:"faculty"`
	Specialization               string  `json:"specialization"`
	TypeEducation                string  `json:"type_education"`
	EducationalInstitution       string  `json:"educational_institution"`
	EducationSpecialization      string  `json:"education_specialization"`
	GraduatedInstitutionYear     int64   `json:"graduated_institution_year"`
	Subject1_Name                string  `json:"subject_1_name"`
	Subject2_Name                string  `json:"subject_2_name"`
	Subject3_Name                string  `json:"subject_3_name"`
	Subject1_Rating              float64   `json:"subject_1_rating"`
	Subject2_Rating              float64   `json:"subject_2_rating"`
	Subject3_Rating              float64   `json:"subject_3_rating"`
	Subject1_Setificate          string  `json:"subject_1_setificate"`
	Subject2_Setificate          string  `json:"subject_2_setificate"`
	Subject3_Setificate          string  `json:"subject_3_setificate"`
	ForigenLangName              string  `json:"forigen_lang_name"`
	ForigenLangRating            float64   `json:"forigen_lang_rating"`
	DocumentTypeEducation        string  `json:"document_type_education"`
	IsSelskayaShkola             bool    `json:"is_selskaya_shkola"`
	IsGorodskaya                 bool    `json:"is_gorodskaya"`
	A2019                        bool    `json:"a2019"`
	Medal                        bool    `json:"medal"`
	Ssuz                         bool    `json:"ssuz"`
	Ptu                          bool    `json:"ptu"`
	DocumentEducationNumberNomer string  `json:"document_education_number_nomer"`
	DocumentEducationAt          int64   `json:"document_education_at"`
	AverageMark                  float64 `json:"average_mark"`
	SumMarks                     float64 `json:"sum_marks"`
	Olimp                        bool    `json:"olimp"`
	SpecFondRb                   bool    `json:"spec_fond_rb"`
	OlimpSport                   bool    `json:"olimp_sport"`
	PerviyRazrad                 bool    `json:"perviy_razrad"`
	LgotniyKredit                bool    `json:"lgotniy_kredit"`
	Hostel                       bool    `json:"hostel"`
	Disabled                     bool    `json:"disabled"`
	Chernobil                    bool    `json:"chernobil"`
	AdditionalInformation        string  `json:"additional_information"`
	Email                        string  `json:"email"`
	Privileges                   bool    `json:"privileges"`
	InfoPrivileges               string  `json:"info_privileges"`
	HasFather                    bool    `json:"has_father"`
	SecondNameFather             string  `json:"second_name_father"`
	ThirdNameFather              string  `json:"third_name_father"`
	AdressFather                 string  `json:"adress_father"`
	WorkAdressFather             string  `json:"work_adress_father"`
	ContactFather                string  `json:"contact_father"`
	HasMother                    bool    `json:"has_mother"`
	SecondMotherName             string  `json:"second_mother_name"`
	MotherFirstName              string  `json:"mother_first_name"`
	ThirdMotherName              string  `json:"third_mother_name"`
	WorkMotherAddress            string  `json:"work_mother_address"`
	MotherContactInfo            string  `json:"mother_contact_info"`
	OwnerPassport                string  `json:"owner_passport"`
	OwnerPassportNumber          string  `json:"owner_passport_number"`
	OwnerPassportID              string  `json:"owner_passport_id"`
	OwnerPassportKtoVidal        string  `json:"owner_passport_kto_vidal"`
	OwnerPassportKogdaVidal      int64   `json:"owner_passport_kogda_vidal"`
	OwnerPassportSrok            int64   `json:"owner_passport_srok"`
	PasportPhoto                 string  `json:"pasport_photo"`
	CT1_Photo                    string  `json:"ct_1_photo"`
	CT2_Photo                    string  `json:"ct_2_photo"`
	CT3_Photo                    string  `json:"ct_3_photo"`
	CT4_Photo                    string  `json:"ct_4_photo"`
	OwnerPasportPhoto            string  `json:"owner_pasport_photo"`
	EducationDocumentPhoto       string  `json:"education_document_photo"`
	FirstNameFather string `json:"first_name_father"`
	AdressMother string `json:"adress_mother"`
}