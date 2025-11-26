package main

import (
    "fmt"
    "os"

    "github.com/cloudflare-manager/commands"
    "github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
    Use:   "cfm",
    Short: "Cloudflare Multi-Account Manager",
    Long: `Cloudflare Multi-Account Manager - A powerful CLI tool to manage multiple 
Cloudflare accounts, zones, DNS records, Workers, and Pages projects.

Features:
  • Multi-account management with easy switching
  • Zone/domain management
  • DNS record operations (create, list, update, delete, import, export)
  • Worker deployment and routing
  • Pages project management
  • Cache purging
  • And more...`,
    Version: version,
}

func main() {
    rootCmd.AddCommand(commands.AccountCmd)
    rootCmd.AddCommand(commands.ZoneCmd)
    rootCmd.AddCommand(commands.DNSCmd)
    rootCmd.AddCommand(commands.WorkerCmd)
    rootCmd.AddCommand(commands.PagesCmd)
    rootCmd.AddCommand(commands.KVCmd)
    rootCmd.AddCommand(commands.R2Cmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
