# Contributing to kleat

## Protocol buffers

### Overview

Kleat's data model is expressed as
[protocol buffers](https://developers.google.com/protocol-buffers) in the
[/api/proto](api/proto) directory. For example, the
[hal config proto](/api/proto/config/halconfig.proto) represents the source of
truth for what fields are valid in a hal config.

These protocol buffers are then transformed to generate:

- A [go library](/api/client) used by kleat to interact with its data model
- [Documentation](/docs/docs.md) of the data model

### Regenerating dependent code

In order to keep the go library and documentation in sync with the protocol
buffers, any commit that changes a protocol buffer must regenerate any dependent
code by running the following and committing the result:

```bash
make proto
```

This command requires that `docker` be installed, but has no other dependencies.
The fist time it is run, it may take a minute or two as it builds the container
image it will use to compile the protocol buffers; subsequent runs will be
significantly faster.

By compiling protocol buffers in a container with explicitly-specified
dependencies, we avoid having this compilation depend on the versions of tools
on a particular contributor's machine (and in fact don't require contributors to
have these tools at all, in exchange for requiring that they have `docker`).

### Upgrading protocol buffer libraries

The versions of all tools used to compile the protocol buffers are specified in
the [Dockerfile](/build/protoc/Dockerfile). To update the version of one of
these dependencies, make the appropriate change to the Dockerfile, run
`make proto`, and commit any changes.

### Adding a new field

Adding a new field to the [hal config](/api/proto/config/halconfig.proto)
involves three steps:

1. **Define and Document.** Define the field using protocol buffers.
   [Add comments](https://developers.google.com/protocol-buffers/docs/proto3#adding-comments)
   to document each new message and field. Ensure the field is included in the
   [hal config](/api/proto/config/halconfig.proto) and each consuming service's
   config.

2. **Translate.** If necessary, update the code that
   [translates](/pkg/transform/transform.go) the hal config to each individual
   microservice's config. Add unit tests to the appropriate service config tests
   in the `parse_hal` package.

3. **Test.** Update the end-to-end tests to exercise the inclusion of a
   non-default value for your field by updating the
   [test hal config](/testdata/halconfig.yml) and the appropriate test service
   configs.

### Choosing field type

#### Overview

We are using proto3 syntax, which by default does not distinguish between fields
that are omitted and fields that are set to their
[default value](https://developers.google.com/protocol-buffers/docs/proto3#default_values).
If a field is set to its default value (`""`, `0`, `false`), it will be omitted
from the generated output (just as if the field were not set at all).

In some cases, it is important to distinguish between a field that has been set
to its default value and a field that has not been set at all.

For example, suppose we have a feature that defaults to enabled, but allow it to
be disabled by setting `isEnabled: false`. In this case `kleat` _must_
distinguish between users who have not set the flag (and should omit including
it in the output) and users who have explicitly set the flag to `false` (and
should include this in the output).

These cases should be handled by using the appropriate wrapper type from
[google.protobuf](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf):

- `BoolValue`
- `StringValue`
- `Int32Value`

When using one of these types, it is possible to distinguish between a field
that has been omitted from the config and a field that has has been explicitly
set to its default value. The output config will include the field even when it
has been set to its default value (`""`, `0`, `false`).

#### General rules

The general rule is that `string` and `int32` fields should rarely be wrapped,
as in general the zero value is not a valid value for the field with any meaning
distinct from omitting the field.

On the other hand, it generally is meaningful to distinguish between omitting a
`bool` field and setting it to `false`. This is because explicitly setting
something to `false` is meaningful and distinct from not setting it at all (even
if the consuming code currently defaults the field to `false`). For this reason,
all `bool` fields should be wrapped as `BoolValue`.
