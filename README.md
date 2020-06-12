# kleat

![Kleat CI](https://github.com/spinnaker/kleat/workflows/Kleat%20CI/badge.svg)

Kleat is tool for managing Spinnaker configuration files that is intended to
replace Halyard. Please see the
[RFC](https://github.com/spinnaker/governance/blob/master/rfc/halyard-lite.md)
for an overview of what this tool will do.

It is currently under active development, but is ready for early adopters
to try in development clusters! Please share any feedback in the #kleat
channel in [Spinnaker slack](join.spinnaker.io).

### Migrating from Halyard

Please see the [migration README](/migration/README.md) for details on migrating
from Halyard.

### Installation

Download the appropriate binary from the
[releases page](https://github.com/spinnaker/kleat/releases/latest).

### Usage

- `kleat generate /path/to/halconfig /path/to/output/directory`
- `kleat help`
- `kleat version`
