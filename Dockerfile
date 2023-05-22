# Start by building the application.
FROM node:18 as build-node

WORKDIR /node/src
COPY . .

RUN cd internal/web \
      && corepack enable && corepack prepare \
      && pnpm i \
      && pnpm run build

FROM golang:1 as build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .
COPY --from=build-node /node/src/internal/web/dist /go/src/app/internal/web/dist

RUN go mod download
RUN go build -o /go/bin/quiz-app -ldflags="-s -w -X 'github.com/michaelcoll/quiz-app/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11:nonroot

COPY --from=build-go /go/bin/quiz-app /bin/quiz-app

EXPOSE 8080

CMD ["quiz-app", "serve"]
