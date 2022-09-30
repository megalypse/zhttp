package services

import (
	"bytes"
	"net/http"
	"testing"

	utils "github.com/megalypse/zhttp/internal"
	"github.com/megalypse/zhttp/zmodels"
)

var test *testing.T

func TestHaveDefaultContentTypeHeader(t *testing.T) {
	request := zmodels.ZRequest[zmodels.Void]{
		Url: "https://mock-url.com",
		Headers: map[string]string{
			"Test-header1": "value1",
			"Test-header2": "value2",
			"Test-header3": "value3",
		},
	}

	httpRequest, _ := http.NewRequest(
		"POST",
		utils.ParseUrl(request),
		bytes.NewBuffer(nil),
	)

	setRequestHeaders(&request, httpRequest)

	assertEquals(4, len(httpRequest.Header))
	assertEquals("application/json", httpRequest.Header.Get("Content-type"))
}

func TestDefaultContentTypeHeaderGotOverriden(t *testing.T) {
	request := zmodels.ZRequest[zmodels.Void]{
		Url: "https://mock-url.com",
		Headers: map[string]string{
			"Test-header1": "value1",
			"Test-header2": "value2",
			"Test-header3": "value3",
			"Content-type": "value4",
		},
	}

	httpRequest, _ := http.NewRequest(
		"POST",
		utils.ParseUrl(request),
		bytes.NewBuffer(nil),
	)

	setRequestHeaders(&request, httpRequest)

	assertEquals(4, len(httpRequest.Header))
	assertEquals("value4", httpRequest.Header.Get("Content-type"))
}

func assertEquals[T comparable](expected, actual T) {
	if expected != actual {
		test.Errorf("Was expecting %v, got %v", expected, actual)
	}
}
