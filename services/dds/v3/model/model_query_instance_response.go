package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryInstanceResponse struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Remark string `json:"remark"`

	Status string `json:"status"`

	Port string `json:"port"`

	Mode string `json:"mode"`

	Region string `json:"region"`

	Datastore *DatastoreItem `json:"datastore"`

	Engine string `json:"engine"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	DbUserName string `json:"db_user_name"`

	Ssl int32 `json:"ssl"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	BackupStrategy *BackupStrategyForItemResponse `json:"backup_strategy"`

	PayMode *string `json:"pay_mode,omitempty"`

	MaintenanceWindow string `json:"maintenance_window"`

	Groups []GroupResponseItem `json:"groups"`

	DiskEncryptionId string `json:"disk_encryption_id"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	TimeZone string `json:"time_zone"`

	DssPoolId *string `json:"dss_pool_id,omitempty"`

	Actions []string `json:"actions"`

	OrderId *string `json:"order_id,omitempty"`

	Tags []TagResponse `json:"tags"`
}

func (o QueryInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInstanceResponse struct{}"
	}

	return strings.Join([]string{"QueryInstanceResponse", string(data)}, " ")
}
