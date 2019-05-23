FROM golang:1.12.5-alpine3.9

RUN apk --no-cache add ca-certificates git
ENV TRIVY_VERSION=0.1.1 \
    TRIVY_CHECKSUM=c3d4b1e6b1ff9a93934a9992a8deccad6771cb54a52a17e31d7b019ab8e373e4
WORKDIR /tmp
RUN repo="github.com/knqyf263/trivy" \
    && base_url="https://${repo}/releases/download/v${TRIVY_VERSION}" \
    && wget "${base_url}/trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" \
    && echo "${TRIVY_CHECKSUM}  trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" | sha256sum -c - \
    && tar xvf "trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" \
    && mv trivy /go/bin/ \
    && rm -rf /tmp/*
WORKDIR /go
