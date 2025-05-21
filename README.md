# gocontainer

gocontainer is a Go project that implements a lightweight container using Linux features such as namespaces, chroot, and cgroups. This project aims to provide a simple and efficient way to create and manage containers in a Linux environment.

## Features

- **Namespaces**: Isolate the container's processes, file systems, and network interfaces.
- **Chroot**: Change the root filesystem for the container, providing a secure environment.
- **Cgroups**: Manage and limit the resources (CPU, memory, etc.) allocated to the container.

## Project Structure

```
gocontainer
├── src
│   ├── main.go               # Entry point of the application
│   ├── container
│   │   ├── container.go      # Container struct and lifecycle management
│   │   ├── namespace.go      # Namespace creation and management
│   │   ├── cgroups.go        # Cgroups management for resource allocation
│   │   └── chroot.go         # Chroot functionality implementation
│   ├── utils
│   │   └── fileutils.go      # Utility functions for file operations
│   └── config
│       └── config.go         # Configuration structures and parsing
├── scripts
│   ├── debootstrap.sh        # Script to install debootstrap and set up Debian system
│   └── setup.sh              # Initial setup tasks for the container environment
├── examples
│   ├── simple-container.go    # Example of creating and running a simple container
│   └── network-container.go    # Example of creating a container with network capabilities
├── go.mod                     # Go module definition file
├── go.sum                     # Checksums for module dependencies
├── Makefile                   # Build and run commands for the project
└── README.md                  # Project documentation
```

## Installation

To get started with gocontainer, clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd gocontainer
```

Run the setup script to install necessary dependencies and prepare the environment:

```bash
bash scripts/setup.sh
```

## Usage

To run the container, use the following command:

```bash
go run src/main.go
```

You can also explore the examples provided in the `examples` directory to see how to create and manage containers with different configurations.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.