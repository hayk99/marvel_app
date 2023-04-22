export GOCOVERDIR := "false"
generate: clean
	go generate -x ./...

clean:
    find . -name mocks -type d | xargs rm -fr
    find . -name *.coverprofile | xargs rm -fr

fmt:
    go install mvdan.cc/gofumpt@latest
    go fmt ./...
    gofumpt -w ./

lint:
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    golangci-lint run --timeout 1h

test: deps test-unit test-integration

test-unit:
        @ginkgo -p --randomize-all --randomize-suites -skip-file=infrastructure -trace -race --show-node-events -cover  -r ./pkg/... -r ./internal/...

test-integration: deps
    @ginkgo --randomize-all --randomize-suites -trace -race --show-node-events  -cover -p -r ./pkg/infrastructure/...

deps:
    go install github.com/onsi/ginkgo/v2/ginkgo@latest
    go install github.com/golang/mock/mockgen@latest
