package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupsRespGroupAssignment struct {
	Topic *string `json:"topic,omitempty"`

	Partitions *[]int32 `json:"partitions,omitempty"`
}

func (o ShowGroupsRespGroupAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupAssignment struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupAssignment", string(data)}, " ")
}
