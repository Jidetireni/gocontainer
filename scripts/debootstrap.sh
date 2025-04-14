#!/bin/bash

set -e  

# Check if debootstrap is installed
if ! command -v debootstrap &> /dev/null; then
    echo "debootstrap is not installed. Installing it with: sudo apt install debootstrap"
    echo "Ensure you have root permissions and password-less sudo access"
    sudo apt install -y debootstrap
else
    echo "debootstrap is already installed"
fi

CHROOT_DIR="/rootfs"

# Create a minimal Debian system
sudo debootstrap stable "$CHROOT_DIR" http://deb.debian.org/debian

# Mount necessary filesystems
sudo mount --bind /dev "$CHROOT_DIR/dev"
sudo mount --bind /proc "$CHROOT_DIR/proc"
sudo mount --bind /sys "$CHROOT_DIR/sys"
