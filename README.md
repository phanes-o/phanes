
  
## Phanes 	法那斯 	法那斯 	奥尔弗斯传统中的原始神。 

Quick Start
Version
The version of Phanes must be v0.0.1 or above.

Environment Requirements
These environments and tools must be installed properly.

go
protoc
protoc-gen-go
The GO111MODULE should be enabled.

go env -w GO111MODULE=on
If you faced with network problem (especially you are in China Mainland), please setup GOPROXY

Install Phanes tool
You can do it either way.

1. go install installation
go install github.com/phanes-o/phanes/cmd
2. Source code compilation and installation
git clone https://github.com/phanes-o/phanes
cd phanes
make install
Project Creation
### create project's layout
phanes new helloworld

cd helloworld
### pull dependencies
go mod download
Compilation and Running
### generate all codes of proto and wire etc.
go generate ./...

### run the application
phanes run
Try it out
curl 'http://127.0.0.1:8000/helloworld/phanes'

The response should be
{
  "message": "Hello phanes"
}
Project Layout
phanes CLI always pull the layout project from GitHub for project creation. The layout project is:

Phanes Layout