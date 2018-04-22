FROM golang AS builder
WORKDIR /go/src/github.com/ffddorf/unms_exporter
COPY . .
RUN go get -d
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/unms_exporter

FROM alpine AS certs
RUN apk add --no-cache curl && curl -o /etc/ssl/certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/unms_exporter unms_exporter
CMD ["./unms_exporter"]
