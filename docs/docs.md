# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [artifact/artifacts.proto](#artifact/artifacts.proto)
    - [Artifacts](#proto.artifact.Artifacts)
  
- [artifact/bitbucket.proto](#artifact/bitbucket.proto)
    - [Bitbucket](#proto.artifact.Bitbucket)
    - [BitbucketAccount](#proto.artifact.BitbucketAccount)
  
- [artifact/gcs.proto](#artifact/gcs.proto)
    - [Gcs](#proto.artifact.Gcs)
    - [GcsAccount](#proto.artifact.GcsAccount)
  
- [artifact/github.proto](#artifact/github.proto)
    - [GitHub](#proto.artifact.GitHub)
    - [GitHubAccount](#proto.artifact.GitHubAccount)
  
- [artifact/gitlab.proto](#artifact/gitlab.proto)
    - [GitLab](#proto.artifact.GitLab)
    - [GitLabAccount](#proto.artifact.GitLabAccount)
  
- [artifact/gitrepo.proto](#artifact/gitrepo.proto)
    - [GitRepo](#proto.artifact.GitRepo)
    - [GitRepoAccount](#proto.artifact.GitRepoAccount)
  
- [artifact/helm.proto](#artifact/helm.proto)
    - [Helm](#proto.artifact.Helm)
    - [HelmAccount](#proto.artifact.HelmAccount)
  
- [artifact/http.proto](#artifact/http.proto)
    - [Http](#proto.artifact.Http)
    - [HttpAccount](#proto.artifact.HttpAccount)
  
- [artifact/maven.proto](#artifact/maven.proto)
    - [Maven](#proto.artifact.Maven)
    - [MavenAccount](#proto.artifact.MavenAccount)
  
- [artifact/oracle.proto](#artifact/oracle.proto)
    - [Oracle](#proto.artifact.Oracle)
    - [OracleAccount](#proto.artifact.OracleAccount)
  
- [artifact/s3.proto](#artifact/s3.proto)
    - [S3](#proto.artifact.S3)
    - [S3Account](#proto.artifact.S3Account)
  
- [artifact/template.proto](#artifact/template.proto)
    - [Template](#proto.artifact.Template)
  
- [canary/aws.proto](#canary/aws.proto)
    - [Aws](#proto.canary.Aws)
    - [AwsAccount](#proto.canary.AwsAccount)
  
- [canary/canary.proto](#canary/canary.proto)
    - [Canary](#proto.canary.Canary)
    - [Canary.ServiceIntegrations](#proto.canary.Canary.ServiceIntegrations)
  
- [canary/datadog.proto](#canary/datadog.proto)
    - [Datadog](#proto.canary.Datadog)
    - [DatadogAccount](#proto.canary.DatadogAccount)
    - [DatadogAccount.Endpoint](#proto.canary.DatadogAccount.Endpoint)
  
- [canary/gcs.proto](#canary/gcs.proto)
    - [Gcs](#proto.canary.Gcs)
  
- [canary/google.proto](#canary/google.proto)
    - [Google](#proto.canary.Google)
    - [GoogleAccount](#proto.canary.GoogleAccount)
  
- [canary/newrelic.proto](#canary/newrelic.proto)
    - [NewRelic](#proto.canary.NewRelic)
    - [NewRelicAccount](#proto.canary.NewRelicAccount)
    - [NewRelicAccount.Endpoint](#proto.canary.NewRelicAccount.Endpoint)
  
- [canary/prometheus.proto](#canary/prometheus.proto)
    - [Prometheus](#proto.canary.Prometheus)
    - [PrometheusAccount](#proto.canary.PrometheusAccount)
    - [PrometheusAccount.Endpoint](#proto.canary.PrometheusAccount.Endpoint)
  
- [canary/s3.proto](#canary/s3.proto)
    - [S3](#proto.canary.S3)
  
- [canary/signalfx.proto](#canary/signalfx.proto)
    - [SignalFx](#proto.canary.SignalFx)
    - [SignalFxAccount](#proto.canary.SignalFxAccount)
    - [SignalFxAccount.Endpoint](#proto.canary.SignalFxAccount.Endpoint)
  
- [canary/stackdriver.proto](#canary/stackdriver.proto)
    - [Stackdriver](#proto.canary.Stackdriver)
  
- [canary/supported_type.proto](#canary/supported_type.proto)
    - [SupportedType](#proto.canary.SupportedType)
  
- [ci/ci.proto](#ci/ci.proto)
    - [Ci](#proto.ci.Ci)
  
- [ci/codebuild.proto](#ci/codebuild.proto)
    - [CodeBuild](#proto.ci.CodeBuild)
    - [CodeBuildAccount](#proto.ci.CodeBuildAccount)
  
- [ci/concourse.proto](#ci/concourse.proto)
    - [Concourse](#proto.ci.Concourse)
    - [ConcourseAccount](#proto.ci.ConcourseAccount)
  
- [ci/gcb.proto](#ci/gcb.proto)
    - [GoogleCloudBuild](#proto.ci.GoogleCloudBuild)
    - [GoogleCloudBuildAccount](#proto.ci.GoogleCloudBuildAccount)
  
- [ci/jenkins.proto](#ci/jenkins.proto)
    - [Jenkins](#proto.ci.Jenkins)
    - [JenkinsAccount](#proto.ci.JenkinsAccount)
  
- [ci/travis.proto](#ci/travis.proto)
    - [Travis](#proto.ci.Travis)
    - [TravisAccount](#proto.ci.TravisAccount)
  
- [ci/wercker.proto](#ci/wercker.proto)
    - [Wercker](#proto.ci.Wercker)
    - [WerckerAccount](#proto.ci.WerckerAccount)
  
- [cloudprovider/appengine.proto](#cloudprovider/appengine.proto)
    - [Appengine](#proto.cloudprovider.Appengine)
    - [AppengineAccount](#proto.cloudprovider.AppengineAccount)
  
    - [GcloudReleaseTrack](#proto.cloudprovider.GcloudReleaseTrack)
  
- [cloudprovider/aws.proto](#cloudprovider/aws.proto)
    - [Aws](#proto.cloudprovider.Aws)
    - [AwsAccount](#proto.cloudprovider.AwsAccount)
    - [AwsBakeryDefaults](#proto.cloudprovider.AwsBakeryDefaults)
    - [AwsBaseImage](#proto.cloudprovider.AwsBaseImage)
    - [AwsBaseImageSettings](#proto.cloudprovider.AwsBaseImageSettings)
    - [AwsFeatures](#proto.cloudprovider.AwsFeatures)
    - [AwsFeatures.CloudFormation](#proto.cloudprovider.AwsFeatures.CloudFormation)
    - [AwsLifecycleHook](#proto.cloudprovider.AwsLifecycleHook)
    - [AwsRegion](#proto.cloudprovider.AwsRegion)
    - [AwsVirtualizationSettings](#proto.cloudprovider.AwsVirtualizationSettings)
  
- [cloudprovider/azure.proto](#cloudprovider/azure.proto)
    - [Azure](#proto.cloudprovider.Azure)
    - [AzureAccount](#proto.cloudprovider.AzureAccount)
    - [AzureBakeryDefaults](#proto.cloudprovider.AzureBakeryDefaults)
    - [AzureBaseImage](#proto.cloudprovider.AzureBaseImage)
    - [AzureBaseImageSettings](#proto.cloudprovider.AzureBaseImageSettings)
  
- [cloudprovider/cloudfoundry.proto](#cloudprovider/cloudfoundry.proto)
    - [CloudFoundry](#proto.cloudprovider.CloudFoundry)
    - [CloudFoundryAccount](#proto.cloudprovider.CloudFoundryAccount)
  
- [cloudprovider/dcos.proto](#cloudprovider/dcos.proto)
    - [Dcos](#proto.cloudprovider.Dcos)
    - [DcosAccount](#proto.cloudprovider.DcosAccount)
    - [DcosAccountCluster](#proto.cloudprovider.DcosAccountCluster)
    - [DcosAccountDockerRegistry](#proto.cloudprovider.DcosAccountDockerRegistry)
    - [DcosCluster](#proto.cloudprovider.DcosCluster)
    - [DcosClusterLoadBalancer](#proto.cloudprovider.DcosClusterLoadBalancer)
  
- [cloudprovider/docker_registry.proto](#cloudprovider/docker_registry.proto)
    - [DockerRegistry](#proto.cloudprovider.DockerRegistry)
    - [DockerRegistryAccount](#proto.cloudprovider.DockerRegistryAccount)
  
- [cloudprovider/ecs.proto](#cloudprovider/ecs.proto)
    - [Ecs](#proto.cloudprovider.Ecs)
    - [EcsAccount](#proto.cloudprovider.EcsAccount)
  
- [cloudprovider/google.proto](#cloudprovider/google.proto)
    - [Consul](#proto.cloudprovider.Consul)
    - [GoogleBakeryDefaults](#proto.cloudprovider.GoogleBakeryDefaults)
    - [GoogleBaseImage](#proto.cloudprovider.GoogleBaseImage)
    - [GoogleBaseImageSettings](#proto.cloudprovider.GoogleBaseImageSettings)
    - [GoogleComputeEngine](#proto.cloudprovider.GoogleComputeEngine)
    - [GoogleComputeEngineAccount](#proto.cloudprovider.GoogleComputeEngineAccount)
    - [VirtualizationSettings](#proto.cloudprovider.VirtualizationSettings)
  
- [cloudprovider/huaweicloud.proto](#cloudprovider/huaweicloud.proto)
    - [HuaweiCloud](#proto.cloudprovider.HuaweiCloud)
    - [HuaweiCloudAccount](#proto.cloudprovider.HuaweiCloudAccount)
    - [HuaweiCloudBakeryDefaults](#proto.cloudprovider.HuaweiCloudBakeryDefaults)
    - [HuaweiCloudBaseImage](#proto.cloudprovider.HuaweiCloudBaseImage)
    - [HuaweiCloudBaseImageSettings](#proto.cloudprovider.HuaweiCloudBaseImageSettings)
    - [HuaweiCloudVirtualizationSettings](#proto.cloudprovider.HuaweiCloudVirtualizationSettings)
  
- [cloudprovider/kubernetes.proto](#cloudprovider/kubernetes.proto)
    - [Kubernetes](#proto.cloudprovider.Kubernetes)
    - [KubernetesAccount](#proto.cloudprovider.KubernetesAccount)
    - [KubernetesAccountDockerRegistry](#proto.cloudprovider.KubernetesAccountDockerRegistry)
    - [KubernetesCachingPolicy](#proto.cloudprovider.KubernetesCachingPolicy)
    - [KubernetesCustomResource](#proto.cloudprovider.KubernetesCustomResource)
  
- [cloudprovider/oracle.proto](#cloudprovider/oracle.proto)
    - [Oracle](#proto.cloudprovider.Oracle)
    - [OracleAccount](#proto.cloudprovider.OracleAccount)
    - [OracleBakeryDefaults](#proto.cloudprovider.OracleBakeryDefaults)
    - [OracleBaseImage](#proto.cloudprovider.OracleBaseImage)
    - [OracleBaseImageSettings](#proto.cloudprovider.OracleBaseImageSettings)
    - [OracleVirtualizationSettings](#proto.cloudprovider.OracleVirtualizationSettings)
  
- [cloudprovider/providers.proto](#cloudprovider/providers.proto)
    - [Providers](#proto.cloudprovider.Providers)
  
- [config/clouddriver.proto](#config/clouddriver.proto)
    - [Clouddriver](#proto.config.Clouddriver)
  
- [config/deck.proto](#config/deck.proto)
    - [Deck](#proto.config.Deck)
    - [Deck.Canary](#proto.config.Deck.Canary)
    - [Deck.Changelog](#proto.config.Deck.Changelog)
    - [Deck.Features](#proto.config.Deck.Features)
    - [Deck.Notifications](#proto.config.Deck.Notifications)
    - [Deck.Providers](#proto.config.Deck.Providers)
    - [Deck.Providers.Appengine](#proto.config.Deck.Providers.Appengine)
    - [Deck.Providers.Appengine.Defaults](#proto.config.Deck.Providers.Appengine.Defaults)
    - [Deck.Providers.Aws](#proto.config.Deck.Providers.Aws)
    - [Deck.Providers.Aws.Defaults](#proto.config.Deck.Providers.Aws.Defaults)
    - [Deck.Providers.Azure](#proto.config.Deck.Providers.Azure)
    - [Deck.Providers.Azure.Defaults](#proto.config.Deck.Providers.Azure.Defaults)
    - [Deck.Providers.Cloudfoundry](#proto.config.Deck.Providers.Cloudfoundry)
    - [Deck.Providers.Cloudfoundry.Defaults](#proto.config.Deck.Providers.Cloudfoundry.Defaults)
    - [Deck.Providers.Dcos](#proto.config.Deck.Providers.Dcos)
    - [Deck.Providers.Dcos.Defaults](#proto.config.Deck.Providers.Dcos.Defaults)
    - [Deck.Providers.Ecs](#proto.config.Deck.Providers.Ecs)
    - [Deck.Providers.Ecs.Defaults](#proto.config.Deck.Providers.Ecs.Defaults)
    - [Deck.Providers.Gce](#proto.config.Deck.Providers.Gce)
    - [Deck.Providers.Gce.Defaults](#proto.config.Deck.Providers.Gce.Defaults)
    - [Deck.Providers.HuaweiCloud](#proto.config.Deck.Providers.HuaweiCloud)
    - [Deck.Providers.HuaweiCloud.Defaults](#proto.config.Deck.Providers.HuaweiCloud.Defaults)
    - [Deck.Providers.TencentCloud](#proto.config.Deck.Providers.TencentCloud)
    - [Deck.Providers.TencentCloud.Defaults](#proto.config.Deck.Providers.TencentCloud.Defaults)
  
- [config/echo.proto](#config/echo.proto)
    - [Echo](#proto.config.Echo)
    - [Echo.Scheduler](#proto.config.Echo.Scheduler)
    - [Echo.Scheduler.Cron](#proto.config.Echo.Scheduler.Cron)
  
- [config/fiat.proto](#config/fiat.proto)
    - [Fiat](#proto.config.Fiat)
  
- [config/front50.proto](#config/front50.proto)
    - [Front50](#proto.config.Front50)
    - [Front50.Spinnaker](#proto.config.Front50.Spinnaker)
  
- [config/gate.proto](#config/gate.proto)
    - [Cors](#proto.config.Cors)
    - [Gate](#proto.config.Gate)
    - [Gate.GoogleConfig](#proto.config.Gate.GoogleConfig)
    - [Gate.Integrations](#proto.config.Gate.Integrations)
    - [Gate.Integrations.Gremlin](#proto.config.Gate.Integrations.Gremlin)
    - [Gate.Services](#proto.config.Gate.Services)
    - [ServerConfig](#proto.config.ServerConfig)
    - [SpringSecurity](#proto.config.SpringSecurity)
  
- [config/halconfig.proto](#config/halconfig.proto)
    - [Hal](#proto.config.Hal)
  
- [config/kayenta.proto](#config/kayenta.proto)
    - [Kayenta](#proto.config.Kayenta)
    - [Kayenta.ServiceIntegrations](#proto.config.Kayenta.ServiceIntegrations)
    - [Kayenta.ServiceIntegrations.Aws](#proto.config.Kayenta.ServiceIntegrations.Aws)
    - [Kayenta.ServiceIntegrations.Google](#proto.config.Kayenta.ServiceIntegrations.Google)
  
- [config/orca.proto](#config/orca.proto)
    - [Orca](#proto.config.Orca)
    - [Orca.Defaults](#proto.config.Orca.Defaults)
    - [Orca.Defaults.BakeDefaults](#proto.config.Orca.Defaults.BakeDefaults)
    - [Orca.PipelineTemplates](#proto.config.Orca.PipelineTemplates)
    - [Orca.Services](#proto.config.Orca.Services)
    - [Orca.Tasks](#proto.config.Orca.Tasks)
    - [Orca.Tasks.ExecutionWindow](#proto.config.Orca.Tasks.ExecutionWindow)
  
- [config/rosco.proto](#config/rosco.proto)
    - [Rosco](#proto.config.Rosco)
  
- [config/service_enabled.proto](#config/service_enabled.proto)
    - [ServiceSettings](#proto.config.ServiceSettings)
  
- [config/services.proto](#config/services.proto)
    - [Services](#proto.config.Services)
  
- [features.proto](#features.proto)
    - [Features](#proto.Features)
  
- [notification/bearychat.proto](#notification/bearychat.proto)
    - [BearyChat](#proto.notification.BearyChat)
  
- [notification/email.proto](#notification/email.proto)
    - [Email](#proto.notification.Email)
  
- [notification/github_status.proto](#notification/github_status.proto)
    - [GithubStatus](#proto.notification.GithubStatus)
  
- [notification/google_chat.proto](#notification/google_chat.proto)
    - [GoogleChat](#proto.notification.GoogleChat)
  
- [notification/notifications.proto](#notification/notifications.proto)
    - [Notifications](#proto.notification.Notifications)
  
- [notification/pubsub.proto](#notification/pubsub.proto)
    - [PubSub](#proto.notification.PubSub)
  
- [notification/slack.proto](#notification/slack.proto)
    - [Slack](#proto.notification.Slack)
  
- [notification/twilio.proto](#notification/twilio.proto)
    - [Twilio](#proto.notification.Twilio)
  
- [permissions.proto](#permissions.proto)
    - [Permissions](#proto.Permissions)
  
- [pubsub/google.proto](#pubsub/google.proto)
    - [Google](#proto.pubsub.Google)
    - [GooglePublisher](#proto.pubsub.GooglePublisher)
    - [GoogleSubscriber](#proto.pubsub.GoogleSubscriber)
  
    - [MessageFormat](#proto.pubsub.MessageFormat)
  
- [pubsub/pubsub.proto](#pubsub/pubsub.proto)
    - [Pubsub](#proto.pubsub.Pubsub)
  
- [security/authn/authn.proto](#security/authn/authn.proto)
    - [Authentication](#proto.security.authn.Authentication)
    - [Basic](#proto.security.authn.Basic)
    - [Iap](#proto.security.authn.Iap)
    - [Ldap](#proto.security.authn.Ldap)
    - [OAuth2](#proto.security.authn.OAuth2)
    - [OAuth2.UserInfoRequirementsEntry](#proto.security.authn.OAuth2.UserInfoRequirementsEntry)
    - [OAuth2Client](#proto.security.authn.OAuth2Client)
    - [OAuth2Resource](#proto.security.authn.OAuth2Resource)
    - [OAuth2UserInfoMapping](#proto.security.authn.OAuth2UserInfoMapping)
    - [Saml](#proto.security.authn.Saml)
    - [Saml.UserAttributes](#proto.security.authn.Saml.UserAttributes)
    - [UsernamePassword](#proto.security.authn.UsernamePassword)
    - [X509](#proto.security.authn.X509)
  
    - [OAuth2.OAuth2Provider](#proto.security.authn.OAuth2.OAuth2Provider)
    - [OAuth2Client.AuthenticationScheme](#proto.security.authn.OAuth2Client.AuthenticationScheme)
  
- [security/authz/authz.proto](#security/authz/authz.proto)
    - [Authorization](#proto.security.authz.Authorization)
    - [FileRoleProvider](#proto.security.authz.FileRoleProvider)
    - [GithubRoleProvider](#proto.security.authz.GithubRoleProvider)
    - [GoogleRoleProvider](#proto.security.authz.GoogleRoleProvider)
    - [GroupMembership](#proto.security.authz.GroupMembership)
    - [LdapRoleProvider](#proto.security.authz.LdapRoleProvider)
  
    - [GroupMembership.RoleProviderType](#proto.security.authz.GroupMembership.RoleProviderType)
  
- [security/security.proto](#security/security.proto)
    - [Security](#proto.security.Security)
  
- [security/ssl.proto](#security/ssl.proto)
    - [ApiSecurity](#proto.security.ApiSecurity)
    - [ApiSsl](#proto.security.ApiSsl)
    - [UiSecurity](#proto.security.UiSecurity)
    - [UiSsl](#proto.security.UiSsl)
  
    - [ClientAuth](#proto.security.ClientAuth)
  
- [security/trust_store.proto](#security/trust_store.proto)
    - [TrustStore](#proto.security.TrustStore)
    - [WebhookConfig](#proto.security.WebhookConfig)
  
- [stats.proto](#stats.proto)
    - [Stats](#proto.Stats)
  
- [storage/azs.proto](#storage/azs.proto)
    - [Azs](#proto.storage.Azs)
  
- [storage/gcs.proto](#storage/gcs.proto)
    - [Gcs](#proto.storage.Gcs)
  
- [storage/oracle.proto](#storage/oracle.proto)
    - [Oracle](#proto.storage.Oracle)
  
- [storage/persistent_storage.proto](#storage/persistent_storage.proto)
    - [PersistentStorage](#proto.storage.PersistentStorage)
  
- [storage/s3.proto](#storage/s3.proto)
    - [S3](#proto.storage.S3)
  
    - [S3ServerSideEncryption](#proto.storage.S3ServerSideEncryption)
  
- [Scalar Value Types](#scalar-value-types)



<a name="artifact/artifacts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/artifacts.proto



<a name="proto.artifact.Artifacts"></a>

### Artifacts
Configuration for artifact support.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bitbucket | [Bitbucket](#proto.artifact.Bitbucket) |  |  |
| gcs | [Gcs](#proto.artifact.Gcs) |  |  |
| github | [GitHub](#proto.artifact.GitHub) |  |  |
| gitlab | [GitLab](#proto.artifact.GitLab) |  |  |
| gitrepo | [GitRepo](#proto.artifact.GitRepo) |  |  |
| helm | [Helm](#proto.artifact.Helm) |  |  |
| http | [Http](#proto.artifact.Http) |  |  |
| maven | [Maven](#proto.artifact.Maven) |  |  |
| oracle | [Oracle](#proto.artifact.Oracle) |  |  |
| s3 | [S3](#proto.artifact.S3) |  |  |
| templates | [Template](#proto.artifact.Template) | repeated |  |





 

 

 

 



<a name="artifact/bitbucket.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/bitbucket.proto



<a name="proto.artifact.Bitbucket"></a>

### Bitbucket
Configuration for the Bitbucket artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Bitbucket artifact provider is enabled. |
| accounts | [BitbucketAccount](#proto.artifact.BitbucketAccount) | repeated | The list of configured Bitbucket accounts. |






<a name="proto.artifact.BitbucketAccount"></a>

### BitbucketAccount
Configuration for a Bitbucket artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | The username of the account. Either `username` and `password` should be set, or `usernamePasswordFile` should be set. |
| password | [string](#string) |  | The password of the account. Either `username` and `password` should be set, or `usernamePasswordFile` should be set. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password of the account in the format `${username}:${password}`. Either `username` and `password` should be set, or `usernamePasswordFile` should be set. |





 

 

 

 



<a name="artifact/gcs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/gcs.proto



<a name="proto.artifact.Gcs"></a>

### Gcs
Configuration for the Google Cloud Storage (GCS) artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the GCS artifact provider is enabled. |
| accounts | [GcsAccount](#proto.artifact.GcsAccount) | repeated | The list of configured GCS accounts. |






<a name="proto.artifact.GcsAccount"></a>

### GcsAccount
Configuration for a GCS artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| jsonPath | [string](#string) |  | The path to a JSON key for a GCP service account with which to authenticate. The service account must have the `roles/storage.admin` role enabled. |





 

 

 

 



<a name="artifact/github.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/github.proto



<a name="proto.artifact.GitHub"></a>

### GitHub
Configuration for the GitHub artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the GitHub artifact provider is enabled. |
| accounts | [GitHubAccount](#proto.artifact.GitHubAccount) | repeated | The list of configured GitHub accounts. |






<a name="proto.artifact.GitHubAccount"></a>

### GitHubAccount
Configuration for a GitHub artifact account. Either `username` and
`password`, `usernamePasswordFile`, `token`, or `tokenFile` should be
specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | The GitHub account username. |
| password | [string](#string) |  | The GitHub account password. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password of the account in the format `${username}:${password}`. |
| token | [string](#string) |  | The GitHub access token. |
| tokenFile | [string](#string) |  | The path to a file containing the GitHub access token. |





 

 

 

 



<a name="artifact/gitlab.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/gitlab.proto



<a name="proto.artifact.GitLab"></a>

### GitLab
Configuration for the GitLab artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the GitLab artifact provider is enabled. |
| accounts | [GitLabAccount](#proto.artifact.GitLabAccount) | repeated | The list of configured GitLab accounts. |






<a name="proto.artifact.GitLabAccount"></a>

### GitLabAccount
Configuration for a GitLab artifact account. Either `token` or
`tokenFile` should be specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| token | [string](#string) |  | The GitLab access token. |
| tokenFile | [string](#string) |  | The path to a file containing the GitLab access token. |





 

 

 

 



<a name="artifact/gitrepo.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/gitrepo.proto



<a name="proto.artifact.GitRepo"></a>

### GitRepo
Configuration for the Git repo artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Git repo artifact provider is enabled. |
| accounts | [GitRepoAccount](#proto.artifact.GitRepoAccount) | repeated | The list of configured Git Repo accounts. |






<a name="proto.artifact.GitRepoAccount"></a>

### GitRepoAccount
Configuration for a Git repo artifact account. An account maps to a
credential that is able to authenticate against a Git repository hosted by a
Git hosting service. Either `username` and `password`,
`usernamePasswordFile`, `token`, `tokenFile`, or `sshPrivateKeyFilePath` and
`sshPrivateKeyPassphrase` must be set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | The username of the account. |
| password | [string](#string) |  | The password of the account. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password of the account in the format `${username}:${password}`. |
| token | [string](#string) |  | The Git repo access token. |
| tokenFile | [string](#string) |  | The path to a file containing the Git repo access token. |
| sshPrivateKeyFilePath | [string](#string) |  | The path to an SSH private key to be used when connecting with the Git repo over SSH. |
| sshPrivateKeyPassphrase | [string](#string) |  | The passphrase to an SSH private key to be used when connecting with the Git repo over SSH. |
| sshKnownHostsFilePath | [string](#string) |  | The path to a `known_hosts` file to be used when connecting with a Git repository over SSH. |
| sshTrustUnknownHosts | [bool](#bool) |  | If `true`, Spinnaker can connect with a Git repository over SSH without verifying the server&#39;s IP address against a `known_hosts` file. |





 

 

 

 



<a name="artifact/helm.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/helm.proto



<a name="proto.artifact.Helm"></a>

### Helm
Configuration for the Helm artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Helm artifact provider is enabled. |
| accounts | [HelmAccount](#proto.artifact.HelmAccount) | repeated | The list of configured Helm accounts. |






<a name="proto.artifact.HelmAccount"></a>

### HelmAccount
Configuration for a Helm artifact account. Either `username` and `password`
or `usernamePasswordFile` should be specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| repository | [string](#string) |  | The Helm chart repository URL. |
| username | [string](#string) |  | A username for Helm chart repository basic auth. |
| password | [string](#string) |  | A password for Helm chart repository basic auth. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password for Helm chart repository basic auth. Must be in the format `${username}:${password}`. |





 

 

 

 



<a name="artifact/http.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/http.proto



<a name="proto.artifact.Http"></a>

### Http
Configuration for the HTTP artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the HTTP artifact provider is enabled. |
| accounts | [HttpAccount](#proto.artifact.HttpAccount) | repeated | The list of configured HTTP accounts. |






<a name="proto.artifact.HttpAccount"></a>

### HttpAccount
Configuration for an HTTP artifact account. Either `username` and `password`
or `usernamePasswordFile` should be specified as means of authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | A username for HTTP basic auth. |
| password | [string](#string) |  | A password for HTTP basic auth. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the username and password for HTTP basic auth. Must be in the format `${username}:${password}`. |





 

 

 

 



<a name="artifact/maven.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/maven.proto



<a name="proto.artifact.Maven"></a>

### Maven
Configuration for the Maven artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Maven artifact provider is enabled. |
| accounts | [MavenAccount](#proto.artifact.MavenAccount) | repeated | The list of configured Maven accounts. |






<a name="proto.artifact.MavenAccount"></a>

### MavenAccount
Configuration for a Maven artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| repositoryUrl | [string](#string) |  | (Required) Full URI for the Maven repository (e.g., `http://some.host.com/repository/path`). |





 

 

 

 



<a name="artifact/oracle.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/oracle.proto



<a name="proto.artifact.Oracle"></a>

### Oracle
Configuration for the Oracle artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Oracle artifact provider is enabled. |
| accounts | [OracleAccount](#proto.artifact.OracleAccount) | repeated | The list of configured Oracle artifact accounts. |






<a name="proto.artifact.OracleAccount"></a>

### OracleAccount
Configuration for an Oracle artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| fingerprint | [string](#string) |  | The fingerprint of the public key. |
| namespace | [string](#string) |  | The namespace in which the bucket and objects should be created. |
| privateKeyPassphrase | [string](#string) |  | The passphrase used for the private key, if it is encrypted. |
| region | [string](#string) |  | An Oracle region (e.g., `us-phoenix-1`). |
| sshPrivateKeyFilePath | [string](#string) |  | Path to the private key in PEM format. |
| tenancyId | [string](#string) |  | The OCID of the Oracle Tenancy to use. |
| userId | [string](#string) |  | The OCID of the Oracle User with which to authenticate. |





 

 

 

 



<a name="artifact/s3.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/s3.proto



<a name="proto.artifact.S3"></a>

### S3
Configuration for the S3 artifact provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the S3 artifact provider is enabled. |
| accounts | [S3Account](#proto.artifact.S3Account) | repeated | The list of configured S3 artifact accounts. |






<a name="proto.artifact.S3Account"></a>

### S3Account
Configuration for an S3 artifact account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| apiEndpoint | [string](#string) |  | The S3 API endpoint. Only required when using an S3 clone such as Minio. |
| apiRegion | [string](#string) |  | The S3 API region. Only required when using an S3 clone such as Minio. |
| awsAccessKeyId | [string](#string) |  | The AWS Access Key ID. If not provided, Spinnaker will try to find AWS credentials as described at http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default. |
| awsSecretAccessKey | [string](#string) |  | The AWS Secret Key. |
| region | [string](#string) |  | The S3 region. |





 

 

 

 



<a name="artifact/template.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## artifact/template.proto



<a name="proto.artifact.Template"></a>

### Template
Configuration for a Jinja template for Spinnaker to use for artifact
extraction. For more details, please read the documentation:
https://www.spinnaker.io/reference/artifacts/from-build-triggers/#artifacts-from-build-triggers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the artifact template. |
| templatePath | [string](#string) |  | The path to the artifact template. |





 

 

 

 



<a name="canary/aws.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/aws.proto



<a name="proto.canary.Aws"></a>

### Aws
Configuration for the AWS canary integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| accounts | [AwsAccount](#proto.canary.AwsAccount) | repeated | The list of configured accounts. |
| s3Enabled | [bool](#bool) |  | Whether to enable S3 as a persistent store. |






<a name="proto.canary.AwsAccount"></a>

### AwsAccount
Configuration for an AWS account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| bucket | [string](#string) |  | The name of a storage bucket that this account has access to. If you specify a globally unique bucket name that doesn&#39;t exist yet, Kayenta will create that bucket for you. |
| region | [string](#string) |  | The AWS region to use. |
| rootFolder | [string](#string) |  | The root folder in the chosen bucket in which to store all of the canary service&#39;s persistent data. Defaults to `kayenta`. |
| profileName | [string](#string) |  | The profile name to use when resolving AWS credentials. Typically found in `~/.aws/credentials`. Defaults to `default`. |
| endpoint | [string](#string) |  | The endpoint used to reach the service implementing the AWS API. Typical use is with Minio. |
| accessKeyId | [string](#string) |  | The default access key used to communicate with AWS. |
| supportedTypes | [SupportedType](#proto.canary.SupportedType) | repeated | If enabling S3, include CONFIGURATION_STORE and/or OBJECT_STORE in this list. |





 

 

 

 



<a name="canary/canary.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/canary.proto



<a name="proto.canary.Canary"></a>

### Canary
Configuration for Spinnaker&#39;s canary service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the canary service is enabled. |
| serviceIntegrations | [Canary.ServiceIntegrations](#proto.canary.Canary.ServiceIntegrations) |  | Canary service integrations. You must configure at least one account of each canary.SupportedType (METRICS_STORE, CONFIGURATION_STORE, OBJECT_STORE) in order to use Spinnaker&#39;s canary service. |
| defaultMetricsAccount | [string](#string) |  | Name of the metrics account to use by default. |
| defaultMetricsStore | [string](#string) |  | Name of the metrics store to use by default (e.g., `prometheus`, `datadog`). |
| showAllConfigsEnabled | [bool](#bool) |  | Whether or not to show all canary configs in Deck, or just those scoped to the current application. |
| templatesEnabled | [bool](#bool) |  | Whether or not to enable custom filter templates for canary configs in Deck. |
| defaultJudge | [string](#string) |  | The default canary judge. Defaults to `NetflixACAJudge-v1.0`, which is currently the only open-source judge available by default. |
| storageAccountName | [string](#string) |  | Name of storage account to use by default. |






<a name="proto.canary.Canary.ServiceIntegrations"></a>

### Canary.ServiceIntegrations



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aws | [Aws](#proto.canary.Aws) |  |  |
| datadog | [Datadog](#proto.canary.Datadog) |  |  |
| google | [Google](#proto.canary.Google) |  |  |
| newrelic | [NewRelic](#proto.canary.NewRelic) |  |  |
| prometheus | [Prometheus](#proto.canary.Prometheus) |  |  |
| signalfx | [SignalFx](#proto.canary.SignalFx) |  |  |





 

 

 

 



<a name="canary/datadog.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/datadog.proto



<a name="proto.canary.Datadog"></a>

### Datadog
Configuration for the Datadog canary integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| accounts | [DatadogAccount](#proto.canary.DatadogAccount) | repeated | The list of configured accounts. |






<a name="proto.canary.DatadogAccount"></a>

### DatadogAccount
Configuration for a Datadog account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| endpoint | [DatadogAccount.Endpoint](#proto.canary.DatadogAccount.Endpoint) |  | Configuration for the Datadog server endpoint. |
| apiKey | [string](#string) |  | Your organization&#39;s unique Datadog API key. See https://app.datadoghq.com/account/settings#api. |
| applicationKey | [string](#string) |  | Your Datadog application key. See https://app.datadoghq.com/account/settings#api. |






<a name="proto.canary.DatadogAccount.Endpoint"></a>

### DatadogAccount.Endpoint
Configuration for the Datadog server endpoint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseUrl | [string](#string) |  | (Required) The base URL of the Datadog server. |





 

 

 

 



<a name="canary/gcs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/gcs.proto



<a name="proto.canary.Gcs"></a>

### Gcs
Configuration for the GCS canary integration. If enabled, you must
also configure at least one canary.GoogleAccount with a list of
supportedTypes that includes canary.SupportedType.CONFIGURATION_STORE and/or
canary.SupportedType.OBJECT_STORE.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |





 

 

 

 



<a name="canary/google.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/google.proto



<a name="proto.canary.Google"></a>

### Google
Configuration for the Google canary integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| accounts | [GoogleAccount](#proto.canary.GoogleAccount) | repeated | The list of configured accounts. |
| gcsEnabled | [bool](#bool) |  | Whether GCS is enabled as a persistent store. |
| stackdriverEnabled | [bool](#bool) |  | Whether Stackdriver is enabled as a metrics source. |
| metadataCachingIntervalMS | [int32](#int32) |  | Number of milliseconds to wait in between caching the names of available Stackdriver metric types (used when building canary configs). Defaults to 60000. |






<a name="proto.canary.GoogleAccount"></a>

### GoogleAccount
Configuration for a Google account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| bucket | [string](#string) |  | The name of a storage bucket that this account has access to. If you specify a globally unique bucket name that doesn&#39;t exist yet, Kayenta will create that bucket for you. |
| bucketLocation | [string](#string) |  | This is only required if the bucket you specify doesn’t exist yet. In that case, the bucket will be created in that location. See https://cloud.google.com/storage/docs/managing-buckets#manage-class-location. |
| rootFolder | [string](#string) |  | The root folder in the chosen bucket in which to store all of the canary service&#39;s persistent data in. Defaults to `kayenta`. |
| project | [string](#string) |  | (Required) The Google Cloud Platform project the canary service will use to consume GCS and Stackdriver. |
| supportedTypes | [SupportedType](#proto.canary.SupportedType) | repeated | If enabling Stackdriver, include METRICS_STORE in this list. If enabling GCS, include CONFIGURATION_STORE and/or OBJECT_STORE in this list. |





 

 

 

 



<a name="canary/newrelic.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/newrelic.proto



<a name="proto.canary.NewRelic"></a>

### NewRelic
Configuration for the New Relic canary integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| accounts | [NewRelicAccount](#proto.canary.NewRelicAccount) | repeated | The list of configured accounts. |






<a name="proto.canary.NewRelicAccount"></a>

### NewRelicAccount
Configuration for a New Relic account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| endpoint | [NewRelicAccount.Endpoint](#proto.canary.NewRelicAccount.Endpoint) |  | Configuration for the New Relic Insights server endpoint. |
| apiKey | [string](#string) |  | (Required) Your account&#39;s unique New Relic Insights API key. See https://docs.newrelic.com/docs/insights/insights-api/get-data/query-insights-event-data-api. |
| applicationKey | [string](#string) |  | (Required) Your New Relic account ID. See https://docs.newrelic.com/docs/accounts/install-new-relic/account-setup/account-id. |






<a name="proto.canary.NewRelicAccount.Endpoint"></a>

### NewRelicAccount.Endpoint
Configuration for the New Relic Insights server endpoint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseUrl | [string](#string) |  | The base URL to the New Relic Insights server. |





 

 

 

 



<a name="canary/prometheus.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/prometheus.proto



<a name="proto.canary.Prometheus"></a>

### Prometheus
Configuration for the Prometheus canary integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| accounts | [PrometheusAccount](#proto.canary.PrometheusAccount) | repeated | The list of configured accounts. |
| metadataCachingIntervalMS | [int32](#int32) |  | Number of milliseconds to wait in between caching the names of available metric types (used when building canary configs). Defaults to 60000. |






<a name="proto.canary.PrometheusAccount"></a>

### PrometheusAccount
Configuration for a Prometheus account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| endpoint | [PrometheusAccount.Endpoint](#proto.canary.PrometheusAccount.Endpoint) |  | Configuration for the Prometheus server endpoint. |
| username | [string](#string) |  | A basic auth username. Either `usernamePasswordFile` or `username` and `password` is required. |
| password | [string](#string) |  | A basic auth password. Either `usernamePasswordFile` or `username` and `password` is required. |
| usernamePasswordFile | [string](#string) |  | The path to a file containing the basic auth username and password in the format `${username}:${password}`. Either `usernamePasswordFile` or `username` and `password` is required. |






<a name="proto.canary.PrometheusAccount.Endpoint"></a>

### PrometheusAccount.Endpoint
Configuration for the Prometheus server endpoint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseUrl | [string](#string) |  | (Required) The base URL of the Prometheus server. |





 

 

 

 



<a name="canary/s3.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/s3.proto



<a name="proto.canary.S3"></a>

### S3
Configuration for the S3 canary integration. If enabled, you must
also configure at least one canary.AwsAccount with a list of
supportedTypes that includes canary.SupportedType.CONFIGURATION_STORE and/or
canary.SupportedType.OBJECT_STORE.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |





 

 

 

 



<a name="canary/signalfx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/signalfx.proto



<a name="proto.canary.SignalFx"></a>

### SignalFx
Configuration for the SignalFx canary integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| accounts | [SignalFxAccount](#proto.canary.SignalFxAccount) | repeated | The list of configured accounts. |






<a name="proto.canary.SignalFxAccount"></a>

### SignalFxAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| accessToken | [string](#string) |  | (Required) The SignalFx access token. |
| endpoint | [SignalFxAccount.Endpoint](#proto.canary.SignalFxAccount.Endpoint) |  | The SignalFx server endpoint. |
| defaultScopeKey | [string](#string) |  | Scope key is used to distinguish between base and canary deployments. If omitted, every request must supply the `_scope_key` param in extended scope params. |
| defaultLocationKey | [string](#string) |  | Location key is used to filter by deployment region. If omitted, requests must supply the `_location_key` if it is needed. |






<a name="proto.canary.SignalFxAccount.Endpoint"></a>

### SignalFxAccount.Endpoint
The SignalFx server endpoint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseUrl | [string](#string) |  | The base URL to the SignalFx server. Defaults to https://stream.signalfx.com |





 

 

 

 



<a name="canary/stackdriver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/stackdriver.proto



<a name="proto.canary.Stackdriver"></a>

### Stackdriver
Configuration for the Stackdriver canary integration. If enabled, you must
also configure at least one canary.GoogleAccount with a list of
supportedTypes that includes canary.SupportedType.METRICS_STORE.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the integration is enabled. |
| metadataCachingIntervalMS | [int32](#int32) |  | Number of milliseconds to wait in between caching the names of available Stackdriver metric types (used when building canary configs). Defaults to 60000. |





 

 

 

 



<a name="canary/supported_type.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## canary/supported_type.proto


 


<a name="proto.canary.SupportedType"></a>

### SupportedType


| Name | Number | Description |
| ---- | ------ | ----------- |
| METRICS_STORE | 0 |  |
| CONFIGURATION_STORE | 1 |  |
| OBJECT_STORE | 2 |  |


 

 

 



<a name="ci/ci.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/ci.proto



<a name="proto.ci.Ci"></a>

### Ci
Configuration for integration with continuous integration systems.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gcb | [GoogleCloudBuild](#proto.ci.GoogleCloudBuild) |  |  |
| codebuild | [CodeBuild](#proto.ci.CodeBuild) |  |  |
| concourse | [Concourse](#proto.ci.Concourse) |  |  |
| jenkins | [Jenkins](#proto.ci.Jenkins) |  |  |
| travis | [Travis](#proto.ci.Travis) |  |  |
| wercker | [Wercker](#proto.ci.Wercker) |  |  |





 

 

 

 



<a name="ci/codebuild.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/codebuild.proto



<a name="proto.ci.CodeBuild"></a>

### CodeBuild
Configuration for AWS CodeBuild.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether AWS CodeBuild is enabled. |
| accounts | [CodeBuildAccount](#proto.ci.CodeBuildAccount) | repeated | The list of configured AWS CodeBuild accounts. |
| accessKeyId | [string](#string) |  | Your AWS Access Key ID. If not provided, Spinnaker will try to find AWS credentials as described at http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default |
| secretAccessKey | [string](#string) |  | Your AWS Secret Key. |






<a name="proto.ci.CodeBuildAccount"></a>

### CodeBuildAccount
Configuration for an AWS CodeBuild account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| accountId | [string](#string) |  | The AWS account ID that will be used to trigger CodeBuild builds. |
| assumeRole | [string](#string) |  | If set, Spinnaker will configure a credentials provider that uses the AWS Security Token Service to assume the specified role. |
| region | [string](#string) |  | (Required) The AWS region in which your CodeBuild projects live. |





 

 

 

 



<a name="ci/concourse.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/concourse.proto



<a name="proto.ci.Concourse"></a>

### Concourse
Configuration for Concourse.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Concourse is enabled. |
| masters | [ConcourseAccount](#proto.ci.ConcourseAccount) | repeated | The list of configured Concourse accounts. |






<a name="proto.ci.ConcourseAccount"></a>

### ConcourseAccount
Configuration for a Concourse account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | (Required) The username of the Concourse user to authenticate as. |
| password | [string](#string) |  | (Required) The password of the Concourse user to authenticate as. |
| url | [string](#string) |  | (Required) The URL at which your Concourse search is reachable. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. A user must have at least one of the READ roles in order to view this build account or use it as a trigger source. A user must have at least one of the WRITE roles in order to run jobs on this build account. |





 

 

 

 



<a name="ci/gcb.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/gcb.proto



<a name="proto.ci.GoogleCloudBuild"></a>

### GoogleCloudBuild
Configuration for the Google Cloud Build Provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the Google Cloud Build provider is enabled. |
| accounts | [GoogleCloudBuildAccount](#proto.ci.GoogleCloudBuildAccount) | repeated | The list of configured Google Cloud Build accounts. |






<a name="proto.ci.GoogleCloudBuildAccount"></a>

### GoogleCloudBuildAccount
Configuration for a Google Cloud Build account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| project | [string](#string) |  | The name of the Google Cloud Platform project in which to trigger and monitor builds. |
| subscriptionName | [string](#string) |  | The name of the Pub/Sub subscription on which to listen for build changes. |
| jsonKey | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |





 

 

 

 



<a name="ci/jenkins.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/jenkins.proto



<a name="proto.ci.Jenkins"></a>

### Jenkins
Configuration for Jenkins.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Jenkins is enabled. |
| masters | [JenkinsAccount](#proto.ci.JenkinsAccount) | repeated | The list of configured Jenkins accounts. |






<a name="proto.ci.JenkinsAccount"></a>

### JenkinsAccount
Configuration for a Jenkins account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| username | [string](#string) |  | (Required) The username of the Jenkins user to authenticate as. |
| password | [string](#string) |  | (Required) The password of the Jenkins user to authenticate as. |
| address | [string](#string) |  | (Required) The address at which the Jenkins server is reachable. |
| csrf | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Whether or not to negotiate CSRF tokens when calling Jenkins. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. A user must have at least one of the READ roles in order to view this build account or use it as a trigger source. A user must have at least one of the WRITE roles in order to run jobs on this build account. |





 

 

 

 



<a name="ci/travis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/travis.proto



<a name="proto.ci.Travis"></a>

### Travis
Configuration for Travis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Travis is enabled. |
| masters | [TravisAccount](#proto.ci.TravisAccount) | repeated | The list of configured Travis accounts. |






<a name="proto.ci.TravisAccount"></a>

### TravisAccount
Configuration for a Travis account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| address | [string](#string) |  | (Required) The address of the Travis API (https://api.travis-ci.org). |
| baseUrl | [string](#string) |  | (Required) The base URL to the Travis UI (https://travis-ci.org). |
| githubToken | [string](#string) |  | The GitHub token with which to authenticate against Travis. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. A user must have at least one of the READ roles in order to view this build account or use it as a trigger source. A user must have at least one of the WRITE roles in order to run jobs on this build account. |





 

 

 

 



<a name="ci/wercker.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ci/wercker.proto



<a name="proto.ci.Wercker"></a>

### Wercker
Configuration for Wercker.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Wercker is enabled. |
| masters | [WerckerAccount](#proto.ci.WerckerAccount) | repeated | The list of configured Wercker accounts. |






<a name="proto.ci.WerckerAccount"></a>

### WerckerAccount
Configuration for a Wercker account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| address | [string](#string) |  | (Required) The address at which your Wercker instance is reachable. |
| token | [string](#string) |  | The personal token of the Wercker user to authenticate as. |
| user | [string](#string) |  | The username of the Wercker user to authenticate as. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. A user must have at least one of the READ roles in order to view this build account or use it as a trigger source. A user must have at least one of the WRITE roles in order to run jobs on this build account. |





 

 

 

 



<a name="cloudprovider/appengine.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/appengine.proto



<a name="proto.cloudprovider.Appengine"></a>

### Appengine
Configuration for the Google App Engine (GAE) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AppengineAccount](#proto.cloudprovider.AppengineAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.cloudprovider.AppengineAccount"></a>

### AppengineAccount
Configuration for an App Engine account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cachingIntervalSeconds | [int32](#int32) |  | The interval in seconds at which Spinnaker will poll for updates in your App Engine clusters. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| gcloudReleaseTrack | [GcloudReleaseTrack](#proto.cloudprovider.GcloudReleaseTrack) |  | The gcloud release track that Spinnaker will use when deploying to App Engine. |
| gitHttpsUsername | [string](#string) |  | A username to be used when connecting to a remote git repository server over HTTPS. If set, `gitHttpsPassword` must also be set. |
| gitHttpsPassword | [string](#string) |  | A password to be used when connecting to a remote git repository server over HTTPS. If set, `gitHttpsUsername` must also be set. |
| githubOAuthAccessToken | [string](#string) |  | An OAuth token provided by Github for connecting to a git repository over HTTPS. See https://help.github.com/articles/creating-an-access-token-for-command-line-use for more information. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| localRepositoryDirectory | [string](#string) |  | A local directory to be used to stage source files for App Engine deployments within Clouddriver. Defaults to `/var/tmp/clouddriver`. |
| omitServices | [string](#string) | repeated | A list of regular expressions. Any service matching one of these regexes will be ignored by Spinnaker. |
| omitVersions | [string](#string) | repeated | A list of regular expressions. Any version matching one of these regexes will be ignored by Spinnaker. |
| project | [string](#string) |  | The Google Cloud Platform project this Spinnaker account will manage. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| services | [string](#string) | repeated | A list of regular expressions. Any service matching one of these regexes will be indexed by Spinnaker (unless the service also matches a regex in `omitServices`). |
| sshKnownHostsFilePath | [string](#string) |  | The path to a `known_hosts` file to be used when connecting with a remote git repository over SSH. |
| sshPrivateKeyFilePath | [string](#string) |  | The path to an SSH private key to be used when connecting with a remote git repository over SSH. If set, `sshPrivateKeyPassphrase` must also be set. |
| sshPrivateKeyPassphrase | [string](#string) |  | The passphrase to an SSH private key to be used when connecting with a remote git repository over SSH. If set, `sshPrivateKeyFilePath` must also be set. |
| sshTrustUnknownHosts | [bool](#bool) |  | Enabling this flag will allow Spinnaker to connect with a remote git repository over SSH without verifying the server&#39;s IP address against a `known_hosts` file. Defaults to false. |
| versions | [string](#string) | repeated | A list of regular expressions. Any version matching one of these regexes will be indexed by Spinnaker (unless the version also matches a regex in `omitVersions`). |
| name | [string](#string) |  | The name of the account. |





 


<a name="proto.cloudprovider.GcloudReleaseTrack"></a>

### GcloudReleaseTrack
Represents a release track of the gcloud tool.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STABLE | 0 | Standard release track; runs commands via `gcloud...` |
| BETA | 1 | Alpha release track; runs commands via `gcloud beta...` |
| ALPHA | 2 | Alpha release track; runs commands via `gcloud alpha...` |


 

 

 



<a name="cloudprovider/aws.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/aws.proto



<a name="proto.cloudprovider.Aws"></a>

### Aws
Configuration for the AWS provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AwsAccount](#proto.cloudprovider.AwsAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| accessKeyId | [string](#string) |  | Your AWS Access Key ID. Note that if you are baking AMIs with Rosco, you may also need to set `AwsBakeryDefaults.awsAccessKey`. |
| secretAccessKey | [string](#string) |  | Your AWS Secret Key. Note that if you are baking AMIs with Rosco, you may also need to set `AwsBakeryDefaults.awsSecretKey`. |
| defaultRegions | [AwsRegion](#proto.cloudprovider.AwsRegion) | repeated | List of default regions. |
| features | [AwsFeatures](#proto.cloudprovider.AwsFeatures) |  | Configuration for AWS-specific features. |
| bakeryDefaults | [AwsBakeryDefaults](#proto.cloudprovider.AwsBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.cloudprovider.AwsAccount"></a>

### AwsAccount
Configuration for an AWS account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountId | [string](#string) |  | The AWS account ID to manage. See http://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html for more information. |
| assumeRole | [string](#string) |  | If set, Spinnaker will configure a credentials provider that uses AWS Security Token Service to assume the specified role. Examples: `user/spinnaker`, `role/spinnakerManaged`. |
| defaultKeyPair | [string](#string) |  | The name of the AWS key-pair to use. See http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html for more information. |
| discovery | [string](#string) |  | The endpoint at which your Eureka discovery system is reachable. See https://github.com/Netflix/eureka for more information. Example: `http://.eureka.url.to.use:8080/eureka-server/v2`. Using will make Spinnaker use AWS regions in the hostname to access discovery so that you can have discovery for multiple regions. |
| edda | [string](#string) |  | The endpoint at which Edda is reachable. Edda is not a hard dependency of Spinnaker, but is helpful for reducing the request volume against AWS. See https://github.com/Netflix/edda for more information. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| lifecycleHooks | [AwsLifecycleHook](#proto.cloudprovider.AwsLifecycleHook) | repeated | List of configured AWS lifecycle hooks. |
| regions | [AwsRegion](#proto.cloudprovider.AwsRegion) | repeated | List of configured AWS regions. |
| name | [string](#string) |  | The name of the account. |






<a name="proto.cloudprovider.AwsBakeryDefaults"></a>

### AwsBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| awsAccessKey | [string](#string) |  | The default access key used to communicate with AWS. |
| awsSecretKey | [string](#string) |  | The secret key used to communicate with AWS. |
| awsSubnetId | [string](#string) |  | If using VPC, the default ID of the subnet, such as `subnet-12345def`, where Packer will launch the EC2 instance. This field is required if you are using a non-default VPC. |
| awsVpcId | [string](#string) |  | If launching into a VPC subnet, Packer needs the VPC ID in order to create a temporary security group within the VPC. Requires `subnet_id` to be set. If this default value is left blank, Packer will try to get the VPC ID from `awsSubnetId`. |
| awsAssociatePublicIpAddress | [bool](#bool) |  | If using a non-default VPC, public IP addresses are not provided by default. If this is enabled, your new instance will get a Public IP. |
| defaultVirtualizationType | [string](#string) |  | The default type of virtualization for the AMI you are building. This option must match the supported virtualization type of `AwsVirtualizationSettings.sourceAmi`. Acceptable values: `pv`, `hvm`. |
| baseImages | [AwsBaseImageSettings](#proto.cloudprovider.AwsBaseImageSettings) | repeated | List of configured base images. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.cloudprovider.AwsBaseImage"></a>

### AwsBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | This is the identifier used by AWS to find this base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.cloudprovider.AwsBaseImageSettings"></a>

### AwsBaseImageSettings
Configuration for a base image for the AWS provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [AwsBaseImage](#proto.cloudprovider.AwsBaseImage) |  | Base image configuration. |
| virtualizationSettings | [AwsVirtualizationSettings](#proto.cloudprovider.AwsVirtualizationSettings) |  | Base image virtualization settings. |






<a name="proto.cloudprovider.AwsFeatures"></a>

### AwsFeatures
Configuration for AWS-specific features.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cloudFormation | [AwsFeatures.CloudFormation](#proto.cloudprovider.AwsFeatures.CloudFormation) |  | Configuration for AWS CloudFormation. |






<a name="proto.cloudprovider.AwsFeatures.CloudFormation"></a>

### AwsFeatures.CloudFormation
Configuration for AWS CloudFormation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether AWS CloudFormation is enabled. |






<a name="proto.cloudprovider.AwsLifecycleHook"></a>

### AwsLifecycleHook
Configuration for AWS Auto Scaling Lifecycle Hooks. For more information, see:
https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaultResult | [string](#string) |  | Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Acceptable values: `CONTINUE`, `ABANDON`. |
| heartbeatTimeout | [int32](#int32) |  | Set the heartbeat timeout in seconds for the lifecycle hook. Instances can remain in a wait state for a finite period of time. Must be greater than or equal to 30 and less than or equal to 7200. The default is 3600 (one hour). |
| lifecycleTransition | [string](#string) |  | Type of lifecycle transition. Acceptable values: `autoscaling:EC2_INSTANCE_LAUNCHING`, `autoscaling:EC2_INSTANCE_TERMINATING` |
| notificationTargetARN | [string](#string) |  | The ARN of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. This target can be either an SQS queue or an SNS topic. |
| roleARN | [string](#string) |  | The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue. |






<a name="proto.cloudprovider.AwsRegion"></a>

### AwsRegion
An AWS region.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the region. |






<a name="proto.cloudprovider.AwsVirtualizationSettings"></a>

### AwsVirtualizationSettings
Base image virtualization settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | The name of the region in which to launch the EC2 instance to create the AMI. |
| virtualizationType | [string](#string) |  | The type of virtualization for the AMI you are building. This option must match the supported virtualization type of `sourceAmi`. Acceptable values: `pv`, `hvm`. |
| instanceType | [string](#string) |  | The EC2 instance type to use while building the AMI, such as `t2.small`. |
| sourceAmi | [string](#string) |  | The source AMI whose root volume will be copied and provisioned on the currently running instance. This must be an EBS-backed AMI with a root volume snapshot that you have access to. |
| sshUserName | [string](#string) |  | The username to connect to SSH with. Required if using SSH. |
| winRmUserName | [string](#string) |  | The username to use to connect to WinRM. |
| spotPrice | [string](#string) |  | The maximum hourly price to pay for a spot instance to create the AMI. Spot instances are a type of instance that EC2 starts when the current spot price is less than the maximum price you specify. Spot price will be updated based on available spot instance capacity and current spot instance requests. It may save you some costs. You can set this to `auto` for Packer to automatically discover the best spot price or to &#34;0&#34; to use an on demand instance (default). |
| spotPriceAutoProduct | [string](#string) |  | Required if `spotPrice` is set to `auto`. This tells Packer what sort of AMI you are launching to find the best spot price. This must be one of: `Linux/UNIX`, `SUSE Linux`, `Windows`, `Linux/UNIX (Amazon VPC)`, `SUSE Linux (Amazon VPC)`, `Windows (Amazon VPC)`. |





 

 

 

 



<a name="cloudprovider/azure.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/azure.proto



<a name="proto.cloudprovider.Azure"></a>

### Azure
Configuration for the Azure provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AzureAccount](#proto.cloudprovider.AzureAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [AzureBakeryDefaults](#proto.cloudprovider.AzureBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.cloudprovider.AzureAccount"></a>

### AzureAccount
Configuration for an Azure account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| appKey | [string](#string) |  | (Required) The `appKey` (password) of your service principal. |
| clientId | [string](#string) |  | (Required) The `clientId` (also called `appId`) of your service principal. |
| defaultKeyVault | [string](#string) |  | (Required) The name of a KeyVault that contains the user name, password, and ssh public key used to create VMs |
| defaultResourceGroup | [string](#string) |  | (Required) The default resource group to contain any non-application specific resources. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| objectId | [string](#string) |  | The `objectId` of your service principal. This is only required if using Packer to bake Windows images. |
| packerResourceGroup | [string](#string) |  | The resource group to use if baking images with Packer. |
| packerStorageAccount | [string](#string) |  | The storage account to use if baking images with Packer. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| regions | [string](#string) | repeated | The Azure regions this Spinnaker account will manage. |
| subscriptionId | [string](#string) |  | (Required) The `subscriptionId` to which your service principal is assigned. |
| tenantId | [string](#string) |  | (Required) The `tenantId` to which your service principal is assigned. |
| useSshPublicKey | [bool](#bool) |  | If true, the SSH public key is used to provision the linux VM. If false, the password is used instead. |






<a name="proto.cloudprovider.AzureBakeryDefaults"></a>

### AzureBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImages | [AzureBaseImageSettings](#proto.cloudprovider.AzureBaseImageSettings) | repeated | List of configured base images. |






<a name="proto.cloudprovider.AzureBaseImage"></a>

### AzureBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| publisher | [string](#string) |  | (Required) The Publisher name for your base image. See https://aka.ms/azspinimage to get a list of images. |
| offer | [string](#string) |  | (Required) The offer for your base image. See https://aka.ms/azspinimage to get a list of images. |
| sku | [string](#string) |  | (Required) The SKU for your base image. See https://aka.ms/azspinimage to get a list of images. |
| version | [string](#string) |  | The version of your base image. This defaults to `latest` if not specified. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying `deb` indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.cloudprovider.AzureBaseImageSettings"></a>

### AzureBaseImageSettings
Configuration for a base image for the Azure provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [AzureBaseImage](#proto.cloudprovider.AzureBaseImage) |  | Base image configuration. |





 

 

 

 



<a name="cloudprovider/cloudfoundry.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/cloudfoundry.proto



<a name="proto.cloudprovider.CloudFoundry"></a>

### CloudFoundry
Configuration for the Cloud Foundry provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [CloudFoundryAccount](#proto.cloudprovider.CloudFoundryAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.cloudprovider.CloudFoundryAccount"></a>

### CloudFoundryAccount
Configuration for a Spinnaker Cloud Foundry account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| api | [string](#string) |  | (Required) Host of the Cloud Foundry Foundation API endpoint (e.g., `api.sys.somesystem.com`). |
| appsManagerUri | [string](#string) |  | HTTP(S) URL of the Apps Manager application for the Cloud Foundry Foundation (e.g., `https://apps.sys.somesystem.com`). |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| metricsUri | [string](#string) |  | HTTP(S) URL of the metrics application for the Cloud Foundry Foundation (e.g., `https://metrics.sys.somesystem.com`). |
| password | [string](#string) |  | (Required) Password for the account to use for this Cloud Foundry Foundation. |
| skipSslValidation | [bool](#bool) |  | (Default: `false`) Skip SSL server certificate validation of the API endpoint. |
| user | [string](#string) |  | (Required) User name for the account to use for this Cloud Foundry Foundation. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |





 

 

 

 



<a name="cloudprovider/dcos.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/dcos.proto



<a name="proto.cloudprovider.Dcos"></a>

### Dcos
Configuration for the DC/OS (Distributed Cloud Operating System) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [DcosAccount](#proto.cloudprovider.DcosAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| clusters | [DcosCluster](#proto.cloudprovider.DcosCluster) | repeated | The list of configured clusters. |






<a name="proto.cloudprovider.DcosAccount"></a>

### DcosAccount
Credentials to authenticate against one or more DC/OS clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | (Required) The name of the account. |
| clusters | [DcosAccountCluster](#proto.cloudprovider.DcosAccountCluster) | repeated | (Required) The clusters against which this account will authenticate. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| dockerRegistries | [DcosAccountDockerRegistry](#proto.cloudprovider.DcosAccountDockerRegistry) | repeated | (Required) The list of Docker registries to use with this DC/OS account. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.cloudprovider.DcosAccountCluster"></a>

### DcosAccountCluster
Configuration for a DC/OS cluster associated with a `DcosAccount`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | (Required) The name of the cluster. Must match the name of a `DcosCluster` defined for this provider. |
| uid | [string](#string) |  | (Required) User or service account identifier. |
| serviceKeyFile | [string](#string) |  | Path to a file containing the secret key for service account authentication. If set, `password` should not be set. |
| password | [string](#string) |  | Password for a user account. If set, `serviceKeyFile` should not be set. |






<a name="proto.cloudprovider.DcosAccountDockerRegistry"></a>

### DcosAccountDockerRegistry
Configuration for a Docker registry associated with a `DcosAccount`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | The name of the Docker registry. Must be the name of an account configured with the Docker registry provider. |






<a name="proto.cloudprovider.DcosCluster"></a>

### DcosCluster
Configuration for a DC/OS cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | (Required) The name of the cluster. |
| caCertFile | [string](#string) |  | Root certificate file to trust for connections to the cluster. |
| dcosUrl | [string](#string) |  | (Required) URL of the endpoint for the DC/OS cluster&#39;s admin router. |
| loadBalancer | [DcosClusterLoadBalancer](#proto.cloudprovider.DcosClusterLoadBalancer) |  | Configuration for a DC/OS load balancer. |
| insecureSkipTlsVerify | [bool](#bool) |  | If `true`, disables verification of certificates from the cluster (insecure). |






<a name="proto.cloudprovider.DcosClusterLoadBalancer"></a>

### DcosClusterLoadBalancer
Configuration for a DC/OS load balancer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image | [string](#string) |  | Marathon-lb image to use when creating a load balancer with Spinnaker. |
| serviceAccountSecret | [string](#string) |  | Name of the secret to use for allowing marathon-lb to authenticate with the cluster. Only necessary for clusters with strict or permissive security. |





 

 

 

 



<a name="cloudprovider/docker_registry.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/docker_registry.proto



<a name="proto.cloudprovider.DockerRegistry"></a>

### DockerRegistry
Configuration for the Docker Registry provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [DockerRegistryAccount](#proto.cloudprovider.DockerRegistryAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.cloudprovider.DockerRegistryAccount"></a>

### DockerRegistryAccount
A credential able to authenticate against a set of Docker repositories.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| address | [string](#string) |  | (Required) The registry address from which to pull and deploy images (e.g., `https://index.docker.io`). |
| cacheIntervalSeconds | [int32](#int32) |  | The number of seconds between polling the Docker registry. Certain registries are sensitive to over-polling, and larger intervals (e.g., 10 minutes = 600 seconds) are desirable if you experience rate limiting. Defaults to `30`. |
| cacheThreads | [int32](#int32) |  | The number of threads on which to cache all provided repositories. Really only useful if you have a ton of repos. Defaults to 1. |
| clientTimeoutMillis | [int32](#int32) |  | Timeout in milliseconds for provided repositories. Defaults to `60,000`. |
| email | [string](#string) |  | The email associated with your Docker registry. Often this only needs to be well-formed, rather than be a real address. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| insecureRegistry | [bool](#bool) |  | If `true`, Spinnaker will treat the Docker registry as insecure and not validate the SSL certificate. Defaults to `false`. |
| paginateSize | [int32](#int32) |  | Pagination size for the Docker `repository _catalog` endpoint. Defaults to `100`. |
| password | [string](#string) |  | The Docker registry password. Only one of `password`, `passwordCommand`, and `passwordFile` should be specified. |
| passwordCommand | [string](#string) |  | Command to retrieve Docker token/password. The command must be available in the environment. Only one of `password`, `passwordCommand`, and `passwordFile` should be specified. |
| passwordFile | [string](#string) |  | The path to a file containing your Docker password in plaintext (not a Docker `config.json` file). Only one of `password`, `passwordCommand`, and `passwordFile` should be specified. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |
| repositories | [string](#string) | repeated | An optional list of repositories from which to cache images. If not provided, Spinnaker will attempt to read accessible repositories from the `registries _catalog` endpoint. |
| sortTagsByDate | [bool](#bool) |  | If `true`, Spinnaker will sort tags by creation date. Defaults to `false`. Not recommended for use with large registries; sorting performance scales poorly due to limitations of the Docker V2 API. |
| trackDigests | [bool](#bool) |  | If `true`, Spinnaker will track digest changes. This is not recommended because it greatly increases queries to the registry, and most registries are flaky. Defaults to `false`. |
| username | [string](#string) |  | The username associated with this Docker registry. |





 

 

 

 



<a name="cloudprovider/ecs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/ecs.proto



<a name="proto.cloudprovider.Ecs"></a>

### Ecs
Configuration for the ECS provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [EcsAccount](#proto.cloudprovider.EcsAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.cloudprovider.EcsAccount"></a>

### EcsAccount
Configuration for an ECS account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| awsAccount | [string](#string) |  | (Required) Provide the name of the AWS account associated with this ECS account. See https://github.com/spinnaker/clouddriver/blob/master/clouddriver-ecs/README.md for more information. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |





 

 

 

 



<a name="cloudprovider/google.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/google.proto



<a name="proto.cloudprovider.Consul"></a>

### Consul
Configuration for Consul.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Consul is enabled. |
| agentEndpoint | [string](#string) |  | Reachable Consul node endpoint connected to the Consul cluster. Defaults to localhost. |
| agentPort | [int32](#int32) |  | Port consul is running on for every agent. Defaults to 8500. |
| datacenters | [string](#string) | repeated | List of data centers to cache and keep updated. Defaults to all. |






<a name="proto.cloudprovider.GoogleBakeryDefaults"></a>

### GoogleBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| baseImages | [GoogleBaseImageSettings](#proto.cloudprovider.GoogleBaseImageSettings) | repeated | List of configured base images. |
| zone | [string](#string) |  | The default zone in which to bake an image. |
| network | [string](#string) |  | The Google Compute network ID or URL to use for the launched instance. Defaults to default. |
| useInternalIp | [bool](#bool) |  | If true, use the instance&#39;s internal IP instead of its external IP during baking. |
| networkProjectId | [string](#string) |  | The default project ID for the network and subnet to use for the VM baking your image. |






<a name="proto.cloudprovider.GoogleBaseImage"></a>

### GoogleBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | This is the identifier used by GCP to find this base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| imageFamily | [bool](#bool) |  | If set to true, Deck will annotate the popup tooltip to indicate that the selected option represents an image family. |






<a name="proto.cloudprovider.GoogleBaseImageSettings"></a>

### GoogleBaseImageSettings
Configuration for a base image for the Google provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [GoogleBaseImage](#proto.cloudprovider.GoogleBaseImage) |  | Base image configuration. |
| virtualizationSettings | [VirtualizationSettings](#proto.cloudprovider.VirtualizationSettings) |  | Image source configuration. |






<a name="proto.cloudprovider.GoogleComputeEngine"></a>

### GoogleComputeEngine
Configuration for the Google Compute Engine (GCE) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [GoogleComputeEngineAccount](#proto.cloudprovider.GoogleComputeEngineAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [GoogleBakeryDefaults](#proto.cloudprovider.GoogleBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.cloudprovider.GoogleComputeEngineAccount"></a>

### GoogleComputeEngineAccount
Configuration for a Spinnaker Google account. An account maps to a
credential that can authenticate against a GCP project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| project | [string](#string) |  | The GCP project this Spinnaker account will manage. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| alphaListed | [bool](#bool) |  | Enable this flag if your GCP project has access to alpha features and you want Spinnaker to take advantage of them. |
| imageProjects | [string](#string) | repeated | A list of GCP projects from which Spinnaker will be able to cache and deploy images. When this is omitted, it defaults to the current project. Each project must have granted the IAM role compute.imageUser to the service account associated with the JSON key used by this account, as well as to the Google APIs service account automatically created for the project being managed (should look similar to 12345678912@cloudservices.gserviceaccount.com). See https://cloud.google.com/compute/docs/images/sharing-images-across-projects for more information about sharing images across GCP projects. |
| consul | [Consul](#proto.cloudprovider.Consul) |  | Configuration for Consul. |
| regions | [string](#string) | repeated | A list of regions for caching and mutating calls. This overwrites any default regions set on the provider. |
| userDataFile | [string](#string) |  | The path to user data template file. Spinnaker has the ability to inject userdata into generated instance templates. The mechanism is via a template file that is token replaced to provide some specifics about the deployment. See https://github.com/spinnaker/clouddriver/blob/master/clouddriver-aws/UserData.md for more information. |






<a name="proto.cloudprovider.VirtualizationSettings"></a>

### VirtualizationSettings
Image source configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceImage | [string](#string) |  | The source image. If both sourceImage and sourceImageFamily are set, sourceImage will take precedence. |
| sourceImageFamily | [string](#string) |  | The source image family to create the image from. The newest, non-deprecated image is used. If both sourceImage and sourceImageFamily are set, sourceImage will take precedence. |





 

 

 

 



<a name="cloudprovider/huaweicloud.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/huaweicloud.proto



<a name="proto.cloudprovider.HuaweiCloud"></a>

### HuaweiCloud
Configuration for the Huawei Cloud provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [HuaweiCloudAccount](#proto.cloudprovider.HuaweiCloudAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [HuaweiCloudBakeryDefaults](#proto.cloudprovider.HuaweiCloudBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.cloudprovider.HuaweiCloudAccount"></a>

### HuaweiCloudAccount
Configuration for a Huawei Cloud account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| accountType | [string](#string) |  | The type of account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| authUrl | [string](#string) |  | (Required) The auth URL of the cloud. |
| domainName | [string](#string) |  | (Required) The domain name of the cloud. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| insecure | [bool](#bool) |  | If `true`, disables certificate validation on SSL connections. Needed if certificates are self-signed. Defaults to `false`. |
| password | [string](#string) |  | (Required) The password used to access the cloud. |
| projectName | [string](#string) |  | (Required) The name of the project within the cloud. |
| regions | [string](#string) | repeated | (Required) The region(s) of the cloud. |
| username | [string](#string) |  | (Required) The username used to access the cloud. |






<a name="proto.cloudprovider.HuaweiCloudBakeryDefaults"></a>

### HuaweiCloudBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImages | [HuaweiCloudBaseImageSettings](#proto.cloudprovider.HuaweiCloudBaseImageSettings) | repeated | List of configured base images. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| authUrl | [string](#string) |  | (Required) The default auth URL in which images will be baked. |
| username | [string](#string) |  | (Required) The default username with which images will be baked. |
| password | [string](#string) |  | (Required) The default password with which images will be baked. |
| projectName | [string](#string) |  | The name of the default project in which images will be baked. |
| domainName | [string](#string) |  | (Required) The default domain name in which images will be baked. |
| insecure | [bool](#bool) |  | The security setting for connecting to the Huawei Cloud account. Defaults to `false`. |
| vpcId | [string](#string) |  | (Required) The VPC in which images will be baked. |
| subnetId | [string](#string) |  | (Required) The subnet in which images will be baked. |
| securityGroup | [string](#string) |  | (Required) The default security group in which images will be baked. |
| eipBandwidthSize | [int32](#int32) |  | (Required) The bandwidth size of EIP in which images will be baked. |






<a name="proto.cloudprovider.HuaweiCloudBaseImage"></a>

### HuaweiCloudBaseImage
Huawei Cloud base image settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The name of the base image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying `deb` indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |






<a name="proto.cloudprovider.HuaweiCloudBaseImageSettings"></a>

### HuaweiCloudBaseImageSettings
Configuration for a base image for the Huawei Cloud provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [HuaweiCloudBaseImage](#proto.cloudprovider.HuaweiCloudBaseImage) |  | Base image configuration. |
| virtualizationSettings | [HuaweiCloudVirtualizationSettings](#proto.cloudprovider.HuaweiCloudVirtualizationSettings) | repeated | Image source configuration. |






<a name="proto.cloudprovider.HuaweiCloudVirtualizationSettings"></a>

### HuaweiCloudVirtualizationSettings
Huawei Cloud virtualization settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | (Required) The region for the baking configuration. |
| instanceType | [string](#string) |  | (Required) The instance type for the baking configuration. |
| sourceImageId | [string](#string) |  | (Required) The source image ID for the baking configuration. |
| sshUserName | [string](#string) |  | (Required) The SSH username for the baking configuration. |
| eipType | [string](#string) |  | (Required) The EIP type for the baking configuration. See the API doc to get its value. |





 

 

 

 



<a name="cloudprovider/kubernetes.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/kubernetes.proto



<a name="proto.cloudprovider.Kubernetes"></a>

### Kubernetes
Configuration for the Kubernetes provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [KubernetesAccount](#proto.cloudprovider.KubernetesAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.cloudprovider.KubernetesAccount"></a>

### KubernetesAccount
Configuration for a Spinnaker Kubernetes account. An account maps to a
credential that can authenticate against your Kubernetes cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| kinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account can deploy and will cache. When no kinds are configured, this defaults to all kinds described here: https://spinnaker.io/reference/providers/kubernetes-v2/. This can only be set when omitKinds is empty or not set. |
| omitKinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account cannot deploy to or cache. This can only be set when kinds is empty or not set. |
| context | [string](#string) |  | The kubernetes context to be managed by Spinnaker. See http://kubernetes.io/docs/user-guide/kubeconfig-file/#context for more information. When no context is configured for an account the `current-context` in your kubeconfig is assumed. |
| cacheThreads | [int32](#int32) |  | Number of caching agents for this kubernetes account. Each agent handles a subset of the namespaces available to this account. By default, only 1 agent caches all kinds for all namespaces in the account. |
| namespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account can deploy to and will cache. When no namespaces are configured, this defaults to all namespaces. |
| omitNamespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account cannot deploy to or cache. This can only be set when namespaces is empty or not set. |
| customResources | [KubernetesCustomResource](#proto.cloudprovider.KubernetesCustomResource) | repeated | The list of custom resources Clouddriver will manage and make available for use in Patch and Delete (Manifest) stages. |
| cachingPolicies | [KubernetesCachingPolicy](#proto.cloudprovider.KubernetesCachingPolicy) | repeated | The list of kind-specific caching policies. |
| dockerRegistries | [KubernetesAccountDockerRegistry](#proto.cloudprovider.KubernetesAccountDockerRegistry) | repeated | The list of the Spinnaker docker registry account names this Spinnaker account can use as image sources. These docker registry accounts must be registered in your halconfig before you can add them here. |
| oAuthScopes | [string](#string) | repeated | The list of OAuth scopes used by kubectl to fetch an OAuth token. |
| kubeconfigFile | [string](#string) |  | The path to your kubeconfig file. By default, it will be under the Spinnaker user&#39;s home directory in the typical .kube/config location. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.cloudprovider.KubernetesAccountDockerRegistry"></a>

### KubernetesAccountDockerRegistry
Configuration for a Docker registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | The configured name of the Docker registry. |
| namespaces | [string](#string) | repeated | The list of Docker registry namespaces usable as image sources. |






<a name="proto.cloudprovider.KubernetesCachingPolicy"></a>

### KubernetesCachingPolicy
Configuration for a kind-specific caching policy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind to which the policy applies. |
| maxEntriesPerAgent | [int32](#int32) |  | The maximum number of resources an agent will cache of the specified Kubernetes kind. |






<a name="proto.cloudprovider.KubernetesCustomResource"></a>

### KubernetesCustomResource
Configuration for a CRD to be managed by Spinnaker. If Spinnaker does not
have permission to list CRDs but you need Spinnaker to manage CRDs, you
need to explicitly register each CRD.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind of the custom resource. |
| spinnakerKind | [string](#string) |  | The Spinnaker kind to which you would like the custom resource to map. |
| deployPriority | [string](#string) |  | An integer representing the deployment priority of this resource. Resources with lower values are deployed before resources with higher values. |
| versioned | [bool](#bool) |  | Whether Spinnaker should manage versioning this resource. |
| namespaced | [bool](#bool) |  | Whether the resource is namespaced. |





 

 

 

 



<a name="cloudprovider/oracle.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/oracle.proto



<a name="proto.cloudprovider.Oracle"></a>

### Oracle
Configuration for the Oracle provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [OracleAccount](#proto.cloudprovider.OracleAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [OracleBakeryDefaults](#proto.cloudprovider.OracleBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.cloudprovider.OracleAccount"></a>

### OracleAccount
Configuration for an Oracle account. An account maps to an Oracle Cloud
Infrastructure (OCI) user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated) List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [proto.Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| compartmentId | [string](#string) |  | (Required) The OCID of the Oracle Compartment to use. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| fingerprint | [string](#string) |  | (Required) Fingerprint of the public key. |
| privateKeyPassphrase | [string](#string) |  | Passphrase used for the private key, if it is encrypted. |
| region | [string](#string) |  | (Required) An Oracle region (e.g., `us-phoenix-1`). |
| sshPrivateKeyFilePath | [string](#string) |  | (Required) Path to the private key in PEM format. |
| tenancyId | [string](#string) |  | (Required) The OCID of the Oracle Tenancy to use. |
| userId | [string](#string) |  | (Required) The OCID of the Oracle User with which to authenticate. |






<a name="proto.cloudprovider.OracleBakeryDefaults"></a>

### OracleBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| baseImages | [OracleBaseImageSettings](#proto.cloudprovider.OracleBaseImageSettings) | repeated | List of configured base images. |
| availabilityDomain | [string](#string) |  | (Required) The name of the Availability Domain within which a new instance is launched and provisioned. |
| instanceShape | [string](#string) |  | (Required) The shape for a newly created instance. |
| subnetId | [string](#string) |  | (Required) The name of the subnet within which a new instance is launched and provisioned. |






<a name="proto.cloudprovider.OracleBaseImage"></a>

### OracleBaseImage
Oracle base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The name of the base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker&#39;s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.cloudprovider.OracleBaseImageSettings"></a>

### OracleBaseImageSettings
Configuration for a base image for the Oracle provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [OracleBaseImage](#proto.cloudprovider.OracleBaseImage) |  | Oracle base image configuration. |
| virtualizationSettings | [OracleVirtualizationSettings](#proto.cloudprovider.OracleVirtualizationSettings) |  | Oracle virtualization settings. |






<a name="proto.cloudprovider.OracleVirtualizationSettings"></a>

### OracleVirtualizationSettings
Oracle virtualization settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImageId | [string](#string) |  | (Required) The OCID of the base image ID for the baking configuration. |
| sshUserName | [string](#string) |  | (Required) The ssh username for the baking configuration. |





 

 

 

 



<a name="cloudprovider/providers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cloudprovider/providers.proto



<a name="proto.cloudprovider.Providers"></a>

### Providers
Configuration for cloud provider integrations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [Kubernetes](#proto.cloudprovider.Kubernetes) |  |  |
| google | [GoogleComputeEngine](#proto.cloudprovider.GoogleComputeEngine) |  |  |
| appengine | [Appengine](#proto.cloudprovider.Appengine) |  |  |
| aws | [Aws](#proto.cloudprovider.Aws) |  |  |
| azure | [Azure](#proto.cloudprovider.Azure) |  |  |
| cloudfoundry | [CloudFoundry](#proto.cloudprovider.CloudFoundry) |  |  |
| dcos | [Dcos](#proto.cloudprovider.Dcos) |  |  |
| dockerRegistry | [DockerRegistry](#proto.cloudprovider.DockerRegistry) |  |  |
| ecs | [Ecs](#proto.cloudprovider.Ecs) |  |  |
| huaweicloud | [HuaweiCloud](#proto.cloudprovider.HuaweiCloud) |  |  |
| oracle | [Oracle](#proto.cloudprovider.Oracle) |  |  |





 

 

 

 



<a name="config/clouddriver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/clouddriver.proto



<a name="proto.config.Clouddriver"></a>

### Clouddriver
Configuration for the clouddriver microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [proto.cloudprovider.Kubernetes](#proto.cloudprovider.Kubernetes) |  |  |
| google | [proto.cloudprovider.GoogleComputeEngine](#proto.cloudprovider.GoogleComputeEngine) |  |  |
| appengine | [proto.cloudprovider.Appengine](#proto.cloudprovider.Appengine) |  |  |
| aws | [proto.cloudprovider.Aws](#proto.cloudprovider.Aws) |  |  |
| azure | [proto.cloudprovider.Azure](#proto.cloudprovider.Azure) |  |  |
| cloudfoundry | [proto.cloudprovider.CloudFoundry](#proto.cloudprovider.CloudFoundry) |  |  |
| dcos | [proto.cloudprovider.Dcos](#proto.cloudprovider.Dcos) |  |  |
| dockerRegistry | [proto.cloudprovider.DockerRegistry](#proto.cloudprovider.DockerRegistry) |  |  |
| ecs | [proto.cloudprovider.Ecs](#proto.cloudprovider.Ecs) |  |  |
| huaweicloud | [proto.cloudprovider.HuaweiCloud](#proto.cloudprovider.HuaweiCloud) |  |  |
| oracle | [proto.cloudprovider.Oracle](#proto.cloudprovider.Oracle) |  |  |
| artifacts | [proto.artifact.Artifacts](#proto.artifact.Artifacts) |  |  |





 

 

 

 



<a name="config/deck.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/deck.proto



<a name="proto.config.Deck"></a>

### Deck



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gateUrl | [string](#string) |  | The endpoint at which Deck communicates with Gate. |
| authEnabled | [bool](#bool) |  | Whether authn is enabled. |
| authEndpoint | [string](#string) |  | The Gate authn endpoint. |
| bakeryDetailUrl | [string](#string) |  | Baking details URL used in Bake stage execution details. |
| canary | [Deck.Canary](#proto.config.Deck.Canary) |  | Configuration for the canary UI. |
| changelog | [Deck.Changelog](#proto.config.Deck.Changelog) |  | Configuration for surfacing the Spinnaker changelog in Deck. |
| notifications | [Deck.Notifications](#proto.config.Deck.Notifications) |  | Configuration for notifications providers. |
| providers | [Deck.Providers](#proto.config.Deck.Providers) |  | Configuration for cloud provider defaults. |
| version | [string](#string) |  | Spinnaker version. |
| defaultTimeZone | [string](#string) |  | Default time zone in which to display timestamps in the UI. |
| features | [Deck.Features](#proto.config.Deck.Features) |  | Configuration for UI-related feature flags. |






<a name="proto.config.Deck.Canary"></a>

### Deck.Canary
Configuration for the canary UI.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaultJudge | [string](#string) |  | The default canary judge. Defaults to `NetflixACAJudge-v1.0`, which is currently the only open-source judge available by default. |
| featureDisabled | [bool](#bool) |  | Whether the canary UI is disabled. |
| metricsAccountName | [string](#string) |  | Name of the canary metrics account to use by default. |
| metricStore | [string](#string) |  | Name of the metrics store to use by default (e.g., `prometheus`, `datadog`). |
| showAllConfigs | [bool](#bool) |  | Whether or not to show all canary configs in Deck, or just those scoped to the current application. |
| storageAccountName | [string](#string) |  | Name of storage account to use by default. |
| templatesEnabled | [bool](#bool) |  | Whether or not to enable custom filter templates for canary configs in Deck. |






<a name="proto.config.Deck.Changelog"></a>

### Deck.Changelog
Configuration for surfacing the Spinnaker changelog in Deck.
TODO: can we do one of the following?
1. Change the component in Deck to fall back to just displaying
a link to the spinnaker.io versions page, bookmarked to correct version?
2. Somehow derive the gist ID from the top-level version
without reaching out to the internet?


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fileName | [string](#string) |  | The Spinnaker changelog gist file name. |
| gistId | [string](#string) |  | The Spinnaker changelog gist ID. |






<a name="proto.config.Deck.Features"></a>

### Deck.Features
Configuration for UI-related feature flags.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipelineTemplates | [bool](#bool) |  | Enable pipeline template support. Read more at https://github.com/spinnaker/dcd-spec. |
| canary | [bool](#bool) |  | Enable canary UI support. |
| chaosMonkey | [bool](#bool) |  | Enable Chaos Monkey support. For this to work, you&#39;ll need a running Chaos Monkey deployment. See https://github.com/Netflix/chaosmonkey/wiki. |
| fiatEnabled | [bool](#bool) |  | Whether authz is enabled. |
| roscoMode | [bool](#bool) |  | Whether Rosco UI support is enabled. TODO: handle per comments in feature.proto |
| managedPipelineTemplatesV2UI | [bool](#bool) |  | Enable managed pipeline templates v2 UI support. |






<a name="proto.config.Deck.Notifications"></a>

### Deck.Notifications
Configuration for notifications providers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bearychat | [proto.notification.BearyChat](#proto.notification.BearyChat) |  | Configuration for the BearyChat notification provider. |
| email | [proto.notification.Email](#proto.notification.Email) |  | Configuration for the email notification provider. |
| githubStatus | [proto.notification.GithubStatus](#proto.notification.GithubStatus) |  | Configuration for the GitHub Status notification provider. |
| googleChat | [proto.notification.GoogleChat](#proto.notification.GoogleChat) |  | Configuration for the Google Chat notification provider. |
| pubsub | [proto.notification.PubSub](#proto.notification.PubSub) |  | Configuration for the Pub/Sub notification provider. |
| slack | [proto.notification.Slack](#proto.notification.Slack) |  | Configuration for the Slack notification provider. |
| sms | [proto.notification.Twilio](#proto.notification.Twilio) |  | Configuration for the SMS notification provider. |






<a name="proto.config.Deck.Providers"></a>

### Deck.Providers
UI-specific provider default settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appengine | [Deck.Providers.Appengine](#proto.config.Deck.Providers.Appengine) |  | Appengine provider defaults. |
| aws | [Deck.Providers.Aws](#proto.config.Deck.Providers.Aws) |  | AWS provider defaults. |
| azure | [Deck.Providers.Azure](#proto.config.Deck.Providers.Azure) |  | Azure provider defaults. |
| cloudfoundry | [Deck.Providers.Cloudfoundry](#proto.config.Deck.Providers.Cloudfoundry) |  | Cloudfoundry provider defaults. |
| dcos | [Deck.Providers.Dcos](#proto.config.Deck.Providers.Dcos) |  | DC/OS provider defaults. |
| ecs | [Deck.Providers.Ecs](#proto.config.Deck.Providers.Ecs) |  | ECS provider defaults. |
| gce | [Deck.Providers.Gce](#proto.config.Deck.Providers.Gce) |  | GCE provider defaults. |
| huaweicloud | [Deck.Providers.HuaweiCloud](#proto.config.Deck.Providers.HuaweiCloud) |  | Huawei Cloud provider defaults. |
| tencentcloud | [Deck.Providers.TencentCloud](#proto.config.Deck.Providers.TencentCloud) |  | Tencent Cloud provider defaults. |






<a name="proto.config.Deck.Providers.Appengine"></a>

### Deck.Providers.Appengine



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Appengine.Defaults](#proto.config.Deck.Providers.Appengine.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Appengine.Defaults"></a>

### Deck.Providers.Appengine.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.Aws"></a>

### Deck.Providers.Aws



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Aws.Defaults](#proto.config.Deck.Providers.Aws.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Aws.Defaults"></a>

### Deck.Providers.Aws.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| region | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.Azure"></a>

### Deck.Providers.Azure



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Azure.Defaults](#proto.config.Deck.Providers.Azure.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Azure.Defaults"></a>

### Deck.Providers.Azure.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| region | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.Cloudfoundry"></a>

### Deck.Providers.Cloudfoundry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Cloudfoundry.Defaults](#proto.config.Deck.Providers.Cloudfoundry.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Cloudfoundry.Defaults"></a>

### Deck.Providers.Cloudfoundry.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.Dcos"></a>

### Deck.Providers.Dcos



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Dcos.Defaults](#proto.config.Deck.Providers.Dcos.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Dcos.Defaults"></a>

### Deck.Providers.Dcos.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.Ecs"></a>

### Deck.Providers.Ecs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Ecs.Defaults](#proto.config.Deck.Providers.Ecs.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Ecs.Defaults"></a>

### Deck.Providers.Ecs.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.Gce"></a>

### Deck.Providers.Gce



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.Gce.Defaults](#proto.config.Deck.Providers.Gce.Defaults) |  |  |






<a name="proto.config.Deck.Providers.Gce.Defaults"></a>

### Deck.Providers.Gce.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| region | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.HuaweiCloud"></a>

### Deck.Providers.HuaweiCloud



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.HuaweiCloud.Defaults](#proto.config.Deck.Providers.HuaweiCloud.Defaults) |  |  |






<a name="proto.config.Deck.Providers.HuaweiCloud.Defaults"></a>

### Deck.Providers.HuaweiCloud.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| region | [string](#string) |  |  |






<a name="proto.config.Deck.Providers.TencentCloud"></a>

### Deck.Providers.TencentCloud
TODO: add TencentCloud as a cloud provider so we can write
config to Deck and Clouddriver.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaults | [Deck.Providers.TencentCloud.Defaults](#proto.config.Deck.Providers.TencentCloud.Defaults) |  |  |






<a name="proto.config.Deck.Providers.TencentCloud.Defaults"></a>

### Deck.Providers.TencentCloud.Defaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| region | [string](#string) |  |  |





 

 

 

 



<a name="config/echo.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/echo.proto



<a name="proto.config.Echo"></a>

### Echo
Configuration for the echo microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slack | [proto.notification.Slack](#proto.notification.Slack) |  |  |
| twilio | [proto.notification.Twilio](#proto.notification.Twilio) |  |  |
| githubStatus | [proto.notification.GithubStatus](#proto.notification.GithubStatus) |  |  |
| artifacts | [proto.artifact.Artifacts](#proto.artifact.Artifacts) |  |  |
| pubsub | [proto.pubsub.Pubsub](#proto.pubsub.Pubsub) |  |  |
| gcb | [proto.ci.GoogleCloudBuild](#proto.ci.GoogleCloudBuild) |  |  |
| stats | [proto.Stats](#proto.Stats) |  |  |
| scheduler | [Echo.Scheduler](#proto.config.Echo.Scheduler) |  |  |






<a name="proto.config.Echo.Scheduler"></a>

### Echo.Scheduler
Echo scheduler configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cron | [Echo.Scheduler.Cron](#proto.config.Echo.Scheduler.Cron) |  | Cron configuration. |






<a name="proto.config.Echo.Scheduler.Cron"></a>

### Echo.Scheduler.Cron
Cron configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timezone | [string](#string) |  | Default timezone. Defaults to `America/Los_Angeles`. |





 

 

 

 



<a name="config/fiat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/fiat.proto



<a name="proto.config.Fiat"></a>

### Fiat
Configuration for the fiat microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [proto.security.authz.Authorization](#proto.security.authz.Authorization) |  | Configuration for what resources users of Spinnaker can read and modify. |





 

 

 

 



<a name="config/front50.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/front50.proto



<a name="proto.config.Front50"></a>

### Front50
Configuration for the front50 microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spinnaker | [Front50.Spinnaker](#proto.config.Front50.Spinnaker) |  |  |






<a name="proto.config.Front50.Spinnaker"></a>

### Front50.Spinnaker



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gcs | [proto.storage.Gcs](#proto.storage.Gcs) |  |  |
| azs | [proto.storage.Azs](#proto.storage.Azs) |  |  |
| oracle | [proto.storage.Oracle](#proto.storage.Oracle) |  |  |
| s3 | [proto.storage.S3](#proto.storage.S3) |  |  |





 

 

 

 



<a name="config/gate.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/gate.proto



<a name="proto.config.Cors"></a>

### Cors
Configuration for cross-origin resource sharing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowedOriginsPattern | [string](#string) |  | A regex matching all URLs authentication redirects may come from. |






<a name="proto.config.Gate"></a>

### Gate
Configuration for the gate microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [ServerConfig](#proto.config.ServerConfig) |  | Web server configuration. |
| cors | [Cors](#proto.config.Cors) |  | Configuration for cross-origin resource sharing. |
| security | [SpringSecurity](#proto.config.SpringSecurity) |  | Wrapper for Spring configuration properties (including OAuth2 authentication). |
| saml | [proto.security.authn.Saml](#proto.security.authn.Saml) |  | Configuration for SAML authentication. |
| ldap | [proto.security.authn.Ldap](#proto.security.authn.Ldap) |  | Configuration for LDAP authentication. |
| x509 | [proto.security.authn.X509](#proto.security.authn.X509) |  | Configuration for X509 authentication. |
| google | [Gate.GoogleConfig](#proto.config.Gate.GoogleConfig) |  | Wrapper for Google-specific authentication (ex: IAP). |
| services | [Gate.Services](#proto.config.Gate.Services) |  | Configuration for the status of non-core services. |
| integrations | [Gate.Integrations](#proto.config.Gate.Integrations) |  |  |






<a name="proto.config.Gate.GoogleConfig"></a>

### Gate.GoogleConfig
Wrapper for Google-specific authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iap | [proto.security.authn.Iap](#proto.security.authn.Iap) |  | Configuration for Identity-Aware Proxy authentication. |






<a name="proto.config.Gate.Integrations"></a>

### Gate.Integrations
Wrapper for Gate integrations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gremlin | [Gate.Integrations.Gremlin](#proto.config.Gate.Integrations.Gremlin) |  |  |






<a name="proto.config.Gate.Integrations.Gremlin"></a>

### Gate.Integrations.Gremlin
Configuration for Gremlin fault-injection support.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Gremlin is enabled. |






<a name="proto.config.Gate.Services"></a>

### Gate.Services



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kayenta | [ServiceSettings](#proto.config.ServiceSettings) |  |  |
| deck | [ServiceSettings](#proto.config.ServiceSettings) |  |  |






<a name="proto.config.ServerConfig"></a>

### ServerConfig
Web server configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ssl | [proto.security.ApiSsl](#proto.security.ApiSsl) |  | SSL configuration. |






<a name="proto.config.SpringSecurity"></a>

### SpringSecurity
Wrapper for Spring security configuration properties.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oauth2 | [proto.security.authn.OAuth2](#proto.security.authn.OAuth2) |  | Configuration for OAuth2 authentication. |
| basic | [proto.security.authn.Basic](#proto.security.authn.Basic) |  | Configuration for basic authentication. |





 

 

 

 



<a name="config/halconfig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/halconfig.proto



<a name="proto.config.Hal"></a>

### Hal
Configuration for a Spinnaker installation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStorage | [proto.storage.PersistentStorage](#proto.storage.PersistentStorage) |  |  |
| providers | [proto.cloudprovider.Providers](#proto.cloudprovider.Providers) |  |  |
| artifacts | [proto.artifact.Artifacts](#proto.artifact.Artifacts) |  |  |
| notifications | [proto.notification.Notifications](#proto.notification.Notifications) |  |  |
| pubsub | [proto.pubsub.Pubsub](#proto.pubsub.Pubsub) |  |  |
| ci | [proto.ci.Ci](#proto.ci.Ci) |  |  |
| stats | [proto.Stats](#proto.Stats) |  |  |
| features | [proto.Features](#proto.Features) |  |  |
| webhook | [proto.security.WebhookConfig](#proto.security.WebhookConfig) |  |  |
| security | [proto.security.Security](#proto.security.Security) |  |  |
| canary | [proto.canary.Canary](#proto.canary.Canary) |  |  |
| timezone | [string](#string) |  | The timezone in which your Spinnaker instance runs. This affects what the UI will display as well as how CRON triggers are run. |
| version | [string](#string) |  | Top-level Spinnaker version. |





 

 

 

 



<a name="config/kayenta.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/kayenta.proto



<a name="proto.config.Kayenta"></a>

### Kayenta
Configuration for the Kayenta microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kayenta | [Kayenta.ServiceIntegrations](#proto.config.Kayenta.ServiceIntegrations) |  |  |






<a name="proto.config.Kayenta.ServiceIntegrations"></a>

### Kayenta.ServiceIntegrations



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| google | [Kayenta.ServiceIntegrations.Google](#proto.config.Kayenta.ServiceIntegrations.Google) |  |  |
| stackdriver | [proto.canary.Stackdriver](#proto.canary.Stackdriver) |  |  |
| gcs | [proto.canary.Gcs](#proto.canary.Gcs) |  |  |
| prometheus | [proto.canary.Prometheus](#proto.canary.Prometheus) |  |  |
| datadog | [proto.canary.Datadog](#proto.canary.Datadog) |  |  |
| aws | [Kayenta.ServiceIntegrations.Aws](#proto.config.Kayenta.ServiceIntegrations.Aws) |  |  |
| s3 | [proto.canary.S3](#proto.canary.S3) |  |  |
| signalfx | [proto.canary.SignalFx](#proto.canary.SignalFx) |  |  |
| newrelic | [proto.canary.NewRelic](#proto.canary.NewRelic) |  |  |






<a name="proto.config.Kayenta.ServiceIntegrations.Aws"></a>

### Kayenta.ServiceIntegrations.Aws



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| accounts | [proto.canary.AwsAccount](#proto.canary.AwsAccount) | repeated |  |






<a name="proto.config.Kayenta.ServiceIntegrations.Google"></a>

### Kayenta.ServiceIntegrations.Google



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| accounts | [proto.canary.GoogleAccount](#proto.canary.GoogleAccount) | repeated |  |





 

 

 

 



<a name="config/orca.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/orca.proto



<a name="proto.config.Orca"></a>

### Orca
Configuration for the Orca microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipelineTemplates | [Orca.PipelineTemplates](#proto.config.Orca.PipelineTemplates) |  |  |
| webhook | [proto.security.WebhookConfig](#proto.security.WebhookConfig) |  |  |
| default | [Orca.Defaults](#proto.config.Orca.Defaults) |  |  |
| services | [Orca.Services](#proto.config.Orca.Services) |  |  |
| tasks | [Orca.Tasks](#proto.config.Orca.Tasks) |  |  |






<a name="proto.config.Orca.Defaults"></a>

### Orca.Defaults
Defaults applicable to the orca microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bake | [Orca.Defaults.BakeDefaults](#proto.config.Orca.Defaults.BakeDefaults) |  | Configuration of bakery defaults. |






<a name="proto.config.Orca.Defaults.BakeDefaults"></a>

### Orca.Defaults.BakeDefaults
Configuration of bakery defaults.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | The default account to use for baking. |






<a name="proto.config.Orca.PipelineTemplates"></a>

### Orca.PipelineTemplates
Configuration for pipeline templates.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether pipeline templates are enabled. |






<a name="proto.config.Orca.Services"></a>

### Orca.Services
Configuration for the status of non-core services.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kayenta | [ServiceSettings](#proto.config.ServiceSettings) |  |  |






<a name="proto.config.Orca.Tasks"></a>

### Orca.Tasks
Configuration for Orca tasks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| executionWindow | [Orca.Tasks.ExecutionWindow](#proto.config.Orca.Tasks.ExecutionWindow) |  | Execution window configuration. |






<a name="proto.config.Orca.Tasks.ExecutionWindow"></a>

### Orca.Tasks.ExecutionWindow
Execution window configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timezone | [string](#string) |  | Default timezone. Defaults to `America/Los_Angeles`. |





 

 

 

 



<a name="config/rosco.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/rosco.proto



<a name="proto.config.Rosco"></a>

### Rosco
Configuration for the rosco microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| google | [proto.cloudprovider.GoogleComputeEngine](#proto.cloudprovider.GoogleComputeEngine) |  |  |
| aws | [proto.cloudprovider.Aws](#proto.cloudprovider.Aws) |  |  |
| azure | [proto.cloudprovider.Azure](#proto.cloudprovider.Azure) |  |  |
| huaweicloud | [proto.cloudprovider.HuaweiCloud](#proto.cloudprovider.HuaweiCloud) |  |  |
| oracle | [proto.cloudprovider.Oracle](#proto.cloudprovider.Oracle) |  |  |





 

 

 

 



<a name="config/service_enabled.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/service_enabled.proto



<a name="proto.config.ServiceSettings"></a>

### ServiceSettings
Configuration for a particular microservice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the microservice is enabled. |
| baseUrl | [string](#string) |  | The base URL for the microservice. In general, this should be the externally-resolvable URL for services that are exposed externally (ie, deck and gate). |





 

 

 

 



<a name="config/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config/services.proto



<a name="proto.config.Services"></a>

### Services
Configuration for Spinnaker&#39;s microservices.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clouddriver | [Clouddriver](#proto.config.Clouddriver) |  |  |
| echo | [Echo](#proto.config.Echo) |  |  |
| front50 | [Front50](#proto.config.Front50) |  |  |
| orca | [Orca](#proto.config.Orca) |  |  |
| gate | [Gate](#proto.config.Gate) |  |  |
| fiat | [Fiat](#proto.config.Fiat) |  |  |
| kayenta | [Kayenta](#proto.config.Kayenta) |  |  |
| rosco | [Rosco](#proto.config.Rosco) |  |  |
| deck | [Deck](#proto.config.Deck) |  |  |





 

 

 

 



<a name="features.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## features.proto



<a name="proto.Features"></a>

### Features
Feature flags


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pipelineTemplates | [bool](#bool) |  | Enable pipeline template support. Read more at https://github.com/spinnaker/dcd-spec. |
| mineCanary | [bool](#bool) |  | Enable canary support. For this to work, you&#39;ll need a canary judge configured. |
| chaos | [bool](#bool) |  | Enable Chaos Monkey support. For this to work, you&#39;ll need a running Chaos Monkey deployment. See https://github.com/Netflix/chaosmonkey/wiki. |
| managedPipelineTemplatesV2UI | [bool](#bool) |  | Enable managed pipeline templates v2 UI support. |
| gremlin | [bool](#bool) |  | Enable Gremlin fault-injection support. |
| roscoMode | [bool](#bool) |  | Enable Rosco UI support. TODO: one of the following: 1. Safely default this to true in Deck 2. Document among breaking changes that users will need to manually set this to true. 3. Have Kleat set this to true if there is at least one configured and enabled VM-based cloud provider consumed by Rosco. |





 

 

 

 



<a name="notification/bearychat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/bearychat.proto



<a name="proto.notification.BearyChat"></a>

### BearyChat
Configuration for BearyChat notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether BearyChat notifications are enabled. |





 

 

 

 



<a name="notification/email.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/email.proto



<a name="proto.notification.Email"></a>

### Email
Configuration for Email notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Email notifications are enabled. |





 

 

 

 



<a name="notification/github_status.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/github_status.proto



<a name="proto.notification.GithubStatus"></a>

### GithubStatus
Configuration for Github status notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether GitHub status notifications are enabled. |
| token | [string](#string) |  | Your GitHub account token. |





 

 

 

 



<a name="notification/google_chat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/google_chat.proto



<a name="proto.notification.GoogleChat"></a>

### GoogleChat
Configuration for Google Chat notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Google Chat notifications are enabled. |





 

 

 

 



<a name="notification/notifications.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/notifications.proto



<a name="proto.notification.Notifications"></a>

### Notifications
Configuration for notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slack | [Slack](#proto.notification.Slack) |  |  |
| twilio | [Twilio](#proto.notification.Twilio) |  |  |
| githubStatus | [GithubStatus](#proto.notification.GithubStatus) |  |  |
| bearychat | [BearyChat](#proto.notification.BearyChat) |  |  |
| email | [Email](#proto.notification.Email) |  |  |
| googlechat | [GoogleChat](#proto.notification.GoogleChat) |  |  |
| pubsub | [PubSub](#proto.notification.PubSub) |  |  |





 

 

 

 



<a name="notification/pubsub.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/pubsub.proto



<a name="proto.notification.PubSub"></a>

### PubSub
Configuration for Pub/Sub notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Pub/Sub notifications are enabled. |





 

 

 

 



<a name="notification/slack.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/slack.proto



<a name="proto.notification.Slack"></a>

### Slack
Configuration for Slack notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Slack notifications are enabled. |
| botName | [string](#string) |  | The name of your Slack bot. |
| token | [string](#string) |  | Your Slack bot token. |
| baseUrl | [string](#string) |  | Slack endpoint. Optional, can only be set if using a compatible API. |
| forceUseIncomingWebhook | [bool](#bool) |  | Force usage of incoming webhooks endpoint for Slack. Optional, only set if using a compatible API. |





 

 

 

 



<a name="notification/twilio.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/twilio.proto



<a name="proto.notification.Twilio"></a>

### Twilio
Configuration for Twilio notifications.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Twilio notifications are enabled. |
| account | [string](#string) |  | Your Twilio account SID. |
| token | [string](#string) |  | Your Twilio auth token. |
| baseUrl | [string](#string) |  | The endpoint of the Twilio API. Optional, only set if overriding the default. |
| from | [string](#string) |  | The phone number from which the SMS will be sent (e.g., &#43;1234-567-8910). |





 

 

 

 



<a name="permissions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## permissions.proto



<a name="proto.Permissions"></a>

### Permissions
A Fiat permissions configuration object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| READ | [string](#string) | repeated | A user must have at least one of these roles in order to view this account&#39;s cloud resources. |
| WRITE | [string](#string) | repeated | A user must have at least one of these roles in order to make changes to this account&#39;s cloud resources. |
| EXECUTE | [string](#string) | repeated | A user must have at least one of these roles in order to execute pipelines. |





 

 

 

 



<a name="pubsub/google.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pubsub/google.proto



<a name="proto.pubsub.Google"></a>

### Google
Configuration for Google Cloud Pub/Sub integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Google Cloud Pub/Sub is enabled. |
| subscriptions | [GoogleSubscriber](#proto.pubsub.GoogleSubscriber) | repeated | The list of configured subscriptions. |
| publishers | [GooglePublisher](#proto.pubsub.GooglePublisher) | repeated | The list of configured publishers. |






<a name="proto.pubsub.GooglePublisher"></a>

### GooglePublisher
Configuration for a Google Cloud Pub/Sub publisher.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the publisher account. |
| project | [string](#string) |  | The name of the GCP project your topic lives in. |
| topicName | [string](#string) |  | The name of the topic to publish to. This identifier does not include the name of the project, and must already be configured. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| content | [string](#string) |  | The content to publish to the topic. Must be one of ALL or NOTIFICATIONS. |






<a name="proto.pubsub.GoogleSubscriber"></a>

### GoogleSubscriber
Configuration for a Google Cloud Pub/Sub subscriber.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the subscriber account. |
| project | [string](#string) |  | The name of the GCP project your subscription lives in. |
| subscriptionName | [string](#string) |  | The name of the subscription to listen to. This identifier does not include the name of the project, and must already be configured. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| ackDeadlineSeconds | [int32](#int32) |  | The acknowledgement deadline as configured on the Pub/Sub subscription. |
| messageFormat | [MessageFormat](#proto.pubsub.MessageFormat) |  | The format of the incoming message. Used to translate the incoming message into Spinnaker artifacts. |
| templatePath | [string](#string) |  | A path to a jinja template that specifies how artifacts from this pubsub system are interpreted and transformed into Spinnaker artifacts. Only used if messageFormat is set to CUSTOM. |





 


<a name="proto.pubsub.MessageFormat"></a>

### MessageFormat
Represents the format of an incoming pub/sub message.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CUSTOM | 0 |  |
| GCB | 1 |  |
| GCS | 2 |  |
| GCR | 3 |  |


 

 

 



<a name="pubsub/pubsub.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pubsub/pubsub.proto



<a name="proto.pubsub.Pubsub"></a>

### Pubsub
Configuration for Pub/Sub integration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Pub/Sub is enabled. |
| google | [Google](#proto.pubsub.Google) |  | Configuration for the Google Cloud Pub/Sub integration. |





 

 

 

 



<a name="security/authn/authn.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## security/authn/authn.proto



<a name="proto.security.authn.Authentication"></a>

### Authentication
Configuration of how users authenticate against Spinnaker.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether to enabled authentication. |
| oauth2 | [OAuth2](#proto.security.authn.OAuth2) |  | OAuth 2.0 configuration. |
| saml | [Saml](#proto.security.authn.Saml) |  | SAML configuration. |
| ldap | [Ldap](#proto.security.authn.Ldap) |  | LDAP configuration. |
| x509 | [X509](#proto.security.authn.X509) |  | X509 configuration. |
| iap | [Iap](#proto.security.authn.Iap) |  | Google Cloud Identity-Aware Proxy configuration. |
| basic | [Basic](#proto.security.authn.Basic) |  | Basic username/password authentication. |






<a name="proto.security.authn.Basic"></a>

### Basic
Configuration for basic username/password authentication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the authentication method is enabled. |
| user | [UsernamePassword](#proto.security.authn.UsernamePassword) |  | The username and password used to log in via basic authentication. |






<a name="proto.security.authn.Iap"></a>

### Iap
Configuration for authentication via Google Cloud Identity-Aware Proxy.
Google Cloud Identity-Aware Proxy (IAP) is an authentication model that utilizes
Google OAuth 2.0 and an authorization service to provide access control for users
of GCP. After a user has been authenticated and authorized by IAP&#39;s service, a
JWT token is passed along which Spinnaker uses to check for authenticity and to
get the user email from the payload and sign the user in. To configure IAP, set
the audience field retrieved from the IAP console.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the authentication method is enabled. |
| jwtHeader | [string](#string) |  | The HTTP request header that contains the JWT token. |
| issuerId | [string](#string) |  | The Issuer from the ID token payload. |
| audience | [string](#string) |  | The Audience from the ID token payload. You can retrieve this field from the IAP console: https://cloud.google.com/iap/docs/signed-headers-howto#verify_the_id_token_header. |
| iapVerifyKeyUrl | [string](#string) |  | The URL containing the Cloud IAP public keys in JWK format. |






<a name="proto.security.authn.Ldap"></a>

### Ldap
Configuration for authentication via LDAP.
Lightweight Directory Access Protocol (LDAP) is a standard way many organizations
maintain user credentials and group memberships. Spinnaker uses the standard
&#39;bind&#39; approach for user authentication. This is a fancy way of saying that
Gate uses your username and password to login to the LDAP server, and if the
connection is successful, you&#39;re considered authenticated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the authentication method is enabled. |
| url | [string](#string) |  | ldap:// or ldaps:// url of the LDAP server. |
| userDnPattern | [string](#string) |  | The pattern for finding a user&#39;s DN using simple pattern matching. For example, if your LDAP server has the URL ldap://mysite.com/dc=spinnaker,dc=org, and you have the pattern &#39;uid={0},ou=members&#39;, &#39;me&#39; will map to a DN uid=me,ou=members,dc=spinnaker,dc=org. If no match is found, will try to find the user using user-search-filter, if set. |
| userSearchBase | [string](#string) |  | The part of the directory tree under which user searches should be performed. If user-search-base isn&#39;t supplied, the search will be performed from the root. |
| userSearchFilter | [string](#string) |  | The filter to use when searching for a user&#39;s DN. Will search either from user-search-base (if specified) or root for entires matching the filter, then attempt to bind as that user with the login password. For example, the filter &#39;uid={0}&#39; would apply to any user where uid matched the user&#39;s login name. If -user-dn-pattern is also specified, will attempt to find a match using the specified pattern first, before searching with the specified search filter if no match is found from the pattern. |
| managerDn | [string](#string) |  | An LDAP manager user is required for binding to the LDAP server for the user authentication process. This property refers to the DN of that entry. I.e. this is not the user which will be authenticated when logging into DHIS2, rather the user which binds to the LDAP server in order to do the authentication. |
| managerPassword | [string](#string) |  | The password for the LDAP manager user. |
| groupSearchBase | [string](#string) |  | The part of the directory tree under which group searches should be performed. |






<a name="proto.security.authn.OAuth2"></a>

### OAuth2
Configuration for authentication via OAuth 2.0.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the authentication method is enabled. |
| client | [OAuth2Client](#proto.security.authn.OAuth2Client) |  | Configuration for your OAuth 2.0 client. |
| userInfoRequirements | [OAuth2.UserInfoRequirementsEntry](#proto.security.authn.OAuth2.UserInfoRequirementsEntry) | repeated | The map of requirements the userInfo request must have. This is used to restrict user login to specific domains or having a specific attribute. |
| resource | [OAuth2Resource](#proto.security.authn.OAuth2Resource) |  | Configuration for OAuth 2.0 resources. |
| userInfoMapping | [OAuth2UserInfoMapping](#proto.security.authn.OAuth2UserInfoMapping) |  | Mapping of user attributes to fields returned by your OAuth 2.0 provider. This field controls how the fields returned from the OAuth 2.0 provider&#39;s user info endpoint are translated into a Spinnaker user. |
| provider | [OAuth2.OAuth2Provider](#proto.security.authn.OAuth2.OAuth2Provider) |  | The OAuth 2.0 provider handling authentication. |






<a name="proto.security.authn.OAuth2.UserInfoRequirementsEntry"></a>

### OAuth2.UserInfoRequirementsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="proto.security.authn.OAuth2Client"></a>

### OAuth2Client
Configuration for an OAuth 2.0 client.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clientId | [string](#string) |  | The OAuth client ID you have configured with your OAuth 2.0 provider. |
| clientSecret | [string](#string) |  | The OAuth client secret you have configured with your OAuth provider. |
| accessTokenUri | [string](#string) |  | The access token uri for your OAuth provider. |
| userAuthorizationUri | [string](#string) |  | The user authorization uri for your OAuth 2.0 provider. |
| clientAuthenticationScheme | [OAuth2Client.AuthenticationScheme](#proto.security.authn.OAuth2Client.AuthenticationScheme) |  | The method used to transmit authentication credentials to your OAuth 2.0 provider. |
| scope | [string](#string) |  | The scope to request when obtaining an access token from your OAuth 2.0 provider. |
| preEstablishedRedirectUri | [string](#string) |  | The externally accessible URL for Gate. For use with load balancers that do any kind of address manipulation for Gate traffic, such as an SSL terminating load balancer. |
| useCurrentUri | [bool](#bool) |  | Whether the current URI in the request should be preferred over the pre-established redirect URI. |






<a name="proto.security.authn.OAuth2Resource"></a>

### OAuth2Resource
Configuration for OAuth 2.0 resources.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userInfoUri | [string](#string) |  | The user info URI for your OAuth 2.0 provider. |






<a name="proto.security.authn.OAuth2UserInfoMapping"></a>

### OAuth2UserInfoMapping
Mapping of user attributes to fields returned by an OAuth 2.0 provider.
This field controls how the fields returned from the OAuth 2.0 provider&#39;s user
info endpoint are translated into a Spinnaker user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email. |
| firstName | [string](#string) |  | First name. |
| lastName | [string](#string) |  | Last name. |
| username | [string](#string) |  | Username. |






<a name="proto.security.authn.Saml"></a>

### Saml
Configuration for authentication via SAML.
SAML authenticates users by passing cryptographically signed XML documents
between the Gate server and an identity provider. Gate&#39;s key is stored and
accessed via the -keystore parameters, while the identity provider&#39;s keys are
included in the metadata.xml. Finally, the identity provider must redirect the
control flow (through the user&#39;s browser) back to Gate by way of the
-serviceAddressUrl. This is likely the address of Gate&#39;s load balancer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the authentication method is enabled. |
| metadataUrl | [string](#string) |  | The path to a local file containing identity provider&#39;s metadata XML file; can be either a local file or a URI. |
| issuerId | [string](#string) |  | The identity of the Spinnaker application registered with the SAML provider. |
| keyStore | [string](#string) |  | Path to the keystore that contains this server&#39;s private key. This key is used to cryptographically sign SAML AuthNRequest objects. |
| keyStorePassword | [string](#string) |  | The password used to access the file specified in -keystore. |
| keyStoreAliasName | [string](#string) |  | The name of the alias under which this server&#39;s private key is stored in the -keystore file. |
| redirectHostname | [string](#string) |  | The host name of the gate server as accessible by the SAML identity provider. If deployed behind a load balancer, this would be the load balancer&#39;s address. (Ex: gate.org.com:8084) |
| redirectBasePath | [string](#string) |  | The base path on the gate server to which redirects will be sent. Defaults to &#39;/&#39; if absent. |
| redirectProtocol | [string](#string) |  | The protocol to use to when redirecting back to the Gate server. Defaults to &#39;https&#39; if absent. |
| userAttributeMapping | [Saml.UserAttributes](#proto.security.authn.Saml.UserAttributes) |  | Configuration for fields returned from your SAML provider. |






<a name="proto.security.authn.Saml.UserAttributes"></a>

### Saml.UserAttributes
Configuration for fields returned from your SAML provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firstName | [string](#string) |  | First name. |
| lastName | [string](#string) |  | Last name. |
| roles | [string](#string) |  | Roles. |
| rolesDelimiter | [string](#string) |  | Roles delimiter. |
| username | [string](#string) |  | Username. |
| email | [string](#string) |  | Email. |






<a name="proto.security.authn.UsernamePassword"></a>

### UsernamePassword
Configuration for a username/password combination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | Username. |
| password | [string](#string) |  | Password. |






<a name="proto.security.authn.X509"></a>

### X509
Configuration for authentication via X509 certificates.
X509 authenticates users via client certificate and a corresponding private key.
These certificates optionally provide authorization information via custom OIDs
with corresponding group information for the user. This can be configured via -roleOid.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the authentication method is enabled. |
| roleOid | [string](#string) |  | The OID that encodes roles that the user specified in the x509 certificate belongs to. |
| subjectPrincipalRegex | [string](#string) |  | The regex used to parse the subject principal name embedded in the x509 certificate if necessary. |





 


<a name="proto.security.authn.OAuth2.OAuth2Provider"></a>

### OAuth2.OAuth2Provider
Supported OAuth 2.0 providers.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OTHER | 0 | Other OAuth 2.0 provider. |
| AZURE | 1 | Azure OAuth 2.0 provider. |
| GITHUB | 2 | Github OAuth 2.0 provider. |
| ORACLE | 3 | Oracle OAuth 2.0 provider. |
| GOOGLE | 4 | Google OAuth 2.0 provider. |



<a name="proto.security.authn.OAuth2Client.AuthenticationScheme"></a>

### OAuth2Client.AuthenticationScheme
Methods to transmit authentication tokens to an OAuth 2.0 provider.

| Name | Number | Description |
| ---- | ------ | ----------- |
| header | 0 | Token is sent in the request header. |
| query | 1 | Token is sent as a query parameter. |
| form | 2 | Token is sent in the form body. |
| none | 3 | Token is not sent at all. |


 

 

 



<a name="security/authz/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## security/authz/authz.proto



<a name="proto.security.authz.Authorization"></a>

### Authorization
Configuration for what resources users of Spinnaker can read and modify.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Spinnaker&#39;s role-based authorization is enabled. |
| groupMembership | [GroupMembership](#proto.security.authz.GroupMembership) |  | Configuration role providers that map users to groups. |






<a name="proto.security.authz.FileRoleProvider"></a>

### FileRoleProvider
Configuration for the file-based role provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | A path to a file describing the roles of each user. |






<a name="proto.security.authz.GithubRoleProvider"></a>

### GithubRoleProvider
Configuration for the GitHub role provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseUrl | [string](#string) |  | Used if using GitHub enterprise some other non github.com GitHub installation. |
| accessToken | [string](#string) |  | A personal access token of an account with access to your organization&#39;s GitHub Teams structure. |
| organization | [string](#string) |  | The GitHub organization under which to query for GitHub Teams. |






<a name="proto.security.authz.GoogleRoleProvider"></a>

### GoogleRoleProvider
Configuration for the Google role provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentialPath | [string](#string) |  | A path to a valid json service account that can authenticate against the Google role provider. |
| adminUsername | [string](#string) |  | Your role provider&#39;s admin username e.g. admin@myorg.net. |
| domain | [string](#string) |  | The domain your role provider is configured for e.g. myorg.net. |






<a name="proto.security.authz.GroupMembership"></a>

### GroupMembership
Configuration role providers that map users to groups.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [GroupMembership.RoleProviderType](#proto.security.authz.GroupMembership.RoleProviderType) |  | Configuration for which role provider to use for authorization decisions. Each role provider has a corresponding field; configuration specific to the role provider you are using should be added to the appropriate field. |
| google | [GoogleRoleProvider](#proto.security.authz.GoogleRoleProvider) |  | Configuration for the Google role provider. |
| github | [GithubRoleProvider](#proto.security.authz.GithubRoleProvider) |  | Configuration for the GitHub role provider. |
| file | [FileRoleProvider](#proto.security.authz.FileRoleProvider) |  | Configuration for the file-based role provider. |
| ldap | [LdapRoleProvider](#proto.security.authz.LdapRoleProvider) |  | Configuration for the LDAP role provider. |






<a name="proto.security.authz.LdapRoleProvider"></a>

### LdapRoleProvider
Configuration for the LDAP role provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | ldap:// or ldaps:// url of the LDAP server. |
| managerDn | [string](#string) |  | The manager user&#39;s distinguished name (principal) to use for querying LDAP groups. |
| managerPassword | [string](#string) |  | The manager user&#39;s password to use for querying LDAP groups. |
| userDnPattern | [string](#string) |  | The pattern for finding a user&#39;s DN using simple pattern matching. For example, if your LDAP server has the URL ldap://mysite.com/dc=spinnaker,dc=org, and you have the pattern &#39;uid={0},ou=members&#39;, &#39;me&#39; will map to a DN uid=me,ou=members,dc=spinnaker,dc=org. If no match is found, will try to find the user using -user-search-filter, if set. |
| userSearchBase | [string](#string) |  | The part of the directory tree under which user searches should be performed. If -user-search-base isn&#39;t supplied, the search will be performed from the root. |
| groupSearchBase | [string](#string) |  | The part of the directory tree under which group searches should be performed. |
| userSearchFilter | [string](#string) |  | The filter to use when searching for a user&#39;s DN. Will search either from -user-search-base (if specified) or root for entries matching the filter. |
| groupSearchFilter | [string](#string) |  | The filter which is used to search for group membership. The default is &#39;uniqueMember={0}&#39;, corresponding to the groupOfUniqueMembers LDAP class. In this case, the substituted parameter is the full distinguished name of the user. The parameter &#39;{1}&#39; can be used if you want to filter on the login name. |
| groupRoleAttributes | [string](#string) |  | The attribute which contains the name of the authority defined by the group entry. Defaults to &#39;cn&#39;. |





 


<a name="proto.security.authz.GroupMembership.RoleProviderType"></a>

### GroupMembership.RoleProviderType
Configuration for which role provider to use for authorization decisions.

| Name | Number | Description |
| ---- | ------ | ----------- |
| EXTERNAL | 0 | External role provider. |
| FILE | 1 | File-based role provider. |
| GOOGLE | 2 | Google role provider. |
| GITHUB | 3 | GitHub role provider. |
| LDAP | 4 | LDAP role provider. |


 

 

 



<a name="security/security.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## security/security.proto



<a name="proto.security.Security"></a>

### Security
Configuration for security settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiSecurity | [ApiSecurity](#proto.security.ApiSecurity) |  | Configuration for the API server&#39;s addressable URL and CORS policies. |
| uiSecurity | [UiSecurity](#proto.security.UiSecurity) |  | Configuration for the UI server&#39;s addressable URL. |
| authn | [authn.Authentication](#proto.security.authn.Authentication) |  | Configuration of how users authenticate against Spinnaker. |
| authz | [authz.Authorization](#proto.security.authz.Authorization) |  | Configuration for what resources users of Spinnaker can read and modify. |





 

 

 

 



<a name="security/ssl.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## security/ssl.proto



<a name="proto.security.ApiSecurity"></a>

### ApiSecurity
Configuration for the API server&#39;s addressable URL and CORS policies.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| corsAccessPattern | [string](#string) |  | If you have authentication enabled, are accessing Spinnaker remotely, and are logging in from sources other than the UI, provide a regex matching all URLs authentication redirects may come from. |
| ssl | [ApiSsl](#proto.security.ApiSsl) |  | If you want the API server to do SSL termination, it must be enabled and configured here. If you are doing your own SSL termination, leave this disabled. |
| overrideBaseUrl | [string](#string) |  | If you are accessing the API server remotely, provide the full base URL of whatever proxy or load balancer is fronting the API requests |






<a name="proto.security.ApiSsl"></a>

### ApiSsl
Configuration for SSL termination by the API server.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether SSL is enabled. |
| keyAlias | [string](#string) |  | Name of your keystore entry as generated with your keytool. |
| keyStore | [string](#string) |  | Path to the keystore holding your security certificates. |
| keyStoreType | [string](#string) |  | The type of your keystore. Examples include JKS, and PKCS12. |
| keyStorePassword | [string](#string) |  | The password to unlock your keystore. Due to a limitation in Tomcat, this must match your key&#39;s password in the keystore. |
| trustStore | [string](#string) |  | Path to the truststore holding your trusted certificates. |
| trustStoreType | [string](#string) |  | The type of your truststore. Examples include JKS, and PKCS12. |
| trustStorePassword | [string](#string) |  | The password to unlock your truststore. |
| clientAuth | [ClientAuth](#proto.security.ClientAuth) |  | Whether to require or allow client authentication. |






<a name="proto.security.UiSecurity"></a>

### UiSecurity
Configuration for the UI server&#39;s addressable URL.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ssl | [UiSsl](#proto.security.UiSsl) |  | Configuration for SSL termination by the UI gateway. |
| overrideBaseUrl | [string](#string) |  | If you are accessing the UI server remotely, provide the full base URL of whatever proxy or load balancer is fronting the UI requests. |






<a name="proto.security.UiSsl"></a>

### UiSsl
Configuration for SSL termination by the UI gateway.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether SSL is enabled. |
| sslCertificateFile | [string](#string) |  | Path to your .crt file. |
| sslCertificateKeyFile | [string](#string) |  | Path to your .key file. |
| sslCACertificateFile | [string](#string) |  | Path to the .crt file for the CA that issued your SSL certificate. This is only needed for local git deployments that serve the UI using webpack dev server. |
| sslCertificatePassphrase | [string](#string) |  | The passphrase needed to unlock your SSL certificate. This will be provided to Apache on startup. |





 


<a name="proto.security.ClientAuth"></a>

### ClientAuth
Setting for client authentication.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | No client authentication. |
| WANT | 1 | Client authentication is optional. |
| NEED | 2 | Client authentication is required. |


 

 

 



<a name="security/trust_store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## security/trust_store.proto



<a name="proto.security.TrustStore"></a>

### TrustStore
Configuration for a custom trust store.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether this custom trust store is enabled. |
| trustStore | [string](#string) |  | The path to a key store in JKS format containing certification authorities that should be trusted. |
| trustStorePassword | [string](#string) |  | The password for the supplied trustStore. |






<a name="proto.security.WebhookConfig"></a>

### WebhookConfig
Configuration for webhooks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trust | [TrustStore](#proto.security.TrustStore) |  | A custom trust store to use for outgoing webhook connections. |





 

 

 

 



<a name="stats.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stats.proto



<a name="proto.Stats"></a>

### Stats
Configuration for optional collection of usage metrics.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Whether to send usage metrics. |





 

 

 

 



<a name="storage/azs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/azs.proto



<a name="proto.storage.Azs"></a>

### Azs
Configuration for an Azure Storage persistent store.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether this persistent store is enabled. |
| storageAccountName | [string](#string) |  | The name of an Azure Storage Account. |
| storageAccountKey | [string](#string) |  | The key to access the Azure Storage Account. |
| storageContainerName | [string](#string) |  | The container name in the chosen storage account to place Spinnaker&#39;s persistent data. Defaults to &#39;spinnaker&#39; if unspecified. |





 

 

 

 



<a name="storage/gcs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/gcs.proto



<a name="proto.storage.Gcs"></a>

### Gcs
Configuration for a Google Cloud Storage persistent store


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether this persistent store is enabled. |
| jsonPath | [string](#string) |  | A path to a JSON service account with permission to read and write to the bucket to be used as a backing store. |
| project | [string](#string) |  | The Google Cloud Platform project you are using to host the GCS bucket as a backing store. |
| bucket | [string](#string) |  | The name of a storage bucket that your specified account has access to. |
| rootFolder | [string](#string) |  | The root folder in the chosen bucket to place all of Spinnaker&#39;s persistent data in. |
| bucketLocation | [string](#string) |  | This is only required if the bucket you specify does not exist yet. |





 

 

 

 



<a name="storage/oracle.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/oracle.proto



<a name="proto.storage.Oracle"></a>

### Oracle
Configuration for an Oracle persistent store.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether this persistent store is enabled. |
| bucketName | [string](#string) |  | The bucket name to store persistent state object in. |
| namespace | [string](#string) |  | The namespace the bucket and objects should be created in. |
| region | [string](#string) |  | An Oracle region (e.g., us-phoenix-1). |
| userId | [string](#string) |  | The OCID of the Oracle User you&#39;re authenticating as. |
| fingerprint | [string](#string) |  | Fingerprint of the public key. |
| sshPrivateKeyFilePath | [string](#string) |  | Path to the private key in PEM format. |
| privateKeyPassphrase | [string](#string) |  | Passphrase used for the private key, if it is encrypted. |
| tenancyId | [string](#string) |  | The OCID of the Oracle Tenancy to use. |





 

 

 

 



<a name="storage/persistent_storage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/persistent_storage.proto



<a name="proto.storage.PersistentStorage"></a>

### PersistentStorage
Configuration of Spinnaker&#39;s persistent storage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gcs | [Gcs](#proto.storage.Gcs) |  |  |
| azs | [Azs](#proto.storage.Azs) |  |  |
| oracle | [Oracle](#proto.storage.Oracle) |  |  |
| s3 | [S3](#proto.storage.S3) |  |  |





 

 

 

 



<a name="storage/s3.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/s3.proto



<a name="proto.storage.S3"></a>

### S3
Configuration for an Amazon S3 persistent store.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether this persistent store is enabled. |
| bucket | [string](#string) |  | The name of a storage bucket that your specified account has access to. |
| rootFolder | [string](#string) |  | The root folder in the chosen bucket to place all of Spinnaker&#39;s persistent data in. |
| region | [string](#string) |  | This is only required if the bucket you specify doesn&#39;t exist yet. In that case, the bucket will be created in that region. See http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region. |
| pathStyleAccess | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | When true, use path-style to access bucket; when false, use virtual hosted-style to access bucket. See https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingExamples. |
| endpoint | [string](#string) |  | An alternate endpoint that your S3-compatible storage can be found at. This is intended for self-hosted storage services with S3-compatible APIs, e.g. Minio. |
| accessKeyId | [string](#string) |  | Your AWS Access Key ID. If not provided, Spinnaker will try to find AWS credentials as described at http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default |
| serverSideEncryption | [S3ServerSideEncryption](#proto.storage.S3ServerSideEncryption) |  | Configuration for S3 server-size encryption. |
| secretAccessKey | [string](#string) |  | Your AWS Secret Key. |





 


<a name="proto.storage.S3ServerSideEncryption"></a>

### S3ServerSideEncryption
Configuration for S3 server-side encryption; values correspond to values of
the &#39;x-amz-server-side-encryption&#39; header.

| Name | Number | Description |
| ---- | ------ | ----------- |
| AES256 | 0 | Amazon S3-managed encryption keys, equivalent to a header value of &#39;AES256&#39;. |
| AWSKMS | 1 | AWS KMS-managed encryption keys, equivalent to a header value of &#39;aws:kms&#39;. |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

