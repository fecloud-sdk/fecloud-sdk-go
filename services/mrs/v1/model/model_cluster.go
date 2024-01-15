package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Cluster struct {
	ClusterId *string `json:"clusterId,omitempty"`

	ClusterName *string `json:"clusterName,omitempty"`

	MasterNodeNum *string `json:"masterNodeNum,omitempty"`

	CoreNodeNum *string `json:"coreNodeNum,omitempty"`

	TotalNodeNum *string `json:"totalNodeNum,omitempty"`

	ClusterState *string `json:"clusterState,omitempty"`

	CreateAt *string `json:"createAt,omitempty"`

	UpdateAt *string `json:"updateAt,omitempty"`

	DataCenter *string `json:"dataCenter,omitempty"`

	Vpc *string `json:"vpc,omitempty"`

	VpcId *string `json:"vpcId,omitempty"`

	Fee *string `json:"fee,omitempty"`

	HadoopVersion *string `json:"hadoopVersion,omitempty"`

	ComponentList *[]ComponentAmb `json:"componentList,omitempty"`

	ExternalIp *string `json:"externalIp,omitempty"`

	ExternalAlternateIp *string `json:"externalAlternateIp,omitempty"`

	InternalIp *string `json:"internalIp,omitempty"`

	DeploymentId *string `json:"deploymentId,omitempty"`

	Remark *string `json:"remark,omitempty"`

	OrderId *string `json:"orderId,omitempty"`

	AzId *string `json:"azId,omitempty"`

	AzName *string `json:"azName,omitempty"`

	AzCode *string `json:"azCode,omitempty"`

	AvailabilityZoneId *string `json:"availabilityZoneId,omitempty"`

	InstanceId *string `json:"instanceId,omitempty"`

	Vnc *string `json:"vnc,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	VolumeSize *int32 `json:"volumeSize,omitempty"`

	VolumeType *string `json:"volumeType,omitempty"`

	SubnetId *string `json:"subnetId,omitempty"`

	SubnetName *string `json:"subnetName,omitempty"`

	SecurityGroupsId *string `json:"securityGroupsId,omitempty"`

	SlaveSecurityGroupsId *string `json:"slaveSecurityGroupsId,omitempty"`

	BootstrapScripts *[]BootstrapScript `json:"bootstrapScripts,omitempty"`

	StageDesc *string `json:"stageDesc,omitempty"`

	SafeMode *int32 `json:"safeMode,omitempty"`

	ClusterVersion *string `json:"clusterVersion,omitempty"`

	NodePublicCertName *string `json:"nodePublicCertName,omitempty"`

	MasterNodeIp *string `json:"masterNodeIp,omitempty"`

	PrivateIpFirst *string `json:"privateIpFirst,omitempty"`

	ErrorInfo *string `json:"errorInfo,omitempty"`

	Tags *string `json:"tags,omitempty"`

	MasterNodeSize *string `json:"masterNodeSize,omitempty"`

	CoreNodeSize *string `json:"coreNodeSize,omitempty"`

	MasterNodeProductId *string `json:"masterNodeProductId,omitempty"`

	MasterNodeSpecId *string `json:"masterNodeSpecId,omitempty"`

	CoreNodeProductId *string `json:"coreNodeProductId,omitempty"`

	CoreNodeSpecId *string `json:"coreNodeSpecId,omitempty"`

	MasterDataVolumeType *string `json:"masterDataVolumeType,omitempty"`

	MasterDataVolumeSize *int32 `json:"masterDataVolumeSize,omitempty"`

	MasterDataVolumeCount *int32 `json:"masterDataVolumeCount,omitempty"`

	CoreDataVolumeType *string `json:"coreDataVolumeType,omitempty"`

	CoreDataVolumeSize *int32 `json:"coreDataVolumeSize,omitempty"`

	CoreDataVolumeCount *int32 `json:"coreDataVolumeCount,omitempty"`

	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	IsMrsManagerFinish *bool `json:"isMrsManagerFinish,omitempty"`

	ClusterType *int32 `json:"clusterType,omitempty"`

	LogCollection *int32 `json:"logCollection,omitempty"`

	Scale *string `json:"scale,omitempty"`

	NodeGroups *[]NodeGroupV10 `json:"nodeGroups,omitempty"`

	TaskNodeGroups *[]NodeGroupV10 `json:"taskNodeGroups,omitempty"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
