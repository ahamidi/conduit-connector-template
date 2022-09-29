.PHONY: build test test-integration

build:
	go build -o conduit-connector-template cmd/connector/main.go

test:
	go test $(GOTEST_FLAGS) -v -race ./...

test-integration:
	# run required docker containers, execute integration tests, stop containers after tests
    # Example:
	# docker compose -f test/docker-compose-template.yml up --quiet-pull -d --wait
	# go test $(GOTEST_FLAGS) -v -race ./...; ret=$$?; \
	#	docker compose -f test/docker-compose-template.yml down; \
	#	exit $$ret
