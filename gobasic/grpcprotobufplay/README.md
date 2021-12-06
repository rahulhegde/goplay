https://developers.google.com/protocol-buffers/docs/gotutorial

## protobuffer specification
- 


## Compiler & Geneation 

### background
protoc - protoc compiler
    - written in different language, requires language specifies protoc runtime to generates corresponding specific files.
    - c++ is default implementation - https://github.com/protocolbuffers/protobuf
    - Binary availble - https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip

protoc runtime for protobuf, grpc go - 
    - for go related ouput  -protoc-go-gen must be present in path
    - go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    - go get google.golang.org/grpc/cmd/protoc-gen-go-grpc


### Compilation 
protoc -I=./ --go_out=/home/developer/workspace/go-ws/src --go-grpc_out=/home/developer/workspace/go-ws/src ./addressbook.proto
protoc -I=./ --go_out=/home/developer/workspace/go-ws/src --go-grpc_out=/home/developer/workspace/go-ws/src ./serviceapis.proto

## Example