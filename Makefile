build: build-web build-go

prepare:
	cd internal/web \
		&& corepack enable && corepack prepare \
		&& pnpm i

dep-upgrade: dep-upgrade-go dep-upgrade-node

dep-upgrade-go:
	@go get -u
	@go mod tidy

dep-upgrade-node:
	cd internal/web \
  		&& pnpm update --latest

build-go:
	@go build -v -ldflags="-s -w -X 'github.com/michaelcoll/quiz-app/cmd.version=v0.0.0'" .

build-web:
	cd internal/web \
		&& pnpm run build

build-docker:
	@docker build . -t michaelcoll/quiz-app:latest --pull --build-arg VERSION=v0.0.1

.PHONY: test
test:
	@go test -v ./...

.PHONY: vet
vet: ## check go code
	@go vet ./...

run:
	@go run . serve

run-vue:
	cd internal/web \
  		&& pnpm run dev

vue-lint:
	cd internal/web \
  		&& pnpm run lint

run-docker:
	docker run -ti --rm -p 8080:8080 web:latest

.PHONY: generate
generate:
	@go generate internal/back/domain/repositories.go
	@go generate internal/back/domain/callers.go
	@sqlc generate
	@sqlc-addon generate --quiet
