package commands

import (
    "fmt"

    "github.com/cloudflare-manager/client"
    "github.com/cloudflare-manager/config"
    "github.com/cloudflare-manager/utils"
    "github.com/cloudflare/cloudflare-go"
    "github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
    Use:   "account",
    Short: "Manage Cloudflare accounts",
    Long:  "Add, list, switch, and remove Cloudflare accounts",
}

var accountAddCmd = &cobra.Command{
    Use:   "add [name]",
    Short: "Add a new Cloudflare account",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := args[0]
        apiToken, _ := cmd.Flags().GetString("token")
        email, _ := cmd.Flags().GetString("email")

        if apiToken == "" {
            return fmt.Errorf("API token is required")
        }

        cfg, err := config.Load()
        if err != nil {
            return err
        }

        account := config.Account{
            Name:     name,
            APIToken: apiToken,
            Email:    email,
        }

        c, err := client.New(&account)
        if err != nil {
            return fmt.Errorf("failed to verify API token: %w", err)
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return fmt.Errorf("failed to get account ID: %w", err)
        }
        account.AccountID = accountID

        if err := cfg.AddAccount(account); err != nil {
            return err
        }

        fmt.Printf("✓ Account '%s' added successfully (ID: %s)\n", name, accountID)
        return nil
    },
}

var accountListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all configured accounts",
    RunE: func(cmd *cobra.Command, args []string) error {
        cfg, err := config.Load()
        if err != nil {
            return err
        }

        if len(cfg.Accounts) == 0 {
            fmt.Println("No accounts configured. Use 'account add' to add an account.")
            return nil
        }

        headers := []string{"CURRENT", "NAME", "EMAIL", "ACCOUNT_ID"}
        var rows [][]string

        for _, acc := range cfg.Accounts {
            current := " "
            if acc.Name == cfg.CurrentAccount {
                current = "*"
            }
            rows = append(rows, []string{
                current,
                acc.Name,
                acc.Email,
                acc.AccountID,
            })
        }

        utils.PrintTable(headers, rows)
        return nil
    },
}

var accountSwitchCmd = &cobra.Command{
    Use:   "switch [name]",
    Short: "Switch to a different account",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := args[0]
        cfg, err := config.Load()
        if err != nil {
            return err
        }

        if err := cfg.SetCurrentAccount(name); err != nil {
            return err
        }

        fmt.Printf("✓ Switched to account '%s'\n", name)
        return nil
    },
}

var accountRemoveCmd = &cobra.Command{
    Use:   "remove [name]",
    Short: "Remove an account",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := args[0]
        cfg, err := config.Load()
        if err != nil {
            return err
        }

        if err := cfg.RemoveAccount(name); err != nil {
            return err
        }

        fmt.Printf("✓ Account '%s' removed\n", name)
        return nil
    },
}

var accountInfoCmd = &cobra.Command{
    Use:   "info",
    Short: "Show current account information",
    RunE: func(cmd *cobra.Command, args []string) error {
        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accounts, _, err := c.API.Accounts(c.Context, cloudflare.AccountsListParams{})
        if err != nil {
            return fmt.Errorf("failed to get account info: %w", err)
        }

        if len(accounts) == 0 {
            return fmt.Errorf("no account found")
        }

        acc := accounts[0]
        fmt.Printf("Account Information:\n")
        fmt.Printf("  Name:    %s\n", acc.Name)
        fmt.Printf("  ID:      %s\n", acc.ID)
        fmt.Printf("  Type:    %s\n", acc.Type)
        fmt.Printf("  Status:  %s\n", acc.Settings.EnforceTwoFactor)

        return nil
    },
}

func init() {
    accountAddCmd.Flags().StringP("token", "t", "", "Cloudflare API token (required)")
    accountAddCmd.Flags().StringP("email", "e", "", "Email address (optional)")
    accountAddCmd.MarkFlagRequired("token")

    AccountCmd.AddCommand(accountAddCmd)
    AccountCmd.AddCommand(accountListCmd)
    AccountCmd.AddCommand(accountSwitchCmd)
    AccountCmd.AddCommand(accountRemoveCmd)
    AccountCmd.AddCommand(accountInfoCmd)
}
