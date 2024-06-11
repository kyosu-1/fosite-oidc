package main

import (
	"github.com/kyosu-1/fosite-oidc/pkg/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the OIDC server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
