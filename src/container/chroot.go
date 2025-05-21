package container

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// ChrootContainer represents a container that uses chroot.
type ChrootContainer struct {
	Rootfs string
}

// NewChrootContainer creates a new ChrootContainer with the specified root filesystem.
func NewChrootContainer(rootfs string) *ChrootContainer {
	return &ChrootContainer{Rootfs: rootfs}
}

// ChangeRoot changes the root filesystem of the current process to the container's root filesystem.
func (c *ChrootContainer) ChangeRoot() error {
	// Ensure the root filesystem exists
	if _, err := os.Stat(c.Rootfs); os.IsNotExist(err) {
		return err
	}

	// Change the root filesystem
	if err := syscall.Chroot(c.Rootfs); err != nil {
		return err
	}

	// Change the current working directory to the new root
	if err := os.Chdir("/"); err != nil {
		return err
	}

	return nil
}

// SetupEnvironment sets up the necessary environment for the container.
func (c *ChrootContainer) SetupEnvironment() error {
	// Example: Create necessary directories in the new root
	dirs := []string{"/dev", "/proc", "/sys"}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(c.Rootfs, dir), 0755); err != nil {
			return err
		}
	}

	// Mount necessary filesystems
	if err := exec.Command("mount", "--bind", "/dev", filepath.Join(c.Rootfs, "dev")).Run(); err != nil {
		return err
	}
	if err := exec.Command("mount", "--bind", "/proc", filepath.Join(c.Rootfs, "proc")).Run(); err != nil {
		return err
	}
	if err := exec.Command("mount", "--bind", "/sys", filepath.Join(c.Rootfs, "sys")).Run(); err != nil {
		return err
	}

	return nil
}
