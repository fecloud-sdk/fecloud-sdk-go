package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateBigkeyAutoscanConfigResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	EnableAutoScan *bool `json:"enable_auto_scan,omitempty"`

	ScheduleAt *[]string `json:"schedule_at,omitempty"`

	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBigkeyAutoscanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBigkeyAutoscanConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateBigkeyAutoscanConfigResponse", string(data)}, " ")
}
