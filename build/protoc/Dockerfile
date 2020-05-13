FROM alpine:3.11.6 as curl
RUN apk add --no-cache curl unzip

FROM curl as protoc-gen-go
ARG PROTOC_GEN_GO_TAG=v1.22.0
RUN curl -L https://github.com/protocolbuffers/protobuf-go/releases/download/$PROTOC_GEN_GO_TAG/protoc-gen-go.$PROTOC_GEN_GO_TAG.linux.amd64.tar.gz \
 | tar -xz

FROM curl as protoc-gen-doc
ARG PROTOC_GEN_DOC_TAG=1.3.1
RUN curl -L https://github.com/pseudomuto/protoc-gen-doc/releases/download/v$PROTOC_GEN_DOC_TAG/protoc-gen-doc-$PROTOC_GEN_DOC_TAG.linux-amd64.go1.12.6.tar.gz \
 | tar --strip-components=1 -xz

FROM curl as protoc
ARG PROTOC_VERSION=3.11.4
RUN curl -vLo protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
 && unzip protoc.zip

FROM debian:buster-slim
COPY --from=protoc /bin/protoc /usr/local/bin/protoc
COPY --from=protoc /include/google /usr/local/include/google
COPY --from=protoc-gen-go /protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=protoc-gen-doc /protoc-gen-doc /usr/local/bin/protoc-gen-doc
RUN mkdir -p /output /staging/go /staging/doc
COPY genproto.sh .
ENTRYPOINT ["./genproto.sh"]
