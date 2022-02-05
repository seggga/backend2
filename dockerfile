FROM golang:1.15 as modules

ADD go.mod go.sum /module/
RUN cd /module && go mod download


FROM golang:1.15 as builder
COPY --from=modules /go/pkg /go/pkg
RUN mkdir -p /src
ADD . /src
WORKDIR /src
RUN useradd -u 10001 simpleuser
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o /myapp ./

# Готовим пробный файл статики
# RUN mkdir -p /test_static && touch /test_static/index.html
# RUN echo "Hello, world!" > /test_static/index.html


FROM busybox
ENV PORT 8080
ENV STATICS_PATH /static
# COPY --from=builder /test_static /test_static
RUN mkdir -p /static
ADD ./static/index.html /static
COPY --from=builder /etc/passwd /etc/passwd
USER simpleuser
COPY --from=builder /myapp /myapp
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

CMD ["/myapp"]