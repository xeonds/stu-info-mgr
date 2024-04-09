NAME=stu-info-mgr
BINDIR=build
VERSION=1.0.0
BUILDTIME=$(shell date -u)
GOBUILD=go build
GOFLAGS=-ldflags '-s -w -X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'

all: gen linux-amd64

gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/student.proto

linux-amd64: 
	GOOS=linux GOARCH=amd64 cd client && $(GOBUILD) -o ../$(BINDIR)/$(NAME)-client-$@
	GOOS=linux GOARCH=amd64 cd server && $(GOBUILD) -o ../$(BINDIR)/$(NAME)-server-$@
