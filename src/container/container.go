package container

import (
	"fmt"
	"os"
	"syscall"
)

type Container struct {
	rootfs string
}

func NewContainer(rootfs string) *Container {
	return &Container{rootfs: rootfs}
}

func (c *Container) Setup() error {
	// Setup prepares the container environment but doesn't chroot yet
	if err := c.setupNamespaces(); err != nil {
		return err
	}
	if err := c.setupCgroups(); err != nil {
		return err
	}
	return nil
}

func (c *Container) Start() error {
	// Setup proper mounts before chroot
	if err := SetupMounts(c.rootfs); err != nil {
		return fmt.Errorf("failed to setup mounts: %w", err)
	}

	// Fully start the container including chroot
	if err := c.setupChroot(); err != nil {
		return err
	}
	return nil
}

func (c *Container) Stop() error {
	// Logic to stop the container
	return nil
}

func (c *Container) setupChroot() error {
	if err := syscall.Chroot(c.rootfs); err != nil {
		return fmt.Errorf("failed to change root: %w", err)
	}
	if err := os.Chdir("/"); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}
	return nil
}

func (c *Container) setupNamespaces() error {
	// Create namespaces using the imported namespace functionality
	return CreateNamespace()
}

func (c *Container) setupCgroups() error {
	// Set up cgroups using cgroup manager
	cgroupPath := "/sys/fs/cgroup/unified/gocontainer"
	cgroupManager := NewCgroupManager(cgroupPath)

	// Create the cgroup
	if err := cgroupManager.Create(); err != nil {
		return fmt.Errorf("failed to create cgroup: %w", err)
	}

	// Add the current process to the cgroup
	if err := cgroupManager.AddProcess(os.Getpid()); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %w", err)
	}

	// Set resource limits (optional)
	if err := cgroupManager.SetMemoryLimit("512M"); err != nil {
		return fmt.Errorf("failed to set memory limit: %w", err)
	}

	return nil
}

func (c *Container) Exec(command string, args ...string) error {
	// Use the ExecuteCommand function from namespace.go to run the command
	// in the proper namespaces
	return ExecuteCommand(command, args)
}
