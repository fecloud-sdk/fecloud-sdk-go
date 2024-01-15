package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteTransitIpTagRequest struct {
	Key string `json:"key"`

	ResourceId string `json:"resource_id"`
}

func (o DeleteTransitIpTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpTagRequest", string(data)}, " ")
}
