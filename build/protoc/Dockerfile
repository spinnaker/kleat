FROM alpine:3.11.6 as curl
RUN apk add --no-cache curl unzip

FROM curl as protoc-gen-go
ARG PROTOC_GEN_GO_TAG=v1.23.0
RUN curl -L https://github.com/protocolbuffers/protobuf-go/releases/download/$PROTOC_GEN_GO_TAG/protoc-gen-go.$PROTOC_GEN_GO_TAG.linux.amd64.tar.gz \
 | tar -xz

FROM golang:1.14.4-buster as protoc-gen-doc
WORKDIR /
ARG PROTOC_GEN_DOC_TAG=v1.3.2-proto3optional.1
RUN git init protoc-gen-doc \
 && cd protoc-gen-doc \
 && git remote add origin https://github.com/ezimanyi/protoc-gen-doc \
 && git fetch --no-tags --prune --depth=1 origin $PROTOC_GEN_DOC_TAG \
 && git checkout FETCH_HEAD \
 && go install ./cmd/protoc-gen-doc

FROM curl as protoc
ARG PROTOC_VERSION=3.12.3
RUN curl -vLo protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
 && unzip protoc.zip

FROM debian:buster-slim
COPY --from=protoc /bin/protoc /usr/local/bin/protoc
COPY --from=protoc /include/google /usr/local/include/google
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=protoc-gen-doc /go/bin/protoc-gen-doc /usr/local/bin/protoc-gen-doc
RUN mkdir -p /output /staging/go /staging/doc
COPY genproto.sh .
ENTRYPOINT ["./genproto.sh"]
