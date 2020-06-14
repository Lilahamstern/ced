# Base image
FROM golang:1.14.4-alpine as build-env

# Install git
RUN apk add --no-cache  git

# Env
# Vendor support to speed up build proccess
ENV GOFLAGS="-mod=vendor"

# Build Delve
RUN go get github.com/go-delve/delve/cmd/dlv

# Set workdir
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy source
COPY . .

# Using vendor so no need to download dependencies
# Install dependencies
#RUN go mod download

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -gcflags="all=-N -l" -o main .


# New stage
FROM alpine:latest

# Args
ARG app_env=development
ARG port=8080
# Env
ENV APP_ENV $app_env
ENV PORT $port

# Support file system access
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache coreutils

# Expose ports
EXPOSE ${port} 40000

# Set workdir
WORKDIR /app/

# Copy binary file from build-env
COPY --from=build-env /go/bin/dlv .
COPY --from=build-env /app/main .

# Execute app

#CMD ["./main"]
CMD ["./dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./main"]
