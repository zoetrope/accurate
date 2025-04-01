# Build the manager binary. This always executes in the native architecture of the building machine.
FROM --platform=$BUILDPLATFORM ghcr.io/cybozu/golang:1.23-jammy AS builder

COPY ./ .

# Build the binary, cross-compiling if necessary
ARG TARGETPLATFORM
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETPLATFORM \
	go build -ldflags="-w -s" -o accurate-controller ./cmd/accurate-controller

# the controller image, this is in the target architecture.
FROM --platform=$BUILDPLATFORM ghcr.io/cybozu/ubuntu-debug:24.04
LABEL org.opencontainers.image.source https://github.com/zoetrope/accurate

COPY --from=builder /work/accurate-controller ./
USER 10000:10000

ENTRYPOINT ["/usr/bin/bash"]
