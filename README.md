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
- Halyard used the field `persistentStorage.persistentStoreType` to select which
  persistent store type is enabled. Kleat does not recognize this field and
  instead requires that the flag `enabled: true` be set on the persistent store
  that should be used.
- Halyard specified `canary.serviceIntegrations` as a list, with each
  integration specifying its type under a `name` key. Kleat requires
  `canary.serviceIntegrations` to be specified as a map with each type as a key.
  You are only required to specify an entry for each integration type you wish
  to configure. For example, if you are using GCS as your canary persistent
  store and Prometheus as your metric source, your `canary.serviceIntegrations`
  block is only required to have `google` and `prometheus` entries.
- There are two changes to SAML config in the `security.authn.saml` block in the
  halconfig:
  - Halyard has two separate fields for the metadata depending on whether it
    represents a local file (`metadataLocal`) or a remote file
    (`metadataRemote`). Kleat does not make this distinction and requires the
    metadata location to be specified under the `metadataUrl` field. Users
    should update their SAML config to remove whichever of `metadataLocal` or
    `metadataRemote` was specified, and instead set `metadataUrl` to this value.
    A path should be prefixed with `file://` (though currently Gate will add
    this prefix if it is missing for a local file).
  - Halyard requires the `redirectUrl` to be configured as a full URL (ex:
    `https://gate.org.com:8084/mypath`). Kleat will require that the redirect
    host name, path, and protocol be separately configured. Users will need to
    remove the `redirectUrl` field and instead specify the following fields:
    - `redirectProtocol` (ex: `https`)
    - `redirectHostname` (ex: `gate.org.com:8084`)
    - `redirectBasePath` (ex: `/mypath`)
- Halyard allowed basic username/password authentication to be configured for
  Gate by adding entries to the `~/.hal/default/service-settings/gate.yml` file.
  Kleat will move this configuration to be alongside the configuration of other
  authentication methods in the halconfig. Users who have configured basic auth
  for Gate in the above file will need to instead specify their basic auth
  settings in `security.authn.basic` in their halconfig.
- In the `security.authz.groupMembership` block, there is a field
  `roleProviderType` under each configurable role provider. This field can be
  removed as it was only used internally by Halyard and does not affect the
  generated config.
- The `[provider].bakeryDefaults.images` field in the `halconfig` allows a user
  to specify the images that will be available to use as base images. Halyard
  would merge any images specified here with the default images defined in the
  microservice, which is not how any of the other list entries behaved. Kleat
  will treat this field as all other fields; if this field is specified, its
  value will override the default images.
- Halyard only exposed commands to enable the Slack, Twilio, and GitHub Status
  notification types, but defaulted the remaining notification types (BearyChat,
  Email, Google Chat, and Pub/Sub) to enabled in Deck. Kleat requires you to
  explicitly enable BearyChat, Email, Google Chat, and Pub/Sub notification
  types.
- Kleat requires you to explicitly specify a primaryAccount for each enabled
  cloud provider for which you expect Deck to receive a statically configured
  default account and/or region. (Halyard will fall back to the first configured
  account if no primaryAccount is configured.)  

#### Optional Changes

- Halyard provided a default value for every possible member of the halconfig.
  Kleat neither requires nor supplies any default values, which means it is
  acceptable to remove all values for which the microservice's deafult value is
  acceptable. For example, if you do not use the Google App Engine provider, you
  can remove the `Providers.AppengineProvider` block from your halconfig.
