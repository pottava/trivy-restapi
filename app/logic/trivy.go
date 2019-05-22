// Package logic defines business logic
package logic

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/go-openapi/swag"
	"github.com/pottava/trivy-restapi/app/generated/models"
	"github.com/pottava/trivy-restapi/app/lib"
)

var IsReady bool

func MakeVulnerabilityDatabase() {
	cmd := exec.Command("sh", "-c", fmt.Sprintf(
		"trivy --debug --cache-dir %s alpine:3.9",
		lib.Config.CacheDir,
	))
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		panic("Cannot execute trivy command")
	}
	bytes := make([]byte, 100)
	year := fmt.Sprintf("%d", time.Now().Year())
	for {
		_, err := stdout.Read(bytes)
		if err != nil {
			break
		}
		line, _, _ := bufio.NewReader(stdout).ReadLine()
		if strings.HasPrefix(string(line), year) {
			fmt.Println(string(line))
		}
	}
	cmd.Wait()
	IsReady = true
	lib.Logger.Info("Vulnerability database has been built")
}

func Scan(ctx context.Context, id, severities string, ignoreUnfixed, skipUpdate bool) (*models.Vulnerabilities, error) {
	options := ""
	if ignoreUnfixed {
		options += " --ignore-unfixed"
	}
	if skipUpdate {
		options += " --skip-update"
	}
	commands := fmt.Sprintf(
		"set -o pipefail && trivy --format=json --severity=%s %s "+
			"--cache-dir %s --exit-code 1 --quiet "+
			"%s | grep -v \"%s\" | jq",
		severities,
		options,
		lib.Config.CacheDir,
		id,
		"$( date '+%Y-%m-%dT' )",
	)
	result := &models.Vulnerabilities{
		Vulnerabilities: []*models.Vulnerability{},
		Count:           swag.Int64(0),
	}
	out, err := exec.CommandContext(ctx, "sh", "-c", commands).Output()

	vulnerabilities := map[string][]*models.Vulnerability{}
	if e := json.Unmarshal(out, &vulnerabilities); e == nil {
		for key := range vulnerabilities {
			result.Vulnerabilities = append(result.Vulnerabilities, vulnerabilities[key]...)
			result.Count = swag.Int64(swag.Int64Value(result.Count) + int64(len(vulnerabilities[key])))
		}
	}
	return result, err
}
