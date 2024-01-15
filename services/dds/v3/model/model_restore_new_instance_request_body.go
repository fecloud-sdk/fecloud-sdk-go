package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreNewInstanceRequestBody struct {
	Name string `json:"name"`

	AvailabilityZone string `json:"availability_zone"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	Password *string `json:"password,omitempty"`

	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	Configurations *[]RestoreNewInstanceConfigurationsOption `json:"configurations,omitempty"`

	Flavor []RestoreNewInstanceFlavorOption `json:"flavor"`

	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SslOption *string `json:"ssl_option,omitempty"`

	DssPoolId *string `json:"dss_pool_id,omitempty"`

	ServerGroupPolicies *[]string `json:"server_group_policies,omitempty"`

	RestorePoint *RestorePoint `json:"restore_point"`

	ChargeInfo *ChargeInfoOption `json:"charge_info,omitempty"`
}

func (o RestoreNewInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceRequestBody", string(data)}, " ")
}
