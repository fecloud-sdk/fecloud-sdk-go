package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreNewInstanceResponse struct {
	Id *string `json:"id,omitempty"`

	Datastore *Datastore `json:"datastore,omitempty"`

	Name *string `json:"name,omitempty"`

	Created *string `json:"created,omitempty"`

	Status *string `json:"status,omitempty"`

	Region *string `json:"region,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	Mode *string `json:"mode,omitempty"`

	Configurations *[]RestoreNewInstanceConfigurationsOption `json:"configurations,omitempty"`

	Flavor *[]RestoreNewInstanceFlavorOption `json:"flavor,omitempty"`

	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SslOption *string `json:"ssl_option,omitempty"`

	DssPoolId *string `json:"dss_pool_id,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ChargeInfo     *ChargeInfoResult `json:"charge_info,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RestoreNewInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceResponse", string(data)}, " ")
}
