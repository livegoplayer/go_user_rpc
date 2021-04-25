# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

COMMIT_SHA=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags)
BRANCH=$(shell git symbolic-ref --short -q HEAD)
LDX= "-X '$(BINARY_NAME)/config.Version=$(VERSION)' -X '$(BINARY_NAME)/config.LastCommit=$(COMMIT_SHA)' -X '$(BINARY_NAME)/config.Branch=$(BRANCH)'"

all: run
# 生成自动code
code:
	$(GOCMD) run ./cli/private_model --app-path "`pwd`/build"

buildfile:
	make code
	$(GOBUILD) -o ./build/$(BINARY_NAME) -ldflags $(LDX) -tags=jsoniter -v ./

test:
	make code
	cd ./models/ && $(GOTEST) --app-path "`pwd`/../build/"

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)
	rm -f ./build/$(BINARY_UNIX)

model_temp:
	$(GOCMD) run ./cli/parse_database/parse_database.go -prefix=us_ -out=repository_template_model

model_query:
	$(GOCMD) run ./cli/private_model/parse_private_model.go

collect:
	$(GOCMD) run ./cli/parse_collect/parse_collect.go

models:
	make model_temp
	make model_query
	make collect

exec:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOCMD) build main.go --app-path "`pwd`"



