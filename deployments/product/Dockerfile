FROM golang:1.21-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /product

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/product ./cmd/productServer/main.go

# Start fresh from a smaller image
FROM alpine:3.9

COPY --from=build_base /product/out/product /app/product

COPY --from=build_base /product/.env /.env

# This container exposes port 8080 to the outside world
EXPOSE 5001

# Run the binary program produced by `go install`
CMD ["/app/product"]