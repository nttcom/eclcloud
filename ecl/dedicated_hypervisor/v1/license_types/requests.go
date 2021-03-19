package license_types

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// List retrieves a list of LicenseTypes.
func List(client *eclcloud.ServiceClient) pagination.Pager {
	url := listURL(client)
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return LicenseTypePage{pagination.LinkedPageBase{PageResult: r}}
	})
}
