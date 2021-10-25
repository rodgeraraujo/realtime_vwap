echo "> Test"

GO111MODULE=on go test -race -timeout=2m ./... --v -ginkgo.progress $@  | grep -v 'no test files'