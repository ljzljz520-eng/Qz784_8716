package main

import "time"

type Config struct {
	Address  string
	Database string
	Shutdown time.Duration
}

func DefaultConfig() Config {
	return Config{Address: ":8080", Database: "meeting.db", Shutdown: 5 * time.Second}
}
func (c Config) Valid() bool { return c.Address != "" && c.Database != "" && c.Shutdown > 0 }
