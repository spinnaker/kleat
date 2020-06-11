#!/bin/sh
set -e

gen_proto() {
  output_dir=$1
  find /proto -name *.proto | sort | xargs protoc \
    --proto_path=/proto \
    --go_out=paths=source_relative:$1/go \
    --doc_out=$1/doc \
    --doc_opt=markdown,docs.md
}

update_proto() {
  gen_proto /staging
  rm -rf /output/go/* /output/doc/*
  cp -r /staging/doc/* /output/doc/
  cp -r /staging/go/* /output/go/
}

check_proto() {
  gen_proto /staging
  if ! diff -r /staging /output; then
    echo "Protocol buffer compilation out of date. Please run 'make proto' and commit the changes."
    echo "To help with debugging, the difference between the committed and generated code is printed above."
    exit 1
  fi
}

command=$1
case $command in
  update)
    update_proto
    ;;
  check)
    check_proto
    ;;
  *)
    echo "Unrecognized command: $command"
    exit 2
    ;;
esac
