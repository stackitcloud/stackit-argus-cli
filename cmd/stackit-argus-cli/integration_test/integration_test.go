package integration_test

import (
	"github.com/spf13/viper"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func runCmd(url, key, value string, t *testing.T, f *os.File, noErr bool) {
	viper.SetConfigFile("./.test.yaml")
	viper.Set("base_url", url)
	if key != "" {
		viper.Set(key, value)
	}

	stdout := os.Stdout
	os.Stdout = f

	err := cmd.Execute()
	if noErr {
		assert.NoError(t, err)
	} else {
		assert.Error(t, err)
	}
	viper.Reset()

	os.Stdout = stdout
}

// these tests are general for all get subcommands
func TestGetCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)

	defer func(nullFile *os.File) {
		err := nullFile.Close()
		assert.NoError(t, err)
	}(nullFile)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/projects/project/instances/instance/acl" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"value":"fixed"}`)); err != nil {
			t.Errorf("Cannot write a response, err: %s", err.Error())
		}
	}))
	defer server.Close()

	// get root command and set arguments
	c := cmd.NewArgusCliCmd()
	c.SetArgs([]string{"get", "acl"})
	err = c.Flags().Set("config", "./.test.yaml")
	assert.NoError(t, err)
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test correct case with output flag
	c.SetArgs([]string{"get", "acl", "-o", "wide"})
	runCmd(server.URL, "", "", t, nullFile, true)

	// test case with wrong output flag value
	c.SetArgs([]string{"get", "acl", "-o", "wrong"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// next tests are the same as for other commands, so they will be tested only with get command

	// test with wrong flag
	c.SetArgs([]string{"get", "acl", "-q"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// test with project flag
	c.SetArgs([]string{"get", "acl", "-p", "project"})
	// changing the value in config file to a wrong one to see if cli actually took project id value from the flag
	runCmd(server.URL, "project_id", "not_project", t, nullFile, true)

	// test with instance flag
	c.SetArgs([]string{"get", "acl", "-i", "instance"})
	// changing the value in config file to a wrong one to see if cli actually took instance id value from the flag
	runCmd(server.URL, "instance_id", "not_instance", t, nullFile, true)

	// test with multiple flags
	c.SetArgs([]string{"get", "acl", "-d", "-i", "instance", "-o", "wide"})
	runCmd(server.URL, "", "", t, nullFile, true)
}

// these tests are general for all delete subcommands
func TestDeleteCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)

	defer func(nullFile *os.File) {
		err := nullFile.Close()
		assert.NoError(t, err)
	}(nullFile)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/projects/project/instances/instance/alertgroups/group" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// get root command and set arguments
	c := cmd.NewArgusCliCmd()
	c.SetArgs([]string{"delete", "alert-group", "group"})
	err = c.Flags().Set("config", "./.test.yaml")
	assert.NoError(t, err)
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)
}

