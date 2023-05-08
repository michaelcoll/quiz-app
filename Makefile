build: build-web build-go

prepare:
	cd internal/web \
	&& corepack enable && corepack prepare \
	&& pnpm i

dep-upgrade: dep-upgrade-go dep-upgrade-node

dep-upgrade-go:
	go get -u
	go mod tidy

dep-upgrade-node:
	cd internal/web \
  && pnpm update --latest

build-go:
	go build -v -ldflags="-s -w -X 'github.com/schoolbyhiit/quizz-app/cmd.version=v0.0.0'" .

build-web:
	cd internal/web \
	&& pnpm run build

build-docker:
	docker build . -t web --pull --build-arg VERSION=v0.0.1

.PHONY: test
test:
	go test -v ./...

run:
	go run . serve

run-vue:
	cd internal/web \
  && pnpm run dev

vue-lint:
	cd internal/web \
  && pnpm run lint

run-docker:
	docker run -ti --rm -p 8080:8080 web:latest

.PHONY: sqlc
sqlc:
	sqlc generate
