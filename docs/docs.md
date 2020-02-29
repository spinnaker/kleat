# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [appengine.proto](#appengine.proto)
    - [Appengine](#proto.Appengine)
    - [AppengineAccount](#proto.AppengineAccount)
  
  
  
  

- [aws.proto](#aws.proto)
    - [Aws](#proto.Aws)
    - [AwsAccount](#proto.AwsAccount)
    - [AwsBakeryDefaults](#proto.AwsBakeryDefaults)
    - [AwsBaseImage](#proto.AwsBaseImage)
    - [AwsBaseImageSettings](#proto.AwsBaseImageSettings)
    - [AwsDefaults](#proto.AwsDefaults)
    - [AwsFeatures](#proto.AwsFeatures)
    - [AwsFeatures.CloudFormation](#proto.AwsFeatures.CloudFormation)
    - [AwsLifecycleHook](#proto.AwsLifecycleHook)
    - [AwsRegion](#proto.AwsRegion)
    - [AwsVirtualizationSettings](#proto.AwsVirtualizationSettings)
  
  
  
  

- [clouddriver.proto](#clouddriver.proto)
    - [ClouddriverConfig](#proto.ClouddriverConfig)
  
  
  
  

- [front50.proto](#front50.proto)
    - [Front50Config](#proto.Front50Config)
    - [Front50Config.Spinnaker](#proto.Front50Config.Spinnaker)
  
  
  
  

- [google.proto](#google.proto)
    - [Consul](#proto.Consul)
    - [Google](#proto.Google)
    - [GoogleAccount](#proto.GoogleAccount)
    - [GoogleBakeryDefaults](#proto.GoogleBakeryDefaults)
    - [GoogleBaseImage](#proto.GoogleBaseImage)
    - [GoogleBaseImageSettings](#proto.GoogleBaseImageSettings)
    - [VirtualizationSettings](#proto.VirtualizationSettings)
  
  
  
  

- [halconfig.proto](#halconfig.proto)
    - [HalConfig](#proto.HalConfig)
    - [HalConfig.PersistentStorage](#proto.HalConfig.PersistentStorage)
    - [HalConfig.Providers](#proto.HalConfig.Providers)
  
  
  
  

- [kubernetes.proto](#kubernetes.proto)
    - [DockerRegistry](#proto.DockerRegistry)
    - [Kubernetes](#proto.Kubernetes)
    - [KubernetesAccount](#proto.KubernetesAccount)
    - [KubernetesCachingPolicy](#proto.KubernetesCachingPolicy)
    - [KubernetesCustomResource](#proto.KubernetesCustomResource)
  
  
  
  

- [permissions.proto](#permissions.proto)
    - [Permissions](#proto.Permissions)
  
  
  
  

- [persistentstorage.proto](#persistentstorage.proto)
    - [GCS](#proto.GCS)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="appengine.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## appengine.proto



<a name="proto.Appengine"></a>

### Appengine
Configuration for the Google App Engine (GAE) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AppengineAccount](#proto.AppengineAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.AppengineAccount"></a>

### AppengineAccount
Configuration for an App Engine account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cachingIntervalSeconds | [int32](#int32) |  | The interval in seconds at which Spinnaker will poll for updates in your App Engine clusters. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| gcloudReleaseTrack | [string](#string) |  | The gcloud release track (`ALPHA`, `BETA`, or `STABLE`) that Spinnaker will use when deploying to App Engine. |
| gitHttpsUsername | [string](#string) |  | A username to be used when connecting to a remote git repository server over HTTPS. If set, `gitHttpsPassword` must also be set. |
| gitHttpsPassword | [string](#string) |  | A password to be used when connecting to a remote git repository server over HTTPS. If set, `gitHttpsUsername` must also be set. |
| githubOAuthAccessToken | [string](#string) |  | An OAuth token provided by Github for connecting to a git repository over HTTPS. See https://help.github.com/articles/creating-an-access-token-for-command-line-use for more information. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| localRepositoryDirectory | [string](#string) |  | A local directory to be used to stage source files for App Engine deployments within Clouddriver. Defaults to `/var/tmp/clouddriver`. |
| omitServices | [string](#string) | repeated | A list of regular expressions. Any service matching one of these regexes will be ignored by Spinnaker. |
| omitVersions | [string](#string) | repeated | A list of regular expressions. Any version matching one of these regexes will be ignored by Spinnaker. |
| project | [string](#string) |  | The Google Cloud Platform project this Spinnaker account will manage. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| services | [string](#string) | repeated | A list of regular expressions. Any service matching one of these regexes will be indexed by Spinnaker (unless the service also matches a regex in `omitServices`). |
| sshKnownHostsFilePath | [string](#string) |  | The path to a `known_hosts` file to be used when connecting with a remote git repository over SSH. |
| sshPrivateKeyFilePath | [string](#string) |  | The path to an SSH private key to be used when connecting with a remote git repository over SSH. If set, `sshPrivateKeyPassphrase` must also be set. |
| sshPrivateKeyPassphrase | [string](#string) |  | The passphrase to an SSH private key to be used when connecting with a remote git repository over SSH. If set, `sshPrivateKeyFilePath` must also be set. |
| sshTrustUnknownHosts | [bool](#bool) |  | Enabling this flag will allow Spinnaker to connect with a remote git repository over SSH without verifying the server’s IP address against a `known_hosts` file. Defaults to false. |
| versions | [string](#string) | repeated | A list of regular expressions. Any version matching one of these regexes will be indexed by Spinnaker (unless the version also matches a regex in `omitVersions`). |
| name | [string](#string) |  | The name of the account. |





 

 

 

 



<a name="aws.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## aws.proto



<a name="proto.Aws"></a>

### Aws
Configuration for the AWS provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [AwsAccount](#proto.AwsAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| accessKeyId | [string](#string) |  | Your AWS Access Key ID. Note that if you are baking AMIs with Rosco, you may also need to set `AwsBakeryDefaults.awsAccessKey`. |
| secretAccessKey | [string](#string) |  | Your AWS Secret Key. Note that if you are baking AMIs with Rosco, you may also need to set `AwsBakeryDefaults.awsSecretKey`. |
| defaultKeyPairTemplate | [string](#string) |  | Halyard does not expose a command to edit this field, but defaults it to `{{name}}-keypair`. TODO(mneterval): Can we move to Clouddriver? |
| defaultRegions | [AwsRegion](#proto.AwsRegion) | repeated | List of default regions. |
| defaults | [AwsDefaults](#proto.AwsDefaults) |  | Defaults relevant to the AWS provider. TODO(mneterval): Can we move to Clouddriver? |
| features | [AwsFeatures](#proto.AwsFeatures) |  | Configuration for AWS-specific features. |
| bakeryDefaults | [AwsBakeryDefaults](#proto.AwsBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.AwsAccount"></a>

### AwsAccount
Configuration for an AWS account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountId | [string](#string) |  | Your AWS account ID to manage. See http://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html for more information. |
| assumeRole | [string](#string) |  | If set, Spinnaker will configure a credentials provider that uses AWS Security Token Service to assume the specified role. Examples: `user/spinnaker`, `role/spinnakerManaged`. |
| defaultKeyPair | [string](#string) |  | The name of the AWS key-pair to use. See http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html for more information. |
| discovery | [string](#string) |  | The endpoint at which your Eureka discovery system is reachable. See https://github.com/Netflix/eureka for more information. Example: `http://.eureka.url.to.use:8080/eureka-server/v2`. Using will make Spinnaker use AWS regions in the hostname to access discovery so that you can have discovery for multiple regions. |
| edda | [string](#string) |  | The endpoint at which Edda is reachable. Edda is not a hard dependency of Spinnaker, but is helpful for reducing the request volume against AWS. See https://github.com/Netflix/edda for more information. |
| environment | [string](#string) |  | The environment name for the account. Many accounts can share the same environment (e.g., dev, test, prod). |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| lifecycleHooks | [AwsLifecycleHook](#proto.AwsLifecycleHook) | repeated | List of configured AWS lifecycle hooks. |
| regions | [AwsRegion](#proto.AwsRegion) | repeated | List of configured AWS regions. |
| name | [string](#string) |  | The name of the account. |






<a name="proto.AwsBakeryDefaults"></a>

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
| baseImages | [AwsBaseImageSettings](#proto.AwsBaseImageSettings) | repeated | List of configured base images. |
| templateFile | [string](#string) |  | This is the name of the packer template that will be used to bake images from this base image. The template file must be found in this list https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.AwsBaseImage"></a>

### AwsBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | This is the identifier used by AWS to find this base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker’s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |






<a name="proto.AwsBaseImageSettings"></a>

### AwsBaseImageSettings
Configuration for a base image for the AWS provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [AwsBaseImage](#proto.AwsBaseImage) |  | Base image configuration. |
| virtualizationSettings | [AwsVirtualizationSettings](#proto.AwsVirtualizationSettings) |  | Base image virtualization settings. |






<a name="proto.AwsDefaults"></a>

### AwsDefaults
Defaults relevant to the AWS provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iamRole | [string](#string) |  | Halyard does not expose a command to edit this field, but defaults it to `BaseIAMRole`. TODO(mneterval): Can we move to Clouddriver? |






<a name="proto.AwsFeatures"></a>

### AwsFeatures
Configuration for AWS-specific features.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cloudFormation | [AwsFeatures.CloudFormation](#proto.AwsFeatures.CloudFormation) |  | Configuration for AWS CloudFormation. |






<a name="proto.AwsFeatures.CloudFormation"></a>

### AwsFeatures.CloudFormation
Configuration for AWS CloudFormation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether AWS CloudFormation is enabled. |






<a name="proto.AwsLifecycleHook"></a>

### AwsLifecycleHook
Configuration for AWS Auto Scaling Lifecycle Hooks. For more information, see:
https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| defaultResult | [string](#string) |  | Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. This parameter can be either CONTINUE or ABANDON. The default value is ABANDON. |
| heartbeatTimeout | [int32](#int32) |  | Set the heartbeat timeout in seconds for the lifecycle hook. Instances can remain in a wait state for a finite period of time. Must be greater than or equal to 30 and less than or equal to 7200. The default is 3600 (one hour). |
| lifecycleTransition | [string](#string) |  | Type of lifecycle transition. Acceptable values: `autoscaling:EC2_INSTANCE_LAUNCHING`, `autoscaling:EC2_INSTANCE_TERMINATING` |
| notificationTargetARN | [string](#string) |  | The ARN of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. This target can be either an SQS queue or an SNS topic. |
| roleARN | [string](#string) |  | The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue. |






<a name="proto.AwsRegion"></a>

### AwsRegion
An AWS region.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the region. |






<a name="proto.AwsVirtualizationSettings"></a>

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





 

 

 

 



<a name="clouddriver.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## clouddriver.proto



<a name="proto.ClouddriverConfig"></a>

### ClouddriverConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetes | [Kubernetes](#proto.Kubernetes) |  |  |
| google | [Google](#proto.Google) |  |  |
| appengine | [Appengine](#proto.Appengine) |  |  |
| aws | [Aws](#proto.Aws) |  |  |





 

 

 

 



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





 

 

 

 



<a name="google.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## google.proto



<a name="proto.Consul"></a>

### Consul
Configuration for Consul.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether Consul is enabled. |
| agentEndpoint | [string](#string) |  | Reachable Consul node endpoint connected to the Consul cluster. Defaults to localhost. |
| agentPort | [int32](#int32) |  | Port consul is running on for every agent. Defaults to 8500. |
| datacenters | [string](#string) | repeated | List of data centers to cache and keep updated. Defaults to all. |






<a name="proto.Google"></a>

### Google
Configuration for the Google Compute Engine (GCE) provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [GoogleAccount](#proto.GoogleAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |
| bakeryDefaults | [GoogleBakeryDefaults](#proto.GoogleBakeryDefaults) |  | Configuration for Spinnaker&#39;s image bakery. |






<a name="proto.GoogleAccount"></a>

### GoogleAccount
Configuration for a Spinnaker Google account. An account maps to a
credential that can authenticate against a GCP project.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| project | [string](#string) |  | The GCP project this Spinnaker account will manage. |
| jsonPath | [string](#string) |  | The path to a JSON service account that Spinnaker will use as credentials. This is only needed if Spinnaker is not deployed on a Google Compute Engine VM, or needs permissions not afforded to the VM it is running on. See https://cloud.google.com/compute/docs/access/service-accounts for more information. |
| alphaListed | [bool](#bool) |  | Enable this flag if your GCP project has access to alpha features and you want Spinnaker to take advantage of them. |
| imageProjects | [string](#string) | repeated | A list of GCP projects from which Spinnaker will be able to cache and deploy images. When this is omitted, it defaults to the current project. Each project must have granted the IAM role compute.imageUser to the service account associated with the JSON key used by this account, as well as to the Google APIs service account automatically created for the project being managed (should look similar to 12345678912@cloudservices.gserviceaccount.com). See https://cloud.google.com/compute/docs/images/sharing-images-across-projects for more information about sharing images across GCP projects. |
| consul | [Consul](#proto.Consul) |  | Configuration for Consul. |
| regions | [string](#string) | repeated | A list of regions for caching and mutating calls. This overwrites any default regions set on the provider. |
| userDataFile | [string](#string) |  | The path to user data template file. Spinnaker has the ability to inject userdata into generated instance templates. The mechanism is via a template file that is token replaced to provide some specifics about the deployment. See https://github.com/spinnaker/clouddriver/blob/master/clouddriver-aws/UserData.md for more information. |






<a name="proto.GoogleBakeryDefaults"></a>

### GoogleBakeryDefaults
Configuration for Spinnaker&#39;s image bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateFile | [string](#string) |  | The name of the Packer template that will be used to bake images from this base image. The template file must be found in this list: https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or supplied as described here: https://spinnaker.io/setup/bakery/. |
| baseImages | [GoogleBaseImageSettings](#proto.GoogleBaseImageSettings) | repeated | List of configured base images. |
| zone | [string](#string) |  | The default zone in which to bake an image. |
| network | [string](#string) |  | The Google Compute network ID or URL to use for the launched instance. Defaults to default. |
| useInternalIp | [bool](#bool) |  | If true, use the instance&#39;s internal IP instead of its external IP during baking. |
| networkProjectId | [string](#string) |  | The default project ID for the network and subnet to use for the VM baking your image. |






<a name="proto.GoogleBaseImage"></a>

### GoogleBaseImage
Base image configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | This is the identifier used by GCP to find this base image. |
| shortDescription | [string](#string) |  | A short description to help human operators identify the image. |
| detailedDescription | [string](#string) |  | A long description to help human operators identify the image. |
| packageType | [string](#string) |  | This is used to help Spinnaker’s bakery download the build artifacts you supply it with. For example, specifying deb indicates that your artifacts will need to be fetched from a debian repository. |
| imageFamily | [bool](#bool) |  | If set to true, Deck will annotate the popup tooltip to indicate that the selected option represents an image family. |






<a name="proto.GoogleBaseImageSettings"></a>

### GoogleBaseImageSettings
Configuration for a base image for the Google provider&#39;s bakery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| baseImage | [GoogleBaseImage](#proto.GoogleBaseImage) |  | Base image configuration. |
| virtualizationSettings | [VirtualizationSettings](#proto.VirtualizationSettings) |  | Image source configuration. |






<a name="proto.VirtualizationSettings"></a>

### VirtualizationSettings
Image source configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceImage | [string](#string) |  | The source image. If both sourceImage and sourceImageFamily are set, sourceImage will take precedence. |
| sourceImageFamily | [string](#string) |  | The source image family to create the image from. The newest, non-deprecated image is used. If both sourceImage and sourceImageFamily are set, sourceImage will take precedence. |





 

 

 

 



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
| appengine | [Appengine](#proto.Appengine) |  |  |
| aws | [Aws](#proto.Aws) |  |  |





 

 

 

 



<a name="kubernetes.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## kubernetes.proto



<a name="proto.DockerRegistry"></a>

### DockerRegistry
Configuration for a Docker registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accountName | [string](#string) |  | The configured name of the Docker registry. |
| namespaces | [string](#string) | repeated | The list of Docker registry namespaces usable as image sources. |






<a name="proto.Kubernetes"></a>

### Kubernetes
Configuration for the Kubernetes provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  | Whether the provider is enabled. |
| accounts | [KubernetesAccount](#proto.KubernetesAccount) | repeated | The list of configured accounts. |
| primaryAccount | [string](#string) |  | The name of the primary account. |






<a name="proto.KubernetesAccount"></a>

### KubernetesAccount
Configuration for a Spinnaker Kubernetes account. An account maps to a
credential that can authenticate against your Kubernetes cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the account. |
| providerVersion | [string](#string) |  | Some providers support multiple versions/release tracks. This allows you to pick the version of the provider (not the resources it manages) to run within Spinnaker. |
| kinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account can deploy and will cache. When no kinds are configured, this defaults to all kinds described here: https://spinnaker.io/reference/providers/kubernetes-v2/. This can only be set when omitKinds is empty or not set. |
| omitKinds | [string](#string) | repeated | A list of resource kinds this Spinnaker account cannot deploy to or cache. This can only be set when kinds is empty or not set. |
| context | [string](#string) |  | The kubernetes context to be managed by Spinnaker. See http://kubernetes.io/docs/user-guide/kubeconfig-file/#context for more information. When no context is configured for an account the ‘current-context’ in your kubeconfig is assumed. |
| cacheThreads | [int32](#int32) |  | Number of caching agents for this kubernetes account. Each agent handles a subset of the namespaces available to this account. By default, only 1 agent caches all kinds for all namespaces in the account. |
| namespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account can deploy to and will cache. When no namespaces are configured, this defaults to all namespaces. |
| omitNamespaces | [string](#string) | repeated | A list of namespaces this Spinnaker account cannot deploy to or cache. This can only be set when namespaces is empty or not set. |
| customResources | [KubernetesCustomResource](#proto.KubernetesCustomResource) | repeated | The list of custom resources Clouddriver will manage and make available for use in Patch and Delete (Manifest) stages. |
| cachingPolicies | [KubernetesCachingPolicy](#proto.KubernetesCachingPolicy) | repeated | The list of kind-specific caching policies. |
| dockerRegistries | [DockerRegistry](#proto.DockerRegistry) | repeated | The list of the Spinnaker docker registry account names this Spinnaker account can use as image sources. These docker registry accounts must be registered in your halconfig before you can add them here. |
| oAuthScopes | [string](#string) | repeated | The list of OAuth scopes used by kubectl to fetch an OAuth token. |
| kubeconfigFile | [string](#string) |  | The path to your kubeconfig file. By default, it will be under the Spinnaker user’s home directory in the typical .kube/config location. todo: document new var/secrets convention. |
| permissions | [Permissions](#proto.Permissions) |  | Fiat permissions configuration. |
| requiredGroupMemberships | [string](#string) | repeated | (Deprecated): List of required Fiat permission groups. Configure `permissions` instead. |






<a name="proto.KubernetesCachingPolicy"></a>

### KubernetesCachingPolicy
Configuration for a kind-specific caching policy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubernetesKind | [string](#string) |  | The Kubernetes kind to which the policy applies. |
| maxEntriesPerAgent | [int32](#int32) |  | The maximum number of resources an agent will cache of the specified Kubernetes kind. |






<a name="proto.KubernetesCustomResource"></a>

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





 

 

 

 



<a name="permissions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## permissions.proto



<a name="proto.Permissions"></a>

### Permissions
A Fiat permissions configuration object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| READ | [string](#string) | repeated | A user must have at least one of these roles in order to view this account’s cloud resources. |
| WRITE | [string](#string) | repeated | A user must have at least one of these roles in order to make changes to this account’s cloud resources. |
| EXECUTE | [string](#string) | repeated | A user must have at least one of these roles in order to execute pipelines. |





 

 

 

 



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

