package host_based

import (
	"github.com/nttcom/eclcloud"
)

// GetOptsBuilder allows extensions to add additional parameters to
// the order progress API request
type GetOptsBuilder interface {
	ToServiceOrderQuery() (string, error)
}

// GetOpts represents result of host based security API response.
type GetOpts struct {
	TenantID string `q:"tenant_id"`
	// Locale   string `q:"locale"`
	// SoID     string `q:"soid"`
}

// ToServiceOrderQuery formats a GetOpts into a query string.
func (opts GetOpts) ToServiceOrderQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details of an order progress, by SoId.
func Get(client *eclcloud.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	url := getURL(client)
	if opts != nil {
		query, _ := opts.ToServiceOrderQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToHADeviceCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents parameters used to create a Host based security.
type CreateOpts struct {
	SOKind              string `json:"sokind" required:"true"`
	TenantID            string `json:"tenant_id" required:"true"`
	Locale              string `json:"locale,omitempty"`
	ServiceOrderService string `json:"service_order_service" required:"true"`
	MaxAgentValue       string `json:"max_agent_value" required:"true"`
	MailAddress         string `json:"mailaddress" required:"true"`
	DSMLang             string `json:"dsm_lang" required:"true"`
	TimeZone            string `json:"time_zone" required:"true"`
}

// ToHADeviceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToHADeviceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new Host based security.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToHADeviceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteOptsBuilder allows extensions to add additional parameters to
// the Delete request.
type DeleteOptsBuilder interface {
	ToHADeviceDeleteMap() (map[string]interface{}, error)
}

// DeleteOpts represents parameters used to cancel Host Based Security.
type DeleteOpts struct {
	SOKind      string `json:"sokind" required:"true"`
	TenantID    string `json:"tenant_id" required:"true"`
	Locale      string `json:"locale,omitempty"`
	MailAddress string `json:"mailaddress" required:"true"`
}

// ToHADeviceDeleteMap formats a DeleteOpts into a delete request.
func (opts DeleteOpts) ToHADeviceDeleteMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Delete deletes a device.
func Delete(client *eclcloud.ServiceClient, opts DeleteOptsBuilder) (r DeleteResult) {
	b, err := opts.ToHADeviceDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return

}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToHADeviceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a Host Based Security.
type UpdateOpts struct {
	SOKind      string `json:"sokind" required:"true"`
	TenantID    string `json:"tenant_id" required:"true"`
	Locale      string `json:"locale,omitempty"`
	MailAddress string `json:"mailaddress" required:"true"`
	// Set this in case of Type M1 Change
	ServiceOrderService *string `json:"service_order_service,omitempty"`
	// Set this in case of Type M2 Change
	MaxAgentValue *string `json:"max_agent_value,omitempty"`
}

// ToHADeviceUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToHADeviceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Update modifies the attributes of a Host Based Security.
func Update(client *eclcloud.ServiceClient, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToHADeviceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(updateURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
