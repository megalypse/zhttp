package zclient

import (
	"github.com/megalypse/zhttp"
	utils "github.com/megalypse/zhttp/internal"
	"github.com/megalypse/zhttp/zmodels"
)

// Get function that uses predermined data from the given client
func Get[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Get[Response](request)
}

// Post function that makes the request using basic information from the given client
func Post[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Post[Response](request)
}

func Patch[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Patch[Response](request)
}

func Put[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Put[Response](request)
}

func Delete[Response any, Request any](
	client zmodels.ZClient,
	request zmodels.ZRequest[Request],
) zmodels.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Delete[Response](request)
}
