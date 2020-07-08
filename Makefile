protobuilder:
	docker build -q -t kleat-protobuilder:latest build/protoc/

proto_command = docker run \
  --rm \
  --mount type=bind,source=$(abspath api/proto/),target=/mnt/proto,readonly \
  --mount type=bind,source=$(abspath api/client/),target=/mnt/output/go \
  --mount type=bind,source=$(abspath docs),target=/mnt/output/doc \
  kleat-protobuilder:latest

proto: protobuilder
	$(proto_command) update

checkproto: protobuilder
	$(proto_command) check
