package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceRequestBody struct {
	Name string `json:"name"`

	Datastore *DatastoreOption `json:"datastore"`

	Region string `json:"region"`

	AvailabilityZone string `json:"availability_zone"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	Password string `json:"password"`

	Mode string `json:"mode"`

	Flavor []CreateInstanceFlavorOption `json:"flavor"`

	ConfigurationId *string `json:"configuration_id,omitempty"`

	BackupStrategy *BackupStrategyOption `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	SslOption *string `json:"ssl_option,omitempty"`

	ChargeInfo *ChargeInfoOption `json:"charge_info,omitempty"`

	RestoreInfo *RestoreInfo `json:"restore_info,omitempty"`

	Port *string `json:"port,omitempty"`

	AvailabilityZoneDetail *AvailabilityZoneDetail `json:"availability_zone_detail,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
