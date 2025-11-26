package commands

import (
    "fmt"
    "os"

    "github.com/cloudflare-manager/client"
    "github.com/cloudflare-manager/utils"
    "github.com/cloudflare/cloudflare-go"
    "github.com/spf13/cobra"
)

var WorkerCmd = &cobra.Command{
    Use:   "worker",
    Short: "Manage Cloudflare Workers",
    Long:  "Deploy, list, and manage Cloudflare Workers",
}

var workerDeployCmd = &cobra.Command{
    Use:   "deploy [name] [script-file]",
    Short: "Deploy a Worker script",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := args[0]
        scriptFile := args[1]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        scriptContent, err := os.ReadFile(scriptFile)
        if err != nil {
            return fmt.Errorf("failed to read script file: %w", err)
        }

        params := cloudflare.CreateWorkerParams{
            ScriptName: name,
            Script:     string(scriptContent),
        }

        rc := cloudflare.AccountIdentifier(accountID)
        _, err = c.API.UploadWorker(c.Context, rc, params)
        if err != nil {
            return fmt.Errorf("failed to deploy worker: %w", err)
        }

        fmt.Printf("✓ Worker '%s' deployed successfully\n", name)
        return nil
    },
}

var workerDeleteCmd = &cobra.Command{
    Use:   "delete [name]",
    Short: "Delete a Worker",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        name := args[0]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        rc := cloudflare.AccountIdentifier(accountID)
        err = c.API.DeleteWorker(c.Context, rc, cloudflare.DeleteWorkerParams{ScriptName: name})
        if err != nil {
            return fmt.Errorf("failed to delete worker: %w", err)
        }

        fmt.Printf("✓ Worker '%s' deleted successfully\n", name)
        return nil
    },
}

var workerRouteCmd = &cobra.Command{
    Use:   "route",
    Short: "Manage Worker routes",
}

var workerRouteListCmd = &cobra.Command{
    Use:   "list [zone-id or domain]",
    Short: "List Worker routes for a zone",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        rc := cloudflare.ZoneIdentifier(zoneID)
        routes, err := c.API.ListWorkerRoutes(c.Context, rc, cloudflare.ListWorkerRoutesParams{})
        if err != nil {
            return fmt.Errorf("failed to list worker routes: %w", err)
        }

        if len(routes.Routes) == 0 {
            fmt.Println("No worker routes found.")
            return nil
        }

        headers := []string{"PATTERN", "WORKER", "ID"}
        var rows [][]string

        for _, route := range routes.Routes {
            worker := "-"
            rows = append(rows, []string{
                route.Pattern,
                worker,
                route.ID,
            })
        }

        utils.PrintTable(headers, rows)
        return nil
    },
}

var workerRouteCreateCmd = &cobra.Command{
    Use:   "create [zone-id or domain] [pattern] [worker-name]",
    Short: "Create a Worker route",
    Long:  "Route a URL pattern to a Worker. Example pattern: example.com/*",
    Args:  cobra.ExactArgs(3),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        pattern := args[1]
        workerName := args[2]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        rc := cloudflare.ZoneIdentifier(zoneID)
        params := cloudflare.CreateWorkerRouteParams{
            Pattern: pattern,
            Script:  workerName,
        }

        route, err := c.API.CreateWorkerRoute(c.Context, rc, params)
        if err != nil {
            return fmt.Errorf("failed to create worker route: %w", err)
        }

        fmt.Printf("✓ Worker route created successfully\n")
        fmt.Printf("  Route ID: %s\n", route.ID)
        fmt.Printf("  Pattern:  %s\n", pattern)
        fmt.Printf("  Worker:   %s\n", workerName)
        return nil
    },
}

var workerRouteDeleteCmd = &cobra.Command{
    Use:   "delete [zone-id or domain] [route-id]",
    Short: "Delete a Worker route",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        routeID := args[1]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        rc := cloudflare.ZoneIdentifier(zoneID)
        _, err = c.API.DeleteWorkerRoute(c.Context, rc, routeID)
        if err != nil {
            return fmt.Errorf("failed to delete worker route: %w", err)
        }

        fmt.Printf("✓ Worker route deleted successfully\n")
        return nil
    },
}

var workerSubdomainCmd = &cobra.Command{
    Use:   "subdomain",
    Short: "Manage Workers subdomain (workers.dev)",
}

var workerSubdomainGetCmd = &cobra.Command{
    Use:   "get",
    Short: "Get Workers subdomain",
    RunE: func(cmd *cobra.Command, args []string) error {
        fmt.Printf("⚠  Workers subdomain management is not fully implemented in this version\n")
        fmt.Printf("   Please use the Cloudflare Dashboard\n")
        return nil
    },
}

var workerSubdomainSetCmd = &cobra.Command{
    Use:   "set [subdomain]",
    Short: "Set Workers subdomain",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        fmt.Printf("⚠  Workers subdomain management is not fully implemented in this version\n")
        fmt.Printf("   Please use the Cloudflare Dashboard\n")
        return nil
    },
}

func init() {
    workerRouteCmd.AddCommand(workerRouteListCmd)
    workerRouteCmd.AddCommand(workerRouteCreateCmd)
    workerRouteCmd.AddCommand(workerRouteDeleteCmd)

    workerSubdomainCmd.AddCommand(workerSubdomainGetCmd)
    workerSubdomainCmd.AddCommand(workerSubdomainSetCmd)

    WorkerCmd.AddCommand(workerDeployCmd)
    WorkerCmd.AddCommand(workerDeleteCmd)
    WorkerCmd.AddCommand(workerRouteCmd)
    WorkerCmd.AddCommand(workerSubdomainCmd)
}
