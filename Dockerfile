# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.20@sha256:6e835db45c7d88e12b057c0638814c2b266f69143437e4110b8bec5cfc7fa53b AS base
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
FROM gcr.io/distroless/static:nonroot@sha256:5c9728ecd8683960ffff84837b21ac9f8f1fabbe437221a486fb93ccc4e05d88 as prd
COPY --from=build /app/stackit-argus-cli /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/stackit-argus-cli"]
