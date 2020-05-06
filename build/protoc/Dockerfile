FROM golang:1.14.1-alpine as builder
RUN apk add --no-cache git

FROM builder as protoc-gen-go
ARG PROTOC_GEN_GO_TAG=v1.22.0
RUN go get -d google.golang.org/protobuf/cmd/protoc-gen-go \
 && git -C "$(go env GOPATH)"/src/google.golang.org/protobuf checkout $PROTOC_GEN_GO_TAG \
 && go install google.golang.org/protobuf/cmd/protoc-gen-go

FROM builder as protoc-gen-doc
ARG PROTOC_GEN_DOC_TAG=v1.3.1
RUN go get -d github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
 && git -C "$(go env GOPATH)"/src/github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc checkout $PROTOC_GEN_DOC_TAG \
 && go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

FROM alpine:3.11.6
RUN apk add protoc
COPY --from=protoc-gen-go /go/bin/protoc-gen-go /bin/
COPY --from=protoc-gen-doc /go/bin/protoc-gen-doc /bin/
RUN mkdir -p /output /staging/go /staging/doc
COPY genproto.sh .
ENTRYPOINT ["./genproto.sh"]
