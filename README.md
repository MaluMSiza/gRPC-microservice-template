GO GRPC MICROSERVICE TEMPLATE
=============================

A production-ready template for building gRPC microservices in Go.

TABLE OF CONTENTS
-----------------
- Features
- Prerequisites
- Getting Started
- Project Structure
- Testing
- Docker Deployment
- Makefile Commands
- Contributing
- License
- About the Test Directory

FEATURES
--------
- gRPC server/client implementation
- Protocol Buffers (proto3) setup
- Docker containerization
- Docker Compose orchestration
- Makefile automation
- Unit testing structure
- Generated code management

PREREQUISITES
-------------
- Go 1.21+
- Docker 20.10+
- Docker Compose 2.20+
- Protocol Buffers Compiler (protoc)
- Go plugins:
  $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

GETTING STARTED
---------------
1. Clone the template:
````shell
  git clone https://github.com/MaluMSiza/gRPC-microservice-template.git
  cd grpc-microservice-template
````


2. Generate gRPC code:
````shell
  $ make generate
````

3. Build the project:
````shell
  $ make build-server build-client
````

PROJECT STRUCTURE
-----------------
grpc-microservice-template/  
├── .gitignore  
├── Makefile  
├── Dockerfile  
├── docker-compose.yml  
├── go.mod  
├── proto/  
│   └── greet.proto  
├── gen/  
├── service/  
│   └── greeter.go  
├── cmd/  
│   ├── server/  
│   └── client/  
└── test/  
└── greeter_test.go  

TESTING
-------
The test/ directory contains unit and integration tests. Example test:
````go
// test/greeter_test.go
package test

import (
"context"
"testing"
"time"
// ... imports ...
)

func TestGreeterService(t *testing.T) {
	// Test setup and cases
}
````

1. Run tests with:
   ```shell
   $ make test
   ```

DOCKER DEPLOYMENT
-----------------
1. Build containers:
   ```shell
   $ make docker-build
   ```
   

2. Start services:
   ```shell
   $ make docker-run
   ```

MAKEFILE COMMANDS
-----------------

````json
generate       - Generate gRPC code   
build-server   - Build server binary   
build-client   - Build client binary  
clean          - Remove generated files  
test           - Run all tests  
docker-build   - Build Docker images  
docker-run     - Start Docker containers 
````

CONTRIBUTING
------------
1. Fork the repository
2. Create a feature branch
3. Submit a Pull Request

LICENSE
-------
MIT License

ABOUT THE TEST DIRECTORY
------------------------
- Contains unit/integration tests
- Includes example test for gRPC service
- Uses Go's testing package
- Supports table-driven tests
- Provides in-memory server testing