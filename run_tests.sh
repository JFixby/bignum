#!/usr/bin/env bash

GO=go
  $GO version
  $GO clean -testcache
  $GO build -v ./...
  $GO test -v ./...
