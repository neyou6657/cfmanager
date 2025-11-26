package commands

import (
    "fmt"
    "time"

    "github.com/cloudflare-manager/client"
    "github.com/cloudflare-manager/utils"
    "github.com/cloudflare/cloudflare-go"
    "github.com/spf13/cobra"
)

var ZoneCmd = &cobra.Command{
    Use:   "zone",
    Short: "Manage Cloudflare zones (domains)",
    Long:  "Create, list, and delete zones on Cloudflare",
}

var zoneListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all zones",
    RunE: func(cmd *cobra.Command, args []string) error {
        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zones, err := c.API.ListZones(c.Context)
        if err != nil {
            return fmt.Errorf("failed to list zones: %w", err)
        }

        if len(zones) == 0 {
            fmt.Println("No zones found. Use 'zone create' to add a zone.")
            return nil
        }

        headers := []string{"NAME", "ID", "STATUS", "NAME_SERVERS"}
        var rows [][]string

        for _, zone := range zones {
            nameServers := ""
            if len(zone.NameServers) > 0 {
                nameServers = zone.NameServers[0]
                if len(zone.NameServers) > 1 {
                    nameServers += " ..."
                }
            }

            rows = append(rows, []string{
                zone.Name,
                zone.ID,
                zone.Status,
                nameServers,
            })
        }

        utils.PrintTable(headers, rows)
        return nil
    },
}

var zoneCreateCmd = &cobra.Command{
    Use:   "create [domain]",
    Short: "Create a new zone",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        domain := args[0]
        jumpStart, _ := cmd.Flags().GetBool("jump-start")

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        accountID, err := c.GetAccountID()
        if err != nil {
            return err
        }

        zone, err := c.API.CreateZone(c.Context, domain, jumpStart, cloudflare.Account{ID: accountID}, "full")
        if err != nil {
            return fmt.Errorf("failed to create zone: %w", err)
        }

        fmt.Printf("✓ Zone '%s' created successfully\n", domain)
        fmt.Printf("  Zone ID: %s\n", zone.ID)
        fmt.Printf("  Status:  %s\n", zone.Status)
        fmt.Printf("\nNameservers:\n")
        for _, ns := range zone.NameServers {
            fmt.Printf("  - %s\n", ns)
        }
        fmt.Printf("\nUpdate your domain's nameservers to the ones listed above.\n")

        return nil
    },
}

var zoneDeleteCmd = &cobra.Command{
    Use:   "delete [zone-id or domain]",
    Short: "Delete a zone",
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

        if _, err := c.API.DeleteZone(c.Context, zoneID); err != nil {
            return fmt.Errorf("failed to delete zone: %w", err)
        }

        fmt.Printf("✓ Zone deleted successfully\n")
        return nil
    },
}

var zoneInfoCmd = &cobra.Command{
    Use:   "info [zone-id or domain]",
    Short: "Show zone information",
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

        zone, err := c.API.ZoneDetails(c.Context, zoneID)
        if err != nil {
            return fmt.Errorf("failed to get zone info: %w", err)
        }

        fmt.Printf("Zone Information:\n")
        fmt.Printf("  Name:              %s\n", zone.Name)
        fmt.Printf("  ID:                %s\n", zone.ID)
        fmt.Printf("  Status:            %s\n", zone.Status)
        fmt.Printf("  Plan:              %s\n", zone.Plan.Name)
        fmt.Printf("  Development Mode:  %s\n", utils.BoolToString(zone.DevMode != 0))
        fmt.Printf("  Created On:        %s\n", zone.CreatedOn.Format(time.RFC3339))
        fmt.Printf("  Modified On:       %s\n", zone.ModifiedOn.Format(time.RFC3339))
        fmt.Printf("\nNameservers:\n")
        for _, ns := range zone.NameServers {
            fmt.Printf("  - %s\n", ns)
        }

        return nil
    },
}

var zonePurgeCmd = &cobra.Command{
    Use:   "purge [zone-id or domain]",
    Short: "Purge cache for a zone",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        everything, _ := cmd.Flags().GetBool("everything")
        files, _ := cmd.Flags().GetStringSlice("files")

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        var purgeReq cloudflare.PurgeCacheRequest
        if everything {
            purgeReq.Everything = true
        } else if len(files) > 0 {
            purgeReq.Files = files
        } else {
            purgeReq.Everything = true
        }

        if _, err := c.API.PurgeCache(c.Context, zoneID, purgeReq); err != nil {
            return fmt.Errorf("failed to purge cache: %w", err)
        }

        fmt.Printf("✓ Cache purged successfully\n")
        return nil
    },
}

func getZoneID(c *client.Client, identifier string) (string, error) {
    zones, err := c.API.ListZones(c.Context)
    if err != nil {
        return "", fmt.Errorf("failed to list zones: %w", err)
    }

    for _, zone := range zones {
        if zone.ID == identifier || zone.Name == identifier {
            return zone.ID, nil
        }
    }

    return "", fmt.Errorf("zone not found: %s", identifier)
}

func init() {
    zoneCreateCmd.Flags().BoolP("jump-start", "j", true, "Automatically scan for DNS records")

    zonePurgeCmd.Flags().Bool("everything", false, "Purge everything")
    zonePurgeCmd.Flags().StringSlice("files", []string{}, "Specific files to purge")

    ZoneCmd.AddCommand(zoneListCmd)
    ZoneCmd.AddCommand(zoneCreateCmd)
    ZoneCmd.AddCommand(zoneDeleteCmd)
    ZoneCmd.AddCommand(zoneInfoCmd)
    ZoneCmd.AddCommand(zonePurgeCmd)
}
