FROM golang:latest

RUN go get github.com/githubnemo/CompileDaemon && \
  go get github.com/go-delve/delve/cmd/dlv  
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY ./ ./

EXPOSE 5000 2345

ENTRYPOINT CompileDaemon --build="echo" --command="dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient cmd/ced/main.go"