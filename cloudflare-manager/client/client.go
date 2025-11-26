package client

import (
	"context"
	"fmt"

	"github.com/cloudflare-manager/config"
	"github.com/cloudflare/cloudflare-go"
)

type Client struct {
	API     *cloudflare.API
	Account *config.Account
	Context context.Context
}

func New(account *config.Account) (*Client, error) {
	api, err := cloudflare.NewWithAPIToken(account.APIToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create Cloudflare API client: %w", err)
	}

	return &Client{
		API:     api,
		Account: account,
		Context: context.Background(),
	}, nil
}

func NewFromConfig() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	account, err := cfg.GetCurrentAccount()
	if err != nil {
		return nil, fmt.Errorf("failed to get current account: %w", err)
	}

	return New(account)
}

func (c *Client) GetAccountID() (string, error) {
	if c.Account.AccountID != "" {
		return c.Account.AccountID, nil
	}

	accounts, _, err := c.API.Accounts(c.Context, cloudflare.AccountsListParams{})
	if err != nil {
		return "", fmt.Errorf("failed to list accounts: %w", err)
	}

	if len(accounts) == 0 {
		return "", fmt.Errorf("no accounts found")
	}

	c.Account.AccountID = accounts[0].ID
	return c.Account.AccountID, nil
}
