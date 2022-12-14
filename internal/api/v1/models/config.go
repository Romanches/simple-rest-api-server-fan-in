package models

import (
	"log"
)

type (
	Config struct {
		// ListenAddr defines HTTP section of the API server configuration
		ListenAddr string      `yaml:"listen_addr"`

	}
)

// Config validator
func (c *Config) Validate() {
	var errs []string

	if c.ListenAddr == "" {
		errs = append(errs, "listen_addr field is required")
	}

	if len(errs) > 0 {
		log.Fatal("Wrong config parameters! ", errs )
	}
}
