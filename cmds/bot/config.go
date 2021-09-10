package main

import (
	"errors"
	"flag"
)

type Config struct {
	Server struct {
		Address string
	}
	Bot struct {
		Name string
	}
	Client struct {
		EventBufferSize uint64
		EventRate       uint
		DebugRead       bool
		DebugWrite      bool
	}
}

func (c Config) Validate() error {
	if c.Server.Address == "" || c.Bot.Name == "" {
		return errors.New("Invalid config")
	}
	return nil
}

func ParseConfig() (*Config, error) {
	var c Config

	flag.StringVar(&c.Server.Address, "server.addr", "ws://localhost:8091", "Tokyo server address")
	flag.StringVar(&c.Bot.Name, "bot.name", "foo", "Bot name")
	flag.Uint64Var(&c.Client.EventBufferSize, "client.eventBuf", 10000, "Event buffer size")
	flag.UintVar(&c.Client.EventRate, "client.eventRate", 0, "Receive event limit rate")
	flag.BoolVar(&c.Client.DebugRead, "client.debugRead", false, "Debug client read")
	flag.BoolVar(&c.Client.DebugWrite, "client.debugWrite", false, "Debug client read")
	flag.Parse()

	return &c, c.Validate()
}
