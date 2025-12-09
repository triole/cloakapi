# Cloakapi

<!-- toc -->

- [Synopsis](#synopsis)
- [Features](#features)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Commands](#commands)
  - [Flags](#flags)
- [Examples](#examples)
- [Template System](#template-system)
  - [Template Variables](#template-variables)
  - [Example Templates](#example-templates)
- [Help](#help)
- [Keycloak config](#keycloak-config)

<!-- /toc -->

## Synopsis

Cloakapi is a command-line tool that provides an abstraction layer for the Keycloak Admin API. It allows users to easily list, fetch, and manipulate Keycloak entities such as users, federated identities, and identity providers through a simple CLI interface.

## Features

- List Keycloak entities (users, federated identities, identity providers)
- Execute templates to transform user data
- Support for multiple configuration file formats (JSON, TOML, YAML)
- Support for multiple output formats (JSON, TOML, YAML, table)
- Proxy and insecure connection support
- Dry-run mode for testing commands

## Configuration

Cloakapi supports multiple configuration file formats (JSON, TOML, YAML). The configuration file should contain the following fields:

```toml
url = "http://localhost:8080"
realm = "master"
client_id = "myclient"
client_secret = "mysecret"
# proxy = "socks5://proxy:3333"
insecure = true
```

The configuration file can be specified using the `-c` or `--conf` flag. The tool will search for configuration files in the following locations:

1. Current directory
2. `$HOME/.config/cloakapi`
3. `$HOME/.conf/cloakapi`
4. Executable directory

The flag value can either be:

- An explicit file path to a configuration file (e.g., `--conf=/path/to/myconfig.toml`)
- A string that is used as a matcher to detect the configuration file (e.g., `--conf=myconfig` would match files named myconfig.toml, myconfig.json, etc.)

## Usage

```bash
cloakapi --conf=STRING <command> [flags]
```

### Commands

- `ls` - list entities, available commands: fed, usr, att, idp
- `tpl` - execute template string or load from file
- `var` - list available template variables

### Flags

- `-h, --help` - Show context-sensitive help
- `-c, --conf=STRING` - config file detection expression
- `-o, --output="table"` - output format (json, toml, yaml, table)
- `--log-file="/dev/stdout"` - log file
- `--log-level="info"` - log level (trace, debug, info, error)
- `--log-no-colors` - disable output colours, print plain text
- `--log-json` - enable json log, instead of text one
- `-n, --dry-run` - dry run, just print operations that would run
- `-V, --version-flag` - display version

## Examples

```bash
# list Users
cloakapi --conf=myconfig.toml ls usr

# list Federated Identities
cloakapi --conf=myconfig.toml ls fed

# list Identity Providers
cloakapi --conf=myconfig.toml ls idp

# execute Template
cloakapi --conf=myconfig.toml tpl -f template.tpl
cloakapi --conf=myconfig.toml tpl -s "{{.username}}:{{.id}}"

# list Available Template Variables
cloakapi --conf=myconfig.toml var
```

## Template System

Cloakapi leverages Go template syntax to transform and format user data. Templates can be provided as command-line arguments or loaded from files, enabling flexible data processing and output generation.

### Template Variables

The template system provides access to all fields from Keycloak's User and FederatedIdentityRepresentation structures, automatically converted to snake_case for easy reference. This includes common user attributes such as `id`, `username`, `email`, `first_name`, `last_name`, and federated identity information.

### Example Templates

Example template 1 (`examples/template1.tpl`):

```json
{"{{.id}}":"{{.username}}"}
```

This template creates a JSON mapping of user IDs to usernames.

Example template 2 (`examples/template2.tpl`):

```go
{{- if eq .remote_id ""}}{{.username}}{{end}}
```

This template outputs usernames for users without a remote ID.

Additional usage examples include:

- Creating CSV output from user data
- Generating JSON mappings for external systems
- Filtering users based on specific criteria
- Formatting user information for display purposes

## Help

Call help by using the `-h` CLI parameter.

```go mdox-exec="r -h"
Usage: cloakapi --conf=STRING <command> [flags]

an abstraction layer for the Keycloak Admin API

Flags:
  -h, --help                      Show context-sensitive help.
  -c, --conf=STRING               config file detection expression
  -o, --output="table"            output format
      --log-file="/dev/stdout"    log file
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -n, --dry-run                   dry run, just print operations that would run
  -V, --version-flag              display version

Commands:
  ls     list entities, available commands: fed, usr, att, idp
  tpl    execute template string or load from file
  var    list available template variables

Run "cloakapi <command> --help" for more information on a command.
```

## Keycloak config

To use Cloakapi with Keycloak, you need to configure a service account client with appropriate permissions. Follow these steps:

1. **Create a client**:

   - Navigate to your Keycloak admin console
   - Go to "Clients" section
   - Click "Create client"
   - Set the client ID to "service-account" (or any name you prefer)
   - Select "Service Account" as the client type
2. **Configure client capabilities**: Under the "Capability Config" section, ensure the following settings are enabled:

   - **Client Authentication**: Enabled
   - **Authorization**: Enabled
   - **Authentication Flow, Standard Flow**: Enabled
   - **Authentication Flow, Direct Access Grants**: Enabled

   The rest of the capability config options can remain disabled.
3. **Assign required roles**:

   - Navigate to the "Service accounts roles" tab in the client settings
   - Assign the necessary Keycloak roles that your service account needs to perform operations
   - For basic functionality, you'll typically need at least "view-users" or "manage-users" roles
   - You may also need "view-realm" or "view-identity-providers" depending on your specific use cases
4. **Obtain client credentials**:

   - After creating the client, go to the "Credentials" tab
   - Note down the client ID and client secret
   - These will be used in your Cloakapi configuration file
