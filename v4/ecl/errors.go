package ecl

import (
	"fmt"
	"github.com/nttcom/eclcloud/v4"
	tokens3 "github.com/nttcom/eclcloud/v4/ecl/identity/v3/tokens"
)

// ErrEndpointNotFound is the error when no suitable endpoint can be found
// in the user's catalog
type ErrEndpointNotFound struct{ eclcloud.BaseError }

func (e ErrEndpointNotFound) Error() string {
	return "No suitable endpoint could be found in the service catalog."
}

// ErrInvalidAvailabilityProvided is the error when an invalid endpoint
// availability is provided
type ErrInvalidAvailabilityProvided struct{ eclcloud.ErrInvalidInput }

func (e ErrInvalidAvailabilityProvided) Error() string {
	return fmt.Sprintf("Unexpected availability in endpoint query: %s", e.Value)
}

// ErrMultipleMatchingEndpointsV3 is the error when more than one endpoint
// for the given options is found in the v3 catalog
type ErrMultipleMatchingEndpointsV3 struct {
	eclcloud.BaseError
	Endpoints []tokens3.Endpoint
}

func (e ErrMultipleMatchingEndpointsV3) Error() string {
	return fmt.Sprintf("Discovered %d matching endpoints: %#v", len(e.Endpoints), e.Endpoints)
}

// ErrNoAuthURL is the error when the OS_AUTH_URL environment variable is not
// found
type ErrNoAuthURL struct{ eclcloud.ErrInvalidInput }

func (e ErrNoAuthURL) Error() string {
	return "Environment variable OS_AUTH_URL needs to be set."
}

// ErrNoUsername is the error when the OS_USERNAME environment variable is not
// found
type ErrNoUsername struct{ eclcloud.ErrInvalidInput }

func (e ErrNoUsername) Error() string {
	return "Environment variable OS_USERNAME needs to be set."
}

// ErrNoPassword is the error when the OS_PASSWORD environment variable is not
// found
type ErrNoPassword struct{ eclcloud.ErrInvalidInput }

func (e ErrNoPassword) Error() string {
	return "Environment variable OS_PASSWORD needs to be set."
}
