# Contributing to kleat

## Protocol buffers

### Overview

Kleat's data model is expressed as [protocol buffers](https://developers.google.com/protocol-buffers)
in the [/api/proto](api/proto) directory. For example, the [hal config proto](/api/proto/config/halconfig.proto)
represents the source of truth for what fields are valid in a hal config.

These protocol buffers are then transformed to generate:

- A [go library](/api/client) used by kleat to interact with its data model
- [Documentation](/docs/docs.md) of the data model

### Regenerating dependent code

In order to keep the go library and documentation in sync with the protocol buffers, any commit that
changes a protocol buffer must regenerate any dependent code by running the following and committing
the result:

```bash
make proto
```

This command requires that `docker` be installed, but has no other dependencies. The fist time it
is run, it may take a minute or two as it builds the container image it will use to compile the
protocol buffers; subsequent runs will be significantly faster.

By compiling protocol buffers in a container with explicitly-specified dependencies, we avoid
having this compilation depend on the versions of tools on a particular contributor's machine
(and in fact don't require contributors to have these tools at all, in exchange for requiring that
they have `docker`).

### Upgrading protocol buffer libraries

The versions of all tools used to compile the protocol buffers are specified in the
[Dockerfile](/build/protoc/Dockerfile). To update the version of one of these dependencies,
make the appropriate change to the Dockerfile, run `make proto`, and commit any changes.

### Adding a new field

Adding a new field to the [hal config](/api/proto/config/halconfig.proto)
involves three steps:

1. **Define and Document.** Define the field using protocol buffers.
[Add comments](https://developers.google.com/protocol-buffers/docs/proto3#adding-comments)
to document each new message and field. Ensure the field is included in the
[hal config](/api/proto/config/halconfig.proto) and each
consuming service's config.

2. **Translate.** If necessary, update the code that [translates](/pkg/transform/transform.go)
the hal config to each individual microservice's config. Add unit tests to the
appropriate service config tests in the `parse_hal` package.

3. **Test.** Update the end-to-end tests to exercise the inclusion of a
non-default value for your field by updating the [test hal config](/testdata/halconfig.yml)
and the appropriate test service configs.
