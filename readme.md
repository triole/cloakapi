# CloakAPI

## Synopsis

An abstraction layer for the Keycloak API.

## Conf

```go mdox-exec="cat examples/conf.toml"
url = "http://localhost:8080"
realm = "master"
client_id = "myclient"
client_secret = "mysecret"
proxy = ""
insecure = true
```

## Help

Call help by using the `-h` CLI parameter.

```go mdox-exec="r -h"
Usage: cloakAPI <command> [flags]

an abstraction layer for the Keycloak Admin API

Flags:
  -h, --help                      Show context-sensitive help.
  -c, --conf="conf.toml"          path to config file
      --log-file="/dev/stdout"    log file
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -n, --dry-run                   dry run, just print operations that would run
  -V, --version-flag              display version

Commands:
  ls    list entities, can be: users,idps

Run "cloakAPI <command> --help" for more information on a command.
```
