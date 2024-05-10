# a CLI and go SDK for the Synology: `synoctl` and `synology-go`

## Why make this?

Synology NAS devices are popular.

In a home setting, a Synology device is particularly handy as a small, always-on Linux device with abundant local storage.

Out of the box, you can interact with a Synology device through a web-based UI. The Web UI is great for trying new features,
but managing containers or network drives from the UI can be slow and error-prone - classic motivations for a scriptable command-line utility.

The `synoctl` command name follows the tradition of `kubectl`, `journalctl`, `sysctl` and other cli utilities for
managing things.
Hopefully it is easy to type and easy to remember!

## Getting Started with `synoctl`

You can download a precompiled `synoctl` binary from the releases tab in this repository.

Run the `synoctl login` and follow the prompts to authenticate with your Synology host

## Commands for the Container Manager App

![animated demo](docs/demo.gif)

`synoctl` supports these commands. You'll need to install the Container Manager app from Package Center first.

| app    | resource   | command | description                                   |
|--------|------------|---------|-----------------------------------------------|
| docker |            | logs    | prints logs from the COntainer Manger app     |
| docker | containers | list    | lists all containers                          |
|        |            | restart | restarts container specified with --name flag |
|        |            | stop    | stops container specified with --name flag    |
|        |            | start   | starts container specified with --name flag   |
|        |            | logs    | prints container logs from target container   |

All commands provide a `--help` command explaining their use.