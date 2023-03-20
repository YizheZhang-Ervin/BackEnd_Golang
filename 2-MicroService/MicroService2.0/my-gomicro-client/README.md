# MyGomicroClient Service

This is the MyGomicroClient service

Generated with

```
micro new --namespace=go.micro --type=web my-gomicro-client
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.my-gomicro-client
- Type: web
- Alias: my-gomicro-client

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./my-gomicro-client-web
```

Build a docker image
```
make docker
```