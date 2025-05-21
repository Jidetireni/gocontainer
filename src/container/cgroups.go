package container

import (
	"fmt"
	"os"
	"path/filepath"
)

// CgroupManager manages control groups for resource allocation and limiting.
type CgroupManager struct {
	path string
}

// NewCgroupManager creates a new CgroupManager for the specified path.
func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{path: path}
}

// Create creates a new cgroup at the specified path.
func (c *CgroupManager) Create() error {
	return os.MkdirAll(c.path, 0755)
}

// AddProcess adds a process to the cgroup.
func (c *CgroupManager) AddProcess(pid int) error {
	cgroupProcsPath := filepath.Join(c.path, "cgroup.procs")
	return writeToFile(cgroupProcsPath, fmt.Sprintf("%d", pid))
}

// SetMemoryLimit sets the memory limit for the cgroup.
func (c *CgroupManager) SetMemoryLimit(limit string) error {
	return writeToFile(filepath.Join(c.path, "memory.limit_in_bytes"), limit)
}

// SetCPULimit sets the CPU limit for the cgroup.
func (c *CgroupManager) SetCPULimit(limit string) error {
	return writeToFile(filepath.Join(c.path, "cpu.cfs_quota_us"), limit)
}

// writeToFile is a helper function to write a string to a file.
func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// Cleanup removes the cgroup.
func (c *CgroupManager) Cleanup() error {
	return os.RemoveAll(c.path)
}
