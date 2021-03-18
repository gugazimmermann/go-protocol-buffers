# Protocol Buffers with Golang

Protocol Buffers (also called Protobuf) is a method of serializing structured data. It is language-neutral, platform-neutral, extensible mechanism for serializing structured data. It is like XML or JSON, but smaller, faster, and simpler. You define how you want your data to be structured once like a template, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages. It is useful in developing programs to communicate with each other over a wire or for storing data.

Why should we use it?

Protobuf is -

-  Smaller: It is 3-10 times smaller than XML or JSON. If you are storing data then it will reduce storage space. If you are using it as RPC then it will reduce latency as message size will be much smaller.
- Faster: Serialization and deserializations is 20-100 times faster which will increase performance of your system as it is using less time in serialization and deserializations (source).
- Simpler yet extensible: It is simple to implement and start with. It is very easy to update or extend an existing protobuf message file.
- Language and platform-neutral: Protobuf messages are language and platform independent. That gives you freedom to choose technology of your choice for different entity using this protocol. For example your client can be using C# but it can still communicate with your server written in Java through protobuf.

Protobuf with Golang -

Both Protobuf and Golang are from Google and Golang has official support of Protobuf. There are 3 simple steps to use Protobuf in any language -

* Define message formats in a .proto file.
* Use the protocol buffer compiler (protoc).
* Use the Go protocol buffer API to write and read messages. 

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