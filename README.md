# Protocol Buffers 3 with Golang

https://www.linkedin.com/pulse/google-protocol-buffers-3-go-jos%C3%A9-augusto-zimmermann-negreiros

[linkedin article](linkedin.md)

## Dep

* protoc: `https://github.com/google/protobuf`
* `go get -u github.com/golang/protobuf/protoc-gen-go`
* `go get -u github.com/golang/protobuf/proto`

## To see: Protocol Buffers for Go with Gadgets

https://github.com/gogo/protobuf



gogoprotobuf is a fork of golang/protobuf with extra code generation features.

This code generation is used to achieve:

- fast marshalling and unmarshalling

- more canonical Go structures

- goprotobuf compatibility

- less typing by optionally generating extra helper code

- peace of mind by optionally generating test and benchmark code

- other serialization formats

There's a really nice article describing how to get started: https://jbrandhorst.com/post/gogoproto/