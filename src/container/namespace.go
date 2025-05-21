package container

import (
    "os"
    "os/exec"
    "syscall"
)

// CreateNamespace sets up the necessary namespaces for the container.
func CreateNamespace() error {
    // Create a new mount namespace
    if err := syscall.Unshare(syscall.CLONE_NEWNS); err != nil {
        return err
    }

    // Create a new user namespace
    if err := syscall.Unshare(syscall.CLONE_NEWUSER); err != nil {
        return err
    }

    // Create a new network namespace
    if err := syscall.Unshare(syscall.CLONE_NEWNET); err != nil {
        return err
    }

    return nil
}

// SetupMounts sets up the necessary mounts for the container.
func SetupMounts(rootfs string) error {
    // Mount the root filesystem
    if err := syscall.Mount(rootfs, rootfs, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
        return err
    }

    // Ensure the root filesystem is mounted as read-only
    if err := syscall.Mount("", rootfs, "", syscall.MS_REMOUNT|syscall.MS_RDONLY, ""); err != nil {
        return err
    }

    return nil
}

// ExecuteCommand runs a command in the container's namespace.
func ExecuteCommand(command string, args []string) error {
    cmd := exec.Command(command, args...)
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
    }
    return cmd.Run()
}