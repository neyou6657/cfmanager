package commands

import (
	"fmt"

	"github.com/cloudflare-manager/client"
	"github.com/cloudflare-manager/utils"
	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var PagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Manage Cloudflare Pages projects",
	Long:  "Create, list, and manage Cloudflare Pages projects",
}

var pagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Pages projects",
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
		projects, _, err := c.API.ListPagesProjects(c.Context, rc, cloudflare.ListPagesProjectsParams{})
		if err != nil {
			return fmt.Errorf("failed to list pages projects: %w", err)
		}

		if len(projects) == 0 {
			fmt.Println("No Pages projects found.")
			return nil
		}

		headers := []string{"NAME", "SUBDOMAIN", "DOMAINS", "CREATED_ON"}
		var rows [][]string

		for _, project := range projects {
			domainCount := fmt.Sprintf("%d", len(project.Domains))
			rows = append(rows, []string{
				project.Name,
				project.SubDomain,
				domainCount,
				project.CreatedOn.Format("2006-01-02"),
			})
		}

		utils.PrintTable(headers, rows)
		return nil
	},
}

var pagesInfoCmd = &cobra.Command{
	Use:   "info [project-name]",
	Short: "Show Pages project information",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		project, err := c.API.GetPagesProject(c.Context, rc, projectName)
		if err != nil {
			return fmt.Errorf("failed to get pages project: %w", err)
		}

		fmt.Printf("Pages Project Information:\n")
		fmt.Printf("  Name:         %s\n", project.Name)
		fmt.Printf("  Subdomain:    %s.pages.dev\n", project.SubDomain)
		fmt.Printf("  Created On:   %s\n", project.CreatedOn.Format("2006-01-02 15:04:05"))
		fmt.Printf("\nDomains:\n")
		if len(project.Domains) == 0 {
			fmt.Printf("  No custom domains configured\n")
		} else {
			for _, domain := range project.Domains {
				fmt.Printf("  - %s\n", domain)
			}
		}

		return nil
	},
}

var pagesDeleteCmd = &cobra.Command{
	Use:   "delete [project-name]",
	Short: "Delete a Pages project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		if err := c.API.DeletePagesProject(c.Context, rc, projectName); err != nil {
			return fmt.Errorf("failed to delete pages project: %w", err)
		}

		fmt.Printf("✓ Pages project '%s' deleted successfully\n", projectName)
		return nil
	},
}

var pagesDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Manage Pages deployments",
}

var pagesDeploymentListCmd = &cobra.Command{
	Use:   "list [project-name]",
	Short: "List deployments for a Pages project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		deployments, _, err := c.API.ListPagesDeployments(c.Context, rc, cloudflare.ListPagesDeploymentsParams{
			ProjectName: projectName,
		})
		if err != nil {
			return fmt.Errorf("failed to list pages deployments: %w", err)
		}

		if len(deployments) == 0 {
			fmt.Println("No deployments found.")
			return nil
		}

		headers := []string{"ID", "ENVIRONMENT", "STATUS", "CREATED_ON"}
		var rows [][]string

		for _, deployment := range deployments {
			rows = append(rows, []string{
				utils.Truncate(deployment.ID, 12),
				deployment.Environment,
				deployment.LatestStage.Status,
				deployment.CreatedOn.Format("2006-01-02 15:04:05"),
			})
		}

		utils.PrintTable(headers, rows)
		return nil
	},
}

var pagesDeploymentInfoCmd = &cobra.Command{
	Use:   "info [project-name] [deployment-id]",
	Short: "Show deployment information",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		deploymentID := args[1]

		c, err := client.NewFromConfig()
		if err != nil {
			return err
		}

		accountID, err := c.GetAccountID()
		if err != nil {
			return err
		}

		rc := cloudflare.AccountIdentifier(accountID)
		deployment, err := c.API.GetPagesDeploymentInfo(c.Context, rc, projectName, deploymentID)
		if err != nil {
			return fmt.Errorf("failed to get deployment info: %w", err)
		}

		fmt.Printf("Deployment Information:\n")
		fmt.Printf("  ID:          %s\n", deployment.ID)
		fmt.Printf("  Environment: %s\n", deployment.Environment)
		fmt.Printf("  Status:      %s\n", deployment.LatestStage.Status)
		fmt.Printf("  URL:         %s\n", deployment.URL)
		fmt.Printf("  Created On:  %s\n", deployment.CreatedOn.Format("2006-01-02 15:04:05"))

		return nil
	},
}

func init() {
	pagesDeploymentCmd.AddCommand(pagesDeploymentListCmd)
	pagesDeploymentCmd.AddCommand(pagesDeploymentInfoCmd)

	PagesCmd.AddCommand(pagesListCmd)
	PagesCmd.AddCommand(pagesInfoCmd)
	PagesCmd.AddCommand(pagesDeleteCmd)
	PagesCmd.AddCommand(pagesDeploymentCmd)
}
