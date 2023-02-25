package dtos

type CompanyRequestDTO struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Address   struct {
		AddressLine string `json:"AddressLine"`
		CityName    string `json:"cityName"`
		StateProv   string `json:"stateProv"`
		CountryCode string `json:"countryCode"`
		PostalCode  int    `json:"postalCode"`
	} `json:"Address"`
	Company struct {
		CompanyName string `json:"CompanyName"`
		CompanyType string `json:"CompanyType"`
	} `json:"Company"`
}
