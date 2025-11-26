package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Account struct {
	Name      string `yaml:"name"`
	APIToken  string `yaml:"api_token"`
	AccountID string `yaml:"account_id,omitempty"`
	Email     string `yaml:"email,omitempty"`
}

type Config struct {
	CurrentAccount string    `yaml:"current_account"`
	Accounts       []Account `yaml:"accounts"`
}

var configPath string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath = filepath.Join(home, ".cloudflare-manager.yaml")
}

func GetConfigPath() string {
	return configPath
}

func Load() (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Accounts: []Account{}}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Save() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0600)
}

func (c *Config) AddAccount(account Account) error {
	for i, acc := range c.Accounts {
		if acc.Name == account.Name {
			c.Accounts[i] = account
			return c.Save()
		}
	}
	c.Accounts = append(c.Accounts, account)
	if c.CurrentAccount == "" {
		c.CurrentAccount = account.Name
	}
	return c.Save()
}

func (c *Config) RemoveAccount(name string) error {
	for i, acc := range c.Accounts {
		if acc.Name == name {
			c.Accounts = append(c.Accounts[:i], c.Accounts[i+1:]...)
			if c.CurrentAccount == name {
				if len(c.Accounts) > 0 {
					c.CurrentAccount = c.Accounts[0].Name
				} else {
					c.CurrentAccount = ""
				}
			}
			return c.Save()
		}
	}
	return fmt.Errorf("account %s not found", name)
}

func (c *Config) GetCurrentAccount() (*Account, error) {
	if c.CurrentAccount == "" {
		return nil, fmt.Errorf("no current account set")
	}
	for _, acc := range c.Accounts {
		if acc.Name == c.CurrentAccount {
			return &acc, nil
		}
	}
	return nil, fmt.Errorf("current account %s not found", c.CurrentAccount)
}

func (c *Config) GetAccount(name string) (*Account, error) {
	for _, acc := range c.Accounts {
		if acc.Name == name {
			return &acc, nil
		}
	}
	return nil, fmt.Errorf("account %s not found", name)
}

func (c *Config) SetCurrentAccount(name string) error {
	for _, acc := range c.Accounts {
		if acc.Name == name {
			c.CurrentAccount = name
			return c.Save()
		}
	}
	return fmt.Errorf("account %s not found", name)
}
