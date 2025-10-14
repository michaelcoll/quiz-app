build: build-web build-go ## Build the app

clear-front: ## remove front build paths
	cd internal/web \
		&& rm -fr .output \
		&& rm -fr .nuxt \
		&& rm -fr node_modules

prepare: clear-front ## Prepares the frontend
	cd internal/web \
		&& corepack enable && corepack prepare \
		&& pnpm i \
		&& pnpm dedupe \
		&& pnpm install @openapitools/openapi-generator-cli -g

dep-upgrade: dep-upgrade-go dep-upgrade-node ## Upgrades dependencies

dep-upgrade-go:
	@go get -u
	@go mod tidy

dep-upgrade-node:
	pnpm update --latest \
		&& cd internal/web \
  		&& pnpm update --latest

build-go:
	@go build -v -ldflags="-s -w -X 'github.com/michaelcoll/quiz-app/cmd.version=v0.0.0'" .

build-web:
	cd internal/web \
		&& pnpm run generate

.PHONY: test
test: ## Launch go tests and linter
	@go test -vet=all ./...

.PHONY: coverage
coverage: ## Launch go tests with coverage
	@go test -vet=all -covermode atomic -coverprofile=coverage.out ./...

run-back: ## Launch the backend
	@go run . serve

run-front: ## Launch the frontend
	cd internal/web \
  		&& pnpm run dev

lint-front: ## Run the linter for the frontend
	cd internal/web \
  		&& pnpm run lint

run-docker:
	docker compose up --build

gen: sqlc generate ## Generate all the code

.PHONY: generate
generate:
	@go generate internal/back/domain/repositories.go

.PHONY: sqlc
sqlc:
	@sqlc generate
	@sqlc-addon generate --quiet

.PHONY: ts-model
ts-model: ## Generate typescript api model
	rm -fr api/model \
		&& openapi-generator-cli \
			generate \
			-i doc/openapi/spec.yml \
			-g typescript-axios \
			-o internal/web/api \
			--additional-properties=apiPackage=api \
			--additional-properties=modelPackage=model \
			--additional-properties=withSeparateModelsAndApi=true \
			--additional-properties=enablePostProcessFile=true \
		&& find ./internal/web/api/model -type f -exec sed -i '' 's:/\* eslint-disable \*/::g' {} \; \
		&& find ./internal/web/api/model -type f -exec sed -i '' 's@/\* tslint:disable \*/@@g' {} \; \
		&& find ./internal/web/api/model -type f -exec sed -i '' '/@ts-ignore/d' {} \; \
		&& cd internal/web \
		&& pnpm run lint

.PHONY: help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'