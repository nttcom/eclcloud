package testing

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/imagestorage/v2/imagedata"
	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestUpload(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandlePutImageDataSuccessfully(t)

	err := imagedata.Upload(
		fakeclient.ServiceClient(),
		"0cb9328d-dd8c-41bb-b378-404b854b93b9",
		readSeekerOfBytes([]byte{5, 3, 7, 24})).ExtractErr()

	th.AssertNoErr(t, err)
}

/*
func TestStage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleStageImageDataSuccessfully(t)

	err := imagedata.Stage(
		fakeclient.ServiceClient(),
		"da3b75d9-3f4a-40e7-8a2c-bfab23927dea",
		readSeekerOfBytes([]byte{5, 3, 7, 24})).ExtractErr()

	th.AssertNoErr(t, err)
}
*/

func readSeekerOfBytes(bs []byte) io.ReadSeeker {
	return &RS{bs: bs}
}

// implements io.ReadSeeker
type RS struct {
	bs     []byte
	offset int
}

func (rs *RS) Read(p []byte) (int, error) {
	leftToRead := len(rs.bs) - rs.offset

	if 0 < leftToRead {
		bytesToWrite := min(leftToRead, len(p))
		for i := 0; i < bytesToWrite; i++ {
			p[i] = rs.bs[rs.offset]
			rs.offset++
		}
		return bytesToWrite, nil
	}
	return 0, io.EOF
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (rs *RS) Seek(offset int64, whence int) (int64, error) {
	var offsetInt = int(offset)
	if whence == 0 {
		rs.offset = offsetInt
	} else if whence == 1 {
		rs.offset = rs.offset + offsetInt
	} else if whence == 2 {
		rs.offset = len(rs.bs) - offsetInt
	} else {
		return 0, fmt.Errorf("for parameter `whence`, expected value in {0,1,2} but got: %#v", whence)
	}

	return int64(rs.offset), nil
}

func TestDownload(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetImageDataSuccessfully(t)

	rdr, err := imagedata.Download(fakeclient.ServiceClient(), "100f4d2d-dcb5-472e-b93f-b4e13d888604").Extract()
	th.AssertNoErr(t, err)

	bs, err := ioutil.ReadAll(rdr)
	th.AssertNoErr(t, err)

	th.AssertByteArrayEquals(t, []byte{34, 87, 0, 23, 23, 23, 56, 255, 254, 0}, bs)
}
