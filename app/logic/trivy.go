// Package logic defines business logic
package logic

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/go-openapi/swag"
	"github.com/pottava/trivy-restapi/app/generated/models"
)

func Scan(id, severities string, ignoreUnfixed, skipUpdate bool) (*models.Vulnerabilities, error) {
	options := ""
	if ignoreUnfixed {
		options += " --ignore-unfixed"
	}
	if skipUpdate {
		options += " --skip-update"
	}
	commands := fmt.Sprintf(
		"set -o pipefail && trivy --format=json --severity=%s %s "+
			"--quiet %s | grep -v \"%s\" | jq",
		severities,
		options,
		id,
		"$( date '+%Y-%m-%dT' )",
	)
	out, err := exec.Command("sh", "-c", commands).Output()
	if err != nil {
		return nil, err
	}
	vulnerabilities := map[string][]*models.Vulnerability{}
	err = json.Unmarshal(out, &vulnerabilities)
	if err != nil {
		return nil, err
	}
	result := &models.Vulnerabilities{
		Vulnerabilities: []*models.Vulnerability{},
		Count:           swag.Int64(0),
	}
	for key := range vulnerabilities {
		result.Vulnerabilities = append(result.Vulnerabilities, vulnerabilities[key]...)
		result.Count = swag.Int64(swag.Int64Value(result.Count) + int64(len(vulnerabilities[key])))
	}
	return result, nil
}
