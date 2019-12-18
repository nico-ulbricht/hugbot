FROM golang:1.13-alpine AS build

WORKDIR /build
RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/server


FROM alpine:3.10 AS final

WORKDIR /bin
RUN apk update && apk add ca-certificates
COPY --from=build /build/server .
COPY --from=build /build/migrations ./migrations

# TODO: limit runtime user permissions
CMD ./server
