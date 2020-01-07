# hactar-daemon
Hactar daemon app built using Go for communication with Filecoin lotus-node.

```
hactar-deamon is a command line interface for ....

Usage:
  hactar [command]

Available Commands:
  start       Starts hactar deamon app
  token       Display lotus token
  help        Help about any command
  version     Show the current version of Hactar deamon app

Start:

    Starts hactar deamon app, connects with hactar service and starts monitoring process.

    Flags:
      -u, --username string         Username
      -p, --password string         Password
      -h, --help                    help for start command

Use "hactar [command] --help" for more information about a command.
```

####Example
Starting hactar without using flags.
```
    hactar start
```
```
    Enter lotus username: XXX
    Enter lotus password: XXX
    Starting hactar.....
```