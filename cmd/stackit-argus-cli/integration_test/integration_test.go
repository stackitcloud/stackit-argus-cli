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

func TestGetCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)

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
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test correct case
	runCmd(server.URL, "", "", t, nullFile, true)

	// test correct case with output flag
	c.SetArgs([]string{"get", "alertGroup", "-o", "wide"})
	runCmd(server.URL, "", "", t, nullFile, true)

	// test with wrong command arguments
	c.SetArgs([]string{"get", "acl", "sdf"})
	runCmd(server.URL, "", "", t, nullFile, false)

	// test with wrong command flag
	c.SetArgs([]string{"get", "acl", "-t"})
	runCmd(server.URL, "", "", t, nullFile, false)
}

func TestDeleteCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)

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
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test correct case
	runCmd(server.URL, "", "", t, nullFile, true)

	// test without delete target
	c.SetArgs([]string{"delete", "alertGroup"})
	runCmd(server.URL, "", "", t, nullFile, false)
}

func TestCreateCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)

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
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test correct case
	runCmd(server.URL, "", "", t, nullFile, true)

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

func TestUpdateCommand(t *testing.T) {
	// Open /dev/null as a file
	nullFile, err := os.Open(os.DevNull)
	assert.NoError(t, err)

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
	defer c.SetArgs([]string{})

	// test bad auth
	runCmd(server.URL, "token", "wrong_token", t, nullFile, false)

	// test with wrong url
	runCmd(server.URL, "instance_id", "wrong_instance_id", t, nullFile, false)

	// test correct case
	runCmd(server.URL, "", "", t, nullFile, true)

	// test without body file
	c.SetArgs([]string{"update", "acl"})
	runCmd(server.URL, "", "", t, nullFile, false)
}
