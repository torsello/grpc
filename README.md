# gRPC

Application of airport consultations

# Server instalation

Install [Python 2.7.15](https://www.python.org/downloads/release/python-2715/).

## Install gRPC (Python)
```bash
$ python -m pip install grpcio
```

## Install gRPC Tools(Python)
```bash
$ python -m pip install grpcio-tools
```

## Run gRPC Server application
```bash
$ python server.py
```

# Client installation

Install [Golang](https://golang.org/dl/). (Requires Go 1.6 or higher).

## Install Protocol Buffers v3

Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(**protoc-<version>-<platform>.zip**) from here: [github Protoc](https://github.com/google/protobuf/releases)

- Unzip this file.  
- Update the environment variable PATH to include the path to the protoc binary file.

Next, install the protoc plugin for Go

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go
```
The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.


```bash
$ export PATH=$PATH:$GOPATH/bin
```

## Run gRPC Client application

```bash
$ go run client.go
```
