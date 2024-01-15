package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RecycleInstance struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Mode *string `json:"mode,omitempty"`

	Datastore *RecycleDatastore `json:"datastore,omitempty"`

	PayMode *string `json:"pay_mode,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	DeletedAt *string `json:"deleted_at,omitempty"`

	RetainedUntil *string `json:"retained_until,omitempty"`
}

func (o RecycleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstance struct{}"
	}

	return strings.Join([]string{"RecycleInstance", string(data)}, " ")
}
