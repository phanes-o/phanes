## Phanes

[Quick Start](https://learnku.com/articles/69363)

Version
The version of Phanes must be v0.1.0 or above.

Environment Requirements
These environments and tools must be installed properly.

go
protoc
protoc-gen-go

The GO111MODULE should be enabled.
```shell
go env -w GO111MODULE=on
```

If you faced with network problem (especially you are in China Mainland), please setup GOPROXY

Install Phanes tool
You can do it either way.

1. go install installation
```
go install github.com/phanes-o/phanes@latest
```
2. Source code compilation and installation
```sehll
git clone https://github.com/phanes-o/phanes
cd phanes
make build
```
Project Creation
### create project's layout

```
phanes new helloworld
cd helloworld
```

### pull dependencies
```sehll
go mod download 
```
Compilation and Running
### generate all codes of proto or wire etc.
```sehll
go generate ./...
```

### generate proto
```shell
phanes proto client hello.proto
phanes proto server hello.proto -t internal/service
```
### run the application
```shell
phanes run
```
Project Layout
phanes CLI always pull the layout project from GitHub for project creation. The layout project is:

[Phanes Layout](https://github.com/phanes-o/phanes-layout)