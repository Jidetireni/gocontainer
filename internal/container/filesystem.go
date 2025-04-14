package container

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/joho/godotenv"
)

var (
	rootFS     string
	scriptPath string
)

// SetUpRootFS makes the bootstrap script executable and runs it to prepare the chroot environment.
func SetUpRootFS() error {

	if err := godotenv.Load(".env"); err != nil {
		return errors.New("no env file found")
	}

	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		return errors.New("PROJECT_ROOT not set in .env")
	}
	scriptPath = filepath.Join(projectRoot, "scripts", "debootstrap.sh")

	// Make the script executable
	if err := exec.Command("chmod", "+x", scriptPath).Run(); err != nil {
		return fmt.Errorf("failed to make script executable: %v", err)
	}

	// Execute the bootstrap script
	cmd := exec.Command(scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("script execution failed: %v", err)
	}

	return nil
}

// ChrootIntoRootFS change chroot and drops into a new shell inside the new root filesystem.
func ChrootIntoRootFS() error {
	if _, err := os.Stat(rootFS); os.IsNotExist(err) {
		return fmt.Errorf("root path does not exist: %s", rootFS)
	}

	// change chroot
	if err := syscall.Chroot(rootFS); err != nil {
		return fmt.Errorf("chroot failed: %v", err)
	}

	// Change working directory to "/"
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir failed: %v", err)
	}

	// Replace current process with bash inside chroot
	if err := syscall.Exec("/bin/bash", []string{"bash"}, os.Environ()); err != nil {
		return fmt.Errorf("failed to exec bash: %v", err)
	}

	return nil
}
