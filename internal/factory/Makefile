CVPKG=go list ./... | grep -v mocks | grep -v internal/
GO_TEST=go test `$(CVPKG)` -race
COVERAGE_FILE="coverage.out"

clean.test:
	go clean --testcache

test:
	go test

test.clean: clean.test test

test.coverage:
	$(GO_TEST) -covermode=atomic -coverprofile=$(COVERAGE_FILE)
