# Migrating from Halyard

A Halyard-generated halconfig requires several manual adjustments before it can
be used as input to Kleat.

## Required changes

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
- Using `redis` as a persistent store for `front50` is no longer supported. All
  other uses of `redis` continue to be supported.
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
  - Halyard requires the `serviceAddress` to be configured as a full URL (ex:
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
- In the `security.authn.oauth2` block, there is a field `provider`. This field
  can be removed as it was only used internally by Halyard.
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
- The `security.enabled` field is no longer used by `kleat`; the
  `security.authn.enabled` and `security.authz.enabled` fields individually
  control whether authentication and authorization are enabled, respectively.

## Optional changes

- Halyard provided a default value for every possible member of the halconfig.
  Kleat neither requires nor supplies any default values, which means it is
  acceptable to remove all values for which the microservice's deafult value is
  acceptable. For example, if you do not use the Google App Engine provider, you
  can remove the `Providers.AppengineProvider` block from your halconfig.
- Halyard supplies a default `providerVersion` of `v1` for each provider
  account. You can safely remove this field.
- Halyard set `metricStores.enabled` to true if one or more monitoring metric
  stores was enabled. This field is not read by spinnaker-monitoring and is safe
  to remove.

## Unsupported Halyard functionality

Since the only objective of Kleat is to expose a CLI for transforming a
top-level Spinnaker config file into individual Spinnaker microservice config
files, there are several `hal` commands with no `kleat` analog.

### Backing up your hal config

- `hal backup`, `hal deploy diff`, `hal deploy rollback`: Instead, version your
  config in source control.

### Customizing your Spinnaker microservice deployments

When using kleat, you can remove the `deploymentEnvironment` block of your hal
config. Kleat only handles Spinnaker application configuration and is intended
for use with the Spinnaker
[kustomization-base](https://github.com/spinnaker/kustomization-base) and
[spinnaker-config](https://github.com/spinnaker/spinnaker-config) projects. To
customize components of your Spinnaker microservice Kubernetes resource YAML
(e.g., custom liveness probes), update your kustomization to include overlays
for any non-default resource properties.

- `hal config deploy component-sizing`: Instead, specify non-default component
  sizing parameters (e.g., replicas, container CPU requests and limits) in your
  Kustomization.
- `hal config deploy ha`: The Spinnaker
  [kustomization-base](https://github.com/spinnaker/kustomization-base) includes
  sharded Clouddriver and Echo ("HA") by default. To deploy a non-sharded
  Clouddriver and Echo instead, update your kustomization.
- `hal config deploy edit`: Since Kleat will not be deploying Spinnaker, all
  aspects of the deployment (cluster, namespace, image variant) must be handled
  downstream in your kustomization or in the context in which you run
  `kubectl apply`.

### Documenting supported Spinnaker versions

- `hal version latest`, `hal version list`: Instead, use the
  [documentation](https://www.spinnaker.io/community/releases/versions/) for
  information about supported Spinnaker versions.

### Installing spin CLI

- `hal spin install`: Instead, follow these
  [instructions](https://www.spinnaker.io/setup/spin/#install-and-configure-spin-cli)
  to install spin CLI.
