package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status string `json:"status"`

	Port string `json:"port"`

	Mode string `json:"mode"`

	Region string `json:"region"`

	Datastore *ListInstancesDatastoreResult `json:"datastore"`

	Engine string `json:"engine"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	DbUserName string `json:"db_user_name"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	BackupStrategy *ListInstancesBackupStrategyResult `json:"backup_strategy"`

	PayMode string `json:"pay_mode"`

	MaintenanceWindow string `json:"maintenance_window"`

	Groups []ListInstancesGroupResult `json:"groups"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	TimeZone string `json:"time_zone"`

	Actions []string `json:"actions"`

	LbIpAddress *string `json:"lb_ip_address,omitempty"`

	LbPort *string `json:"lb_port,omitempty"`
}

func (o ListInstancesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResult struct{}"
	}

	return strings.Join([]string{"ListInstancesResult", string(data)}, " ")
}
