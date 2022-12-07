COMMONENVVAR = GOOS=linux GOARCH=amd64
BUILDENVVAR = CGO_ENABLED=0
RUNTIME ?= docker
REPOOWNER ?= swsehgal
IMAGENAME ?= devlocator
IMAGETAG ?= latest
DEVLOCATOR_CONTAINER_IMAGE ?= quay.io/${REPOOWNER}/${IMAGENAME}:${IMAGETAG}

.PHONY: all
all: build

.PHONY: build
build: gofmt
	 $(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./bin/devlocator main.go

.PHONY: gofmt
gofmt:
	@echo "Running gofmt"
	gofmt -s -w `find . -path ./vendor -prune -o -type f -name '*.go' -print`

.PHONY: govet
govet:
	@echo "Running go vet"
	go vet ./cmd/... ./pkg/...

outdir:
	@mkdir -p _out || :

.PHONY: image
image: build
	@echo "building image"
	$(RUNTIME) build -f Dockerfile_devlocator -t $(DEVLOCATOR_CONTAINER_IMAGE) .

.PHONY: push
push: image
	@echo "pushing image"
	$(RUNTIME) push $(DEVLOCATOR_CONTAINER_IMAGE)

.PHONY: deploy
deploy:
	@echo "Deploying test application"
	kubectl create -f deployment.yaml

.PHONY: undeploy
undeploy:
	@echo "Removing test application"
	kubectl delete -f deployment.yaml

clean-binaries:
	rm -f ./usr/local/bin/devlocator

clean: clean-binaries
	kubectl delete -f deployment.yaml
