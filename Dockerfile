FROM golang:1.7 as builder
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN make build

FROM gcr.io/distroless/base
COPY --from=builder  /go/src/app/bin /opt/empresa/Service
WORKDIR /opt/empresa
VOLUME /opt/empresa/data
CMD ["/opt/empresa/Service"]
