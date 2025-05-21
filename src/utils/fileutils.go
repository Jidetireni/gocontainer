package utils

import (
    "io/ioutil"
    "os"
    "path/filepath"
)

// CreateDir creates a directory if it does not exist.
func CreateDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, os.ModePerm)
    }
    return nil
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
    input, err := ioutil.ReadFile(src)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(dst, input, os.ModePerm)
}

// CheckFilePermissions checks if the file has the specified permissions.
func CheckFilePermissions(filePath string, mode os.FileMode) (bool, error) {
    info, err := os.Stat(filePath)
    if err != nil {
        return false, err
    }
    return info.Mode().Perm() == mode, nil
}

// GetAbsolutePath returns the absolute path of the given relative path.
func GetAbsolutePath(relativePath string) (string, error) {
    return filepath.Abs(relativePath)
}