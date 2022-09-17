package pkg

import (
	imodels "github.com/megalypse/zhttp/internal/models"
	iservices "github.com/megalypse/zhttp/internal/services"
	utils "github.com/megalypse/zhttp/pkg/internal"
	models "github.com/megalypse/zhttp/pkg/models"
)

func Get[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("GET", request)
}

func Post[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("POST", request)
}

func Patch[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("PATCH]", request)
}

func Put[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("PUT", request)
}

func Delete[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("DELETE", request)
}

func ClientGet[Response any, Request any](
	client models.ZHttpClient,
	request imodels.ZRequest[Request],
) imodels.ZResponse[Response] {
	context := client.Context
	request.Url = utils.GenerateRequestUrl(context, request.Url)

	return Get[Response](request)
}

func ClientPost[Response any, Request any](
	client models.ZHttpClient,
	request imodels.ZRequest[Request],
) imodels.ZResponse[Response] {
	context := client.Context
	request.Url = utils.GenerateRequestUrl(context, request.Url)

	return Post[Response](request)
}

func ClientPatch[Response any, Request any](
	client models.ZHttpClient,
	request imodels.ZRequest[Request],
) imodels.ZResponse[Response] {
	context := client.Context
	request.Url = utils.GenerateRequestUrl(context, request.Url)

	return Patch[Response](request)
}

func ClientPut[Response any, Request any](
	client models.ZHttpClient,
	request imodels.ZRequest[Request],
) imodels.ZResponse[Response] {
	context := client.Context
	request.Url = utils.GenerateRequestUrl(context, request.Url)

	return Put[Response](request)
}

func ClientDelete[Response any, Request any](
	client models.ZHttpClient,
	request imodels.ZRequest[Request],
) imodels.ZResponse[Response] {
	context := client.Context
	request.Url = utils.GenerateRequestUrl(context, request.Url)

	return Delete[Response](request)
}
