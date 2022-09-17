package internal

import (
	"fmt"
	"net/http"

	"github.com/megalypse/zhttp/models"
)

func PrepareClientRequest[T any](client *models.ZClient, request *models.ZRequest[T]) {
	context := client.Context
	request.Url = generateRequestUrl(context, request.Url)

	authenticateRequest(client, request)
}

func MakeFailResponse[T any](message string, httpResponse *http.Response) models.ZResponse[T] {
	return makeResponse[T](nil, httpResponse, false, message)
}

func authenticateRequest[T any](client *models.ZClient, request *models.ZRequest[T]) {
	if client.BearerToken != "" {
		authorizationString := fmt.Sprintf("Bearer %v", client.BearerToken)
		request.Headers = append(request.Headers, models.ZHeader{
			Key:   "Authorization",
			Value: authorizationString,
		})
	}
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
