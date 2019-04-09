# Build container
FROM golang:1.11 as builder

WORKDIR /go/src/uptime-checker

COPY . .

RUN \
  make setup && \
  make deps && \
  make lint

RUN make build-docker

# Release container
FROM busybox:1.29

WORKDIR /usr/local/bin

COPY --from=builder /go/src/uptime-checker/uptime-checker /usr/local/bin

ENTRYPOINT ["/usr/local/bin/uptime-checker"]
CMD [""]
