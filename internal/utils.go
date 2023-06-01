package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/megalypse/zhttp/zmodels"
)

// This function do everything that is needed to the URL url before it is ready
// to be used.
func PrepareClientRequest[T any](client *zmodels.ZClient, request *zmodels.ZRequest[T]) {
	context := client.ContextUrl
	request.Url = generateRequestUrl(context, request.Url)
	request.Headers = prepareRequestHeaders(client, request)
}

// This function makes a ZResponse value enforcing an error template.
func MakeFailResponse[T any](message string, httpResponse *http.Response) zmodels.ZResponse[T] {
	return makeResponse[T](nil, httpResponse, false, message)
}

// It takes the UrlParams and QueryParams maps to generate the parsed string
// with the url params replaced by its values and the query params appended
// to the end of the string.
func ParseUrl[T any](request zmodels.ZRequest[T]) string {
	uri := request.Url
	urlParams := request.UrlParams
	queryParams := request.QueryParams

	for key, value := range urlParams {
		curlyBracketsParam := fmt.Sprintf("{%v}", key)
		colonParam := fmt.Sprintf(":%v", key)

		uri = strings.ReplaceAll(uri, curlyBracketsParam, value)
		uri = strings.ReplaceAll(uri, colonParam, value)
	}

	urlLastIndex := len(uri) - 1

	if string(uri[urlLastIndex]) == "/" {
		uri = uri[:urlLastIndex]
	}

	uri += "?"

	isFirstParam := true
	for key, valueList := range queryParams {

		for _, value := range valueList {
			var param string

			if !isFirstParam {
				param += "&"
				isFirstParam = false
			}

			param += fmt.Sprintf("%v=%v", key, url.QueryEscape(value))
			uri += param
		}

		if isFirstParam {
			isFirstParam = false
		}
	}

	return uri
}

// It takes a host URL and an URI and put them together to form
// a valid URL.
func generateRequestUrl(context, uri string) string {
	contextLastIndex := len(context) - 1

	if string(context[contextLastIndex]) == "/" {
		context = context[:contextLastIndex]
	}

	if string(uri[0]) != "/" {
		uri = "/" + uri
	}

	return context + uri
}

func makeResponse[T any](
	content *T,
	response *http.Response,
	isSuccess bool,
	errorMessage string,
) zmodels.ZResponse[T] {
	return zmodels.ZResponse[T]{
		Content:      content,
		Response:     response,
		IsSuccess:    isSuccess,
		ErrorMessage: errorMessage,
	}
}

func prepareRequestHeaders[T any](client *zmodels.ZClient, request *zmodels.ZRequest[T]) map[string]string {
	clientHeadersSize := len(client.Headers)
	requestHeadersSize := len(request.Headers)

	newHeaders := make(map[string]string, clientHeadersSize+requestHeadersSize)

	for k, v := range client.Headers {
		newHeaders[k] = v
	}

	for k, v := range request.Headers {
		newHeaders[k] = v
	}

	return newHeaders
}
