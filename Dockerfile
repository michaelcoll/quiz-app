# Start by building the application.
FROM golang:1.20 as build-go

ARG BUILDTIME
ARG VERSION
ARG REVISION

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/quizz-app -ldflags="-s -w -X 'github.com/schoolbyhiit/quizz-app/cmd.version=$VERSION'"

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11:nonroot

COPY --from=build-go /go/bin/quizz-app /bin/quizz-app

EXPOSE 8080
EXPOSE 9000

CMD ["quizz-app", "serve"]
