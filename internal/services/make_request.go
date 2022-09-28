package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	utils "github.com/megalypse/zhttp/internal"
	"github.com/megalypse/zhttp/zmodels"
)

// Response can be of any desired type.
// Request can also be of any type.
// `MakeRequest` uses "encoding/json" lib, so feel free to use struct tagging on your response and request types.
func MakeRequest[Response any, Request any](method string, request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	if method == "POSTFORM" {
		converted := any(request).(zmodels.ZRequest[map[string][]string])
		return postFormBehavior[Response](converted)
	} else {
		return defaultBehavior[Response](method, request)
	}
}

func postFormBehavior[Response any](request zmodels.ZRequest[map[string][]string]) zmodels.ZResponse[Response] {
	responseHolder := new(Response)
	res, err := http.PostForm(request.Url, request.Body)

	if err != nil {
		return utils.MakeFailResponse[Response](err.Error(), res)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return utils.MakeFailResponse[Response](err.Error(), res)
	}

	json.Unmarshal(body, &responseHolder)

	statusCode := res.StatusCode
	return zmodels.ZResponse[Response]{
		Content:   responseHolder,
		Response:  res,
		IsSuccess: statusCode >= 200 && statusCode < 300,
	}
}

func defaultBehavior[Response any, Request any](method string, request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	responseHolder := new(Response)
	client := http.Client{}

	var bytesBody []byte
	if request.IsBytesBody {
		bytesBody = any(request.Body).([]byte)
	} else {
		bodyBuffer, marshalErr := json.Marshal(request.Body)

		if marshalErr != nil {
			return utils.MakeFailResponse[Response](marshalErr.Error(), nil)
		}

		bytesBody = bodyBuffer
	}

	httpRequest, _ := http.NewRequest(
		method,
		utils.ParseUrl(request),
		bytes.NewBuffer(bytesBody),
	)

	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}

	httpResponse, _ := client.Do(httpRequest)

	responseBuffer, readErr := io.ReadAll(httpResponse.Body)

	if readErr != nil {
		return utils.MakeFailResponse[Response](readErr.Error(), nil)
	}

	unmarshalError := json.Unmarshal(responseBuffer, &responseHolder)

	if unmarshalError != nil {
		return utils.MakeFailResponse[Response](unmarshalError.Error(), nil)
	}

	statusCode := httpResponse.StatusCode
	return zmodels.ZResponse[Response]{
		Content:   responseHolder,
		Response:  httpResponse,
		IsSuccess: statusCode >= 200 && statusCode < 300,
	}
}
