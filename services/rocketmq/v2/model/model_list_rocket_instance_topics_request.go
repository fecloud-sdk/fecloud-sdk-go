package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRocketInstanceTopicsRequest struct {
	InstanceId string `json:"instance_id"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRocketInstanceTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketInstanceTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListRocketInstanceTopicsRequest", string(data)}, " ")
}
