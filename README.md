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
