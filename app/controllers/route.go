// Package controllers defines application's routes.
package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pottava/trivy-restapi/app/generated/models"
	"github.com/pottava/trivy-restapi/app/generated/restapi/operations"
	"github.com/pottava/trivy-restapi/app/generated/restapi/operations/image"
)

// Routes set API handlers
func Routes(api *operations.TrivyRestapiAPI) {
	api.ImageGetImageVulnerabilitiesHandler = image.GetImageVulnerabilitiesHandlerFunc(getVulnerabilities)
}

func getVulnerabilities(params image.GetImageVulnerabilitiesParams) middleware.Responder {
	result := []*models.Vulnerability{}
	return image.NewGetImageVulnerabilitiesOK().WithPayload(result)
}
