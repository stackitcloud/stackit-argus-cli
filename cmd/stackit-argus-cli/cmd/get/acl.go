package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Get acl for the instance.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get acl called")
	},
}

func init() {

}
