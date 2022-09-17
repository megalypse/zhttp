package services

import (
	imodels "github.com/megalypse/zhttp/internal/models"
	iservices "github.com/megalypse/zhttp/internal/services"
)

func Patch[Response any, Request any](request imodels.ZRequest[Request]) imodels.ZResponse[Response] {
	return iservices.MakeRequest[Response]("PATCH]", request)
}
