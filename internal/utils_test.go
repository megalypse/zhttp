package internal

import (
	"testing"

	"github.com/megalypse/zhttp/models"
)

func TestGenerateRequestUrl(t *testing.T) {
	context := "http://test.com/"
	uri := "/v1/users"

	result := generateRequestUrl(context, uri)
	expected := "http://test.com/v1/users"

	if result != expected {
		t.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expected, result)
	}

	context = "http://test.com"
	uri = "/v1/users"

	result2 := generateRequestUrl(context, uri)

	if result2 != expected {
		t.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expected, result)
	}

	context = "http://test.com/"
	uri = "v1/users"

	result3 := generateRequestUrl(context, uri)

	if result3 != expected {
		t.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expected, result)
	}

	context = "http://test.com"
	uri = "v1/users"

	result4 := generateRequestUrl(context, uri)

	if result4 != expected {
		t.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expected, result)
	}

	context = "http://test.com//"
	uri = "v1/users"

	result5 := generateRequestUrl(context, uri)
	expectedDefective := "http://test.com//v1/users"

	if result5 != expectedDefective {
		t.Errorf("Unexpected behavior on parsing. Expected %q, got %q.", expected, result)
	}
}

func TestParseUrl(t *testing.T) {
	request := models.ZRequest[string]{
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
