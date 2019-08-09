FROM golang:1.12 as builder

COPY . /work
WORKDIR /work/subscriber
RUN useradd app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o subscriber

FROM scratch

COPY --from=builder /work/subscriber/subscriber /usr/sbin/subscriber
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/

USER app

ENTRYPOINT ["/usr/sbin/subscriber"]