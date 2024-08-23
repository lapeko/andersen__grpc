#!/bin/bash

/e/protoc/bin/protoc.exe --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. greet/greetpb/greet.proto
