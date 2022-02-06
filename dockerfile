# copy dependencies
FROM golang:1.15 as modules
ADD go.mod go.sum /module/
RUN cd /module && go mod download

# create user and build a binary
FROM golang:1.15 as builder
COPY --from=modules /go/pkg /go/pkg
RUN mkdir -p /src
ADD . /src
WORKDIR /src
RUN useradd -u 10001 simpleuser
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o /myapp ./

# set environments, copy user
FROM busybox
ENV PORT 8080
ENV STATICS_PATH /static
RUN mkdir -p /static
ADD ./static/index.html /static
COPY --from=builder /etc/passwd /etc/passwd
USER simpleuser
COPY --from=builder /myapp /myapp
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

# start application
CMD ["/myapp"]