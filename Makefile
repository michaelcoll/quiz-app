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
	@go test -vet=all ./...

.PHONY: coverage
coverage:
	@go test -vet=all -covermode atomic -coverprofile=coverage.out ./...

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

gen: sqlc generate

.PHONY: generate
generate:
	@go generate internal/back/domain/repositories.go

.PHONY: sqlc
sqlc:
	@sqlc generate
	@sqlc-addon generate --quiet
