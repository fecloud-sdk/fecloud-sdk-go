package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBrokersRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListBrokersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersRequest struct{}"
	}

	return strings.Join([]string{"ListBrokersRequest", string(data)}, " ")
}
