# Start by building the application.
FROM golang:1.20 as build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o /go/bin/quiz-app -ldflags="-s -w -X 'github.com/michaelcoll/quiz-app/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11:nonroot

COPY --from=build-go /go/bin/quiz-app /bin/quiz-app

EXPOSE 8080
EXPOSE 9000

CMD ["quiz-app", "serve"]
