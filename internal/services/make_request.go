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
		return utils.MakeFailResponse[Response](MakeRequestError+err.Error(), res)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return utils.MakeFailResponse[Response](ResponseBodyReadError+err.Error(), res)
	}

	json.Unmarshal(body, &responseHolder)

	statusCode := res.StatusCode
	isSuccess := statusCode >= 200 && statusCode < 300

	if !isSuccess {
		return utils.MakeFailResponse[Response](HttpStatusFailed+string(body), res)
	}

	return zmodels.ZResponse[Response]{
		Content:   responseHolder,
		Response:  res,
		IsSuccess: true,
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
			return utils.MakeFailResponse[Response](RequestBodyMarshalError+marshalErr.Error(), nil)
		}

		bytesBody = bodyBuffer
	}

	httpRequest, _ := http.NewRequest(
		method,
		utils.ParseUrl(request),
		bytes.NewBuffer(bytesBody),
	)

	setRequestHeaders(&request, httpRequest)

	httpResponse, requestErr := client.Do(httpRequest)

	if requestErr != nil {
		return utils.MakeFailResponse[Response](MakeRequestError+requestErr.Error(), nil)
	}

	responseBuffer, readErr := io.ReadAll(httpResponse.Body)

	if readErr != nil {
		return utils.MakeFailResponse[Response](ResponseBodyReadError+readErr.Error(), nil)
	}

	if len(responseBuffer) > 0 {
		unmarshalError := json.Unmarshal(responseBuffer, &responseHolder)

		if unmarshalError != nil {
			return utils.MakeFailResponse[Response](ResponseBodyUnmarshalError+unmarshalError.Error(), nil)
		}

		statusCode := httpResponse.StatusCode
		isSuccess := statusCode >= 200 && statusCode < 300

		if !isSuccess {
			return utils.MakeFailResponse[Response](HttpStatusFailed+string(responseBuffer), httpResponse)
		}
	}

	return zmodels.ZResponse[Response]{
		Content:   responseHolder,
		Response:  httpResponse,
		IsSuccess: true,
	}
}

func setRequestHeaders[T any](request *zmodels.ZRequest[T], httpRequest *http.Request) {
	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}

	if _, exists := request.Headers["Content-type"]; !exists {
		httpRequest.Header.Set("Content-type", "application/json")
	}
}

const (
	RequestBodyMarshalError    = "[REQUEST BODY MARSHAL ERROR] "
	MakeRequestError           = "[MAKE REQUEST ERROR] "
	ResponseBodyReadError      = "[RESPONSE BODY READ ERROR] "
	ResponseBodyUnmarshalError = "[RESPONSE BODY UNMARSHAL ERROR] "
	HttpStatusFailed           = "[HTTP STATUS FAILED] "
)
