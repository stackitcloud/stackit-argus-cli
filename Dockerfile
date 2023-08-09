# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.21@sha256:fe54f55814913050d8d01cc6ffded67ae350594186b986fdba2ad02af58a6eb0 AS base
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
FROM gcr.io/distroless/static:nonroot@sha256:9ecc53c269509f63c69a266168e4a687c7eb8c0cfd753bd8bfcaa4f58a90876f as prd
COPY --from=build /app/stackit-argus-cli /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/stackit-argus-cli"]
