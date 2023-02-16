# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.20@sha256:2edf6aab2d57644f3fe7407132a0d1770846867465a39c2083770cf62734b05d AS base
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
FROM gcr.io/distroless/static:nonroot@sha256:9ec950c09380320e203369982691eb821df6a6974edf9f4bb8e661d4b77b9d99 as prd
COPY --from=build /app/stackit-argus-cli /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/stackit-argus-cli"]
