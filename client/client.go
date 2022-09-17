package client

import (
	"github.com/megalypse/zhttp"
	utils "github.com/megalypse/zhttp/internal"
	"github.com/megalypse/zhttp/models"
)

func ClientGet[Response any, Request any](
	client models.ZClient,
	request models.ZRequest[Request],
) models.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Get[Response](request)
}

func ClientPost[Response any, Request any](
	client models.ZClient,
	request models.ZRequest[Request],
) models.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Post[Response](request)
}

func ClientPatch[Response any, Request any](
	client models.ZClient,
	request models.ZRequest[Request],
) models.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Patch[Response](request)
}

func ClientPut[Response any, Request any](
	client models.ZClient,
	request models.ZRequest[Request],
) models.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Put[Response](request)
}

func ClientDelete[Response any, Request any](
	client models.ZClient,
	request models.ZRequest[Request],
) models.ZResponse[Response] {
	utils.PrepareClientRequest(&client, &request)

	return zhttp.Delete[Response](request)
}
