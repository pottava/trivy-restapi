// Package controllers defines application's routes.
package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pottava/trivy-restapi/app/generated/models"
	"github.com/pottava/trivy-restapi/app/generated/restapi/operations"
	"github.com/pottava/trivy-restapi/app/generated/restapi/operations/image"
	"github.com/pottava/trivy-restapi/app/logic"
)

// Routes set API handlers
func Routes(api *operations.TrivyRestapiAPI) {
	api.ImageGetImageVulnerabilitiesHandler = image.GetImageVulnerabilitiesHandlerFunc(getVulnerabilities)
}

func getVulnerabilities(params image.GetImageVulnerabilitiesParams) middleware.Responder {
	severities := "UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL"
	if len(params.Severity) > 0 {
		severities = strings.Join(params.Severity, ",")
	}
	ignoreUnfixed := false
	if strings.EqualFold(swag.StringValue(params.IgnoreUnfixed), "yes") {
		ignoreUnfixed = true
	}
	skipUpdate := false
	if strings.EqualFold(swag.StringValue(params.SkipUpdate), "yes") {
		skipUpdate = true
	}
	payload, err := logic.Scan(params.ID, severities, ignoreUnfixed, skipUpdate)
	if err != nil {
		log.Print(err)
		code := http.StatusBadRequest
		return image.NewGetImageVulnerabilitiesDefault(code).WithPayload(newerror(code))
	}
	return image.NewGetImageVulnerabilitiesOK().WithPayload(payload)
}

func newerror(code int) *models.Error {
	return &models.Error{
		Code:    swag.String(fmt.Sprintf("%d", code)),
		Message: swag.String(http.StatusText(code)),
	}
}
