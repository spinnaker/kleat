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



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| accounts | [Kubernetes.Account](#proto.Kubernetes.Account) | repeated |  |
| primaryAccount | [string](#string) |  |  |






<a name="proto.Kubernetes.Account"></a>

### Kubernetes.Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| providerVersion | [string](#string) |  |  |
| kinds | [string](#string) | repeated |  |
| omitKinds | [string](#string) | repeated |  |
| context | [string](#string) |  |  |
| configureImagePullSecrets | [bool](#bool) |  |  |
| cacheThreads | [int32](#int32) |  |  |
| namespaces | [string](#string) | repeated |  |
| omitNamespaces | [string](#string) | repeated |  |
| customResources | [Kubernetes.CustomResource](#proto.Kubernetes.CustomResource) | repeated |  |
| cachingPolicies | [Kubernetes.CachingPolicy](#proto.Kubernetes.CachingPolicy) | repeated |  |
| dockerRegistries | [Kubernetes.DockerRegistry](#proto.Kubernetes.DockerRegistry) | repeated |  |
| oAuthScopes | [string](#string) | repeated |  |
| kubeconfigFile | [string](#string) |  |  |
| permissions | [Permissions](#proto.Permissions) |  |  |
| requiredGroupMemberships | [string](#string) | repeated |  |






<a name="proto.Kubernetes.CachingPolicy"></a>

### Kubernetes.CachingPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  |  |
| maxEntriesPerAgent | [int32](#int32) |  |  |






<a name="proto.Kubernetes.CustomResource"></a>

### Kubernetes.CustomResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  |  |
| spinnakerKind | [string](#string) |  |  |
| deployPriority | [string](#string) |  |  |
| versioned | [bool](#bool) |  |  |
| namespaced | [bool](#bool) |  |  |






<a name="proto.Kubernetes.DockerRegistry"></a>

### Kubernetes.DockerRegistry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  |  |
| namespaces | [string](#string) | repeated |  |






<a name="proto.Permissions"></a>

### Permissions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| READ | [string](#string) | repeated |  |
| WRITE | [string](#string) | repeated |  |
| EXECUTE | [string](#string) | repeated |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

