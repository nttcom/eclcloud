package volumeattach

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// List returns a Pager that allows you to iterate over a collection of
// VolumeAttachments.
func List(client *eclcloud.ServiceClient, serverID string) pagination.Pager {
	return pagination.NewPager(client, listURL(client, serverID), func(r pagination.PageResult) pagination.Page {
		return VolumeAttachmentPage{pagination.SinglePageBase(r)}
	})
}

// CreateOptsBuilder allows extensions to add parameters to the Create request.
type CreateOptsBuilder interface {
	ToVolumeAttachmentCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies volume attachment creation or import parameters.
type CreateOpts struct {
	// Device is the device that the volume will attach to the instance as.
	// Omit for "auto".
	Device string `json:"device,omitempty"`

	// VolumeID is the ID of the volume to attach to the instance.
	VolumeID string `json:"volumeId" required:"true"`
}

// ToVolumeAttachmentCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToVolumeAttachmentCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "volumeAttachment")
}

// Create requests the creation of a new volume attachment on the server.
func Create(client *eclcloud.ServiceClient, serverID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeAttachmentCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client, serverID), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get returns public data about a previously created VolumeAttachment.
func Get(client *eclcloud.ServiceClient, serverID, attachmentID string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, serverID, attachmentID), &r.Body, nil)
	return
}

// Delete requests the deletion of a previous stored VolumeAttachment from
// the server.
func Delete(client *eclcloud.ServiceClient, serverID, attachmentID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, serverID, attachmentID), nil)
	return
}
