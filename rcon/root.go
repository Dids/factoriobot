package rcon

import (
	"fmt"
	"os"

	// rcon "github.com/gtaylor/factorio-rcon"
	"github.com/james4k/rcon"
)

// Connect to the server
func Connect() {
	fmt.Println("Connecting to", os.Getenv("RCON_HOST"), "on port", os.Getenv("RCON_PORT"))

	console, err := rcon.Dial(os.Getenv("RCON_HOST")+":"+os.Getenv("RCON_PORT"), os.Getenv("RCON_PASSWORD"))
	if err != nil {
		panic(err)
	}
	defer console.Close()

	requestID, err := console.Write("status")
	if err != nil {
		panic(err)
	}
	fmt.Println("[WRITE] Request ID:", requestID)

	response, requestID, err := console.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("[READ] Request ID:", requestID)

	fmt.Println("Response:", response)

	fmt.Println("DONE")

	// r, err := rcon.Dial(os.Getenv("RCON_HOST") + ":" + os.Getenv("RCON_PORT"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer r.Close()

	// fmt.Println("Authenticating..")
	// err = r.Authenticate(os.Getenv("RCON_PASSWORD"))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Executing the 'status' command..")
	// response, err := r.Execute("status")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Response: %+v\n", response)
	// fmt.Printf("Response body: %+v\n", response.Body)
}
