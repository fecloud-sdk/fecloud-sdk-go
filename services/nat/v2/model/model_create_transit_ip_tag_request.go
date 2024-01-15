package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTransitIpTagRequest struct {
	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateTransitIpTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTransitIpTagRequest", string(data)}, " ")
}
