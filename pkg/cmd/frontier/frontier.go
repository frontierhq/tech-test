package frontier

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string, commit string, date string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "frontier",
		Short: "Frontier is the command line tool for tech testing",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello from NewRootCmd")
			return nil
		},
	}

	return cmd
}
