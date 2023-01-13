package main

import (
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-argus-cli/cmd/stackit-argus-cli/cmd"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
