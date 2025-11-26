package commands

import (
    "fmt"

    "github.com/cloudflare-manager/client"
    "github.com/cloudflare-manager/utils"
    "github.com/cloudflare/cloudflare-go"
    "github.com/spf13/cobra"
)

var KVCmd = &cobra.Command{
    Use:   "kv",
    Short: "Manage Workers KV Namespaces",
    Long:  "Create, list, and manage Workers KV namespaces and key-value pairs",
}

var kvNamespaceCmd = &cobra.Command{
    Use:   "namespace",
    Short: "Manage KV namespaces",
}

var kvNamespaceListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all KV namespaces",
    RunE: func(cmd *cobra.Command, args []string) error {
        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        namespaces, _, err := c.API.ListWorkersKVNamespaces(c.Context, rc, cloudflare.ListWorkersKVNamespacesParams{})
        if err != nil {
            return fmt.Errorf("failed to list KV namespaces: %w", err)
        }

        if len(namespaces) == 0 {
            fmt.Println("No KV namespaces found. Use 'kv namespace create' to create one.")
            return nil
        }

        headers := []string{"ID", "TITLE"}
        var rows [][]string

        for _, ns := range namespaces {
            rows = append(rows, []string{
                ns.ID,
                ns.Title,
            })
        }

        utils.PrintTable(headers, rows)
        return nil
    },
}

var kvNamespaceCreateCmd = &cobra.Command{
    Use:   "create [title]",
    Short: "Create a new KV namespace",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        title := args[0]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        ns, err := c.API.CreateWorkersKVNamespace(c.Context, rc, cloudflare.CreateWorkersKVNamespaceParams{
            Title: title,
        })
        if err != nil {
            return fmt.Errorf("failed to create KV namespace: %w", err)
        }

        fmt.Printf("✓ KV namespace created successfully\n")
        fmt.Printf("  ID:    %s\n", ns.Result.ID)
        fmt.Printf("  Title: %s\n", ns.Result.Title)
        return nil
    },
}

var kvNamespaceDeleteCmd = &cobra.Command{
    Use:   "delete [namespace-id]",
    Short: "Delete a KV namespace",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        namespaceID := args[0]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        _, err = c.API.DeleteWorkersKVNamespace(c.Context, rc, namespaceID)
        if err != nil {
            return fmt.Errorf("failed to delete KV namespace: %w", err)
        }

        fmt.Printf("✓ KV namespace deleted successfully\n")
        return nil
    },
}

var kvNamespaceRenameCmd = &cobra.Command{
    Use:   "rename [namespace-id] [new-title]",
    Short: "Rename a KV namespace",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        namespaceID := args[0]
        newTitle := args[1]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        _, err = c.API.UpdateWorkersKVNamespace(c.Context, rc, cloudflare.UpdateWorkersKVNamespaceParams{
            NamespaceID: namespaceID,
            Title:       newTitle,
        })
        if err != nil {
            return fmt.Errorf("failed to rename KV namespace: %w", err)
        }

        fmt.Printf("✓ KV namespace renamed to '%s'\n", newTitle)
        return nil
    },
}

var kvKeyCmd = &cobra.Command{
    Use:   "key",
    Short: "Manage KV keys",
}

var kvKeyListCmd = &cobra.Command{
    Use:   "list [namespace-id]",
    Short: "List keys in a KV namespace",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        namespaceID := args[0]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        response, err := c.API.ListWorkersKVKeys(c.Context, rc, namespaceID)
        if err != nil {
            return fmt.Errorf("failed to list KV keys: %w", err)
        }

        if len(response.Result) == 0 {
            fmt.Println("No keys found.")
            return nil
        }

        headers := []string{"NAME"}
        var rows [][]string

        for _, key := range response.Result {
            rows = append(rows, []string{
                key.Name,
            })
        }

        utils.PrintTable(headers, rows)
        fmt.Printf("\nTotal: %d keys\n", len(response.Result))
        return nil
    },
}

var kvKeyGetCmd = &cobra.Command{
    Use:   "get [namespace-id] [key]",
    Short: "Get a value from KV",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        namespaceID := args[0]
        key := args[1]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        value, err := c.API.GetWorkersKV(c.Context, rc, cloudflare.GetWorkersKVParams{
            NamespaceID: namespaceID,
            Key:         key,
        })
        if err != nil {
            return fmt.Errorf("failed to get KV value: %w", err)
        }

        fmt.Println(string(value))
        return nil
    },
}

var kvKeyPutCmd = &cobra.Command{
    Use:   "put [namespace-id] [key] [value]",
    Short: "Put a value into KV",
    Args:  cobra.ExactArgs(3),
    RunE: func(cmd *cobra.Command, args []string) error {
        namespaceID := args[0]
        key := args[1]
        value := args[2]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        _, err = c.API.WriteWorkersKVEntry(c.Context, rc, cloudflare.WriteWorkersKVEntryParams{
            NamespaceID: namespaceID,
            Key:         key,
            Value:       []byte(value),
        })
        if err != nil {
            return fmt.Errorf("failed to put KV value: %w", err)
        }

        fmt.Printf("✓ Value stored successfully\n")
        fmt.Printf("  Namespace: %s\n", namespaceID)
        fmt.Printf("  Key:       %s\n", key)
        return nil
    },
}

var kvKeyDeleteCmd = &cobra.Command{
    Use:   "delete [namespace-id] [key]",
    Short: "Delete a key from KV",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        namespaceID := args[0]
        key := args[1]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        _, err = c.API.DeleteWorkersKVEntry(c.Context, rc, cloudflare.DeleteWorkersKVEntryParams{
            NamespaceID: namespaceID,
            Key:         key,
        })
        if err != nil {
            return fmt.Errorf("failed to delete KV key: %w", err)
        }

        fmt.Printf("✓ Key deleted successfully\n")
        return nil
    },
}

func init() {
    kvNamespaceCmd.AddCommand(kvNamespaceListCmd)
    kvNamespaceCmd.AddCommand(kvNamespaceCreateCmd)
    kvNamespaceCmd.AddCommand(kvNamespaceDeleteCmd)
    kvNamespaceCmd.AddCommand(kvNamespaceRenameCmd)

    kvKeyCmd.AddCommand(kvKeyListCmd)
    kvKeyCmd.AddCommand(kvKeyGetCmd)
    kvKeyCmd.AddCommand(kvKeyPutCmd)
    kvKeyCmd.AddCommand(kvKeyDeleteCmd)

    KVCmd.AddCommand(kvNamespaceCmd)
    KVCmd.AddCommand(kvKeyCmd)
}
