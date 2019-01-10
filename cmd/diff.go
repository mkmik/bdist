package cmd

import (
	"os"

	"github.com/kr/binarydist"
	"github.com/spf13/cobra"
)

func newDiffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff oldfile newfile patchfile",
		Short: "compute a binary diff",

		Args: cobra.ExactArgs(3),

		RunE: func(cmd *cobra.Command, args []string) error {
			old, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer old.Close()
			new, err := os.Open(args[1])
			if err != nil {
				return err
			}
			defer new.Close()
			patch, err := os.Create(args[2])
			if err != nil {
				return err
			}
			defer patch.Close()
			return binarydist.Diff(old, new, patch)
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(newDiffCmd())
}