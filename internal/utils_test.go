package internal

import (
	"testing"

	"github.com/megalypse/zhttp/zmodels"
)

var test *testing.T

func TestGenerateRequestUrl(t *testing.T) {
	expected := "http://test.com/v1/users"

	testUrlGeneration("http://test.com/", "/v1/users", expected)
	testUrlGeneration("http://test.com", "/v1/users", expected)
	testUrlGeneration("http://test.com/", "v1/users", expected)
	testUrlGeneration("http://test.com", "v1/users", expected)

	testUrlGeneration("http://test.com//", "v1/users", "http://test.com//v1/users")
}

func TestParseUrl(t *testing.T) {
	request := zmodels.ZRequest[string]{
		Url: "http://test.com/v1/users/{userId}/:addressId",
		UrlParams: map[string]string{
			"userId":    "1",
			"addressId": "999",
		},
		QueryParams: map[string][]string{
			"firstParam": {"huee"},
			"listParam":  {"value1", "value2"},
		},
	}

	result := ParseUrl(request)
	expectedUrl := "http://test.com/v1/users/1/999?firstParam%3Dhuee%26listParam%3Dvalue1%26listParam%3Dvalue2"

	if result != expectedUrl {
		t.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expectedUrl, result)

	}
}

func testUrlGeneration(host, uri, expected string) {

	result := generateRequestUrl(host, uri)

	if result != expected {
		test.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expected, result)
	}
}
