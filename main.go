package main

import (
	"workspace-organizer/internal/app/cmd"
	"workspace-organizer/internal/app/config"
)

func main() {
	config.InitConfig()
	cmd.Run()
}
