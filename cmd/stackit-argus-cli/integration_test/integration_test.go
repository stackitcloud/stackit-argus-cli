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
	c.SetArgs([]string{"delete", "alertGroup", "group"})
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
	c.SetArgs([]string{"create", "backupRestore", "date", "-t", "alertRules", "-f", ".test.yaml"})
	err = c.Flags().Set("config", "./.test.yaml")
	assert.NoError(t, err)
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test without targets
	c.SetArgs([]string{"create", "backupRestore", "date", "-f", ".test.yaml"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// test without argument
	c.SetArgs([]string{"create", "backupRestore", "-f", ".test.yaml"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// test without body file
	c.SetArgs([]string{"create", "backupRestore", "date", "-t", "alertRules"})
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
			arguments: []string{"get", "alertConfigs"},
			path:      "/projects/project/instances/instance/alertconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alertGroups"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alertGroups", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alertRules", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alertRules", "group", "alert"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     true,
		},
		{
			arguments: []string{"get", "alertRules"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"get", "backupRetentions"},
			path:      "/projects/project/instances/instance/backup-retentions",
			noErr:     true,
		},
		{
			arguments: []string{"get", "backups"},
			path:      "/projects/project/instances/instance/backups",
			noErr:     true,
		},
		{
			arguments: []string{"get", "backupSchedules"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     true,
		},
		{
			arguments: []string{"get", "certCheck"},
			path:      "/projects/project/instances/instance/cert-checks",
			noErr:     true,
		},
		{
			arguments: []string{"get", "credentials"},
			path:      "/projects/project/instances/instance/credentials",
			noErr:     true,
		},
		{
			arguments: []string{"get", "grafanaConfigs"},
			path:      "/projects/project/instances/instance/grafana-configs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "httpCheck"},
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
			arguments: []string{"get", "logsAlertGroups"},
			path:      "/projects/project/instances/instance/logs-alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"get", "logsAlertGroups", "group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"get", "logsConfigs"},
			path:      "/projects/project/instances/instance/logs-configs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "metricsStorageRetention"},
			path:      "/projects/project/instances/instance/metrics-storage-retentions",
			noErr:     true,
		},
		{
			arguments: []string{"get", "networkCheck"},
			path:      "/projects/project/instances/instance/network-checks",
			noErr:     true,
		},
		{
			arguments: []string{"get", "offerings"},
			path:      "/projects/project/offerings",
			noErr:     true,
		},
		{
			arguments: []string{"get", "pingCheck"},
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
			arguments: []string{"get", "scrapeConfigs"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"get", "scrapeConfigs", "job"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     true,
		},
		{
			arguments: []string{"get", "tracesConfigs"},
			path:      "/projects/project/instances/instance/traces-configs",
			noErr:     true,
		},
		// delete commands
		{
			arguments: []string{"delete", "alertGroup", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "alertGroup"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alertRecord", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "alertRecord", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alertRecord"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alertRule", "group", "alert"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "alertRule", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "alertRule"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "certCheck", "source"},
			path:      "/projects/project/instances/instance/cert-checks/source",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "certCheck"},
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
			arguments: []string{"delete", "httpCheck", "url"},
			path:      "/projects/project/instances/instance/http-checks/url",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "httpCheck"},
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
			arguments: []string{"delete", "logsAlertGroup", "group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "logsAlertGroup"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "networkCheck", "address"},
			path:      "/projects/project/instances/instance/network-checks/address",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "networkCheck"},
			path:      "/projects/project/instances/instance/network-checks/address",
			noErr:     false,
		},
		{
			arguments: []string{"delete", "pingCheck", "domain"},
			path:      "/projects/project/instances/instance/ping-checks/domain",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "pingCheck"},
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
			arguments: []string{"delete", "scrapeConfig", "job"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     true,
		},
		{
			arguments: []string{"delete", "scrapeConfig"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     false,
		},
		// create commands
		{
			arguments: []string{"create", "alertGroup", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"create", "alertGroup"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alertRecord", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     true,
		},
		{
			arguments: []string{"create", "alertRecord", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alertRecord"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alertRule", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     true,
		},
		{
			arguments: []string{"create", "alertRule", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "alertRule"},
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
			arguments: []string{"create", "backupRestore", "date", "-t", "alertConfig"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     true,
		},
		{
			arguments: []string{"create", "backupRestore", "date", "-t", "alertConfig", "-t", "alertRules"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backupRestore", "date", "-t", "alertCon"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backupRestore", "date"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backupRestore", "-t", "alertConfig"},
			path:      "/projects/project/instances/instance/backup-restores/date",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backupSchedule", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     true,
		},
		{
			arguments: []string{"create", "backupSchedule", "-t", "alertConfig", "-t", "alertRules", "-t", "scrapeConfig", "-t",
				"grafana", "-f", "./.test.yaml"},
			path:  "/projects/project/instances/instance/backup-schedules",
			noErr: true,
		},
		{
			arguments: []string{"create", "backupSchedule", "-t", "alertCo", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "backupSchedule"},
			path:      "/projects/project/instances/instance/backup-schedules",
			noErr:     false,
		},
		{
			arguments: []string{"create", "certCheck", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/cert-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "certCheck"},
			path:      "/projects/project/instances/instance/cert-checks",
			noErr:     false,
		},
		{
			arguments: []string{"create", "credentials"},
			path:      "/projects/project/instances/instance/credentials",
			noErr:     true,
		},
		{
			arguments: []string{"create", "httpCheck", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/http-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "httpCheck"},
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
			arguments: []string{"create", "logsAlertGroup", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"create", "logsAlertGroup"},
			path:      "/projects/project/instances/instance/logs-alertgroups",
			noErr:     false,
		},
		{
			arguments: []string{"create", "networkCheck", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/network-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "networkCheck"},
			path:      "/projects/project/instances/instance/network-checks",
			noErr:     false,
		},
		{
			arguments: []string{"create", "pingCheck", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/ping-checks",
			noErr:     true,
		},
		{
			arguments: []string{"create", "pingCheck"},
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
			arguments: []string{"create", "scrapeConfig", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"create", "scrapeConfig"},
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
			arguments: []string{"update", "alertConfig", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertconfigs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertConfig"},
			path:      "/projects/project/instances/instance/alertconfigs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertGroups", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertGroups", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertGroups"},
			path:      "/projects/project/instances/instance/alertgroups",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertGroups", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertRecords", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertRecords", "group", "record", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertRecords", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertRecords", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/records/record",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertRecords", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/records",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertRules", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertRules", "group", "alert", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/alert",
			noErr:     true,
		},
		{
			arguments: []string{"update", "alertRules", "group"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertRules", "group", "record"},
			path:      "/projects/project/instances/instance/alertgroups/group/alertrules/record",
			noErr:     false,
		},
		{
			arguments: []string{"update", "alertRules", "-f", "./.test.yaml"},
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
			arguments: []string{"update", "grafanaConfig", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/grafana-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "grafanaConfig"},
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
			arguments: []string{"update", "logsAlertGroup", "group", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     true,
		},
		{
			arguments: []string{"update", "logsAlertGroup", "group"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"update", "logsAlertGroup", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-alertgroups/group",
			noErr:     false,
		},
		{
			arguments: []string{"update", "logsConfigs", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/logs-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "logsConfigs"},
			path:      "/projects/project/instances/instance/logs-configs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "metricsStorageRetention", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/metrics-storage-retentions",
			noErr:     true,
		},
		{
			arguments: []string{"update", "metricsStorageRetention"},
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
			arguments: []string{"update", "scrapeConfigs", "job", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     true,
		},
		{
			// error because of unmarshaling, should be like that
			arguments: []string{"update", "scrapeConfigs", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "scrapeConfigs"},
			path:      "/projects/project/instances/instance/scrapeconfigs",
			noErr:     false,
		},
		{
			arguments: []string{"update", "scrapeConfigs", "job"},
			path:      "/projects/project/instances/instance/scrapeconfigs/job",
			noErr:     false,
		},
		{
			arguments: []string{"update", "tracesConfig", "-f", "./.test.yaml"},
			path:      "/projects/project/instances/instance/traces-configs",
			noErr:     true,
		},
		{
			arguments: []string{"update", "tracesConfig"},
			path:      "/projects/project/instances/instance/traces-configs",
			noErr:     false,
		},
	}

	for _, test := range tests {
		testLoop(test, t)
	}
}
