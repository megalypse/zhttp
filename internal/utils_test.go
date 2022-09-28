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

	assertEquals(expectedUrl, result)
}

func TestHeaderInsertion(t *testing.T) {
	original := "777"
	new := "000"

	client := zmodels.ZClient{
		ContextUrl: "http://hosturl.com/",
		Headers: map[string]string{
			"ClientHeader1": "-",
			"GenericHeader": original,
		},
	}

	request := zmodels.ZRequest[zmodels.Void]{
		Headers: map[string]string{
			"GenericHeader": new,
			"RequestHeader": "--",
		},
	}

	headers := prepareRequestHeaders(&client, &request)

	assertEquals(new, headers["GenericHeader"])
	assertEquals("-", headers["ClientHeader1"])
	assertEquals("--", headers["RequestHeader"])
	assertEquals(3, len(headers))
}

func testUrlGeneration(host, uri, expected string) {
	result := generateRequestUrl(host, uri)

	assertEquals(expected, result)
}

func assertEquals[T comparable](expected, actual T) {
	if expected != actual {
		test.Errorf("Was expecting %v, got %v", expected, actual)
	}
}
