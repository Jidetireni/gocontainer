package main

import (
	"fmt"
	"gocontainer/src/container"
	"log"
	"os"
)

func main() {
	// Create a new container with the specified rootfs
	rootfs := "/rootfs"

	// Ensure rootfs exists (this should be created by debootstrap.sh)
	if _, err := os.Stat(rootfs); os.IsNotExist(err) {
		log.Fatalf("Root filesystem not found at %s. Run debootstrap.sh first.", rootfs)
	}

	// Create a new container instance
	c := container.NewContainer(rootfs)

	// Setup the container (namespaces, cgroups)
	if err := c.Setup(); err != nil {
		log.Fatalf("Failed to setup container: %v", err)
	}

	fmt.Println("Container is set up.")

	// Start the container (performs chroot)
	if err := c.Start(); err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	fmt.Println("Container is running.")

	// Run a command in the container
	command := "/bin/bash"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	fmt.Printf("Running %s in the container\n", command)
	if err := c.Exec(command, []string{}...); err != nil {
		log.Fatalf("Failed to execute command in container: %v", err)
	}
}
