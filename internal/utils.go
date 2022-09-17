package internal

import (
	"fmt"

	imodels "github.com/megalypse/zhttp/internal/models"
	"github.com/megalypse/zhttp/models"
)

func PrepareClientRequest[T any](client *models.ZHttpClient, request *imodels.ZRequest[T]) {
	context := client.Context
	request.Url = generateRequestUrl(context, request.Url)

	authenticateRequest(client, request)
}

func authenticateRequest[T any](client *models.ZHttpClient, request *imodels.ZRequest[T]) {
	if client.BearerToken != "" {
		authorizationString := fmt.Sprintf("Bearer %v", client.BearerToken)
		request.Headers = append(request.Headers, imodels.ZHeader{
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
