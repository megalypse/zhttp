package internal

import (
	"fmt"
	"net/http"
	"strings"

	"net/url"

	"github.com/megalypse/zhttp/models"
)

func PrepareClientRequest[T any](client *models.ZClient, request *models.ZRequest[T]) {
	context := client.Context
	request.Url = generateRequestUrl(context, request.Url)
}

func MakeFailResponse[T any](message string, httpResponse *http.Response) models.ZResponse[T] {
	return makeResponse[T](nil, httpResponse, false, message)
}

func ParseUrl[T any](request models.ZRequest[T]) string {
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

			param += fmt.Sprintf("%v=%v", key, value)
			param = url.QueryEscape(param)
			uri += param
		}

		if isFirstParam {
			isFirstParam = false
		}
	}

	return uri
}

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
) models.ZResponse[T] {
	return models.ZResponse[T]{
		Content:      content,
		Response:     response,
		IsSuccess:    isSuccess,
		ErrorMessage: errorMessage,
	}
}
