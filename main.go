package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Print a banner
	log.Println("---------------------")
	log.Println("---  FactorioBot  ---")
	log.Println("---------------------")
	log.Println("")

	// Initialize our own event handler
	// eventHandler = eventhandler.EventHandler{Name: "rustbot", Listeners: nil}

	// Initialize and open the Discord client
	// discord, discordErr := discord.NewDiscord(&eventHandler, database)
	// if discordErr != nil {
	// 	log.Panic("Failed to initialize Discord:", discordErr)
	// }
	// if discordErr = discord.Open(); discordErr != nil {
	// 	log.Panic("Failed to open Discord:", discordErr)
	// }

	// TODO: Shouldn't we follow the same logic here, so having a separate "Open()" function?
	// Initialize the Webrcon Client (opens the connection automatically)
	webrcon.Initialize(&eventHandler, database)

	// TODO: Implement and setup event handlers for both Discord and Webrcon clients, so they can pass messages between each other

	// TODO: Wait for CTRL-C or something, then call <module>.close() when shutting down
	// Wait here until CTRL-C or other term signal is received.
	log.Println("FactorioBot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Properly dispose of the clients when exiting
	// webrcon.Close()
	// if err := discord.Close(); err != nil {
	// 	log.Panic(err)
	// }
	// if err := database.Close(); err != nil {
	// 	log.Panic(err)
	// }
}
