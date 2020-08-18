FROM golang:latest

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY ./ ./

EXPOSE 5000
ENTRYPOINT CompileDaemon --build="go build cmd/ced/main.go" --command=./main