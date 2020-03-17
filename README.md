# Hactar deamon

> Deamon application for monitoring Filecoin nodes.

This repository contains daemon application for monitoring Filecoin nodes using [Hactar]().

## Installation
You can download already built binaries for your platform from `builds` folder or get `hactar-daemon` golang package and build it locally. Find detailed instructions below.
 
###  Download binaries

Inside the `builds` folder, you can find binaries for all supported platforms. After downloading an appropriate binary file (_in this example it is binary for Linux operating system on intel processor_) you can run daemon app with `hactar-lin-386 [command]`. 

More about different _commands_ can be found in [Usage](#Usage).

### Get `hactar-daemon` package
1. Install [Golang](https://golang.org/doc/install) **1.13 or greater**
2. Run the command below
```
go get github.com/NodeFactoryIo/hactar-daemon
```
3. Run hactar-daemon from your Go bin directory. For linux systems it will likely be:
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

#### Example
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

You can use `--debug` flag for more detailed information.

Login without prompt:
```
$ hactar start --email=user --password=pass
```

## Development
Run daemon app with `go run main.go [command]`.

More about different _commands_ can be found in [Usage](#Usage).

Expected name of the configuration file depends on `ENV` variable. For example, if you run a daemon app with `ENV=test go run main.go start`, expected config file name is `config-test.yaml`

### Using remote lotus node

`config.yaml` allows the use of remote API lotus node and miner worker. In that case paste their tokens in the config, default value is using `$HOMEDIR/.lotus/token` and `$HOMEDIR/.lotusstorage/token`.

## License

This project is dual-licensed under Apache 2.0 and MIT terms:
- Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

![Hactar](hactar-logo.png)
![Filecoin](filecoin-logo.png)