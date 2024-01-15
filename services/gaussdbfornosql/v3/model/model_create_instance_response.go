package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceResponse struct {
	Id *string `json:"id,omitempty"`

	Datastore *DatastoreResult `json:"datastore,omitempty"`

	Name *string `json:"name,omitempty"`

	Created *string `json:"created,omitempty"`

	Status *string `json:"status,omitempty"`

	Region *string `json:"region,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	Mode *string `json:"mode,omitempty"`

	Flavor *[]CreateInstanceFlavorResult `json:"flavor,omitempty"`

	BackupStrategy *BackupStrategyResult `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	SslOption *string `json:"ssl_option,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ChargeInfo     *ChargeInfoResult `json:"charge_info,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}
