package zhttp

import (
	"github.com/megalypse/zhttp/internal/services"
	"github.com/megalypse/zhttp/models"
)

// This function is just a wrapper enforcing the GET method on MakeRequest function
func Get[Response any, Request any](
	request models.ZRequest[Request],
	interceptors models.InterceptorsWrapper[Request, Response],
) models.ZResponse[Response] {
	return services.MakeRequest[Response]("GET", request, interceptors)
}

// This function is just a wrapper enforcing the POST method on MakeRequest function
func Post[Response any, Request any](
	request models.ZRequest[Request],
	interceptors models.InterceptorsWrapper[Request, Response],
) models.ZResponse[Response] {
	return services.MakeRequest[Response]("POST", request, interceptors)
}

// This function is just a wrapper enforcing the PATCH method on MakeRequest function
func Patch[Response any, Request any](
	request models.ZRequest[Request],
	interceptors models.InterceptorsWrapper[Request, Response],
) models.ZResponse[Response] {
	return services.MakeRequest[Response]("PATCH", request, interceptors)
}

// This function is just a wrapper enforcing the PUT method on MakeRequest function
func Put[Response any, Request any](
	request models.ZRequest[Request],
	interceptors models.InterceptorsWrapper[Request, Response],
) models.ZResponse[Response] {
	return services.MakeRequest[Response]("PUT", request, interceptors)
}

// This function is just a wrapper enforcing the DELETE method on MakeRequest function
func Delete[Response any, Request any](
	request models.ZRequest[Request],
	interceptors models.InterceptorsWrapper[Request, Response],
) models.ZResponse[Response] {
	return services.MakeRequest[Response]("DELETE", request, interceptors)
}
