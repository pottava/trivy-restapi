# Trivy v0.1
# docker run -d --rm -p 9000:9000 -v "${HOME}/Library/Caches":/root/.cache pottava/trivy:0.1

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
ENV APP_VERSION=0.1.1
RUN git clone --depth=1 -b ${APP_VERSION} https://github.com/pottava/trivy-restapi.git
WORKDIR /go/src/github.com/pottava/trivy-restapi
RUN swagger_linux_amd64 generate server -f spec.yaml -t app/generated
WORKDIR /go/src/github.com/pottava/trivy-restapi/app
RUN dep ensure
WORKDIR /go/src/github.com/pottava/trivy-restapi
RUN githash=$( git rev-parse --short HEAD 2>/dev/null ) \
    && cd app/generated/cmd/trivy-restapi-server/ \
    && gox -ldflags "-s -w -X github.com/pottava/trivy-restapi/app/lib.date=$(date +%Y-%m-%d) -X github.com/pottava/trivy-restapi/app/lib.version=${APP_VERSION} -X github.com/pottava/trivy-restapi/app/lib.commit=${githash}" \
        --osarch "linux/amd64" -output "/trivy-restapi"

FROM golang:1.12.5-alpine3.9 AS trivy-bin
RUN apk --no-cache add build-base git wget
ENV APP_VERSION=v0.1.2
RUN git clone --depth=1 -b ${APP_VERSION} https://github.com/knqyf263/trivy.git
WORKDIR /go/trivy
RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" cmd/trivy/main.go
RUN mv /go/trivy/main /trivy

FROM alpine:3.9.4
RUN apk --no-cache add ca-certificates git \
    && rm -rf /lib/apk/db
ENV CACHE_DIRECTORY=/root/.cache \
    ACCESS_LOG=true
COPY --from=trivy-bin /trivy /usr/local/bin/
COPY --from=build-app /trivy-restapi /
RUN chmod +x /usr/local/bin/trivy /trivy-restapi
ENTRYPOINT ["/trivy-restapi", "--host", "0.0.0.0", "--port", "9000"]
