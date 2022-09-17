package internal

import (
	"fmt"

	"github.com/megalypse/zhttp/models"
)

func PrepareClientRequest[T any](client *models.ZHttpClient, request *models.ZRequest[T]) {
	context := client.Context
	request.Url = generateRequestUrl(context, request.Url)

	authenticateRequest(client, request)
}

func authenticateRequest[T any](client *models.ZHttpClient, request *models.ZRequest[T]) {
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
