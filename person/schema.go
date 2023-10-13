package person

import (
	"kore_chaincode/core/utils"
	"time"
)

// Person data fields
type Person struct {
	// FirstName      string         `json:"first_name"`
	// MiddleName     string         `json:"middle_name"`
	// LastName       string         `json:"last_name"`
	// Address        string         `json:"address"`
	// Addresses      Address        `json:"addresses"`
	// Country        string         `json:"country"`
	// Email          string         `json:"email"`
	// CountryCode    string         `json:"country_code"`
	// Phone          string         `json:"phone"`
	// DateOfBirth    string         `json:"date_of_birth"`
	// NationalID     []string       `json:"national_id"`
	// DrivingLicense DrivingLicense `json:"driving_license"`
	// Passport       DrivingLicense `json:"passport"`
	PD            string         `json:"pd"`
	Verifications *Verifications `json:"verifications"`
	SourceID      string         `json:"source_id"`
	Status        int32          `json:"status"`
	utils.MetaData
}

// ImportPersonRequest data fields
type ImportPersonRequest struct {
	Data []PersonWithID `json:"data"`
}

// PersonWithID data fields
type PersonWithID struct {
	Data Person `json:"data"`
	ID   string `json:"id"`
}

// Address data fields
type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
}

// DrivingLicense data fields
type DrivingLicense struct {
	Number     string    `json:"number"`
	IssuedDate time.Time `json:"issued_date"`
	ExpiryDate time.Time `json:"expiry_date"`
	Country    string    `json:"country"`
	State      string    `json:"state"`
}

// Verifications data fields
type Verifications struct {
	// AmlVerification                *Verification    `json:"aml_verification"`
	// BadActorVerification           *Verification    `json:"bad_actor_verification"`
	// SuitabilityVerification        *Verification    `json:"suitability_verification"`
	// AccreditedInvestorVerification *Verification    `json:"accredited_investor_verification"`
	KYCVerification *KYCVerification `json:"kyc_verification"`
}

// KYCVerification data fields
type KYCVerification struct {
	ProviderID             string    `json:"provider_id"`
	ProfileID              string    `json:"profile_id"`
	VerificationDate       time.Time `json:"verification_date"`
	VerificationExpiryDate time.Time `json:"verification_expiry_date"`
	TID                    string    `json:"tid"`
	KycReportHash          string    `json:"kyc_report_hash"`
	// Tier          string     `json:"tier"`
	// Rules         []KycRules `json:"rules"`
	OverallStatus string `json:"overall_status"`
}

// KycRules data gields
type KycRules struct {
	RuleGroup   string `json:"rule_group"`
	RuleNumber  string `json:"rule_number"`
	RuleVersion string `json:"rule_version"`
	Status      string `json:"status"`
}

// Verification data fields
type Verification struct {
	TID             string    `json:"tid"`
	ProviderID      string    `json:"provider_id"`
	RequestorID     string    `json:"requestor_id"`
	RequestorSource string    `json:"requestor_source"`
	CertificateID   string    `json:"certificate_id"`
	Status          int       `json:"status"`
	Date            time.Time `json:"date"`
	ExpiryDate      time.Time `json:"expiry_date"`
}

// ShowPersonRequest data fields
type ShowPersonRequest struct {
	ID string `json:"id"`
	// CompanyID   string `json:"company_id"`
	// RequestorID string `json:"requestor_id"`
}
