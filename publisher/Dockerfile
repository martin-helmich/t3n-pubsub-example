FROM golang:1.12 as builder

COPY . /work
WORKDIR /work/publisher
RUN useradd app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o publisher

FROM scratch

COPY --from=builder /work/publisher/publisher /usr/sbin/publisher
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/

USER app

ENTRYPOINT ["/usr/sbin/publisher"]