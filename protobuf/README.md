# Protobuf

[Protocol Buffers](https://protobuf.dev) are language-neutral, platform-netral extensible mechanisms for serializing structured data.

It's the google's language, similar to XML, but smaller, faster and simpler.

You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

Protocol buffers supports _C++_, _C#_, _Dart_, _Go_, _Java_, _Kotlin_, _Objective-C_, _Python_ and _Ruby_. Now, with proto3, you can work with _PHP_.

## 🆚 Protobuf vs JSON

JSON is a data format, widely used in almost all modern REST APIs.

Protobuf is a good alternative for Web APIs, mostly for Microservices.

It has:

- **Efficiency**, are more compact than JSON, because it uses a binary format and thus it uses lesser space, and can be transmitted faster. This is particular benefitial for cases where bandwidth and storage is limited, or when dealing with large amounts of data.
- **Speed**, is faster than JSON in serialization and desarialization, becuase of its binary format.
- **Schema and Type Safety**, requires defining a schema using .protofiles, which provides that strong type safety.
- **Forward and backward compatibility**, you can add new things without breaking code.

## 💽 Installation on Linux

First, download the .tar.gz or .zip package

```zsh
  https://github.com/protocolbuffers/protobuf/releases/latest
```

Or just clone it with git, making sure you run the configure script of the submodules, [learn more](https://github.com/protocolbuffers/protobuf/tree/main/src).

Make sure you have install bazel and g++:

```zsh
  sudo apt-get install g++ bazel
```

This is for ubuntu users, if you have a different distro, you can check the [bazel's official page](https://bazel.build).

Once you have it, you need to build the files with bazel:

```zsh
  bazel build :protoc :protobuf
```

Then you can install the compiler:

```zsh
  cp bazel-bin/protoc /usr/local/bin
```

> [!NOTE]
> Make sure you have go on your system's PATH.

```zsh
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOPATH/bin
```

And then install protobuf package and the protoc-gen-go plugin for generating the go files.

```zsh
  go get google.golang.org/protobuf@latest
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

For executing the proto files:

```sh
  # protoc is the compiler which will generate the code from the proto files
  # the --go_out tells the compiler that we want to generate go files and
  # also specify where to store those files.
  # --go_opt this controls how the path for the generated code are handled
  # by setting the source_relative we tell the compiler to keep the relative
  # dirs of our proto files when generating go files.
  # Finally the folder where the proto files are.
  protoc --go_out=. --go_opt=paths=source_relative models/person.proto
```

## 📫 Making requests

```zsh
  # For adding a person
  curl -X POST --data-binary @tmp/person.bin http://localhost:8080/add

  # For getting a user
  curl -X GET "http://localhost:8080/get?id=1234" --output tmp/retrieved_person.bin
```
