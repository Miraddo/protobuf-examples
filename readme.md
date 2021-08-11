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
    go get -u github.com/golang/protobuf/proto
    go get -u github.com/golang/protobuf/protoc-gen-go
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## learn more
some tips that I taught might be important.

- protocol buffers have full compatibility [read](https://developers.google.com/protocol-buffers/docs/gotutorial#extending-a-protocol-buffer)
    - Rules (There are [some exceptions](https://developers.google.com/protocol-buffers/docs/proto3#updating) to these rules, but they are rarely used.)
        - you must not change the tag numbers of any existing fields.
        - you may delete fields.
        - you may add new fields, but you must use fresh tag numbers (i.e., tag numbers that were never used in this protocol buffer, not even by deleted fields).


### remove a field
while removing a field, you should use `reserved` for both name and the tag

while using reserved, you should know how to use it:
first, you should know we will using reserved to prevent our system crash or get bugs while removing fields

- you can not use the field name and the tag together
- you can use field names together like `reserved "name", "foo", "others...";`
- you can use tags together like `reserved 2,5,8,10 to 15` tip: "10 to 15" is the range of 10,11,12,13,14,15
- don't remove reversed tag **never** :)


```protobuf
message example{
  int32 num = 1;
  string name =2; // now I want to delete this field
}

// try to remove name field here 
message example{
  // always use reserve
  reserved 2;
  reserved "name";
  // then the other part of you message
  int32 num = 1;
}
```

use `reserved` is import to prevent conflicts in the codebase

there is another option we can rename the field name

```protobuf
// try to remove name field here 
message example{
  int32 num = 1;
  string not_will_be_used_name = 2;
}
```
anyway, it would be best if you used `reserved` while removing :)

### Enums
We can use enums for adding or removing as well.

Tip: just use `reserved` while removing :)
```protobuf
message example{
  enum Ext{
    Default = 0;
    Ext1 = 1;
    Ext2 = 3;
  }
  Ext example = 1;
}
```

### Defaults

We should care about default values. They are important.
We should know the default values don't make code get broken. In other words, defaults shouldn't have meaning at all in our code.
