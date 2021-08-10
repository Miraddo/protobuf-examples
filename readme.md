# Protocol Buffers Examples with Go
In this repo, you can find some essential parts of the code that wrote with proto3.

I tried to explain each not part as simply as I could as a comment inside the files.

## Installation Protoc
`Protoc` is the tool that let us compile the `.proto` files to the main code that we want to use

It's so easy to install `protoc` in MacOS/Linux, but for windows it might be a little hard
so here I encourage you to install `choco` as the package manager for windows choco is short of [chocolaty package manager](https://chocolatey.org/)

after installing `choco`, just run your Terminal/PowerShell/CMD then run `choco install protoc` it will be installed automatically

then run `protoc` you should see the result like this:

```text
PS C:\Users\Miraddo> protoc
Usage: C:\ProgramData\chocolatey\lib\protoc\tools\bin\protoc.exe [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
                              If not found in any of the these directories,
                              the --descriptor_set_in descriptors will be
                              checked for required proto file.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --deterministic_output      When using --encode, ensure map fields are
                              deterministically ordered. Note that this order
                              is not canonical, and changes across builds or
                              releases of protoc.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
  
  ...
  
```

## How to use protoc 

while you try to convert the protofiles to go file you should always add `option go_package = "path";` after syntax inside protofiles

The `go_package` option defines the import path of the package which will contain all the generated code for this file. The Go package name will be the last path component of the import path. [more details](https://developers.google.com/protocol-buffers/docs/gotutorial#defining-your-protocol-format)

then follow the steps

- go to inside the project folder for me in this case it go-example folder `cd go-example`
- then run the code below here
```text
    protoc --go_out=. .\protofiles\Person.proto
```
- `--go_out=` is the path that output of golang will be created there. in the "." means current directory
- `.\protofiles\Person.proto` is the file that we want to generate code with protoc 


Tip: if you got any error it might be you don't install protobuf package for your project

so install those packages
```text
    go get -u google.golang.org/grpc
    go get -u github.com/golang/protobuf/proto
    go get -u github.com/golang/protobuf/protoc-gen-go
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
