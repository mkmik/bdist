package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kr/binarydist"
	"github.com/spf13/cobra"
)

func newDiffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff oldfile newfile",
		Short: "compute a binary diff",

		Args: cobra.ExactArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			old, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer old.Close()
			oldHash := sha256.New()

			new, err := os.Open(args[1])
			if err != nil {
				return err
			}
			defer new.Close()
			newHash := sha256.New()

			patch, err := ioutil.TempFile(".", ".patch")
			if err != nil {
				return err
			}
			defer patch.Close()
			defer os.Remove(patch.Name())

			if err := binarydist.Diff(io.TeeReader(old, oldHash), io.TeeReader(new, newHash), patch); err != nil {
				return err
			}

			patchName := fmt.Sprintf("%s-to-%s.bpatch",
				hex.EncodeToString(oldHash.Sum(nil)),
				hex.EncodeToString(newHash.Sum(nil)),
			)

			return os.Rename(patch.Name(), patchName)
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(newDiffCmd())
}