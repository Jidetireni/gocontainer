package main

import (
    "fmt"
    "log"
    "gocontainer/src/container"
)

func main() {
    // Initialize the container
    c, err := container.NewContainer()
    if err != nil {
        log.Fatalf("Failed to create container: %v", err)
    }

    // Set up namespaces, cgroups, and chroot
    if err := c.Setup(); err != nil {
        log.Fatalf("Failed to set up container: %v", err)
    }

    fmt.Println("Container is set up and ready to run.")
    
    // Start the container
    if err := c.Start(); err != nil {
        log.Fatalf("Failed to start container: %v", err)
    }

    fmt.Println("Container is running.")
}