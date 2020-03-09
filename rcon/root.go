package rcon

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	rcon "github.com/gtaylor/factorio-rcon"
)

// Connect to the server
func Connect() {
	fmt.Println("Connecting to", os.Getenv("RCON_HOST"), "on port", os.Getenv("RCON_PORT"))

	inputReader := bufio.NewReader(os.Stdin)

	r, err := rcon.Dial(os.Getenv("RCON_HOST") + ":" + os.Getenv("RCON_PORT"))
	if err != nil {
		panic(err)
	}
	// defer r.Close()

	fmt.Println("Authenticating..")
	err = r.Authenticate(os.Getenv("RCON_PASSWORD"))
	if err != nil {
		r.Close()
		panic(err)
	}

	// fmt.Println("Executing command..")
	// response, err := r.Execute("/players")
	// if err != nil {
	// 	r.Close()
	// 	panic(err)
	// }

	// fmt.Printf("Response: %+v\n", response)
	// fmt.Printf("Response body: %+v\n", response.Body)

	go readPacket(r)
	go handleInput(r, inputReader)
}

// TODO: Use the following mods to enable RCON chat
// https://mods.factorio.com/mod/rconchat
// https://mods.factorio.com/mod/krmqueue
// https://gitlab.com/FishBus/chat-discord

func handleInput(r *rcon.RCON, reader *bufio.Reader) {
	for {
		// Wait for input
		fmt.Println("Waiting for input..")
		input, err := reader.ReadString('\n')
		if err != nil {
			r.Close()
			panic(err)
		}

		// Remove trailing newline from input
		input = strings.TrimSuffix(input, "\n")

		// Send input to server
		fmt.Printf("Sending command '%v'..\n", input)
		cmd := rcon.NewPacket(rcon.ExecCommand, input)
		if err = r.WritePacket(cmd); err != nil {
			r.Close()
			panic(err)
		}
		fmt.Println("Command sent!")
		// response, err := r.Execute(input)
		// if err != nil {
		// 	r.Close()
		// 	panic(err)
		// }
		// fmt.Println("Received a response:", response.Body)
	}
}

func readPacket(r *rcon.RCON) {
	defer r.Close()
	for {
		fmt.Println("Waiting for packets..")
		packet, err := r.ReadPacket()
		if err != nil {
			r.Close()
			panic(err)
		}

		// Remove trailing newlines from packet body
		response := strings.TrimSuffix(packet.Body, "\n")
		response = strings.TrimSuffix(response, "\r")

		fmt.Println("Received packet response:", response)
	}
}
