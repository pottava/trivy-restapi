FROM golang:1.12.5-alpine3.9 AS build-app
RUN apk --no-cache add build-base git wget
RUN go get -u github.com/golang/dep/...
WORKDIR /go/src/github.com/golang/dep/
RUN git checkout v0.5.0
RUN go install github.com/golang/dep/...
WORKDIR /usr/bin/
RUN wget https://github.com/go-swagger/go-swagger/releases/download/v0.17.2/swagger_linux_amd64
RUN chmod +x swagger_linux_amd64
RUN go get github.com/mitchellh/gox
WORKDIR /go/src/github.com/pottava/
ENV APP_VERSION=0.0.16
# RUN git clone --depth=1 -b ${APP_VERSION} https://github.com/pottava/trivy-restapi.git
RUN git clone --depth=1 https://github.com/pottava/trivy-restapi.git
WORKDIR /go/src/github.com/pottava/trivy-restapi
RUN swagger_linux_amd64 generate server -f spec.yaml -t app/generated
WORKDIR /go/src/github.com/pottava/trivy-restapi/app
RUN dep ensure
WORKDIR /go/src/github.com/pottava/trivy-restapi
RUN githash=$( git rev-parse --short HEAD 2>/dev/null ) \
    && cd app/generated/cmd/trivy-restapi-server/ \
    && gox --osarch "linux/amd64" -ldflags "-s -w -X github.com/pottava/trivy-restapi/app/lib.date=$(date +%Y-%m-%d) -X github.com/pottava/trivy-restapi/app/lib.version=${APP_VERSION} -X github.com/pottava/trivy-restapi/app/lib.commit=${githash}" -output "/trivy-restapi"

FROM alpine:3.9 AS trivy-bin
RUN apk --no-cache add wget
ENV TRIVY_VERSION=0.0.16 \
    TRIVY_CHECKSUM=3cb022d850256e477c8d131527900e8d74d251a258852872ffa6949a38f126dc
RUN repo="github.com/knqyf263/trivy" \
    && base_url="https://${repo}/releases/download/v${TRIVY_VERSION}" \
    && wget "${base_url}/trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz"
RUN echo "${TRIVY_CHECKSUM}  trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz" | sha256sum -c -
RUN tar xvf "trivy_${TRIVY_VERSION}_Linux-ARM64.tar.gz"
RUN mv trivy /

FROM alpine:3.9
RUN apk --no-cache add ca-certificates jq
COPY --from=trivy-bin /trivy /usr/bin/
COPY --from=build-app /trivy-restapi /
ENTRYPOINT ["/trivy-restapi", "--host", "0.0.0.0", "--port", "9000"]
