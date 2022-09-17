package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	curl "net/url"
	"strings"

	"github.com/megalypse/zhttp/models"
)

func MakeRequest[Response any, Request any](method string, request models.ZRequest[Request]) models.ZResponse[Response] {
	responseHolder := new(Response)
	client := http.Client{}

	bodyBuffer, marshalErr := json.Marshal(request.Body)

	if marshalErr != nil {
		return models.MakeFailResponse[Response](marshalErr.Error(), nil)
	}

	httpRequest, _ := http.NewRequest(
		method,
		parseUrl(request),
		bytes.NewBuffer(bodyBuffer),
	)

	for _, header := range request.Headers {
		httpRequest.Header.Set(header.Key, header.Value)
	}

	httpResponse, _ := client.Do(httpRequest)

	responseBuffer, readErr := io.ReadAll(httpResponse.Body)

	if readErr != nil {
		return models.MakeFailResponse[Response](marshalErr.Error(), nil)
	}

	unmarshalError := json.Unmarshal(responseBuffer, &responseHolder)

	if unmarshalError != nil {
		return models.MakeFailResponse[Response](unmarshalError.Error(), nil)
	}

	return models.ZResponse[Response]{
		Content:   responseHolder,
		Response:  httpResponse,
		IsSuccess: true,
	}
}

func parseUrl[T any](request models.ZRequest[T]) string {
	url := request.Url
	urlParams := request.UrlParams
	queryParams := request.QueryParams

	for _, param := range urlParams {
		paramInterpolation := fmt.Sprintf("{%v}", param.Key)

		url = strings.ReplaceAll(url, paramInterpolation, param.Value)
	}

	urlLastIndex := len(url) - 1

	if string(url[urlLastIndex]) == "/" {
		url = url[:urlLastIndex]
	}

	url += "?"

	for i, v := range queryParams {
		var param string

		if i > 0 {
			param += "&"
		}

		param += fmt.Sprintf("%v=%v", v.Key, v.Value)
		param = curl.QueryEscape(param)

		url += param
	}

	return url
}
