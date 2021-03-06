#!/usr/bin/env bash

log() {
  echo "STARTSCRIPT: $1"
}

buildServer() {
  log "Building server binary"
  go build -gcflags "all=-N -l" -o /server ./cmd/ced/main.go
}

runServer() {
  log "Run server"

  log "Killing old server"
  killall dlv
  killall server
  log "Run in debug mode"
  #dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2 --accept-multiclient exec /server
  dlv -l 0.0.0.0:2345 --headless --api-version=2 --accept-multiclient exec /server &
}

rerunServer() {
  log "Rerun server"
  buildServer
  runServer
}

liveReloading() {
  log "Run liveReloading"
  inotifywait -e "MODIFY,DELETE,MOVED_TO,MOVED_FROM" -m -r --include '.go$' /go/src/github.com/lilahamstern/ced/server/ | (
    # read changes from inotify, batch results to a second (read -t 1)
    while true; do
      read path action file
      echo "Test"
      ext=${file: -3}
      if [[ "$ext" == ".go" ]]; then
        echo "$file"
      fi
    done
  ) | (
    WAITING=""
    while true; do
      file=""
      read -t 1 file
      if test -z "$file"; then
        if test ! -z "$WAITING"; then
          echo "CHANGED"
          WAITING=""
        fi
      else
        log "File ${file} changed" >>/tmp/filechanges.log
        WAITING=1
      fi
    done
  ) | (
    # read statement release when some file has been changed
    while true; do
      read TMP
      log "File Changed. Reloading..."
      rerunServer
    done
  )
}

initializeFileChangeLogger() {
  echo "" > /tmp/filechanges.log
  tail -f /tmp/filechanges.log &
}

main() {
  initializeFileChangeLogger
  buildServer
  runServer
  liveReloading
}

main