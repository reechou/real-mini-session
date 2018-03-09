package main

import (
	"github.com/reechou/real-mini-session/config"
	"github.com/reechou/real-mini-session/server"
)

func main() {
	server.NewServer(config.NewConfig()).Run()
}
