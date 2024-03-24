/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/litencatt/rlp"
	"github.com/spf13/cobra"
)

var debugMode bool

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run rogue-like poker",
	RunE: func(cmd *cobra.Command, args []string) error {
		poker := rlp.NewRogurLikePoker()
		if debugMode {
			poker.DebugMode = true
		}

		if err := poker.Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&debugMode, "debug", "d", false, "show detail logs")
}
