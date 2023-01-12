# GetCaptcha Service

This is the GetCaptcha service

Generated with

```
micro new bj38web/service/getCaptcha --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.getCaptcha
- Type: srv
- Alias: getCaptcha

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./getCaptcha-srv
```

Build a docker image
```
make docker
```