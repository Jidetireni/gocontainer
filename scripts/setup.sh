#!/bin/bash

set -e

# Create necessary directories for the container environment
mkdir -p /rootfs/{dev,proc,sys}

# Mount necessary filesystems
mount --bind /dev /rootfs/dev
mount --bind /proc /rootfs/proc
mount --bind /sys /rootfs/sys

# Set up a minimal environment for the container
echo "Container environment setup complete."