package qos_options

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type QosOptionPage struct {
	pagination.LinkedPageBase
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the result of Get operations. Call its Extract method to
// interpret it as a QoSOpt.
type GetResult struct {
	commonResult
}

// QoSOpt represents a QoS option.
type QoSOpt struct {
	AWSServiceID      string `json:"aws_service_id"`
	AzureServiceID    string `json:"azure_service_id"`
	Bandwidth         string `json:"bandwidth"`
	Description       string `json:"description"`
	FICServiceID      string `json:"fic_service_id"`
	GCPServiceID      string `json:"gcp_service_id"`
	ID                string `json:"id"`
	InterDCServiceID  string `json:"interdc_service_id"`
	InternetServiceID string `json:"internet_service_id"`
	Name              string `json:"name"`
	QoSType           string `json:"qos_type"`
	ServiceType       string `json:"service_type"`
	Status            string `json:"status"`
	VPNServiceID      string `json:"vpn_service_id"`
}

// IsEmpty checks whether a QosOptionPage struct is empty.
func (r QosOptionPage) IsEmpty() (bool, error) {
	is, err := ExtractQoSOptions(r)
	return len(is) == 0, err
}

// ExtractQoSOptions accepts a Page struct, specifically a QoSOptionPage struct,
// and extracts the elements into a slice of ListOpts structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractQoSOptions(r pagination.Page) ([]QoSOpt, error) {
	var s []QoSOpt
	err := ExtractQoSOptionsInto(r, &s)
	return s, err
}

func ExtractQoSOptionsInto(r pagination.Page, v interface{}) error {
	return r.(QosOptionPage).Result.ExtractIntoSlicePtr(v, "qos_options")
}

// Extract is a function that accepts a result and extracts a QoSOpt.
func (r GetResult) Extract() (*QoSOpt, error) {
	var l QoSOpt
	err := r.ExtractInto(&l)
	return &l, err
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "qos_option")
}
