# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [clouddriver.proto](#clouddriver.proto)
    - [ClouddriverConfig](#proto.ClouddriverConfig)
  
  
  
  

- [front50.proto](#front50.proto)
    - [Front50Config](#proto.Front50Config)
    - [Front50Config.Spinnaker](#proto.Front50Config.Spinnaker)
  
  
  
  

- [halconfig.proto](#halconfig.proto)
    - [HalConfig](#proto.HalConfig)
    - [HalConfig.PersistentStorage](#proto.HalConfig.PersistentStorage)
    - [HalConfig.Providers](#proto.HalConfig.Providers)
  
  
  
  

- [persistentstorage.proto](#persistentstorage.proto)
    - [GCS](#proto.GCS)
  
  
  
  

- [providers.proto](#providers.proto)
    - [Google](#proto.Google)
    - [Google.Account](#proto.Google.Account)
    - [Google.BakeryDefaults](#proto.Google.BakeryDefaults)
    - [Google.BakeryDefaults.BaseImageSettings](#proto.Google.BakeryDefaults.BaseImageSettings)
    - [Google.BakeryDefaults.BaseImageSettings.BaseImage](#proto.Google.BakeryDefaults.BaseImageSettings.BaseImage)
    - [Google.BakeryDefaults.BaseImageSettings.VirtualizationSettings](#proto.Google.BakeryDefaults.BaseImageSettings.VirtualizationSettings)
    - [Google.Consul](#proto.Google.Consul)
    - [Kubernetes](#proto.Kubernetes)
    - [Kubernetes.Account](#proto.Kubernetes.Account)
    - [Kubernetes.CachingPolicy](#proto.Kubernetes.CachingPolicy)
    - [Kubernetes.CustomResource](#proto.Kubernetes.CustomResource)
    - [Kubernetes.DockerRegistry](#proto.Kubernetes.DockerRegistry)
    - [Permissions](#proto.Permissions)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="clouddriver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## clouddriver.proto



<a name="proto.ClouddriverConfig"></a>

### ClouddriverConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [Kubernetes](#proto.Kubernetes) |  |  |
| google | [Google](#proto.Google) |  |  |





 

 

 

 



<a name="front50.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## front50.proto



<a name="proto.Front50Config"></a>

### Front50Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spinnaker | [Front50Config.Spinnaker](#proto.Front50Config.Spinnaker) |  |  |






<a name="proto.Front50Config.Spinnaker"></a>

### Front50Config.Spinnaker



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStoreType | [string](#string) |  |  |
| gcs | [GCS](#proto.GCS) |  |  |





 

 

 

 



<a name="halconfig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## halconfig.proto



<a name="proto.HalConfig"></a>

### HalConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStorage | [HalConfig.PersistentStorage](#proto.HalConfig.PersistentStorage) |  |  |
| providers | [HalConfig.Providers](#proto.HalConfig.Providers) |  |  |






<a name="proto.HalConfig.PersistentStorage"></a>

### HalConfig.PersistentStorage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persistentStoreType | [string](#string) |  |  |
| gcs | [GCS](#proto.GCS) |  |  |






<a name="proto.HalConfig.Providers"></a>

### HalConfig.Providers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [Kubernetes](#proto.Kubernetes) |  |  |
| google | [Google](#proto.Google) |  |  |





 

 

 

 



<a name="persistentstorage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## persistentstorage.proto



<a name="proto.GCS"></a>

### GCS



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| jsonPath | [string](#string) |  | A path to a JSON service account with permission to read and write to the bucket to be used as a backing store. |
| project | [string](#string) |  | The Google Cloud Platform project you are using to host the GCS bucket as a backing store. |
| bucket | [string](#string) |  | The name of a storage bucket that your specified account has access to. If not specified, a random name will be chosen. If you specify a globally unique bucket name that doesn’t exist yet, Halyard will create that bucket for you. |
| rootFolder | [string](#string) |  | The root folder in the chosen bucket to place all of Spinnaker’s persistent data in. |
| bucketLocation | [string](#string) |  | This is only required if the bucket you specify doesn’t exist yet. |





 

 

 

 



<a name="providers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## providers.proto



<a name="proto.Google"></a>

### Google



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| accounts | [Google.Account](#proto.Google.Account) | repeated |  |
| primaryAccount | [string](#string) |  |  |
| bakeryDefaults | [Google.BakeryDefaults](#proto.Google.BakeryDefaults) |  |  |
| requiredGroupMemberships | [string](#string) | repeated |  |
| permissions | [Permissions](#proto.Permissions) |  |  |
| project | [string](#string) |  |  |
| jsonPath | [string](#string) |  |  |
| alphaListed | [bool](#bool) |  |  |
| imageProjects | [string](#string) | repeated |  |
| consul | [Google.Consul](#proto.Google.Consul) |  |  |






<a name="proto.Google.Account"></a>

### Google.Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="proto.Google.BakeryDefaults"></a>

### Google.BakeryDefaults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateFile | [string](#string) |  |  |
| baseImages | [Google.BakeryDefaults.BaseImageSettings](#proto.Google.BakeryDefaults.BaseImageSettings) | repeated |  |
| zone | [string](#string) |  |  |
| network | [string](#string) |  |  |
| useInternalIp | [bool](#bool) |  |  |






<a name="proto.Google.BakeryDefaults.BaseImageSettings"></a>

### Google.BakeryDefaults.BaseImageSettings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImages | [Google.BakeryDefaults.BaseImageSettings.BaseImage](#proto.Google.BakeryDefaults.BaseImageSettings.BaseImage) |  |  |
| virtualizationSettings | [Google.BakeryDefaults.BaseImageSettings.VirtualizationSettings](#proto.Google.BakeryDefaults.BaseImageSettings.VirtualizationSettings) |  |  |






<a name="proto.Google.BakeryDefaults.BaseImageSettings.BaseImage"></a>

### Google.BakeryDefaults.BaseImageSettings.BaseImage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| shortDescription | [string](#string) |  |  |
| detailedDescription | [string](#string) |  |  |
| packageType | [string](#string) |  |  |
| imageFamily | [bool](#bool) |  |  |






<a name="proto.Google.BakeryDefaults.BaseImageSettings.VirtualizationSettings"></a>

### Google.BakeryDefaults.BaseImageSettings.VirtualizationSettings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceImageFamily | [string](#string) |  |  |






<a name="proto.Google.Consul"></a>

### Google.Consul



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| agentEndpoint | [string](#string) |  |  |
| agentPort | [int32](#int32) |  |  |
| datacenters | [string](#string) | repeated |  |






<a name="proto.Kubernetes"></a>

### Kubernetes
Configuration for the Kubernetes provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [Kubernetes.Account](#proto.Kubernetes.Account) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.Kubernetes.Account"></a>

### Kubernetes.Account
Configuration for a Spinnaker Kubernetes account. An account maps to a
credential that can authenticate against your Kubernetes cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| providerVersion | [string](#string) |  | Some providers support multiple versions/release tracks. This allows you to pick the version of the provider (not the resources it manages) to run within Spinnaker. |
| kinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account can deploy to and will cache. When no kinds are configured, this defaults to all kinds described [here](https://spinnaker.io/reference/providers/kubernetes-v2/). This can only be set when omitKinds is empty or not set. |
| omitKinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account cannot deploy to or cache. This can only be set when kinds is empty or not set. |
| context | [string](#string) |  | The kubernetes context to be managed by Spinnaker. See http://kubernetes.io/docs/user-guide/kubeconfig-file/#context for more information. When no context is configured for an account the ‘current-context’ in your kubeconfig is assumed. |
| cacheThreads | [int32](#int32) |  | Number of caching agents for this kubernetes account. Each agent handles a subset of the namespaces available to this account. By default, only 1 agent caches all kinds for all namespaces in the account. |
| namespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account can deploy to and will cache. When no namespaces are configured, this defaults to ‘all namespaces’. |
| omitNamespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account cannot deploy to or cache. This can only be set when –namespaces is empty or not set. |
| customResources | [Kubernetes.CustomResource](#proto.Kubernetes.CustomResource) | repeated | The list of custom resources Clouddriver will manage and make available for use in Patch and Delete (Manifest) stages. |
| cachingPolicies | [Kubernetes.CachingPolicy](#proto.Kubernetes.CachingPolicy) | repeated | The list of kind-specific caching policies. |
| dockerRegistries | [Kubernetes.DockerRegistry](#proto.Kubernetes.DockerRegistry) | repeated | The list of the Spinnaker docker registry account names this Spinnaker account can use as image sources. These docker registry accounts must be registered in your halconfig before you can add them here. |
| oAuthScopes | [string](#string) | repeated | The list of OAuth scopes used by kubectl to fetch an OAuth token. |
| kubeconfigFile | [string](#string) |  | The path to your kubeconfig file. By default, it will be under the Spinnaker user’s home directory in the typical .kube/config location. todo: document new var/secrets convention. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.Kubernetes.CachingPolicy"></a>

### Kubernetes.CachingPolicy
Configuration for a kind-specific caching policy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind to which the policy applies. |
| maxEntriesPerAgent | [int32](#int32) |  | The maximum number of resources an agent will cache of the specified Kubernetes kind. |






<a name="proto.Kubernetes.CustomResource"></a>

### Kubernetes.CustomResource
Configuration for a CRD to be managed by Spinnaker. If Spinnaker does not
have permission to list a CRD but you need Spinnaker to manage it, you
need to explicitly register it.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind of the custom resource. |
| spinnakerKind | [string](#string) |  | The Spinnaker kind to which you would like the custom resource to map. |
| deployPriority | [string](#string) |  | An integer representing the deployment priority of this resource. Resources with lower values are deployed before resources with higher values. |
| versioned | [bool](#bool) |  | Whether Spinnaker should manage versioning this resource. |
| namespaced | [bool](#bool) |  | Whether the CRD is namespaced. |






<a name="proto.Kubernetes.DockerRegistry"></a>

### Kubernetes.DockerRegistry
Configuration for a Docker registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | The configured name of the Docker registry. |
| namespaces | [string](#string) | repeated | The list of Docker registry namespaces usable as image sources. |






<a name="proto.Permissions"></a>

### Permissions
A Fiat permissions configuration object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| READ | [string](#string) | repeated | A user must have at least one of these roles in order to view this account’s cloud resources. |
| WRITE | [string](#string) | repeated | A user must have at least one of these roles in order to make changes to this account’s cloud resources. |
| EXECUTE | [string](#string) | repeated | A user must have at least one of these roles in order to execute pipelines. |





 

 

 

 



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

