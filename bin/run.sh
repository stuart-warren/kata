#!/usr/bin/env bash

## brew install fswatch
while true
do
  fswatch --one-event -e ".*" -i ".*/[^.]*\\.go$" . && \
  go test ./... && git add . && \
  git commit -m "commit $(date)" && \
  echo "success \o/"
done
