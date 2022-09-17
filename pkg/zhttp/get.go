package services

import (
	imodels "github.com/megalypse/zhttp/internal/models"
	iservices "github.com/megalypse/zhttp/internal/services"
)

func Get[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("GET", request)
}
