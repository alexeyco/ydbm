all:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

.PHONY: install
install: ## Install dependencies
	@go mod tidy && go mod vendor && go mod verify

.PHONY: mock
mock: ## Generate mock files
	@mockgen -package=generator_test -source=internal/generator/action.go -destination=internal/generator/action_mock_test.go

.PHONY: lint
lint: ## Run linter
	@golangci-lint --exclude-use-default=false run ./...
