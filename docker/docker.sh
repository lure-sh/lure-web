#!/bin/bash

docker buildx build --platform linux/amd64 --tag arsen6331/lure-web:amd64 --no-cache .
docker buildx build --platform linux/arm64/v8 --tag arsen6331/lure-web:arm64 --no-cache .

docker login
docker push arsen6331/lure-web -a

docker manifest rm arsen6331/lure-web:latest
docker manifest create arsen6331/lure-web:latest --amend arsen6331/lure-web:arm64 --amend arsen6331/lure-web:amd64
docker manifest push arsen6331/lure-web:latest
