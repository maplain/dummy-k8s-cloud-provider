# Cloud provider for Guest Cluster

## Cloud controller manager for Guest Cluster

This repository provides tools and scripts for building and testing `Kubernetes cloud-controller-manager` for Guest Cluster. The project is under development and should not be used in production.

## Building the cloud provider

This section outlines how to build the cloud provider with and without Docker.

### Building locally

Build locally with the following command:

```shell
$ git clone https://gitlab.eng.vmware.com/fangyuanl/cloud-provider-guest-cluster && \
  make -C cloud-provider-guest-cluster
```

The project uses [Go modules](https://github.com/golang/go/wiki/Modules) and:
* Requires Go 1.11+
* Should not be cloned into the `$GOPATH`

### Building with Docker

It is also possible to build the cloud provider with Docker in order to ensure a clean build environment:

```shell
$ git clone https://gitlab.eng.vmware.com/fangyuanl/cloud-provider-guest-cluster && \
  make -C cloud-provider-guest-cluster build-with-docker
```

### Guest Cluster storage support
currently, there is no storage support.
