package container

import (
	"os"
	"os/exec"
	"syscall"
)

// CreateNamespace sets up the necessary namespaces for the container.
func CreateNamespace() error {
	// Create multiple namespaces in one call
	if err := syscall.Unshare(
		syscall.CLONE_NEWNS | // Mount namespace
			syscall.CLONE_NEWUTS | // UTS namespace (hostname)
			syscall.CLONE_NEWIPC | // IPC namespace
			syscall.CLONE_NEWPID | // PID namespace
			syscall.CLONE_NEWNET | // Network namespace
			syscall.CLONE_NEWUSER); err != nil { // User namespace
		return err
	}

	return nil
}

// SetupMounts sets up the necessary mounts for the container.
func SetupMounts(rootfs string) error {
	// First, ensure root is private to avoid propagating mounts
	if err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
		return err
	}

	// Mount the rootfs to itself so we can pivot_root
	if err := syscall.Mount(rootfs, rootfs, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	// Create necessary mount points
	mountPoints := []struct {
		source string
		target string
		fstype string
		flags  int
		data   string
	}{
		{"/proc", rootfs + "/proc", "proc", 0, ""},
		{"/dev", rootfs + "/dev", "tmpfs", 0, ""},
		{"/sys", rootfs + "/sys", "sysfs", 0, ""},
		{"/tmp", rootfs + "/tmp", "tmpfs", 0, ""},
		{"/etc/resolv.conf", rootfs + "/etc/resolv.conf", "bind", syscall.MS_BIND, ""},
	}

	for _, mount := range mountPoints {
		// Create target directory if it doesn't exist
		if err := os.MkdirAll(mount.target, 0755); err != nil {
			return err
		}

		// Perform the mount
		if err := syscall.Mount(mount.source, mount.target, mount.fstype, uintptr(mount.flags), mount.data); err != nil {
			return err
		}
	}

	return nil
}

// ExecuteCommand runs a command in the container's namespace.
func ExecuteCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)

	// Set up process attributes for the new namespaces
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
	}

	// Connect standard I/O
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
