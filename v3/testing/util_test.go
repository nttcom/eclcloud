package testing

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/v3"
	th "github.com/nttcom/eclcloud/v3/testhelper"
)

func TestWaitFor(t *testing.T) {
	err := eclcloud.WaitFor(2, func() (bool, error) {
		return true, nil
	})
	th.CheckNoErr(t, err)
}

func TestWaitForTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	err := eclcloud.WaitFor(1, func() (bool, error) {
		return false, nil
	})
	th.AssertEquals(t, "A timeout occurred", err.Error())
}

func TestWaitForError(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	err := eclcloud.WaitFor(2, func() (bool, error) {
		return false, errors.New("error has occurred")
	})
	th.AssertEquals(t, "error has occurred", err.Error())
}

func TestWaitForPredicateExceed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	err := eclcloud.WaitFor(1, func() (bool, error) {
		time.Sleep(4 * time.Second)
		return false, errors.New("just wasting time")
	})
	th.AssertEquals(t, "A timeout occurred", err.Error())
}

func TestNormalizeURL(t *testing.T) {
	urls := []string{
		"NoSlashAtEnd",
		"SlashAtEnd/",
	}
	expected := []string{
		"NoSlashAtEnd/",
		"SlashAtEnd/",
	}
	for i := 0; i < len(expected); i++ {
		th.CheckEquals(t, expected[i], eclcloud.NormalizeURL(urls[i]))
	}

}

func TestNormalizePathURL(t *testing.T) {
	baseDir := "/test/path"

	rawPath := "template.yaml"
	basePath := "/test/path"
	result, _ := eclcloud.NormalizePathURL(basePath, rawPath)
	expected := strings.Join([]string{"file:/", filepath.ToSlash(baseDir), "template.yaml"}, "/")
	th.CheckEquals(t, expected, result)

	rawPath = "http://www.google.com"
	basePath = "/test/path"
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = "http://www.google.com"
	th.CheckEquals(t, expected, result)

	rawPath = "very/nested/file.yaml"
	basePath = "/test/path"
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = strings.Join([]string{"file:/", filepath.ToSlash(baseDir), "very/nested/file.yaml"}, "/")
	th.CheckEquals(t, expected, result)

	rawPath = "very/nested/file.yaml"
	basePath = "http://www.google.com"
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = "http://www.google.com/very/nested/file.yaml"
	th.CheckEquals(t, expected, result)

	rawPath = "very/nested/file.yaml/"
	basePath = "http://www.google.com/"
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = "http://www.google.com/very/nested/file.yaml"
	th.CheckEquals(t, expected, result)

	rawPath = "very/nested/file.yaml"
	basePath = "http://www.google.com/even/more"
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = "http://www.google.com/even/more/very/nested/file.yaml"
	th.CheckEquals(t, expected, result)

	rawPath = "very/nested/file.yaml"
	basePath = strings.Join([]string{"file:/", filepath.ToSlash(baseDir), "only/file/even/more"}, "/")
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = strings.Join([]string{"file:/", filepath.ToSlash(baseDir), "only/file/even/more/very/nested/file.yaml"}, "/")
	th.CheckEquals(t, expected, result)

	rawPath = "very/nested/file.yaml/"
	basePath = strings.Join([]string{"file:/", filepath.ToSlash(baseDir), "only/file/even/more"}, "/")
	result, _ = eclcloud.NormalizePathURL(basePath, rawPath)
	expected = strings.Join([]string{"file:/", filepath.ToSlash(baseDir), "only/file/even/more/very/nested/file.yaml"}, "/")
	th.CheckEquals(t, expected, result)

}
