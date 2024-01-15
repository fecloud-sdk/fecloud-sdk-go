package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Share struct {
	AvailabilityZone string `json:"availability_zone"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Name string `json:"name"`

	SecurityGroupId string `json:"security_group_id"`

	ShareProto string `json:"share_proto"`

	ShareType string `json:"share_type"`

	Size int32 `json:"size"`

	SubnetId string `json:"subnet_id"`

	VpcId string `json:"vpc_id"`

	BackupId *string `json:"backup_id,omitempty"`
}

func (o Share) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Share struct{}"
	}

	return strings.Join([]string{"Share", string(data)}, " ")
}
