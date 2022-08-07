package main

import (
	"os"
	"fmt"
	"log"

	"github.com/OMO-NOSA/whatsapp-chat-bot/cmd/internal"
)

func main() {
	cfg, err := internal.GetConfig()

	if err != nil {
		log.Fatalf("config error: %s", err)
		os.Exit(1)
	}

	text := fmt.Sprintf("Currently running program on %s mode", cfg.Environment.String())

	fmt.Println(text)
}