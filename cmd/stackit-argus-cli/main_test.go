package main

import (
	"bytes"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func getStringFromStdout(r, w, old *os.File, t *testing.T) string {
	outC := make(chan string)

	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			return
		}
		outC <- buf.String()
	}()

	// back to normal state
	if err := w.Close(); err != nil {
		t.Fatal("cannot close a file")
	}
	os.Stdout = old // restoring the real stdout
	out := <-outC

	return out
}

func TestCommandsCalls(t *testing.T) {
	tests := []struct {
		description    string
		arguments      []string
		err            bool
		expectedOutput string
	}{
		{
			description:    "test get scrape config list command",
			arguments:      []string{"get", "scrapeConfigs"},
			err:            false,
			expectedOutput: "get scrape config\n",
		},
		{
			description:    "test get scrape config command",
			arguments:      []string{"get", "scrapeConfigs", "config"},
			err:            false,
			expectedOutput: "get scrape config\n",
		},
		{
			description:    "test get scrape config command with wrong number of arguments",
			arguments:      []string{"get", "scrapeConfigs", "config", "xd"},
			err:            true,
			expectedOutput: "get scrape config\n",
		},
		{
			description:    "test get scrape config command with output flag",
			arguments:      []string{"get", "scrapeConfigs", "config", "-o", "json"},
			err:            false,
			expectedOutput: "get scrape config\n",
		},
		{
			description:    "test get scrape config command with wrong output flag",
			arguments:      []string{"get", "scrapeConfigs", "config", "-o", "js"},
			err:            true,
			expectedOutput: "get scrape config\n",
		},
		{
			description:    "test create scrape config command without -f flag",
			arguments:      []string{"create", "scrapeConfigs"},
			err:            true,
			expectedOutput: "create scrape config\n",
		},
		{
			description:    "test create scrape config command",
			arguments:      []string{"create", "scrapeConfigs", "-f", "file"},
			err:            false,
			expectedOutput: "create scrape config\n",
		},
		{
			description:    "test update scrape config command",
			arguments:      []string{"update", "scrapeConfigs", "-f", "file"},
			err:            false,
			expectedOutput: "patch scrape config\n",
		},
		{
			description:    "test update scrape config command",
			arguments:      []string{"update", "scrapeConfigs", "-f", "file", "arg"},
			err:            false,
			expectedOutput: "update scrape config\n",
		},
		{
			description:    "test delete scrape config command",
			arguments:      []string{"delete", "scrapeConfig", "config", "-d"},
			err:            false,
			expectedOutput: "delete scrape config command called\n",
		},
		{
			description:    "test delete scrape config command with custom instance id",
			arguments:      []string{"delete", "scrapeConfig", "config", "-i", "id", "-d"},
			err:            false,
			expectedOutput: "delete scrape config command called\n",
		},
		{
			description:    "test delete scrape config command with custom project id",
			arguments:      []string{"delete", "scrapeConfig", "config", "-p", "id", "-d"},
			err:            false,
			expectedOutput: "delete scrape config command called\n",
		},
	}

	// init cmd
	cmdTest := cmd.NewArgusCliCmd()
	b := bytes.NewBufferString("")
	cmdTest.SetOut(b)

	for _, test := range tests {
		old := os.Stdout // keep backup of the real stdout
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal("cannot pipe")
		}
		os.Stdout = w

		cmdTest.SetArgs(test.arguments)

		execErr := cmdTest.Execute()

		out := getStringFromStdout(r, w, old, t)

		if test.err == true {
			assert.Error(t, execErr, test.description)

			continue
		} else {
			assert.NoError(t, err, test.description)
		}

		assert.Contains(t, out, test.expectedOutput, test.description)
	}
}
