# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [clouddriver.proto](#clouddriver.proto)
    - [ClouddriverConfg](#proto.ClouddriverConfg)
  
  
  
  

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
    - [Kubernetes](#proto.Kubernetes)
    - [Kubernetes.Account](#proto.Kubernetes.Account)
  
    - [Kubernetes.ProviderVersion](#proto.Kubernetes.ProviderVersion)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="clouddriver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## clouddriver.proto



<a name="proto.ClouddriverConfg"></a>

### ClouddriverConfg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [Kubernetes](#proto.Kubernetes) |  |  |





 

 

 

 



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



<a name="proto.Kubernetes"></a>

### Kubernetes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [string](#string) |  |  |
| accounts | [Kubernetes.Account](#proto.Kubernetes.Account) | repeated |  |






<a name="proto.Kubernetes.Account"></a>

### Kubernetes.Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| providerVersion | [Kubernetes.ProviderVersion](#proto.Kubernetes.ProviderVersion) |  |  |
| kinds | [string](#string) | repeated |  |
| omitKinds | [string](#string) | repeated |  |





 


<a name="proto.Kubernetes.ProviderVersion"></a>

### Kubernetes.ProviderVersion


| Name | Number | Description |
| ---- | ------ | ----------- |
| V1 | 0 |  |
| V2 | 1 |  |


 

 

 



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

