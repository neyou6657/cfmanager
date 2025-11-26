package commands

import (
    "fmt"
    "strings"

    "github.com/cloudflare-manager/client"
    "github.com/cloudflare-manager/utils"
    "github.com/cloudflare/cloudflare-go"
    "github.com/spf13/cobra"
)

var DNSCmd = &cobra.Command{
    Use:   "dns",
    Short: "Manage DNS records",
    Long:  "Create, list, update, and delete DNS records",
}

var dnsListCmd = &cobra.Command{
    Use:   "list [zone-id or domain]",
    Short: "List DNS records for a zone",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        recordType, _ := cmd.Flags().GetString("type")

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        params := cloudflare.ListDNSRecordsParams{}
        if recordType != "" {
            params.Type = recordType
        }

        records, _, err := c.API.ListDNSRecords(c.Context, cloudflare.ZoneIdentifier(zoneID), params)
        if err != nil {
            return fmt.Errorf("failed to list DNS records: %w", err)
        }

        if len(records) == 0 {
            fmt.Println("No DNS records found.")
            return nil
        }

        headers := []string{"TYPE", "NAME", "CONTENT", "TTL", "PROXIED", "ID"}
        var rows [][]string

        for _, record := range records {
            proxied := " "
            if record.Proxied != nil && *record.Proxied {
                proxied = "✓"
            }

            rows = append(rows, []string{
                record.Type,
                record.Name,
                utils.Truncate(record.Content, 40),
                fmt.Sprintf("%d", record.TTL),
                proxied,
                utils.Truncate(record.ID, 12),
            })
        }

        utils.PrintTable(headers, rows)
        return nil
    },
}

var dnsCreateCmd = &cobra.Command{
    Use:   "create [zone-id or domain] [type] [name] [content]",
    Short: "Create a DNS record",
    Args:  cobra.ExactArgs(4),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        recordType := strings.ToUpper(args[1])
        name := args[2]
        content := args[3]

        ttl, _ := cmd.Flags().GetInt("ttl")
        proxied, _ := cmd.Flags().GetBool("proxied")
        priority, _ := cmd.Flags().GetInt("priority")

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        params := cloudflare.CreateDNSRecordParams{
            Type:    recordType,
            Name:    name,
            Content: content,
            TTL:     ttl,
            Proxied: &proxied,
        }

        if recordType == "MX" || recordType == "SRV" {
            p := uint16(priority)
            params.Priority = &p
        }

        record, err := c.API.CreateDNSRecord(c.Context, cloudflare.ZoneIdentifier(zoneID), params)
        if err != nil {
            return fmt.Errorf("failed to create DNS record: %w", err)
        }

        fmt.Printf("✓ DNS record created successfully\n")
        fmt.Printf("  ID:      %s\n", record.ID)
        fmt.Printf("  Type:    %s\n", record.Type)
        fmt.Printf("  Name:    %s\n", record.Name)
        fmt.Printf("  Content: %s\n", record.Content)
        return nil
    },
}

var dnsUpdateCmd = &cobra.Command{
    Use:   "update [zone-id or domain] [record-id] [content]",
    Short: "Update a DNS record",
    Args:  cobra.ExactArgs(3),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        recordID := args[1]
        content := args[2]

        ttl, _ := cmd.Flags().GetInt("ttl")
        proxied, _ := cmd.Flags().GetBool("proxied")

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        record, err := c.API.GetDNSRecord(c.Context, cloudflare.ZoneIdentifier(zoneID), recordID)
        if err != nil {
            return fmt.Errorf("failed to get DNS record: %w", err)
        }

        params := cloudflare.UpdateDNSRecordParams{
            ID:      recordID,
            Type:    record.Type,
            Name:    record.Name,
            Content: content,
            TTL:     ttl,
            Proxied: &proxied,
        }

        _, err = c.API.UpdateDNSRecord(c.Context, cloudflare.ZoneIdentifier(zoneID), params)
        if err != nil {
            return fmt.Errorf("failed to update DNS record: %w", err)
        }

        fmt.Printf("✓ DNS record updated successfully\n")
        return nil
    },
}

var dnsDeleteCmd = &cobra.Command{
    Use:   "delete [zone-id or domain] [record-id]",
    Short: "Delete a DNS record",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        zoneIdentifier := args[0]
        recordID := args[1]

        c, err := client.NewFromConfig()
        if err != nil {
            return err
        }

        zoneID, err := getZoneID(c, zoneIdentifier)
        if err != nil {
            return err
        }

        if err := c.API.DeleteDNSRecord(c.Context, cloudflare.ZoneIdentifier(zoneID), recordID); err != nil {
            return fmt.Errorf("failed to delete DNS record: %w", err)
        }

        fmt.Printf("✓ DNS record deleted successfully\n")
        return nil
    },
}

var dnsImportCmd = &cobra.Command{
    Use:   "import [zone-id or domain] [bind-file]",
    Short: "Import DNS records from BIND file",
    Args:  cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        fmt.Printf("⚠  DNS import is not fully implemented in this version\n")
        fmt.Printf("   Please use the Cloudflare Dashboard to import DNS records\n")
        return nil
    },
}

var dnsExportCmd = &cobra.Command{
    Use:   "export [zone-id or domain]",
    Short: "Export DNS records to BIND format",
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

        bind, err := c.API.ExportDNSRecords(c.Context, cloudflare.ZoneIdentifier(zoneID), cloudflare.ExportDNSRecordsParams{})
        if err != nil {
            return fmt.Errorf("failed to export DNS records: %w", err)
        }

        fmt.Println(bind)
        return nil
    },
}

func init() {
    dnsListCmd.Flags().StringP("type", "t", "", "Filter by record type (A, AAAA, CNAME, MX, etc.)")

    dnsCreateCmd.Flags().Int("ttl", 1, "TTL in seconds (1 = automatic)")
    dnsCreateCmd.Flags().Bool("proxied", false, "Enable Cloudflare proxy")
    dnsCreateCmd.Flags().Int("priority", 10, "Priority (for MX/SRV records)")

    dnsUpdateCmd.Flags().Int("ttl", 1, "TTL in seconds")
    dnsUpdateCmd.Flags().Bool("proxied", false, "Enable Cloudflare proxy")

    DNSCmd.AddCommand(dnsListCmd)
    DNSCmd.AddCommand(dnsCreateCmd)
    DNSCmd.AddCommand(dnsUpdateCmd)
    DNSCmd.AddCommand(dnsDeleteCmd)
    DNSCmd.AddCommand(dnsImportCmd)
    DNSCmd.AddCommand(dnsExportCmd)
}
