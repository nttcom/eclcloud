package ports

import (
	"github.com/nttcom/eclcloud"
)

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToPortUpdateMap() (map[string]interface{}, error)
}

// SinglePort represents parameters to update a Single Firewall/UTM.
type SinglePort struct {
	EnablePort string `json:"enable_port" required:"true"`
	IPAddress  string `json:"ip_address,omitempty"`
	NetworkID  string `json:"network_id,omitempty"`
	SubnetID   string `json:"subnet_id,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

// UpdateOpts represents options used to update a port.
type UpdateOpts struct {
	Port [7]SinglePort `json:"port" required:"true"`
}

// ToPortUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToPortUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// UpdateQueryOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateQueryOptsBuilder interface {
	ToUpdateQuery() (string, error)
}

// UpdateQueryOpts represents query strings for updating port.
type UpdateQueryOpts struct {
	TenantID  string `q:"tenant_id"`
	UserToken string `q:"usertoken"`
}

// ToUpdateQuery formats a ListOpts into a query string.
func (opts UpdateQueryOpts) ToUpdateQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Update modifies the attributes of a port.
func Update(client *eclcloud.ServiceClient,
	serviceType string,
	hostName string,
	opts UpdateOptsBuilder,
	qOpts UpdateQueryOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPortUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	url := updateURL(client, serviceType, hostName)
	if qOpts != nil {
		query, _ := qOpts.ToUpdateQuery()
		url += query
	}

	_, r.Err = client.Put(url, &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
