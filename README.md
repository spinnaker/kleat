# kleat

![Kleat CI](https://github.com/spinnaker/kleat/workflows/Kleat%20CI/badge.svg)

Kleat is tool for managing Spinnaker configuration files that is intended to
replace Halyard. Please see the
[RFC](https://github.com/spinnaker/governance/blob/master/rfc/halyard-lite.md)
for an overview of what this tool will do.

It is currently under active development and is not ready for production (or
even development) use, but stay tuned for updates!

### Migrating from Halyard

A Halyard-generated halconfig requires several manual adjustments before it can
be used as input to Kleat.

#### Required Changes

- A Halyard-generated halconfig contains two top-level fields:
  `deploymentConfigurations` (a list of `HalConfig`s) and `currentDeployment`,
  (the name of the `deploymentConfiguration` to manage). Kleat accepts a file
  with a single `HalConfig` as input, so any halconfigs with multiple
  `deploymentConfigurations` need to be separated into different files.
- Halyard specified
  `HalConfig.Providers.AzureProvider.AzureAccount.useSshPublicKey` as a `string`
  with `'true'` and `'false'` as acceptable values. Kleat requires this field to
  be specified as a `bool`.
- Halyard required Github status notifications to be configured under
  `notifications.github-status`. To match the casing of other fields, `kleat`
  requires that it be under `notifications.githubStatus`.

#### Optional Changes

- Halyard provided a default value for every possible member of the halconfig.
  Kleat neither requires nor supplies any default values, which means it is
  acceptable to remove all values for which the microservice's deafult value is
  acceptable. For example, if you do not use the Google App Engine provider, you
  can remove the `Providers.AppengineProvider` block from your halconfig.