// these tests are general for all create subcommands
func TestCreateCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)
	defer func(nullFile *os.File) {
		err := nullFile.Close()
		assert.NoError(t, err)
	}(nullFile)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/projects/project/instances/instance/backup-restores/date" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		if r.URL.RawQuery != "restoreTarget=alertRules" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// get root command and set arguments
	c := cmd.NewArgusCliCmd()
	c.SetArgs([]string{"create", "backup-restore", "date", "-t", "alertRules", "-f", ".test.yaml"})
	err = c.Flags().Set("config", "./.test.yaml")
	assert.NoError(t, err)
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test without targets
	c.SetArgs([]string{"create", "backup-restore", "date", "-f", ".test.yaml"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// test without argument
	c.SetArgs([]string{"create", "backup-restore", "-f", ".test.yaml"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// test without body file
	c.SetArgs([]string{"create", "backup-restore", "date", "-t", "alertRules"})
	runCmd(server.URL, "", "", t, nullFile, true)
}

// these tests are general for all update subcommands
func TestUpdateCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)
	defer func(nullFile *os.File) {
		err := nullFile.Close()
		assert.NoError(t, err)
	}(nullFile)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/projects/project/instances/instance/acl" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		if r.Header.Get("Authorization") != "Bearer token" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// get root command and set arguments
	c := cmd.NewArgusCliCmd()
	c.SetArgs([]string{"update", "acl", "-f", ".test.yaml"})
	err = c.Flags().Set("config", "./.test.yaml")
	assert.NoError(t, err)
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test without body file
	c.SetArgs([]string{"update", "acl"})
	runCmd(server.URL, "", "", t, nullFile, false)
}

type testStruct struct {
	arguments []string
	path      string
	noErr     bool
}

func testLoop(test testStruct, t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)
	defer func(nullFile *os.File) {
		err := nullFile.Close()
		assert.NoError(t, err)
	}(nullFile)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != test.path {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"value":"fixed"}`)); err != nil {
			t.Errorf("Cannot write a response, err: %s", err.Error())
		}
	}))
	defer server.Close()

	// get root command and set arguments
	c := cmd.NewArgusCliCmd()
	c.SilenceUsage = true
	c.SetArgs(test.arguments)
	err = c.Flags().Set("config", "./.test.yaml")
	assert.NoError(t, err)
	defer c.SetArgs([]string{})

	// test working case
	runCmd(server.URL, "", "", t, nullFile, test.noErr)

	// test case with wrong number of arguments
	c.SetArgs(append(test.arguments, "some", "useless", "arguments"))
	runCmd(server.URL, "", "", t, nullFile, false)
}

func TestCommands(t *testing.T) {
	tests := []testStruct{
		// get commands
		{
			arguments: []string{"get", "acl"},
			path:      "/projects/project/instances/instance/acl",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alert-configs"},
			path:      "/projects/project/instances/instance/alertconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alert-groups"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alert-groups", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alert-rules", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alert-rules", "group", "alert"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alert-rules"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"get", "backup-retentions"},
			path:      "/projects/project/instances/instance/backup-retentions",
			noErr:     true,
		},
		{
			arguments: []string{"get", "backups"},
			path:      "/projects/project/instances/instance/backups",
			noErr:     true,
		},
		{
			arguments: []string{"get", "backup-schedules"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     true,
		},
		{
			arguments: []string{"get", "cert-check"},
			path:      "/projects/project/instances/instance/cert-checks",
			noErr:     true,
		},
		{
			arguments: []string{"get", "credentials"},
			path:      "/projects/project/instances/instance/credentials",
			noErr:     true,
		},
		{
			arguments: []string{"get", "grafana-configs"},
			path:      "/projects/project/instances/instance/grafana-configs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "http-check"},
			path:      "/projects/project/instances/instance/http-checks",
			noErr:     true,
		},
		{
			arguments: []string{"get", "instances"},
			path:      "/projects/project/instances",
			noErr:     true,
		},
		{
			arguments: []string{"get", "instances", "id"},
			path:      "/projects/project/instances/id",
			noErr:     true,
		},
		{
			arguments: []string{"get", "logs-alert-groups"},
			path:      "/projects/project/instances/instance/logs-alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"get", "logs-alert-groups", "group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"get", "logs-configs"},
			path:      "/projects/project/instances/instance/logs-configs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "metrics-storage-retention"},
			path:      "/projects/project/instances/instance/metrics-storage-retentions",
			noErr:     true,
		},
		{
			arguments: []string{"get", "network-check"},
			path:      "/projects/project/instances/instance/network-checks",
			noErr:     true,
		},
		{
			arguments: []string{"get", "offerings"},
			path:      "/projects/project/offerings",
			noErr:     true,
		},
		{
			arguments: []string{"get", "ping-check"},
			path:      "/projects/project/instances/instance/ping-checks",
			noErr:     true,
		},
		{
			arguments: []string{"get", "plans"},
			path:      "/projects/project/plans",
			noErr:     true,
		},
		{
			arguments: []string{"get", "receivers"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers",
			noErr:     true,
		},
		{
			arguments: []string{"get", "receivers", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers/receiver",
			noErr:     true,
		},
		{
			arguments: []string{"get", "records"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"get", "records", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     true,
		},
		{
			arguments: []string{"get", "records", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     true,
		},
		{
			arguments: []string{"get", "routes"},
			path:      "/projects/project/instances/instance/alertconfigs/routes",
			noErr:     true,
		},
		{
			arguments: []string{"get", "routes", "route"},
			path:      "/projects/project/instances/instance/alertconfigs/routes/route",
			noErr:     true,
		},
		{
			arguments: []string{"get", "scrape-configs"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "scrape-configs", "job"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     true,
		},
		{
			arguments: []string{"get", "traces-configs"},
			path:      "/projects/project/instances/instance/traces-configs",
			noErr:     true,
		},
		// delete commands
		{
			arguments: []string{"delete", "alert-group", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "alert-group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alert-record", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "alert-record", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alert-record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alert-rule", "group", "alert"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "alert-rule", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alert-rule"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "cert-check", "source"},
			path:      "/projects/project/instances/instance/cert-checks/source",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "cert-check"},
			path:      "/projects/project/instances/instance/cert-checks/source",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "credentials", "name"},
			path:      "/projects/project/instances/instance/credentials/name",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "credentials", "name", "-r"},
			path:      "/projects/project/instances/instance/credentials/name/remote-write-limits",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "credentials"},
			path:      "/projects/project/instances/instance/cert-checks/source",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "http-check", "url"},
			path:      "/projects/project/instances/instance/http-checks/url",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "http-check"},
			path:      "/projects/project/instances/instance/http-checks/source",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "instance", "instance"},
			path:      "/projects/project/instances/instance",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "instance"},
			path:      "/projects/project/instances/instance",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "logs-alert-group", "group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "logs-alert-group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "network-check", "address"},
			path:      "/projects/project/instances/instance/network-checks/address",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "network-check"},
			path:      "/projects/project/instances/instance/network-checks/address",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "ping-check", "domain"},
			path:      "/projects/project/instances/instance/ping-checks/domain",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "ping-check"},
			path:      "/projects/project/instances/instance/ping-checks/domain",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "receiver", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers/receiver",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers/receiver",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "route", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/routes/receiver",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "route"},
			path:      "/projects/project/instances/instance/alertconfigs/routes/receiver",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "scrape-config", "job"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "scrape-config"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     false,
		},
		// create commands
		{
			arguments: []string{"create", "alert-group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"create", "alert-group"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alert-record", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     true,
		},
		{
			arguments: []string{"create", "alert-record", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alert-record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alert-rule", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     true,
		},
		{
			arguments: []string{"create", "alert-rule", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alert-rule"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup"},
			path:      "/projects/project/instances/instance/backups",
			noErr:     true,
		},
		{
			arguments: []string{"create", "backup", "-t", "alertConfig", "-t", "alertRules", "-t", "scrapeConfig", "-t",
				"grafana", "-f", "./.test.yaml"},
			path:  "/projects/project/instances/instance/backups",
			noErr: true,
		},
		{
			arguments: []string{"create", "backup", "-t", "alertCo"},
			path:      "/projects/project/instances/instance/backups",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup-restore", "date", "-t", "alertConfig"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     true,
		},
		{
			arguments: []string{"create", "backup-restore", "date", "-t", "alertConfig", "-t", "alertRules"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup-restore", "date", "-t", "alertCon"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup-restore", "date"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup-restore", "-t", "alertConfig"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup-schedule", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     true,
		},
		{
			arguments: []string{"create", "backup-schedule", "-t", "alertConfig", "-t", "alertRules", "-t", "scrapeConfig", "-t",
				"grafana", "-f", "./.test.yaml"},
			path:  "/projects/project/instances/instance/backup-schedules",
			noErr: true,
		},
		{
			arguments: []string{"create", "backup-schedule", "-t", "alertCo", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backup-schedule"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "cert-check", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/cert-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "cert-check"},
			path:      "/projects/project/instances/instance/cert-checks",
			noErr:     false,
		},
		{
			arguments: []string{"create", "credentials"},
			path:      "/projects/project/instances/instance/credentials",
			noErr:     true,
		},
		{
			arguments: []string{"create", "http-check", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/http-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "http-check"},
			path:      "/projects/project/instances/instance/http-checks",
			noErr:     false,
		},
		{
			arguments: []string{"create", "instance", "-f", "./.test.yaml"},
			path:      "/projects/project/instances",
			noErr:     true,
		},
		{
			arguments: []string{"create", "instance"},
			path:      "/projects/project/instances",
			noErr:     false,
		},
		{
			arguments: []string{"create", "logs-alert-group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"create", "logs-alert-group"},
			path:      "/projects/project/instances/instance/logs-alertgroups",
			noErr:     false,
		},
		{
			arguments: []string{"create", "network-check", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/network-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "network-check"},
			path:      "/projects/project/instances/instance/network-checks",
			noErr:     false,
		},
		{
			arguments: []string{"create", "ping-check", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/ping-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "ping-check"},
			path:      "/projects/project/instances/instance/ping-checks",
			noErr:     false,
		},
		{
			arguments: []string{"create", "receiver", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers",
			noErr:     true,
		},
		{
			arguments: []string{"create", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers",
			noErr:     false,
		},
		{
			arguments: []string{"create", "route", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs/routes",
			noErr:     true,
		},
		{
			arguments: []string{"create", "route"},
			path:      "/projects/project/instances/instance/alertconfigs/routes",
			noErr:     false,
		},
		{
			arguments: []string{"create", "scrape-config", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"create", "scrape-config"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     false,
		},
		// update commands
		{
			arguments: []string{"update", "acl", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/acl",
			noErr:     true,
		},
		{
			arguments: []string{"update", "acl"},
			path:      "/projects/project/instances/instance/acl",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-config", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-config"},
			path:      "/projects/project/instances/instance/alertconfigs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-groups", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-groups", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-groups"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-groups", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-records", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-records", "group", "record", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-records", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-records", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-records", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-rules", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-rules", "group", "alert", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alert-rules", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-rules", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/record",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alert-rules", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"update", "credentials", "name", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/credentials/name/remote-write-limits",
			noErr:     true,
		},
		{
			arguments: []string{"update", "credentials", "name"},
			path:      "/projects/project/instances/instance/credentials/name/remote-write-limits",
			noErr:     true,
		},
		{
			arguments: []string{"update", "credentials"},
			path:      "/projects/project/instances/instance/credentials/name/remote-write-limits",
			noErr:     false,
		},
		{
			arguments: []string{"update", "grafana-config", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/grafana-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "grafana-config"},
			path:      "/projects/project/instances/instance/grafana-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "instance", "instance", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance",
			noErr:     true,
		},
		{
			arguments: []string{"update", "instance", "instance"},
			path:      "/projects/project/instances/instance",
			noErr:     false,
		},
		{
			arguments: []string{"update", "instance", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance",
			noErr:     false,
		},
		{
			arguments: []string{"update", "logs-alert-group", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"update", "logs-alert-group", "group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"update", "logs-alert-group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"update", "logs-configs", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "logs-configs"},
			path:      "/projects/project/instances/instance/logs-configs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "metrics-storage-retention", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/metrics-storage-retentions",
			noErr:     true,
		},
		{
			arguments: []string{"update", "metrics-storage-retention"},
			path:      "/projects/project/instances/instance/metrics-storage-retentions",
			noErr:     false,
		},
		{
			arguments: []string{"update", "receiver", "receiver", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers/receiver",
			noErr:     true,
		},
		{
			arguments: []string{"update", "receiver", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers/receiver",
			noErr:     false,
		},
		{
			arguments: []string{"update", "receiver", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs/receivers/receiver",
			noErr:     false,
		},
		{
			arguments: []string{"update", "route", "receiver", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs/routes/receiver",
			noErr:     true,
		},
		{
			arguments: []string{"update", "route", "receiver"},
			path:      "/projects/project/instances/instance/alertconfigs/routes/receiver",
			noErr:     false,
		},
		{
			arguments: []string{"update", "route", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs/routes/receiver",
			noErr:     false,
		},
		{
			arguments: []string{"update", "scrape-configs", "job", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     true,
		},
		{
			// error because of unmarshaling, should be like that
			arguments: []string{"update", "scrape-configs", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "scrape-configs"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "scrape-configs", "job"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     false,
		},
		{
			arguments: []string{"update", "traces-config", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/traces-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "traces-config"},
			path:      "/projects/project/instances/instance/traces-configs",
			noErr:     false,
		},
	}

	for _, test := range tests {
		testLoop(test, t)
	}
}
