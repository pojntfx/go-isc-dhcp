# go-isc-dhcp

Management daemons and CLIs for the ISC DHCP server and client.

[![pipeline status](https://gitlab.com/pojntfx/go-isc-dhcp/badges/master/pipeline.svg)](https://gitlab.com/pojntfx/go-isc-dhcp/commits/master)

## Overview

`go-isc-dhcp` is a collection of management daemons and CLIs for the ISC DHCP server and client. The ISC DHCP server and client are built of two main components:

- `dhcpd`, a DHCP server
- `dhclient`, a DHCP client

In a similar way, `go-isc-dhcp` is built of multiple components. The components are:

- `dhcpdd`, an ISC DHCP server management daemon with a gRPC interface
- `dhclientd`, an ISC DHCP client management daemon with a gRPC interface
- `dhcpdctl`, a CLI for `dhcpdd`
- `dhclientctl`, a CLI for `dhclientd`

`dhcpdd` bundles the `dhcpd` and `dhclient` binaries into it's own binary and extracts them on startup, so there is no need to install the ISC DHCP server and client manually.

## Installation

### Prebuilt Binaries

Prebuilt binaries are available on the [releases page](https://github.com/pojntfx/go-isc-dhcp/releases/latest).

### Go Package

A Go package [is available](https://pkg.go.dev/github.com/pojntfx/go-isc-dhcp).

### Docker Image

#### `dhcpdd`

A Docker image is available on [Docker Hub](https://hub.docker.com/r/pojntfx/dhcpdd).

#### `dhclientd`

A Docker image is available on [Docker Hub](https://hub.docker.com/r/pojntfx/dhclientd).

### Helm Chart

Helm charts for `dhcpdd` and `dhclientd` are available in [@pojntfx's Helm chart repository](https://pojntfx.github.io/charts/).

## Usage

### Daemons

There are two daemons, `dhcpdd` and `dhclientd`; both require root priviledges.

#### `dhcpdd`

You may also set the flags by setting env variables in the format `DHCPDD_[FLAG]` (i.e. `DHCPDD_DHCPDD_CONFIGFILE=examples/dhcpdd.yaml`) or by using a [configuration file](examples/dhcpdd.yaml).

```bash
% dhcpdd --help
dhcpdd is the ISC DHCP server management daemon.

Find more information at:
https://pojntfx.github.io/go-isc-dhcp/

Usage:
  dhcpdd [flags]

Flags:
  -f, --dhcpdd.configFile string       Configuration file to use.
  -l, --dhcpdd.listenHostPort string   TCP listen host:port. (default ":1020")
  -h, --help                           help for dhcpdd
```

#### `dhclientd`

You may also set the flags by setting env variables in the format `DHCLIENTD_[FLAG]` (i.e. `DHCLIENTD_DHCLIENTD_CONFIGFILE=examples/dhclientd.yaml`) or by using a [configuration file](examples/dhclientd.yaml).

```bash
% dhclientd --help
dhclientd is the ISC DHCP client management daemon.

Find more information at:
https://pojntfx.github.io/go-isc-dhcp/

Usage:
  dhclientd [flags]

Flags:
  -f, --dhclientd.configFile string       Configuration file to use.
  -l, --dhclientd.listenHostPort string   TCP listen host:port. (default ":1030")
  -h, --help                              help for dhclientd
```

### Client CLIs

There are two client CLIs, `dhcpdctl` and `dhclientctl`.

#### `dhcpdctl`

You may also set the flags by setting env variables in the format `DHCPD_[FLAG]` (i.e. `DHCPD_DHCPD_CONFIGFILE=examples/dhcpd.yaml`) or by using a [configuration file](examples/dhcpd.yaml). If you want to get started on Kubernetes, see [this configuration file](examples/dhcpd-on-k8s.yaml)

```bash
% dhcpdctl --help
dhcpdctl manages dhcpdd, the ISC DHCP server management daemon.

Find more information at:
https://pojntfx.github.io/go-isc-dhcp/

Usage:
  dhcpdctl [command]

Available Commands:
  apply       Apply a dhcp server
  delete      Delete one or more dhcp server(s)
  get         Get one or all dhcp server(s)
  help        Help about any command

Flags:
  -h, --help   help for dhcpdctl

Use "dhcpdctl [command] --help" for more information about a command.
```

#### `dhclientctl`

You may also set the flags by setting env variables in the format `DHCLIENT_[FLAG]` (i.e. `DHCLIENT_DHCLIENT_CONFIGFILE=examples/dhclient.yaml`) or by using a [configuration file](examples/dhclient.yaml).

```bash
% dhclientctl --help
dhclientctl manages dhclientd, the ISC DHCP client management daemon.

Find more information at:
https://pojntfx.github.io/go-isc-dhcp/

Usage:
  dhclientctl [command]

Available Commands:
  apply       Apply a dhcp client
  delete      Delete one or more dhcp client(s)
  get         Get one or all dhcp client(s)
  help        Help about any command

Flags:
  -h, --help   help for dhclientctl

Use "dhclientctl [command] --help" for more information about a command.
```

## License

go-isc-dhcp (c) 2020 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0
