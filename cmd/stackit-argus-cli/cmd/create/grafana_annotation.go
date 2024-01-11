package create

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-argus-cli/internal/config"
	"github.com/stackitcloud/stackit-argus-cli/internal/utils"

	"github.com/spf13/cobra"
)

// annotation is a struct to marshal an annotation info
type annotation struct {
	DashboardUID string   `json:"dashboardUID,omitempty"`
	PanelId      int      `json:"panelId,omitempty"`
	Time         int64    `json:"time,omitempty"`
	TimeEnd      int64    `json:"timeEnd,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Text         string   `json:"text,omitempty"`
}

// dashboard is a struct to unmarshal dashboards list
type dashboard struct {
	Uid   string `json:"uid"`
	Title string `json:"title"`
}

// GrafanaAnnotationCmd represents the grafanaAnnotation command
var GrafanaAnnotationCmd = &cobra.Command{
	Use:   "grafana-annotation",
	Short: "create a grafana annotation",
	Example: "stackit-argus-cli create grafana-annotation <board name> <description> " +
		"--time=now --time-end=\"2006-01-02T15:04:05Z07:00\" --tags=\"tag1,tag2\" --panel-id=1",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// init authentication header for grafana api calls
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.GetGrafanaUsername()+":"+config.GetGrafanaPassword()))

		// get dashboard uid from a title to add annotation to dashboard by uid
		dashboardUID, err := getDashboardUID(config.GetGrafanaUrl(), authHeader, args[0])
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		a, err := setAnnotationInfo(dashboardUID, args[1])
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}
		body, err := json.Marshal(a)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if err := createAnnotation(config.GetGrafanaUrl(), authHeader, body); err != nil {
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}

func setAnnotationInfo(dashboardUID, text string) (*annotation, error) {
	var a annotation

	a.DashboardUID = dashboardUID
	a.Text = text
	a.Tags = config.GetTags()
	t, err := config.GetTimestamp()
	if err != nil {
		return nil, errors.New("timestamp is incorrect, err: " + err.Error())
	}
	a.Time = t
	t, err = config.GetEndTimestamp()
	if err != nil {
		return nil, errors.New("end timestamp is incorrect, err: " + err.Error())
	}
	if t != -1 {
		a.TimeEnd = t
	}
	panelId := config.GetPanelId()
	if panelId != -1 {
		a.PanelId = int(panelId)
	}

	return &a, nil
}

func createAnnotation(url, authHeader string, body []byte) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodPost, url+"/api/annotations", bytes.NewReader(body))
	if err != nil {
		return errors.New("failed to create a \"create annotation\" request, err: " + err.Error())
	}
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return errors.New("failed to make a \"create annotation\" request, err: " + err.Error())
	}

	if err := utils.ResponseMessage(res.StatusCode, "grafana annotation", req.Method, res.Body); err != nil {
		return err
	}

	if err := res.Body.Close(); err != nil {
		return err
	}

	return nil
}

func getDashboardUID(url, authHeader, title string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url+"/api/search", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if err := res.Body.Close(); err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("cannot get dashboards list, response: %s", res.Status)
	}

	var dashboards []dashboard
	if err := json.Unmarshal(body, &dashboards); err != nil {
		return "", err
	}

	for _, d := range dashboards {
		if d.Title == title {
			return d.Uid, nil
		}
	}
	return "", fmt.Errorf("dashboard \"%s\" doesn't exist", title)
}
