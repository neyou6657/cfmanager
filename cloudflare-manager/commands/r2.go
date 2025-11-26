package commands

import (
	"fmt"

	"github.com/cloudflare-manager/client"
	"github.com/cloudflare-manager/utils"
	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var R2Cmd = &cobra.Command{
	Use:   "r2",
	Short: "Manage R2 Storage Buckets",
	Long:  "Create, list, and manage R2 object storage buckets",
}

var r2ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all R2 buckets",
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
		buckets, err := c.API.ListR2Buckets(c.Context, rc, cloudflare.ListR2BucketsParams{})
		if err != nil {
			return fmt.Errorf("failed to list R2 buckets: %w", err)
		}

		if len(buckets) == 0 {
			fmt.Println("No R2 buckets found. Use 'r2 create' to create one.")
			return nil
		}

		headers := []string{"NAME", "LOCATION", "CREATED_ON"}
		var rows [][]string

		for _, bucket := range buckets {
			rows = append(rows, []string{
				bucket.Name,
				bucket.Location,
				bucket.CreationDate.Format("2006-01-02 15:04:05"),
			})
		}

		utils.PrintTable(headers, rows)
		return nil
	},
}

var r2CreateCmd = &cobra.Command{
	Use:   "create [bucket-name]",
	Short: "Create a new R2 bucket",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName := args[0]
		location, _ := cmd.Flags().GetString("location")

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		bucket, err := c.API.CreateR2Bucket(c.Context, rc, cloudflare.CreateR2BucketParameters{
			Name:             bucketName,
			LocationHint:     location,
		})
		if err != nil {
			return fmt.Errorf("failed to create R2 bucket: %w", err)
		}

		fmt.Printf("✓ R2 bucket created successfully\n")
		fmt.Printf("  Name:     %s\n", bucket.Name)
		fmt.Printf("  Location: %s\n", bucket.Location)
		return nil
	},
}

var r2DeleteCmd = &cobra.Command{
	Use:   "delete [bucket-name]",
	Short: "Delete an R2 bucket",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName := args[0]

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		err = c.API.DeleteR2Bucket(c.Context, rc, bucketName)
		if err != nil {
			return fmt.Errorf("failed to delete R2 bucket: %w", err)
		}

		fmt.Printf("✓ R2 bucket '%s' deleted successfully\n", bucketName)
		return nil
	},
}

var r2InfoCmd = &cobra.Command{
	Use:   "info [bucket-name]",
	Short: "Get R2 bucket information",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketName := args[0]

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		bucket, err := c.API.GetR2Bucket(c.Context, rc, bucketName)
		if err != nil {
			return fmt.Errorf("failed to get R2 bucket info: %w", err)
		}

		fmt.Printf("R2 Bucket Information:\n")
		fmt.Printf("  Name:         %s\n", bucket.Name)
		fmt.Printf("  Location:     %s\n", bucket.Location)
		fmt.Printf("  Created On:   %s\n", bucket.CreationDate.Format("2006-01-02 15:04:05"))
		return nil
	},
}

func init() {
	r2CreateCmd.Flags().String("location", "auto", "Location hint (auto, wnam, enam, weur, eeur, apac)")

	R2Cmd.AddCommand(r2ListCmd)
	R2Cmd.AddCommand(r2CreateCmd)
	R2Cmd.AddCommand(r2DeleteCmd)
	R2Cmd.AddCommand(r2InfoCmd)
}
