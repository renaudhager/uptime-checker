NAME       := uptime-checker
IMAGE_NAME := renaudhager/uptime-checker
VERSION    :=$(shell git describe --abbrev=0 --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
LDFLAGS    := -w -extldflags "-static" -X 'main.version=$(VERSION)'

.PHONY: setup
setup:
	go get -u -v golang.org/x/lint/golint
	go get -u -v github.com/golang/dep/cmd/dep

.PHONY: deps
deps:
	dep ensure -v

.PHONY: lint
lint:
	golint -set_exit_status .
	go vet ./...

.PHONY: build-docker
build-docker:
	CGO_ENABLED=0 go build -o $(NAME) -ldflags "$(LDFLAGS)" .
	strip $(NAME)
	cp $(NAME) $(NAME)_$(VERSION)

.PHONY: build
build:
	docker build --tag="$(IMAGE_NAME):$(VERSION)" --rm .

.PHONY: publish
publish:
	docker push "$(IMAGE_NAME):$(VERSION)"
