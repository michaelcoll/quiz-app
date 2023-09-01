build: build-web build-go

prepare:
	cd internal/web \
		&& corepack enable && corepack prepare \
		&& rm -fr .output \
        && rm -fr .nuxt \
        && rm -fr node_modules \
		&& pnpm i

dep-upgrade: dep-upgrade-go dep-upgrade-node

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

build-docker:
	@docker build . -t michaelcoll/quiz-app:latest --pull --build-arg VERSION=v0.0.1

.PHONY: test
test:
	@go test -vet=all ./...

.PHONY: coverage
coverage:
	@go test -vet=all -covermode atomic -coverprofile=coverage.out ./...

run-back:
	@go run . serve

run-front:
	cd internal/web \
  		&& pnpm run dev

lint-front:
	cd internal/web \
  		&& pnpm run lint

run-docker:
	docker run -ti --rm -p 8080:8080 web:latest

gen: sqlc generate

.PHONY: generate
generate:
	@go generate internal/back/domain/repositories.go

.PHONY: sqlc
sqlc:
	@sqlc generate
	@sqlc-addon generate --quiet

.PHONY: ts-model
ts-model:
	openapi-generator-cli \
		generate \
		-i doc/openapi/spec.yml \
		-g typescript-axios \
		-o internal/web/api \
		--additional-properties=apiPackage=api \
		--additional-properties=modelPackage=model \
		--additional-properties=withSeparateModelsAndApi=true \
		--additional-properties=enablePostProcessFile=true \
		&& find ./internal/web/api/model -type f -exec sed -i 's:/\* eslint-disable \*/::g' {} \; \
		&& find ./internal/web/api/model -type f -exec sed -i 's@/\* tslint:disable \*/@@g' {} \; \
		&& cd internal/web \
		&& pnpm run lint