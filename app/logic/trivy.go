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
		"trivy --debug --auto-refresh --clear-cache --cache-dir %s alpine:3.9",
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

type trivyResult struct {
	Target          string                  `json:"Target"`
	Vulnerabilities []*models.Vulnerability `json:"Vulnerabilities"`
}

func Scan(ctx context.Context, id, severities string, ignoreUnfixed, skipUpdate bool) (*models.Vulnerabilities, error) {
	options := ""
	if ignoreUnfixed {
		options += "--ignore-unfixed "
	}
	if skipUpdate {
		options += "--skip-update "
	} else {
		options += "--auto-refresh "
	}
	commands := fmt.Sprintf(
		"trivy --format=json --severity=%s %s --clear-cache --cache-dir=%s --quiet %s | grep -v \"%s\"",
		severities,
		options,
		lib.Config.CacheDir,
		id,
		"$( date '+%Y-%m' )",
	)
	out, err := exec.CommandContext(ctx, "sh", "-c", commands).Output()
	if err != nil {
		return nil, err
	}
	records := []trivyResult{}
	if err = json.Unmarshal(out, &records); err != nil {
		lib.Logger.Warn(commands)
		lib.Logger.Warn(string(out))
		return nil, err
	}
	result := &models.Vulnerabilities{
		Vulnerabilities: []*models.Vulnerability{},
		Count:           swag.Int64(0),
	}
	for _, record := range records {
		result.Vulnerabilities = append(result.Vulnerabilities, record.Vulnerabilities...)
		result.Count = swag.Int64(swag.Int64Value(result.Count) + int64(len(record.Vulnerabilities)))
	}
	return result, nil
}
