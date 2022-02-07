ARG GIT_COMMIT
ARG VERSION
ARG PROJECT

FROM golang:1.15.1 as builder
ARG GIT_COMMIT
ENV GIT_COMMIT=$GIT_COMMIT

ARG VERSION
ENV VERSION=$VERSION

ARG PROJECT
ENV PROJECT=$PROJECT

ENV GOSUMDB=off
ENV GO111MODULE=on
ENV WORKDIR=${GOPATH}/src/static-srv

COPY . ${WORKDIR}
WORKDIR ${WORKDIR}

RUN set -xe ;\
    go build -ldflags="-X '${PROJECT}/internal/version.Version=${VERSION}' -X '${PROJECT}/internal/version.Commit=${GIT_COMMIT}'"  -o /go/bin/static-srv ./cmd/static-srv/main.go ;\
    ls -lhtr /go/bin/

FROM golang:1.15.1

EXPOSE 8080

WORKDIR /go/bin

COPY --from=builder /go/bin/static-srv .
COPY --from=builder ${GOPATH}/src/static-srv/config/*.env ./config/

ENTRYPOINT ["/go/bin/static-srv"]