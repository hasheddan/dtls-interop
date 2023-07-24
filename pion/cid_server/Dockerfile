# golang:1.20.5
FROM golang@sha256:6b3fa4b908676231b50acbbc00e84d8cee9c6ce072b1175c0ff352c57d8a612f as build

WORKDIR /build

COPY . ./

RUN go mod download
RUN CGO_ENABLED=0 go build -o server main.go

# gcr.io/distroless/static:nonroot
FROM gcr.io/distroless/static@sha256:9ecc53c269509f63c69a266168e4a687c7eb8c0cfd753bd8bfcaa4f58a90876f

COPY --from=build /build/server /usr/local/bin/server
ENTRYPOINT ["server"]
