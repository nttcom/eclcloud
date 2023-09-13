package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4"
	th "github.com/nttcom/eclcloud/v4/testhelper"
)

func TestApplyDefaultsToEndpointOpts(t *testing.T) {
	eo := eclcloud.EndpointOpts{Availability: eclcloud.AvailabilityPublic}
	eo.ApplyDefaults("compute")
	expected := eclcloud.EndpointOpts{Availability: eclcloud.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)

	eo = eclcloud.EndpointOpts{Type: "compute"}
	eo.ApplyDefaults("object-store")
	expected = eclcloud.EndpointOpts{Availability: eclcloud.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)
}
