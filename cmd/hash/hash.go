package hash

import (
	"buildkit_utils/hash"

	"github.com/spf13/cobra"
)

func HashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash",
		Short: "Get hash by docker image uri",
		RunE:  hash.Hash,
	}
	return cmd
}
