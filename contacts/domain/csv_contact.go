package domain

type CsvContact struct {
	DisplayName             string `csv:"display_name"`
	DisplayNameLastFirst    string `csv:"display_name_last_first"`
	FamilyName              string `csv:"family_name"`
	GivenName               string `csv:"given_name"`
	HonorificPrefix         string `csv:"honorific_prefix"`
	HonorificSuffix         string `csv:"honorific_suffix"`
	MiddleName              string `csv:"middle_name"`
	PhoneticFamilyName      string `csv:"phonetic_family_name"`
	PhoneticFullName        string `csv:"phonetic_full_name"`
	PhoneticGivenName       string `csv:"phonetic_given_name"`
	PhoneticHonorificPrefix string `csv:"phonetic_honorific_prefix"`
	PhoneticHonorificSuffix string `csv:"phonetic_honorific_suffix"`
	PhoneticMiddleName      string `csv:"phonetic_middle_name"`
	UnstructuredName        string `csv:"unstructured_name"`

	Nickname string `csv:"nickname"`

	City            string `csv:"city"`
	Country         string `csv:"country"`
	CountryCode     string `csv:"country_code"`
	ExtendedAddress string `csv:"extended_address"`
	PoBox           string `csv:"po_box"`
	PostalCode      string `csv:"postal_code"`
	Region          string `csv:"region"`
	StreetAddress   string `csv:"street_address"`
	AddressType     string `csv:"address_type"`

	BirthDate string `csv:"birth_date"`
	BirthText string `csv:"birth_text"`

	EmailDisplayName string `csv:"email_display_name"`
	EmailType        string `csv:"email_type"`
	Email            string `csv:"email"`

	AddressMeAs string `csv:"address_me_as"`
	Gender      string `csv:"gender"`

	Occupation string `csv:"occupation"`

	PhoneType string `csv:"phone_type"`
	Phone     string `csv:"phone"`

	PhotoDefault string `csv:"photo_default"`
	PhotoUrl     string `csv:"photo_url"`

	Relation     string `csv:"relation"`
	RelationType string `csv:"relation_type"`

	UrlType string `csv:"url_type"`
	Url     string `csv:"url"`

	Department       string `csv:"department"`
	Domain           string `csv:"domain"`
	EndDate          string `csv:"end_date"`
	JobDescription   string `csv:"job_description"`
	Location         string `csv:"location"`
	Name             string `csv:"name"`
	PhoneticName     string `csv:"phonetic_name"`
	StartDate        string `csv:"start_date"`
	Symbol           string `csv:"symbol"`
	Title            string `csv:"title"`
	OrganizationType string `csv:"organization_type"`
	IsCurrent        string `csv:"is_current"`
}
