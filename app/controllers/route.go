// Package controllers defines application's routes.
package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pottava/trivy-restapi/app/generated/models"
	"github.com/pottava/trivy-restapi/app/generated/restapi/operations"
	"github.com/pottava/trivy-restapi/app/generated/restapi/operations/image"
	"github.com/pottava/trivy-restapi/app/logic"
)

func init() {
	go logic.MakeVulnerabilityDatabase()
}

// Routes set API handlers
func Routes(api *operations.TrivyRestapiAPI) {
	api.ImageGetImageVulnerabilitiesHandler = image.GetImageVulnerabilitiesHandlerFunc(getVulnerabilities)
}

func getVulnerabilities(params image.GetImageVulnerabilitiesParams) middleware.Responder {
	if !logic.IsReady {
		code := http.StatusServiceUnavailable
		return image.NewGetImageVulnerabilitiesDefault(code).WithPayload(&models.Error{
			Code:    swag.String(fmt.Sprintf("%d", code)),
			Message: swag.String("Vulnerability database is not ready yet"),
		})
	}
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
	payload, _ := logic.Scan(
		params.HTTPRequest.Context(), params.ID, severities, ignoreUnfixed, skipUpdate,
	)
	return image.NewGetImageVulnerabilitiesOK().WithPayload(payload)
}
