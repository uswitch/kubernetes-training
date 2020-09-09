FROM golang:1.14 AS build-env
WORKDIR /
COPY main.go .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o app .

FROM scratch
COPY --from=build-env /app .
ENTRYPOINT ["/app"]
