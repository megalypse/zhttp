package zhttp

import (
	"github.com/megalypse/zhttp/internal/services"
	imodels "github.com/megalypse/zhttp/models"
)

func Get[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return services.MakeRequest[Response]("GET", request)
}

func Post[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return services.MakeRequest[Response]("POST", request)
}

func Patch[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return services.MakeRequest[Response]("PATCH", request)
}

func Put[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return services.MakeRequest[Response]("PUT", request)
}

func Delete[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return services.MakeRequest[Response]("DELETE", request)
}
