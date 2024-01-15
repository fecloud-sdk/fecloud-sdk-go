package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConfigHistoriesRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigHistoriesRequest", string(data)}, " ")
}
