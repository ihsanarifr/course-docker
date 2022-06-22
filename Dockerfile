# set base image for builder STAGE
FROM golang:latest as builder
# set working directory
WORKDIR /go/src/rest-api
# copy golang module dependencies
COPY go.mod go.mod
COPY go.sum go.sum
# download all golang module dependencies
RUN go mod download
# Copy semua file
COPY . .
# Build app
RUN go build -o /go/bin/rest-api cmd/api/main.go

# deployment STAGE
FROM golang:1.18-alpine as deployment

COPY --from=builder /go/bin/rest-api /go/bin/rest-api

# Entrypoint command
CMD ["/go/bin/rest-api"]
