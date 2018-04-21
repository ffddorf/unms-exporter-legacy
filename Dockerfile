FROM golang AS builder
WORKDIR /go/src/github.com/ffddorf/unms_exporter
COPY . .
RUN go get -d
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/unms_exporter

FROM scratch
COPY --from=builder /go/bin/unms_exporter unms_exporter
CMD ["./unms_exporter"]
