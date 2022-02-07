## Table of Contents

- [Requirements](#requirements)
- [Usage](#usage)

## Requirements

requires the following to run:

- go ^1.17

## Usage

```
    go run main.go
```

## protoc

```
    protoc --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false api/grpc/protos/*.proto
```

## wire

```
   wire cmd/wire.go
```
