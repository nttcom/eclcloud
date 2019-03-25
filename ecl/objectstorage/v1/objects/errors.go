package objects

import "github.com/nttcom/eclcloud"

// ErrWrongChecksum is the error when the checksum generated for an object
// doesn't match the ETAG header.
type ErrWrongChecksum struct {
	eclcloud.BaseError
}

func (e ErrWrongChecksum) Error() string {
	return "Local checksum does not match API ETag header"
}
