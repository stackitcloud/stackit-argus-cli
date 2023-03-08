# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.20@sha256:74a382917f6eaa7cc2d000dc2cd412a7f823f343b3b6268b20d84d057bc56718 AS base
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
FROM gcr.io/distroless/static:nonroot@sha256:775156d9ad597260d6e9cd0a8a400e56f2a860b845e54d8e48ce6da22f90240f as prd
COPY --from=build /app/stackit-argus-cli /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/stackit-argus-cli"]
