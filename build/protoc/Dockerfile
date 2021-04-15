FROM alpine:3.13.5 as curl
RUN apk add --no-cache curl unzip
RUN adduser --disabled-password --gecos '' curl
USER curl
WORKDIR /home/curl

FROM curl as protoc-gen-go
ARG PROTOC_GEN_GO_TAG=v1.25.0
RUN curl -L https://github.com/protocolbuffers/protobuf-go/releases/download/$PROTOC_GEN_GO_TAG/protoc-gen-go.$PROTOC_GEN_GO_TAG.linux.amd64.tar.gz \
 | tar -xz \
 && chmod 755 protoc-gen-go

FROM curl as protoc-gen-doc
ARG PROTOC_GEN_DOC_TAG=1.3.2
RUN curl -L https://github.com/pseudomuto/protoc-gen-doc/releases/download/v$PROTOC_GEN_DOC_TAG/protoc-gen-doc-$PROTOC_GEN_DOC_TAG.linux-amd64.go1.12.6.tar.gz \
 | tar --strip-components=1 -xz \
 && chmod 755 protoc-gen-doc

FROM curl as protoc
ARG PROTOC_VERSION=3.14.0
RUN curl -vLo protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
 && unzip protoc.zip \
 # We want all files/folders in bin/ to be executable
 && chmod -R 755 bin/ \
 # Files under include/ should not be executable so we need
 # to run a separate command for files/folders
 && find include/ -type d -exec chmod 755 {} \+ \
 && find include/ -type f -exec chmod 644 {} \+

FROM debian:buster-slim
COPY --chown=root:root --from=protoc /home/curl/bin/protoc /usr/local/bin/protoc
COPY --chown=root:root --from=protoc /home/curl/include/google /usr/local/include/google
COPY --chown=root:root --from=protoc-gen-go /home/curl/protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --chown=root:root --from=protoc-gen-doc /home/curl/protoc-gen-doc /usr/local/bin/protoc-gen-doc
COPY --chown=root:root genproto.sh /usr/local/bin/genproto.sh
RUN chmod 755 /usr/local/bin/genproto.sh
RUN adduser --disabled-password --gecos '' protoc
USER protoc
ENTRYPOINT ["/usr/local/bin/genproto.sh"]
