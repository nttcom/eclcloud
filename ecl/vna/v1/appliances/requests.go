package appliances

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToVirtualNetworkApplianceListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the virtual network appliance attributes you want to see returned.
type ListOpts struct {
	Name                          string `q:"name"`
	ID                            string `q:"id"`
	ApplianceType                 string `q:"appliance_type`
	Description                   string `q:"description"`
	AvailabilityZone              string `q:"availability_zone"`
	OSMonitoringStatus            string `q:"os_monitoring_status"`
	OSLoginStatus                 string `q:"os_login_status"`
	VMStatus                      string `q:"vm_status"`
	OperationStatus               string `q:"operation_status"`
	VirtualNetworkAppliancePlanID string `q:"virtual_network_appliance_plan_id"`
	TenantID                      string `q:"tenant_id"`
}

// ToVirtualNetworkApplianceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVirtualNetworkApplianceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// virtual network appliances.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToVirtualNetworkApplianceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return AppliancePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific virtual network appliance based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToApplianceCreateMap() (map[string]interface{}, error)
}

/*
Parameters for Create
*/

type CreateOptsFixedIP struct {
	IPAddress string `json:"ip_address" required:"true"`
}

type CreateOptsInterface struct {
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	NetworkID   string               `json:"network_id" required:"true"`
	Tags        map[string]string    `json:"tags,omitempty"`
	FixedIPs    [1]CreateOptsFixedIP `json:"fixed_ips" required:"true"`
}

type CreateOptsInterfaces struct {
	Interface1 CreateOptsInterface `json:"interface_1" required:"true"`
}

// CreateOpts represents options used to create a virtual network appliance.
type CreateOpts struct {
	Name                          string               `json:"name,omitempty"`
	Description                   string               `json:"description,omitempty"`
	DefaultGateway                string               `json:"default_gateway,omitempty"`
	AvailabilityZone              string               `json:"availability_zone,omitempty"`
	VirtualNetworkAppliancePlanID string               `json:"virtual_network_appliance_plan_id" required:"true"`
	TenantID                      string               `json:"tenant_id,omitempty"`
	Tags                          map[string]string    `json:"tags,omitempty"`
	Interfaces                    CreateOptsInterfaces `json:"interfaces,omitempty"`
}

// ToApplianceCreateMap builds a request body from CreateOpts.
// func (opts CreateOpts) ToApplianceCreateMap() (map[string]interface{}, error) {
func (opts CreateOpts) ToApplianceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "virtual_network_appliance")
}

// Create accepts a CreateOpts struct and creates a new virtual network appliance
// using the values provided.
// This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// virtual network appliance.
// An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToApplianceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToApplianceUpdateMap() (map[string]interface{}, error)
}

/*
Update for Allowed Address Pairs
*/

type UpdateAllowedAddressPairAddressInfo struct {
	IPAddress  string       `json:"ip_address" required:"true"`
	MACAddress *string      `json:"mac_address" required:"true"`
	Type       *string      `json:"type" required:"true"`
	VRID       *interface{} `json:"vrid" required:"true"`
}

type UpdateAllowedAddressPairInterface struct {
	AllowedAddressPairs *[]UpdateAllowedAddressPairAddressInfo `json:"allowed_address_pairs,omitempty"`
}

type UpdateAllowedAddressPairInterfaces struct {
	Interface1 interface{} `json:"interface_1,omitempty"`
	Interface2 interface{} `json:"interface_2,omitempty"`
	Interface3 interface{} `json:"interface_3,omitempty"`
	Interface4 interface{} `json:"interface_4,omitempty"`
	Interface5 interface{} `json:"interface_5,omitempty"`
	Interface6 interface{} `json:"interface_6,omitempty"`
	Interface7 interface{} `json:"interface_7,omitempty"`
	Interface8 interface{} `json:"interface_8,omitempty"`
}

type UpdateAllowedAddressPairOpts struct {
	Interfaces interface{} `json:"interfaces,omitempty"`
}

func (opts UpdateAllowedAddressPairOpts) ToApplianceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "virtual_network_appliance")
}

/*
Update for FixedIP (includes network_id)
*/

type UpdateFixedIPAddressInfo struct {
	IPAddress string `json:"ip_address" required:"true"`
}

type UpdateFixedIPInterface struct {
	NetworkID *string                     `json:"network_id,omitempty"`
	FixedIPs  *[]UpdateFixedIPAddressInfo `json:"fixed_ips,omitempty"`
}

type UpdateFixedIPInterfaces struct {
	Interface1 interface{} `json:"interface_1,omitempty"`
	Interface2 interface{} `json:"interface_2,omitempty"`
	Interface3 interface{} `json:"interface_3,omitempty"`
	Interface4 interface{} `json:"interface_4,omitempty"`
	Interface5 interface{} `json:"interface_5,omitempty"`
	Interface6 interface{} `json:"interface_6,omitempty"`
	Interface7 interface{} `json:"interface_7,omitempty"`
	Interface8 interface{} `json:"interface_8,omitempty"`
}

type UpdateFixedIPOpts struct {
	Interfaces interface{} `json:"interfaces,omitempty"`
}

func (opts UpdateFixedIPOpts) ToApplianceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "virtual_network_appliance")
}

/*
Update for Metadata
*/
type UpdateMetadataInterface struct {
	Name        *string            `json:"name,omitempty"`
	Description *string            `json:"description,omitempty"`
	Tags        *map[string]string `json:"tags,omitempty"`
}

type UpdateMetadataInterfaces struct {
	Interface1 *UpdateMetadataInterface `json:"interface_1,omitempty"`
	Interface2 *UpdateMetadataInterface `json:"interface_2,omitempty"`
	Interface3 *UpdateMetadataInterface `json:"interface_3,omitempty"`
	Interface4 *UpdateMetadataInterface `json:"interface_4,omitempty"`
	Interface5 *UpdateMetadataInterface `json:"interface_5,omitempty"`
	Interface6 *UpdateMetadataInterface `json:"interface_6,omitempty"`
	Interface7 *UpdateMetadataInterface `json:"interface_7,omitempty"`
	Interface8 *UpdateMetadataInterface `json:"interface_8,omitempty"`
}

type UpdateMetadataOpts struct {
	Name        *string                  `json:"name,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Tags        *map[string]string       `json:"tags,omitempty"`
	Interfaces  UpdateMetadataInterfaces `json:"interfaces,omitempty"`
}

// ToApplianceUpdateMap builds a request body from UpdateOpts.
func (opts UpdateMetadataOpts) ToApplianceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "virtual_network_appliance")
}

/*
Update Common
*/

// Update accepts a UpdateOpts struct and updates an existing virtual network appliance
// using the values provided. For more information, see the Create function.
func Update(c *eclcloud.ServiceClient, virtualNetworkApplianceID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToApplianceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Patch(updateURL(c, virtualNetworkApplianceID), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete accepts a unique ID and deletes the virtual network appliance associated with it.
func Delete(c *eclcloud.ServiceClient, virtualNetworkApplianceID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, virtualNetworkApplianceID), nil)
	return
}

// IDFromName is a convenience function that returns a virtual network appliance's
// ID, given its name.
func IDFromName(client *eclcloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""

	listOpts := ListOpts{
		Name: name,
	}

	pages, err := List(client, listOpts).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractAppliances(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "virtual_network_appliance"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "virtual_network_appliance"}
	}
}
