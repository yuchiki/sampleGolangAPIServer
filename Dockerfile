FROM golang:alpine AS build-stage
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd ./
RUN go build -o entrypoint ./...


FROM alpine:latest
COPY --from=build-stage /app/entrypoint /usr/local/bin/entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint"]
