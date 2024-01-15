package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupResponse struct {
	Enabled *bool `json:"enabled,omitempty"`

	Broadcast *bool `json:"broadcast,omitempty"`

	Brokers *[]string `json:"brokers,omitempty"`

	Name *string `json:"name,omitempty"`

	GroupDesc *string `json:"group_desc,omitempty"`

	RetryMaxTime float32 `json:"retry_max_time,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	AppName *string `json:"app_name,omitempty"`

	Permissions    *[]string `json:"permissions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupResponse", string(data)}, " ")
}
