# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.20@sha256:403f48633fb5ebd49f9a2b6ad6719f912df23dae44974a0c9445be331e72ff5e AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# build th application
FROM base AS build
# temp mount all files instead of loading into image with COPY
# temp mount module cache
# temp mount go build cache
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-w -s" -o /app/stackit-argus-cli ./cmd/stackit-argus-cli/*.go

# Import the binary from build stage
FROM gcr.io/distroless/static:nonroot@sha256:149531e38c7e4554d4a6725d7d70593ef9f9881358809463800669ac89f3b0ec as prd
COPY --from=build /app/stackit-argus-cli /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/stackit-argus-cli"]
