package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "oidc",
    Short: "oidc is a simple OIDC server built with Fosite and Cobra.",
    Long:  "oidc is a simple OIDC server built with Fosite and Cobra.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(serveCmd)
}
