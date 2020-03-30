protobuilder:
	docker build -q -t kleat-protobuilder:latest build/protoc/

proto_command = docker run --rm -v $(abspath api/proto/):/proto -v $(abspath api/client):/output/go -v $(abspath docs):/output/doc kleat-protobuilder:latest

proto: protobuilder
	$(proto_command) update

checkproto: protobuilder
	$(proto_command) check
