# go-portscanner

A lightweight and efficient port scanner written in Go. This tool allows users to scan specific ports on a target host for multiple protocols, such as TCP and UDP, to determine if they are open or closed.

## Features
- Scan ports for multiple protocols (e.g., TCP, UDP) on users' provided domains.
- Future features
    - Service detection (after identifying open ports, attempt to detect the service running on the port (e.g., HTTP, SSH, FTP).)
    - OS fingerprinting

## Configuration
- Format the codebase
    ```
    go fmt ./...
    ```
- Build the project
    ```
    go build -o port-scanner
    ```
- Run the project
    ```
    ./port-scanner
    ```