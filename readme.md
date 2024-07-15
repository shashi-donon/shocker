
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

## Files

- **rootfs**: Alpine Linux root file system with all necessary files.
- **main.go**: Go program to run Alpine OS.
- **go.mod**: Go module file.
- **readme.md**: This README file.