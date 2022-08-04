package app

import (
	"control-tower-check/internal"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "control-tower-check ORGANIZATIONAL_UNIT...",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Usage()
		}

		mapper := internal.NewChecker()

		for _, organizationalUnit := range args {
			if err := mapper.Check(organizationalUnit); err != nil {
				return err
			}
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
