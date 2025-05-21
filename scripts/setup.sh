#!/bin/bash

set -e

echo "Setting up environment for Go container implementation..."

# Install necessary dependencies
apt-get update
apt-get install -y \
    build-essential \
    iproute2 \
    debootstrap \
    iptables \
    cgroup-tools

# Create necessary directories for the container environment
echo "Creating rootfs directory structure..."
mkdir -p /rootfs/{dev,proc,sys,tmp,etc}

# Check if debootstrap has already been run
if [ -f /rootfs/etc/os-release ]; then
    echo "Root filesystem already exists. Skipping debootstrap."
else
    echo "Running debootstrap to create minimal Debian system..."
    # Run debootstrap (or use the existing script)
    ./debootstrap.sh
fi

# Check if cgroup v2 is available
if [ -d /sys/fs/cgroup/unified ]; then
    echo "Setting up cgroup v2..."
    mkdir -p /sys/fs/cgroup/unified/gocontainer
    echo "+cpu +memory +pids" > /sys/fs/cgroup/unified/gocontainer/cgroup.subtree_control
else
    # Legacy cgroups
    echo "Setting up legacy cgroups..."
    mkdir -p /sys/fs/cgroup/memory/gocontainer
    mkdir -p /sys/fs/cgroup/cpu/gocontainer
fi

# Enable IP forwarding for container networking
echo "Enabling IP forwarding..."
echo 1 > /proc/sys/net/ipv4/ip_forward

echo "Container environment setup complete."
echo "You can now build and run the Go container implementation."