package license_types

import (
	"github.com/nttcom/eclcloud/pagination"
)

// LicenseType represents guest image license information.
type LicenseType struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	HasLicenseKey bool   `json:"has_license_key"`
	Unit          string `json:"unit"`
	LicenseSwitch bool   `json:"license_switch"`
	Description   string `json:"description"`
}

// LicenseTypePage is a single page of LicenseType results.
type LicenseTypePage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of LicenseTypes contains any results.
func (r LicenseTypePage) IsEmpty() (bool, error) {
	licenses, err := ExtractLicenseTypes(r)
	return len(licenses) == 0, err
}

// ExtractLicenseTypes returns a slice of LicenseTypes contained in a single page of results.
func ExtractLicenseTypes(r pagination.Page) ([]LicenseType, error) {
	var s struct {
		LicenseTypes []LicenseType `json:"license_types"`
	}
	err := (r.(LicenseTypePage)).ExtractInto(&s)
	return s.LicenseTypes, err
}
