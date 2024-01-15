package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceRequestBody struct {
	Name string `json:"name"`

	Datastore *Datastore `json:"datastore"`

	Region string `json:"region"`

	AvailabilityZone string `json:"availability_zone"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	Port *string `json:"port,omitempty"`

	Password *string `json:"password,omitempty"`

	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	Mode string `json:"mode"`

	Configurations *[]CreateInstanceConfigurationsOption `json:"configurations,omitempty"`

	Flavor []CreateInstanceFlavorOption `json:"flavor"`

	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SslOption *string `json:"ssl_option,omitempty"`

	DssPoolId *string `json:"dss_pool_id,omitempty"`

	ServerGroupPolicies *[]string `json:"server_group_policies,omitempty"`

	Tags *[]TagWithKeyValue `json:"tags,omitempty"`

	ChargeInfo *ChargeInfoOption `json:"charge_info,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
