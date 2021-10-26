# go-isc-dhcp

Management daemons and CLIs for the ISC DHCP server and client.

[![hydrun CI](https://github.com/pojntfx/go-isc-dhcp/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/go-isc-dhcp/actions/workflows/hydrun.yaml)
[![Matrix](https://img.shields.io/matrix/go-isc-dhcp:matrix.org)](https://matrix.to/#/#go-isc-dhcp:matrix.org?via=matrix.org)
[![Binary Downloads](https://img.shields.io/github/downloads/pojntfx/go-isc-dhcp/total?label=binary%20downloads)](https://github.com/pojntfx/go-isc-dhcp/releases)

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

Static binaries are available on [GitHub releases](https://github.com/pojntfx/go-isc-dhcp/releases).

On Linux, you can install them like so:

```shell
$ curl -L -o /tmp/dhcpdd "https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhcpdd.linux-$(uname -m)"
$ curl -L -o /tmp/dhclientd "https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhclientd.linux-$(uname -m)"
$ curl -L -o /tmp/dhcpdctl "https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhcpdctl.linux-$(uname -m)"
$ curl -L -o /tmp/dhclientctl "https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhclientctl.linux-$(uname -m)"
$ sudo install /tmp/{dhcpdd,dhclientd,dhcpdctl,dhclientctl} /usr/local/bin
```

On macOS, you can use the following (we can't cross-compile for macOS, so only the client CLIs work - if you want to build the daemons locally, see [contributing](#contributing)):

```shell
$ curl -L -o /tmp/dhcpdctl "https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhcpdctl.linux-$(uname -m)"
$ curl -L -o /tmp/dhclientctl "https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhclientctl.linux-$(uname -m)"
$ sudo install /tmp/{dhcpdctl,dhclientctl} /usr/local/bin
```

On Windows, the following should work (we can't cross-compile for Windows, so only the client CLIs work - if you want to build the daemons locally, see [contributing](#contributing)):

```shell
PS> Invoke-WebRequest https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhcpdctl.windows-x86_64.exe -OutFile \Windows\System32\dhcpdctl.exe
PS> Invoke-WebRequest https://github.com/pojntfx/go-isc-dhcp/releases/latest/download/dhclientctl.windows-x86_64.exe -OutFile \Windows\System32\dhclientctl.exe
```

You can find binaries for more operating systems and architectures on [GitHub releases](https://github.com/pojntfx/go-isc-dhcp/releases).

## Usage

### 1. Setting up `dhcpd`

First, start the DHCP server management daemon:

```shell
$ sudo dhcpdd -f examples/dhcpdd.yaml
{"level":"info","timestamp":"2021-10-26T10:45:14Z","message":"Starting server"}
```

Now, in a new terminal, create a DHCP server (be sure to adjust the config file to your requirements first):

```shell
$ dhcpdctl apply -f examples/dhcpd.yaml
dhcp server "c5c55356-d114-48f0-ab99-f1d54b46acb4" created
```

You can retrieve the running DHCP servers with `dhcpdctl get`:

```shell
$ dhcpdctl get
ID                                      DEVICE
887eb0b0-50f5-4609-aede-ea7304a40bbd    edge0
```

### 2. Setting up the `dhclient`

First, start the DHCP client management daemon:

```shell
$ sudo dhclientd -f examples/dhclientd.yaml
{"level":"info","timestamp":"2021-10-26T10:47:26Z","message":"Starting server"}
```

Now, in a new terminal, create the DHCP client (be sure to adjust the config file to your requirements first):

```shell
$ dhclientctl apply -f examples/dhclient.yaml
dhcp client "769329ba-2be5-498a-9281-e6a4aa850973" created
```

You can retrieve the running DHCP clients with `edgectl get`:

```shell
$ dhclientctl get
ID                                      DEVICE
769329ba-2be5-498a-9281-e6a4aa850973    edge1
```

🚀 **That's it!** We've successfully created a DHCP server and client.

Be sure to check out the [reference](#reference) for more information.

## Reference

### Daemons

There are two daemons, `dhcpdd` and `dhclientd`; both require root privileges.

#### `dhcpdd`

You may also set the flags by setting env variables in the format `DHCPDD_[FLAG]` (i.e. `DHCPDD_DHCPDD_CONFIGFILE=examples/dhcpdd.yaml`) or by using a [configuration file](examples/dhcpdd.yaml).

```bash
$ dhcpdd --help
dhcpdd is the ISC DHCP server management daemon.

Find more information at:
https://github.com/pojntfx/go-isc-dhcp

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
$ dhclientd --help
dhclientd is the ISC DHCP client management daemon.

Find more information at:
https://github.com/pojntfx/go-isc-dhcp

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

You may also set the flags by setting env variables in the format `DHCPD_[FLAG]` (i.e. `DHCPD_DHCPD_CONFIGFILE=examples/dhcpd.yaml`) or by using a [configuration file](examples/dhcpd.yaml).

```bash
$ dhcpdctl --help
dhcpdctl manages dhcpdd, the ISC DHCP server management daemon.

Find more information at:
https://github.com/pojntfx/go-isc-dhcp

Usage:
  dhcpdctl [command]

Available Commands:
  apply       Apply a dhcp server
  completion  generate the autocompletion script for the specified shell
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
$ dhclientctl --help
dhclientctl manages dhclientd, the ISC DHCP client management daemon.

Find more information at:
https://github.com/pojntfx/go-isc-dhcp

Usage:
  dhclientctl [command]

Available Commands:
  apply       Apply a dhcp client
  completion  generate the autocompletion script for the specified shell
  delete      Delete one or more dhcp client(s)
  get         Get one or all dhcp client(s)
  help        Help about any command

Flags:
  -h, --help   help for dhclientctl

Use "dhclientctl [command] --help" for more information about a command.
```

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build go-isc-dhcp locally, run:

```shell
$ git clone https://github.com/pojntfx/go-isc-dhcp.git
$ cd go-isc-dhcp
$ make depend
$ make
$ sudo make run/dhclientd # Or run/dhcpdd etc.
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#go-isc-dhcp:matrix.org?via=matrix.org)!

## License

go-isc-dhcp (c) 2021 Felix Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
