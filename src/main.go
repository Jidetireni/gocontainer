package main

import (
	"fmt"
	"gocontainer/src/container"
	"log"
	"os"
)

func main() {
	// Initialize the container with rootfs path
	rootfs := "/rootfs"
	c := container.NewContainer(rootfs)

	// Set up namespaces and cgroups
	if err := c.Setup(); err != nil {
		log.Fatalf("Failed to set up container: %v", err)
	}

	fmt.Println("Container is set up and ready to run.")

	// Start the container (performs chroot)
	if err := c.Start(); err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	fmt.Println("Container is running.")

	// Execute a command in the container
	if len(os.Args) > 1 {
		command := os.Args[1]
		args := []string{}
		if len(os.Args) > 2 {
			args = os.Args[2:]
		}

		fmt.Printf("Executing command in container: %s %v\n", command, args)
		if err := c.Exec(command, args...); err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
	} else {
		fmt.Println("No command specified to run in the container.")
		fmt.Println("Example usage: go run main.go /bin/sh")
	}
}
