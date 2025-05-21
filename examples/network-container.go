package main

import (
	"fmt"
	"gocontainer/src/container"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func mainn() {
	rootfs := "/rootfs"

	// Ensure rootfs exists
	if _, err := os.Stat(rootfs); os.IsNotExist(err) {
		log.Fatalf("Root filesystem not found at %s. Run debootstrap.sh first.", rootfs)
	}

	// Create a container with custom network setup
	c := container.NewContainer(rootfs)

	// Setup the basic container features
	if err := c.Setup(); err != nil {
		log.Fatalf("Failed to setup container: %v", err)
	}

	// Set up a custom network namespace and bridge
	if err := setupCustomNetwork(); err != nil {
		log.Fatalf("Failed to setup network: %v", err)
	}

	fmt.Println("Container with custom network is set up.")

	// Start the container with networking
	if err := c.Start(); err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	fmt.Println("Network-enabled container is running.")

	// Run commands to check network configuration
	commands := []string{
		"/sbin/ip addr",
		"/sbin/ip route",
		"ping -c 1 127.0.0.1", // Should work
		"ping -c 1 8.8.8.8",   // Will fail due to network namespace isolation
	}

	for _, cmd := range commands {
		fmt.Printf("\nRunning: %s\n", cmd)
		err := runInContainer(cmd)
		if err != nil {
			fmt.Printf("Command failed (expected for external network): %v\n", err)
		}
	}
}

func setupCustomNetwork() error {
	// This creates a new network namespace
	if err := syscall.Unshare(syscall.CLONE_NEWNET); err != nil {
		return fmt.Errorf("failed to unshare network namespace: %w", err)
	}

	// Setup loopback interface
	if err := exec.Command("ip", "link", "set", "lo", "up").Run(); err != nil {
		return fmt.Errorf("failed to bring up loopback interface: %w", err)
	}

	return nil
}

func runInContainer(command string) error {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
