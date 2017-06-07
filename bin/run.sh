#!/usr/bin/env bash

## brew install fswatch

fswatch -iE ".*\.[go]+" . | xargs -n1 -I{} sh -c 'go test ./... && git add . && git commit -m "commit $(date)" && echo "success \o/"'
