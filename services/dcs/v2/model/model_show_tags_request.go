package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTagsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTagsRequest", string(data)}, " ")
}
