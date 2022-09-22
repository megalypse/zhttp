package zhttp

import (
	"github.com/megalypse/zhttp/internal/services"
	"github.com/megalypse/zhttp/zmodels"
)

// This function is just a wrapper enforcing the GET method on MakeRequest function
func Get[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return services.MakeRequest[Response]("GET", request)
}

// This function is just a wrapper enforcing the POST method on MakeRequest function
func Post[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return services.MakeRequest[Response]("POST", request)
}

// This function is just a wrapper enforcing the POST method on MakeRequest function
func PostForm[Response any](request zmodels.ZRequest[map[string][]string]) zmodels.ZResponse[Response] {
	return services.MakeRequest[Response]("POSTFORM", request)
}

// This function is just a wrapper enforcing the PATCH method on MakeRequest function
func Patch[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return services.MakeRequest[Response]("PATCH", request)
}

// This function is just a wrapper enforcing the PUT method on MakeRequest function
func Put[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return services.MakeRequest[Response]("PUT", request)
}

// This function is just a wrapper enforcing the DELETE method on MakeRequest function
func Delete[Response any, Request any](request zmodels.ZRequest[Request]) zmodels.ZResponse[Response] {
	return services.MakeRequest[Response]("DELETE", request)
}
