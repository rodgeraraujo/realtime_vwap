###############################
# Rules for local development #
###############################

.PHONY: test
test:
	@./scripts/test.sh ./cmd/... ./pkg/...

.PHONY: test-cov
test-cov:
	@./scripts/test-cov.sh

.PHONY: test-cov-clean
test-cov-clean:
	@./scripts/test-cover-clean.sh

.PHONY: install
install:
	./scripts/install.sh

.PHONY: build
build:
	./scripts/build.sh