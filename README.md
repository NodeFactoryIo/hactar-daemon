# Hactar deamon

> Deamon application for monitoring Filecoin nodes.

This repository contains deamon application for monitoring Filecoin nodes using [Hactar]().

## Installation
1. Install [Golang](https://golang.org/doc/install) **1.13 or greater**
2. Run the command below
```
go get github.com/NodeFactoryIo/hactar-daemon
```
3. Run hactar-deamon from your Go bin directory. For linux systems it will likely be:
```
~/go/bin/hactar-daemon
```
Note that if you need to do this, you probably want to add your Go bin directory to your $PATH to make things easier!

## Usage

Lotus username and password can be passed as flags when calling start, see [examples](#example).

```
$ hactar -h

hactar-deamon is a command line interface for ....

Usage:
  hactar [command]

Available Commands:
  start       Starts hactar deamon app
  token       Display lotus token
  help        Help about any command
  version     Show the current version of Hactar deamon app

Use "hactar [command] --help" for more information about a command.
```

####Example
Starting hactar without using flags.
```
$ hactar start
```
```
Enter lotus username: >> cli input <<
Enter lotus password: >> cli input <<
Starting hactar.....
```

Starting hactar with flags.
```
$ hactar start --username=user --password=pass
```
```
Starting hactar.....
```

### Notice

## License

This project is dual-licensed under Apache 2.0 and MIT terms:
- Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

![Hactar](hactar-logo.png)
![Filecoin](filecoin-logo.png)