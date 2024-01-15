package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteResourceInstanceTagRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Key string `json:"key"`
}

func (o DeleteResourceInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceInstanceTagRequest", string(data)}, " ")
}
