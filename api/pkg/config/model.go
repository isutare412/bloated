package config

import (
	"fmt"

	"github.com/isutare412/bloated/api/pkg/log"
)

type Config struct {
	App         string         `mapstructure:"app"`
	Environment string         `mapstructure:"environment"`
	Version     string         `mapstructure:"version"`
	Log         LogConfig      `mapstructure:"log"`
	Postgres    PostgresConfig `mapstructure:"postgres"`
}

func (c *Config) Validate() error {
	if c.App == "" {
		return fmt.Errorf("app should not be empty")
	}
	if c.Environment == "" {
		return fmt.Errorf("environment should not be empty")
	}
	if c.Version == "" {
		return fmt.Errorf("version should not be empty")
	}
	if err := c.Log.Validate(); err != nil {
		return fmt.Errorf("validating log config: %w", err)
	}
	if err := c.Postgres.Validate(); err != nil {
		return fmt.Errorf("validating postgres config: %w", err)
	}
	return nil
}

type LogConfig struct {
	Development bool       `mapstructure:"development"`
	Format      log.Format `mapstructure:"format"`
	Level       log.Level  `mapstructure:"level"`
	StackTrace  bool       `mapstructure:"stackTrace"`
	Caller      bool       `mapstructure:"caller"`
}

func (c *LogConfig) Validate() error {
	if err := c.Format.Validate(); err != nil {
		return err
	}
	if err := c.Level.Validate(); err != nil {
		return err
	}
	return nil
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func (c *PostgresConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host should not be empty")
	}
	if c.Port == 0 {
		return fmt.Errorf("port should not be empty")
	}
	if c.Database == "" {
		return fmt.Errorf("database should not be empty")
	}
	if c.User == "" {
		return fmt.Errorf("user should not be empty")
	}
	if c.Password == "" {
		return fmt.Errorf("password should not be empty")
	}
	return nil
}
