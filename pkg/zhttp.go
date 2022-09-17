package pkg

import (
	iservices "github.com/megalypse/zhttp/internal/services"

	utils "github.com/megalypse/zhttp/internal"
	"github.com/megalypse/zhttp/zmodels"
)

func Get[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("GET", request)
}

func Post[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("POST", request)
}

func Patch[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("PATCH", request)
}

func Put[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("PUT", request)
}

func Delete[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("DELETE", request)
}

func ClientGet[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return Get[Response](request)
}

func ClientPost[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return Post[Response](request)
}

func ClientPatch[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return Patch[Response](request)
}

func ClientPut[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return Put[Response](request)
}

func ClientDelete[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return Delete[Response](request)
}
