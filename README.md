# Akamai CLI for API Gateway

Akamai CLI for API Gateway allows you manage the Akamai API Gateway.

## Install

To install, use [Akamai CLI](https://github.com/akamai/cli):

```
akamai install api-gateway
```

You may also use this as a stand-alone command by downloading the
[latest release binary](https://github.com/akamai/cli-api-gateway/releases)
for your system, or by cloning this repository and compiling it yourself.

### Compiling from Source

If you want to compile it from source, you will need Go 1.7 or later, and the [Dep](https://golang.github.io/dep/) package manager installed:

1. Fetch the package:  
  `go get github.com/akamai/cli-api-gateway`
2. Change to the package directory:  
  `cd $GOPATH/src/github.com/akamai/cli-api-gateway`
3. Install dependencies using `dep`:  
  `dep ensure`
4. Compile the binaries:  
  - Linux/macOS/*nix: 
    - `go build -o akamai-api-gateway ./api-gateway`
    - `go build -o akamai-api-keys ./api-keys`
    - `go build -o akamai-api-security ./api-security`
  - Windows: 
    - `go build -o akamai-api-gateway.exe ./api-gateway`
    - `go build -o akamai-api-keys.exe ./api-keys`
    - `go build -o akamai-api-security.exe ./api-security`
5. Move the binaries in to your `PATH`

## API Gateway: Usage

```
akamai api-gateway [global flags] <command> [sub-command]
```

## Sub-Commands
- list - List Commands
- help - Displays help information
- create - Create an API Endpoint
- import - Create or Update an API Endpoint using the specified Swagger 2.0 or RAML 0.8 file
- update - Update an API Endpoint version
- list-endpoints - Retrieve a list of APIs running through the Gateway
- list-resources - Retrieve a list resources for a given endpoint.
- activate - Activate an API
- remove - Remove an API Endpoint
- privacy-add - Set an endpoint/resource/method to public/private
- clone - Clone an API Endpoint version
- status - Show the status of the API endpoint

## Global Flags
- `--edgerc value` — Location of the credentials file (default: "/Users/dshafik") [$AKAMAI_EDGERC]
- `--section value` — Section of the credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
- `--version`, `-v` — print the version
