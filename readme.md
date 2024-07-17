
# Shocker - A Simple Fun Docker

Created By Shashi - "**Sh**ashi's **D**ocker - **Sh**ocker"

## Overview

**Shocker** is a minimalistic Docker runtime that leverages namespaces and `chroot` to run an Alpine Linux container on Ubuntu without using `Docker` or `runc`. 

1. This project is inspired by [Liz Rice's GOTO Amsterdam seminar](https://www.youtube.com/watch?v=8fi7uSYlOdc).
2. This guy from [Earthly](https://www.youtube.com/watch?v=JOsWB50LmwQ) YT channel explains what is chroot, namespaces well
He also has git in which he has create his version of [docker](https://www.youtube.com/watch?v=JOsWB50LmwQ)


## Features

- Runs Alpine Linux container in Ubuntu
- Does not require `Docker` or `runc`

## How to Run

1. Open your terminal.
2. (Optional) If you are on an Ubuntu-like distro, type:
   ```bash
   sudo su
   ```
   Note: There are ongoing efforts to run this without root access.
3. Run the following command to start the Alpine OS:
   ```bash
   go run main.go run /bin/sh
   ```

How containers Work????
## Containers: Chroots on Steroids

Containers have revolutionized software development and deployment, providing lightweight, portable, and consistent environments. To understand their power, it's helpful to compare them to chroot, a much older Unix feature that containers build upon and greatly enhance.

### The Basics of chroot

The chroot command changes the apparent root directory for a process and its children. This effectively isolates the process within a subtree of the filesystem, preventing it from accessing files outside of that subtree. While powerful, chroot has limitations:
- *No process isolation:* It doesn't isolate processes; they can still see and interact with all other processes on the system.
- *No resource control:* It doesn't limit resource usage (CPU, memory, I/O).
- *No network isolation:* It doesn't isolate network settings or interfaces.
- *Security vulnerabilities:* It relies on the filesystem's permissions for isolation, which can be bypassed by privileged users.

### Enter Containers

Containers address these limitations by combining chroot with additional Linux kernel features, such as namespaces and cgroups, to create more robust and flexible isolation mechanisms.

#### 1. *Namespaces*
Namespaces provide isolation at multiple levels:
- *PID Namespace:* Isolates the process ID space, meaning processes inside a container can have the same PID as processes outside.
- *Network Namespace:* Isolates network interfaces, allowing containers to have their own IP addresses, routing tables, and firewall rules.
- *Mount Namespace:* Isolates filesystem mounts, providing each container with its own filesystem view.
- *IPC Namespace:* Isolates inter-process communication, ensuring that IPC resources are contained within the container.
- *UTS Namespace:* Isolates the hostname and domain name, allowing containers to have unique names.
- *User Namespace:* Maps user and group IDs inside the container to different IDs outside, enhancing security.

#### 2. *Cgroups (Control Groups)*
Cgroups manage and limit resource usage:
- *CPU Limiting:* Restricts CPU usage for container processes.
- *Memory Limiting:* Sets memory usage limits, preventing one container from consuming all available memory.
- *Disk I/O Limiting:* Controls disk I/O operations to ensure fair usage among containers.
- *Network Bandwidth Limiting:* Limits network bandwidth to prevent any container from monopolizing network resources.

#### 3. *Layered Filesystems*
Containers use layered filesystems like OverlayFS, which allow for efficient image storage and deployment:
- *Image Layers:* Containers are built from layers, where each layer represents a filesystem change (e.g., adding a file). This allows sharing of common layers across containers, saving space.
- *Copy-on-Write:* Changes are only written to a new layer, keeping the base image unchanged and ensuring quick deployments.

## Files

- **rootfs**: Alpine Linux root file system with all necessary files.
- **main.go**: Go program to run Alpine OS.
- **go.mod**: Go module file.
- **readme.md**: This README file.