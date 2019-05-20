FROM golang:1.12.5-alpine3.9

RUN apk --no-cache add ca-certificates jq
ENV TRIVY_VERSION=0.0.16 \
    TRIVY_CHECKSUM=3cb022d850256e477c8d131527900e8d74d251a258852872ffa6949a38f126dc
WORKDIR /tmp
RUN repo="github.com/knqyf263/trivy" \
    && base_url="https://${repo}/releases/download/v${TRIVY_VERSION}" \
    && wget "${base_url}/trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" \
    && echo "${TRIVY_CHECKSUM}  trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" | sha256sum -c - \
    && tar xvf "trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" \
    && mv trivy /go/bin/ \
    && rm -rf /tmp/*
WORKDIR /go
