#!/bin/bash

protoc -I . --go_out=. src/simple/simple.proto
protoc -I . --go_out=. src/enum_example/enum_example.proto
protoc -I . --go_out=. src/complex/complex.proto