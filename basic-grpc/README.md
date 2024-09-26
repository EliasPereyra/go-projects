# gRPC

[gRPC](https://grpc.io/) is an open source framework made by Google for handling RPC calls. It can be run in any platform and support a variety of languages such as Go, C++, Java, etc.

It's more suitable for distributed microservices because of its efficient, structured and secure messages.

It also allows to scale easily to millions of RPCs per second.

## Protocol Buffer
It's a mechanism for serializing structured data, with a syntax similar to JSON but it's faster.

You define a `message` for an object structure and then process it through `protoc` to generate data access in your preffered language.

It provides default methods for the fields like getting a specified field, or modify it, as well as serializing/parsing.

## gRPC vs REST
- Protocol: **gRPC** uses HTTP/2 which provides better performance and less latency, while REST uses HTTP/1.1
- Data format: **gRPC** uses Protocol Buffers, while REST uses the traditional JSON or XML data formats.
- API design: **gRPC** follows the [RPC paradigm](https://en.wikipedia.org/wiki/Remote_procedure_call), while REST adheres to the [Representational State Transfer model](https://en.wikipedia.org/wiki/REST).

>[!NOTE]
>You'll need the go plugin for the proto buffer: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` and for generating code for calling grpc calls in proto buf definitions: `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
